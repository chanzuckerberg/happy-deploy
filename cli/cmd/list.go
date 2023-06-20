package cmd

import (
	"context"
	"fmt"
	"io"

	"github.com/chanzuckerberg/happy/cli/pkg/hapi"
	"github.com/chanzuckerberg/happy/cli/pkg/output"
	stackservice "github.com/chanzuckerberg/happy/cli/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/model"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type StructuredListResult struct {
	Error  string
	Stacks []stackservice.StackInfo
}

var (
	listAll bool
	remote  bool
)

func init() {
	rootCmd.AddCommand(listCmd)
	config.ConfigureCmdWithBootstrapConfig(listCmd)
	listCmd.Flags().StringVar(&OutputFormat, "output", "text", "Output format. One of: json, yaml, or text. Defaults to text, which is the only interactive mode.")
	listCmd.Flags().BoolVar(&listAll, "all", false, "List all stacks, not just those belonging to this app")
	listCmd.Flags().BoolVar(&remote, "remote", false, "List stacks from the remote happy server")
}

var listCmd = &cobra.Command{
	Use:          "list",
	Short:        "List stacks",
	Long:         "Listing stacks in environment '{env}'",
	SilenceUsage: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		checklist := util.NewValidationCheckList()
		return util.ValidateEnvironment(cmd.Context(),
			checklist.TerraformInstalled,
		)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if OutputFormat != "text" {
			logrus.SetOutput(io.Discard)
		}
		happyClient, err := makeHappyClient(cmd, sliceName, "", []string{}, false)
		if err != nil {
			return errors.Wrap(err, "unable to initialize the happy client")
		}

		stackInfos := []stackservice.StackInfo{}
		if remote {
			err = listStacksRemote(cmd.Context(), happyClient)
			if err != nil {
				return err
			}
		} else {
			stackInfos, err = happyClient.StackService.CollectStackInfo(cmd.Context(), listAll, happyClient.HappyConfig.App())
			if err != nil {
				return errors.Wrap(err, "unable to collect stack info")
			}
		}

		printer := output.NewPrinter(OutputFormat)
		err = printer.PrintStacks(cmd.Context(), stackInfos)
		if err != nil {
			return errors.Wrap(err, "unable to print stacks")
		}

		return nil
	},
}

func listStacksRemote(ctx context.Context, happyClient *HappyClient) error {
	api := hapi.MakeApiClient(happyClient.HappyConfig)
	result, err := api.ListStacks(model.MakeAppStackPayload(
		happyClient.HappyConfig.App(),
		happyClient.HappyConfig.GetEnv(),
		"", model.AWSContext{
			AWSProfile:     *happyClient.HappyConfig.AwsProfile(),
			AWSRegion:      *happyClient.HappyConfig.AwsRegion(),
			TaskLaunchType: "k8s",
			K8SNamespace:   happyClient.HappyConfig.K8SConfig().Namespace,
			K8SClusterID:   happyClient.HappyConfig.K8SConfig().ClusterID,
		},
	))
	if err != nil {
		return err
	}

	fmt.Printf("%+v", result.Records[0].AppMetadata)
	return nil
}
