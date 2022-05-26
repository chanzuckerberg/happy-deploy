package cmd

import (
	"sort"

	backend "github.com/chanzuckerberg/happy/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/pkg/config"
	stackservice "github.com/chanzuckerberg/happy/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/pkg/util"
	"github.com/chanzuckerberg/happy/pkg/workspace_repo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/exp/maps"
)

func init() {
	rootCmd.AddCommand(listCmd)
	config.ConfigureCmdWithBootstrapConfig(listCmd)
}

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "list stacks",
	Long:         "Listing stacks in environment '{env}'",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
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

		url := b.Conf().GetTfeUrl()
		org := b.Conf().GetTfeOrg()

		workspaceRepo := workspace_repo.NewWorkspaceRepo(url, org)
		stackSvc := stackservice.NewStackService().WithBackend(b).WithWorkspaceRepo(workspaceRepo)

		stacks, err := stackSvc.GetStacks(ctx)
		if err != nil {
			return err
		}

		logrus.Infof("listing stacks in environment '%s'", happyConfig.GetEnv())

		headings := []string{"Name", "Owner", "Tags", "Status", "URLs", "LastUpdated"}
		tablePrinter := util.NewTablePrinter(headings)

		// Iterate in order
		stackNames := maps.Keys(stacks)
		sort.Strings(stackNames)
		for _, name := range stackNames {
			stack := stacks[name]
			err := stack.Print(ctx, name, tablePrinter)
			if err != nil {
				logrus.Warnf("Error retrieving stack %s:  %s", name, err)
				continue
			}
		}

		tablePrinter.Print()
		return nil
	},
}
