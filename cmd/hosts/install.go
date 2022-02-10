package hosts

import (
	"github.com/chanzuckerberg/happy/pkg/artifact_builder"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/chanzuckerberg/happy/pkg/hostname_manager"
	"github.com/spf13/cobra"
)

func init() {
	config.ConfigureCmdWithBootstrapConfig(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install compose DNS entries",
	Long:  "Install compose DNS entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		bootstrapConfig, err := config.NewBootstrapConfig()
		if err != nil {
			return err
		}
		happyConfig, err := config.NewHappyConfig(cmd.Context(), bootstrapConfig)
		if err != nil {
			return err
		}

		buildConfig := artifact_builder.NewBuilderConfig(bootstrapConfig, happyConfig)
		containers, err := buildConfig.GetContainers()
		if err != nil {
			return err
		}

		hostManager := hostname_manager.NewHostNameManager(hostsFile, containers)
		return hostManager.Install()

	},
}
