package cmd

import (
	"github.com/chanzuckerberg/happy/pkg/artifact_builder"
	"github.com/chanzuckerberg/happy/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(buildCmd)
	config.ConfigureCmdWithBootstrapConfig(buildCmd)
}

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build docker images",
	Long:  "Build docker images using docker-compose",
	RunE: func(cmd *cobra.Command, args []string) error {
		bootstrapConfig, err := config.NewBootstrapConfig()
		if err != nil {
			return err
		}
		happyConfig, err := config.NewHappyConfig(cmd.Context(), bootstrapConfig)
		if err != nil {
			return err
		}
		backend, err := aws.NewAWSBackend(cmd.Context(), happyConfig)
		if err != nil {
			return err
		}

		composeEnv := ""
		builderConfig := artifact_builder.NewBuilderConfig(bootstrapConfig, composeEnv, happyConfig.GetDockerRepo())
		artifactBuilder := artifact_builder.NewArtifactBuilder(builderConfig, happyConfig)
		serviceRegistries := happyConfig.GetRdevServiceRegistries()

		// NOTE  not to login before build for cache to work
		err = artifactBuilder.RegistryLogin(serviceRegistries)
		if err != nil {
			return err
		}

		return artifactBuilder.Build()
	},
}
