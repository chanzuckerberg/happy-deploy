package cmd

import (
	"github.com/chanzuckerberg/happy/cli/pkg/cmd"
	"github.com/chanzuckerberg/happy/cli/pkg/orchestrator"
	backend "github.com/chanzuckerberg/happy/shared/backend/aws"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shellCmd)
	config.ConfigureCmdWithBootstrapConfig(shellCmd)
}

var shellCmd = &cobra.Command{
	Use:          "shell STACK_NAME SERVICE",
	Aliases:      []string{"exec", "sh", "bash"},
	Short:        "Execute into a container",
	Long:         "Execute into a running service task container",
	SilenceUsage: true,
	PreRunE:      cmd.Validate(cobra.ExactArgs(2), cmd.IsStackNameDNSCharset),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		checklist := util.NewValidationCheckList()
		return util.ValidateEnvironment(cmd.Context(),
			[]util.ValidationCallback{
				checklist.AwsInstalled,
				checklist.AwsSessionManagerPluginInstalled,
			})
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		stackName := args[0]
		service := args[1]

		bootstrapConfig, err := config.NewBootstrapConfig(cmd)
		if err != nil {
			return err
		}
		happyConfig, err := config.NewHappyConfig(bootstrapConfig)
		if err != nil {
			return err
		}

		b, err := backend.NewAWSBackend(ctx, happyConfig.GetEnvironmentContext())
		if err != nil {
			return err
		}

		return orchestrator.NewOrchestrator().WithHappyConfig(happyConfig).WithBackend(b).Shell(ctx, stackName, service)
	},
}
