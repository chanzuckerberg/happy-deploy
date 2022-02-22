package artifact_builder

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecr"
	backend "github.com/chanzuckerberg/happy/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/chanzuckerberg/happy/pkg/util"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ArtifactBuilder struct {
	backend        *backend.Backend
	config         *BuilderConfig
	tags           []string
	registryClient util.RegistryClient
}

func NewArtifactBuilder() *ArtifactBuilder {
	return &ArtifactBuilder{
		config:         nil,
		backend:        nil,
		tags:           []string{},
		registryClient: util.NewDefaultRegistryClient(),
	}
}

func (ab *ArtifactBuilder) WithConfig(config *BuilderConfig) *ArtifactBuilder {
	ab.config = config
	return ab
}

func (ab *ArtifactBuilder) WithBackend(backend *backend.Backend) *ArtifactBuilder {
	ab.backend = backend
	return ab
}

func (ab *ArtifactBuilder) WithRegistryClient(registryClient util.RegistryClient) *ArtifactBuilder {
	ab.registryClient = registryClient
	return ab
}

func (ab *ArtifactBuilder) WithTags(tags []string) *ArtifactBuilder {
	if len(tags) > 0 {
		ab.tags = tags
	}
	return ab
}

func (ab *ArtifactBuilder) validate() error {
	if ab.config == nil {
		return errors.New("configuration was not provided")
	}
	if ab.backend == nil {
		return errors.New("backend was not provided")
	}
	if ab.registryClient == nil {
		return errors.New("registry client was not provided")
	}
	return nil
}

func (ab *ArtifactBuilder) CheckImageExists(tag string) (bool, error) {
	err := ab.validate()
	if err != nil {
		return false, errors.Wrap(err, "artifact builder configuration is incomplete")
	}
	serviceRegistries := ab.backend.Conf().GetServiceRegistries()
	images, err := ab.config.GetBuildServicesImage()
	if err != nil {
		return false, errors.Wrap(err, "failed to get service image")
	}

	for serviceName := range images {
		registry, ok := serviceRegistries[serviceName]
		if !ok {
			continue
		}

		parts := strings.Split(registry.GetRepoUrl(), "/")
		if len(parts) < 2 {
			return false, errors.Errorf("invalid registry url format: %s", registry.GetRepoUrl())
		}
		registryId := parts[0]
		repoUrl := parts[1]

		parts = strings.Split(registryId, ".")
		if len(parts) < 6 {
			return false, errors.Errorf("invalid registry id format: %s", registryId)
		}
		registryId = parts[0]

		ecrClient := ab.backend.GetECRClient()

		input := &ecr.BatchGetImageInput{
			RegistryId: &registryId,
			ImageIds: []*ecr.ImageIdentifier{
				{
					ImageTag: aws.String(tag),
				},
			},
			RepositoryName: aws.String(repoUrl),
		}

		result, err := ecrClient.BatchGetImage(input)
		if err != nil {
			return false, errors.Wrap(err, "error getting an image")
		}
		if result == nil || len(result.Images) == 0 {
			return false, nil
		}
	}

	return true, nil
}

func (ab *ArtifactBuilder) RetagImages(
	serviceRegistries map[string]*config.RegistryConfig,
	sourceTag string,
	destTags []string,
	images []string,
) error {
	err := ab.validate()
	if err != nil {
		return errors.Wrap(err, "artifact builder configuration is incomplete")
	}

	ecrClient := ab.backend.GetECRClient()

	imageMap := make(map[string]bool)
	for _, image := range images {
		imageMap[image] = true
	}

	for serviceName, registry := range serviceRegistries {
		if _, ok := imageMap[serviceName]; !ok {
			if len(images) > 0 {
				continue
			}
		}

		repoUrl := strings.Split(registry.GetRepoUrl(), "/")[1]

		log.Infof("retagging %s from '%s' to '%s'", serviceName, sourceTag, strings.Join(destTags, ","))

		input := &ecr.BatchGetImageInput{
			ImageIds: []*ecr.ImageIdentifier{
				{
					ImageTag: aws.String(sourceTag),
				},
			},
			RepositoryName: aws.String(repoUrl),
		}

		result, err := ecrClient.BatchGetImage(input)
		if err != nil {
			log.Errorf("error getting Image: %s", err.Error())
			continue
		}

		if len(result.Images) == 0 {
			continue
		}

		manifest := result.Images[0].ImageManifest

		for _, tag := range destTags {
			input := &ecr.PutImageInput{
				ImageManifest:  manifest,
				ImageTag:       aws.String(tag),
				RepositoryName: aws.String(repoUrl),
			}

			_, err := ecrClient.PutImage(input)
			if err != nil {
				log.Error("error putting image", err)
				continue
			}
		}
	}

	return nil
}

func (ab *ArtifactBuilder) Build() error {
	return ab.config.DockerComposeBuild()
}

func (ab *ArtifactBuilder) RegistryLogin(ctx context.Context) error {
	err := ab.validate()
	if err != nil {
		return errors.Wrap(err, "artifact builder configuration is incomplete")
	}

	ecrAuthorizationToken, err := ab.backend.ECRGetAuthorizationToken(ctx)
	if err != nil {
		return err
	}

	err = ab.registryClient.Login(ecrAuthorizationToken.Username, ecrAuthorizationToken.Password, ecrAuthorizationToken.ProxyEndpoint)
	return errors.Wrap(err, "registry login failed")
}

func (ab *ArtifactBuilder) Push(tags []string) error {
	err := ab.validate()
	if err != nil {
		return errors.Wrap(err, "artifact builder configuration is incomplete")
	}

	serviceRegistries := ab.backend.Conf().GetServiceRegistries()
	servicesImage, err := ab.config.GetBuildServicesImage()
	if err != nil {
		return err
	}

	docker, err := exec.LookPath("docker")
	if err != nil {
		return errors.Wrap(err, "docker not in path")
	}
	for serviceName, registry := range serviceRegistries {
		if _, ok := servicesImage[serviceName]; !ok {
			continue
		}

		image := servicesImage[serviceName]
		for _, currentTag := range tags {
			// re-tag image
			dockerTagArgs := []string{"docker", "tag", fmt.Sprintf("%s:latest", image), fmt.Sprintf("%s:%s", registry.GetRepoUrl(), currentTag)}
			log.WithField("args", dockerTagArgs).Debug("Running shell cmd")
			cmd := &exec.Cmd{
				Path:   docker,
				Args:   dockerTagArgs,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
			if err := ab.config.executor.Run(cmd); err != nil {
				return errors.Errorf("process failure: %v", err)
			}

			// push image
			img := fmt.Sprintf("%s:%s", registry.GetRepoUrl(), currentTag)
			dockerPushArgs := []string{"docker", "push", img}
			log.WithField("args", dockerPushArgs).Debug("Running shell cmd")
			cmd = &exec.Cmd{
				Path:   docker,
				Args:   dockerPushArgs,
				Stdout: os.Stdout,
				Stderr: os.Stderr,
			}
			if err := ab.config.executor.Run(cmd); err != nil {
				return errors.Errorf("process failure: %v", err)
			}
			log.WithField("args", dockerTagArgs).Info("Tagged the image")
		}
	}
	return nil
}

func (ab *ArtifactBuilder) BuildAndPush(
	ctx context.Context,
	opts ...ArtifactBuilderBuildOption,
) error {
	err := ab.validate()
	if err != nil {
		return errors.Wrap(err, "artifact builder configuration is incomplete")
	}

	// calculate defaults
	defaultTag, err := ab.backend.GenerateTag(ctx)
	if err != nil {
		return err
	}
	tags := []string{defaultTag}
	if len(ab.tags) > 0 {
		tags = append(tags, ab.tags...)
	}

	// Get all the options first
	o := &artifactBuilderBuildOptions{
		tags: tags,
	}
	for _, opt := range opts {
		opt(o)
	}

	// Run logic
	err = ab.RegistryLogin(ctx)
	if err != nil {
		return err
	}

	err = ab.Build()
	if err != nil {
		return err
	}

	return ab.Push(o.tags)
}
