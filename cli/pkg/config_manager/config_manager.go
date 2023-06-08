package config_manager

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/asaskevich/govalidator"
	backend "github.com/chanzuckerberg/happy/shared/backend/aws"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/k8s"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Name           string
	ServiceType    string
	Context        string
	DockerfilePath string
	Port           int
	Uri            string
	Priority       int
}

func CreeateHappyConfig(ctx context.Context, bootstrapConfig *config.Bootstrap) (*config.HappyConfig, error) {
	happyConfig, err := config.NewBlankHappyConfig(bootstrapConfig)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get happy config")
	}

	dockerPaths, err := findAllDockerfiles(bootstrapConfig.HappyProjectRoot)
	if err != nil {
		return nil, errors.Wrap(err, "unable to find dockerfiles")
	}

	if len(dockerPaths) == 0 {
		return nil, errors.New("no dockerfiles found in this repo")
	}

	logrus.Info("Welcome to happy bootstrap! We'll ask you a few questions to get started.")

	appName := ""
	prompt1 := &survey.Input{
		Message: "What would you like to name this application?",
		Help:    "This will be the unique name of the application, lowercased and hyphenated",
	}
	err = survey.AskOne(prompt1, &appName, survey.WithValidator(survey.Required), survey.WithValidator(survey.MinLength(5)))
	if err != nil {
		return nil, errors.Wrap(err, "unable to prompt")
	}
	if len(appName) == 0 {
		return nil, errors.New("no application name provided")
	}

	profiles, err := util.GetAwsProfiles()
	if err != nil {
		return nil, errors.Wrap(err, "unable to retrieve aws profiles")
	}
	if len(profiles) == 0 {
		return nil, errors.New("no aws profiles found")
	}

	environmentNames := []string{}
	prompt := []*survey.Question{
		{
			Name: "environments",
			Prompt: &survey.MultiSelect{
				Message: "Your application will be deployed to multiple environments. Which environments would you like to deploy to?",
				Options: []string{"rdev", "dev", "staging", "prod"},
			},
		},
	}

	err = survey.Ask(prompt, &environmentNames)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to prompt")
	}
	if len(environmentNames) == 0 {
		return nil, errors.New("no environments were selected")
	}

	environments := map[string]config.Environment{}
	for _, env := range environmentNames {
		logrus.Infof("A few questions about environment  '%s':", env)

		profile := ""
		region := ""
		clusterId := ""
		happyNamespace := ""

		for {
			prompt := &survey.Select{
				Message: fmt.Sprintf("Which aws profile do you want to use in %s?", env),
				Options: profiles,
				Default: profiles[0],
			}

			err = survey.AskOne(prompt, &profile)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to prompt")
			}
			if len(profile) == 0 {
				continue
			}

			var clusterIds []string
			prompt1 := &survey.Select{
				Message: fmt.Sprintf("Which aws region should we use in %s? (us-west-1 should be avoided, if possible)", env),
				Options: []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2"},
				Default: "us-west-2",
			}

			err = survey.AskOne(prompt1, &region)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to prompt")
			}
			logrus.Info("Checking for eks clusters in this region...")

			clusterIds, err = ListClusterIds(ctx, profile, region)
			if err != nil {
				return nil, errors.Wrap(err, "unable to list eks clusters")
			}
			if len(clusterIds) == 0 {
				logrus.Error("No eks clusters found in this region. Please select a different region or profile.")
				continue
			}

			prompt = &survey.Select{
				Message: fmt.Sprintf("Which EKS cluster should we use in %s?", env),
				Options: clusterIds,
				Default: clusterIds[0],
			}

			err = survey.AskOne(prompt, &clusterId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to prompt")
			}
			var happyNamespaces []string
			happyNamespaces, err = ListHappyNamespaces(ctx, profile, region, clusterId)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to obtain a list of happy namespaces")
			}
			if len(happyNamespaces) == 0 {
				logrus.Error("No happy namespaces were found in the selected cluster, please select a different region, profile or eks cluster.")
				continue
			}

			defaultNamespace := happyNamespaces[0]
			for _, namespace := range happyNamespaces {
				if strings.Contains(namespace, fmt.Sprintf("-%s-happy-env", env)) {
					defaultNamespace = namespace
					break
				}
			}

			prompt2 := &survey.Select{
				Message: fmt.Sprintf("Which happy namespace should we use in %s?", env),
				Options: happyNamespaces,
				Default: defaultNamespace,
			}

			err = survey.AskOne(prompt2, &happyNamespace)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to obtain an eks cluster id")
			}
			break
		}

		environments[env] = config.Environment{
			AWSProfile: util.String(profile),
			AWSRegion:  util.String(region),
			K8S: k8s.K8SConfig{
				Namespace:  happyNamespace,
				ClusterID:  clusterId,
				AuthMethod: "eks",
			},
			TerraformDirectory: fmt.Sprintf(".happy/terraform/envs/%s", env),
			AutoRunMigrations:  false,
			TaskLaunchType:     "k8s",
		}
	}

	services := []Service{}
	logrus.Info("We have found dockerfiles in your project, let's see if you'd like to use them as services in your stack")
	for _, dockerPath := range dockerPaths {
		dockerFileName := filepath.Base(dockerPath)
		contextPath, err := filepath.Rel(bootstrapConfig.HappyProjectRoot, filepath.Dir(dockerPath))
		if err != nil {
			return nil, errors.Wrapf(err, "failed to obtain relative path")
		}

		confirm := false
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("Would you like to use dockerfile %s/%s as a service in your stack?", contextPath, dockerFileName),
		}
		err = survey.AskOne(prompt, &confirm)
		if err != nil {
			return nil, errors.Wrap(err, "unable to prompt")
		}
		if !confirm {
			continue
		}

		serviceName := ""
		prompt1 := &survey.Input{
			Message: fmt.Sprintf("What would you like to name the service for %s?", contextPath),
			Help:    "This will be the name of the service in your stack, lowercased and hyphenated",
			Default: "frontend",
		}
		err = survey.AskOne(prompt1, &serviceName, survey.WithValidator(survey.Required), survey.WithValidator(survey.MinLength(3)))
		if err != nil {
			return nil, errors.Wrap(err, "unable to prompt")
		}
		serviceName = strings.ToLower(serviceName)

		serviceType := "Service is exposed to the internet, and can only be consumed by other services in the stack (PRIVATE)"
		prompt2 := &survey.Select{
			Message: fmt.Sprintf("What kind of service is %s?", serviceName),
			Options: []string{
				"Serivce is exposed to the internet (EXTERNAL)",
				"Service is exposed to the internet, but is protected by Okta (INTERNAL)",
				"Service is exposed to the internet, and can only be consumed by other services in the stack (PRIVATE)",
			},
			Default: serviceType,
		}

		err = survey.AskOne(prompt2, &serviceType)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to obtain an aws profile")
		}

		serviceType = regexp.MustCompile(`\((.*?)\)`).FindStringSubmatch(serviceType)[1]

		port := ""
		prompt3 := &survey.Input{
			Message: fmt.Sprintf("Which port does service %s listen on?", serviceName),
			Default: "3000",
		}
		err = survey.AskOne(prompt3, &port, survey.WithValidator(survey.Required), survey.WithValidator(survey.MinLength(2)), survey.WithValidator(Port))
		if err != nil {
			return nil, errors.Wrap(err, "unable to prompt")
		}

		uri := ""
		prompt3 = &survey.Input{
			Message: fmt.Sprintf("Which uri does %s respond on?", serviceName),
			Help:    "This is the relative path that the service will respond on, e.g. /api/v1",
			Default: "/",
		}
		err = survey.AskOne(prompt3, &uri, survey.WithValidator(survey.Required), survey.WithValidator(survey.MinLength(1)), survey.WithValidator(URI))
		if err != nil {
			return nil, errors.Wrap(err, "unable to prompt")
		}

		uri, _ = strings.CutSuffix(uri, "/")
		portNumber, err := strconv.Atoi(port)

		if err != nil {
			return nil, errors.Wrap(err, "port number is not valid")
		}

		services = append(services, Service{
			Name:           serviceName,
			ServiceType:    serviceType,
			DockerfilePath: dockerFileName,
			Context:        contextPath,
			Port:           portNumber,
			Uri:            uri,
		})
	}

	happyConfig.GetData().Environments = environments
	happyConfig.GetData().FeatureFlags = config.Features{
		EnableDynamoLocking:   true,
		EnableHappyApiUsage:   true,
		EnableECRAutoCreation: true,
	}
	happyConfig.GetData().DefaultEnv = environmentNames[0]
	happyConfig.GetData().DefaultComposeEnvFile = ".env.ecr"

	serviceDefs := map[string]any{}
	stackDefaults := map[string]any{
		"stack_defaults":   "git@github.com:chanzuckerberg/happy//terraform/modules/happy-stack-eks?ref=main",
		"routing_method":   "CONTEXT",
		"create_dashboard": false,
		"app":              appName,
		"services":         serviceDefs,
	}

	serviceNames := []string{}
	// sort services by length of service.Uri in reverse order
	sort.Slice(services, func(i, j int) bool {
		return len(services[i].Uri) > len(services[j].Uri)
	})
	priority := 0
	for _, service := range services {
		service.Priority = priority
		priority++
	}

	for _, service := range services {
		serviceDefs[service.Name] =
			map[string]any{
				"name":                  service.Name,
				"port":                  service.Port,
				"health_check_path":     service.Uri,
				"service_type":          service.ServiceType,
				"path":                  fmt.Sprintf("%s/*", service.Uri),
				"priority":              service.Priority,
				"success_codes":         "200-499",
				"platform_architecture": "arm64",
				"build": map[string]any{
					"context":    service.Context,
					"dockerfile": service.DockerfilePath,
				},
			}
		serviceNames = append(serviceNames, service.Name)
	}
	happyConfig.GetData().StackDefaults = stackDefaults
	happyConfig.GetData().Services = serviceNames

	err = os.MkdirAll(filepath.Dir(bootstrapConfig.HappyConfigPath), 0777)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to create folder %s", filepath.Dir(bootstrapConfig.HappyConfigPath))
	}
	err = happyConfig.Save()

	if err != nil {
		return nil, errors.Wrap(err, "unable to save happy config")
	}

	happyConfig, err = config.NewHappyConfig(bootstrapConfig)

	if err != nil {
		return nil, errors.Wrap(err, "unable to load happy config")
	}

	return happyConfig, nil
}

func ListClusterIds(ctx context.Context, profile, region string) ([]string, error) {
	b, err := backend.NewAWSBackend(ctx, config.EnvironmentContext{
		AWSProfile:     util.String(profile),
		AWSRegion:      util.String(region),
		TaskLaunchType: util.LaunchTypeNull,
	})
	if err != nil {
		return []string{}, errors.Wrap(err, "unable to create an aws backend")
	}
	return b.ListEKSClusterIds(ctx)
}

func ListHappyNamespaces(ctx context.Context, profile, region, clusterId string) ([]string, error) {
	b, err := backend.NewAWSBackend(ctx, config.EnvironmentContext{
		AWSProfile:     util.String(profile),
		AWSRegion:      util.String(region),
		TaskLaunchType: util.LaunchTypeK8S,
		K8S: k8s.K8SConfig{
			AuthMethod: "eks",
			ClusterID:  clusterId,
		},
	}, backend.WithIntegrationSecret(&config.IntegrationSecret{})) // This will prevent the backend from trying to load the integration secret, as we have not selected the namespace yet
	if err != nil {
		return []string{}, errors.Wrap(err, "unable to create an aws backend")
	}
	return b.ListHappyNamespaces(ctx)
}

func findAllDockerfiles(path string) ([]string, error) {
	logrus.Infof("Searching for Dockerfiles in %s", path)
	paths := []string{}

	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == ".terraform" || d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".dockerignore") {
			return nil
		}
		if !strings.HasSuffix(path, "Dockerfile") && !strings.Contains(path, "Dockerfile.") {
			return nil
		}
		paths = append(paths, path)
		return nil
	})

	return paths, err
}

func Port(val interface{}) error {
	if str, ok := val.(string); ok {
		if !govalidator.IsPort(str) {
			return fmt.Errorf("value is not a valid port number")
		}
	} else {
		return fmt.Errorf("cannot enforce a port number check on response of type %v", reflect.TypeOf(val).Name())
	}
	return nil
}

func URI(val interface{}) error {
	if str, ok := val.(string); ok {
		if !govalidator.IsRequestURI(str) {
			return fmt.Errorf("value is not a valid uri")
		}
	} else {
		return fmt.Errorf("cannot enforce a uri check on response of type %v", reflect.TypeOf(val).Name())
	}
	return nil
}
