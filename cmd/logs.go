package cmd

import (
	backend "github.com/chanzuckerberg/happy/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/pkg/cmd"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/chanzuckerberg/happy/pkg/orchestrator"
	"github.com/spf13/cobra"
)

var since string

func init() {
	rootCmd.AddCommand(logsCmd)
	config.ConfigureCmdWithBootstrapConfig(logsCmd)

	logsCmd.Flags().StringVar(&since, "since", "10m", "Length of time to look back in logs")
}

var logsCmd = &cobra.Command{
	Use:          "logs STACK_NAME SERVICE",
	Short:        "Tail logs",
	Long:         "Tail the logs of a service (frontend, backend, upload, migrations)",
	SilenceUsage: true,
	RunE:         runLogs,
	PreRunE:      cmd.Validate(cobra.ExactArgs(2), cmd.CheckStackName),
}

func runLogs(cmd *cobra.Command, args []string) error {
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

	b, err := backend.NewAWSBackend(ctx, happyConfig)
	if err != nil {
		return err
	}

	taskOrchestrator := orchestrator.NewOrchestrator().WithBackend(b)

	return taskOrchestrator.Logs(ctx, stackName, service, since, true)
}
