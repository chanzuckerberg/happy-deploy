package cmd

import (
	"errors"
	"fmt"
	"os"

	// "time"

	"github.com/chanzuckerberg/happy-deploy/pkg/artifact_builder"
	// "github.com/chanzuckerberg/happy-deploy/pkg/backend"
	"github.com/chanzuckerberg/happy-deploy/pkg/config"
	"github.com/chanzuckerberg/happy-deploy/pkg/util"

	"github.com/spf13/cobra"
)

var pushImages []string

// TODO add support for this flag
var useComposeEnv bool

func init() {
	pushCmd.Flags().StringSliceVar(&pushImages, "images", []string{}, "List of images to push to registry.")
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "push docker images",
	Long:  "Push docker images to ECR",
	RunE: func(cmd *cobra.Command, args []string) error {
		tag := ""
		return runPush(tag)
	},
}

type PushOptions struct {
	tag    string
	images []string
}

func runPush(tag string) error {
	return runPushWithOptions(PushOptions{tag: tag})
}

func runPushWithOptions(options PushOptions) error {

	// TODO do not hardcode dev
	env := "rdev"

	dockerComposeConfigPath, ok := os.LookupEnv("DOCKER_COMPOSE_CONFIG_PATH")
	if !ok {
		return errors.New("Please set env var DOCKER_COMPOSE_CONFIG_PATH")
	}

	happyConfigPath, ok := os.LookupEnv("HAPPY_CONFIG_PATH")
	if !ok {
		return errors.New("Please set env var HAPPY_CONFIG_PATH")
	}

	happyConfig, err := config.NewHappyConfig(happyConfigPath, env)
	if err != nil {
		return fmt.Errorf("Failed to get Happy Config: %s", err)
	}

	composeEnv := ""
	if useComposeEnv {
		composeEnv = happyConfig.DefaultComposeEnv()
	}
	buildConfig := artifact_builder.NewBuilderConfig(dockerComposeConfigPath, composeEnv)
	artifactBuilder := artifact_builder.NewArtifactBuilder(buildConfig, happyConfig)
	serviceRegistries, err := happyConfig.GetRdevServiceRegistries()
	if err != nil {
		return err
	}
	// NOTE login before build in order for cache to work
	artifactBuilder.RegistryLogin(serviceRegistries, pushImages)

	servicesImage, err := buildConfig.GetBuildServicesImage()
	if err != nil {
		return fmt.Errorf("Failed to get service image: %s", err)
	}

	for service, reg := range serviceRegistries {
		fmt.Printf("%q: %q\t%q\n", service, reg.GetRepoUrl(), reg.GetRegistryUrl())
	}

	if options.tag == "" {
		options.tag, err = util.GenerateTag(happyConfig)
		if err != nil {
			return err
		}
	}
	tags := []string{options.tag}
	fmt.Println(tags)

	err = artifactBuilder.Build()
	if err != nil {
		return fmt.Errorf("Failed to push image: %s", err)
	}
	fmt.Println("Build complete")

	// TODO add extra tag from input

	artifactBuilder.Push(serviceRegistries, servicesImage, tags)
	return nil
}
