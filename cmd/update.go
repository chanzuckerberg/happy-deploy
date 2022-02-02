package cmd

import (
	"fmt"

	"github.com/chanzuckerberg/happy/pkg/backend"
	"github.com/chanzuckerberg/happy/pkg/config"
	stack_service "github.com/chanzuckerberg/happy/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/pkg/util"
	"github.com/chanzuckerberg/happy/pkg/workspace_repo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	config.ConfigureCmdWithBootstrapConfig(updateCmd)

	updateCmd.Flags().StringVar(&tag, "tag", "", "Tag name for docker image. Leave empty to generate one automatically.")
	updateCmd.Flags().StringVarP(&sliceName, "slice", "s", "", "If you only need to test a slice of the app, specify it here")
	updateCmd.Flags().StringVar(&sliceDefaultTag, "slice-default-tag", "", "For stacks using slices, override the default tag for any images that aren't being built & pushed by the slice")
	updateCmd.Flags().BoolVar(&skipCheckTag, "skip-check-tag", false, "Skip checking that the specified tag exists (requires --tag)")
}

var updateCmd = &cobra.Command{
	Use:     "update STACK_NAME",
	Short:   "update stack",
	Long:    "Update stack mathcing STACK_NAME",
	PreRunE: checkFlags,
	RunE:    runUpdate,
	Args:    cobra.ExactArgs(1),
}

func runUpdate(cmd *cobra.Command, args []string) error {
	stackName := args[0]

	bootstrapConfig, err := config.NewBootstrapConfig()
	if err != nil {
		return err
	}

	happyConfig, err := config.NewHappyConfig(bootstrapConfig)
	if err != nil {
		return err
	}

	url, err := happyConfig.TfeUrl()
	if err != nil {
		return err
	}

	org, err := happyConfig.TfeOrg()
	if err != nil {
		return err
	}

	workspaceRepo, err := workspace_repo.NewWorkspaceRepo(url, org)
	if err != nil {
		return err
	}

	paramStoreBackend := backend.GetAwsBackend(happyConfig)
	stackService := stack_service.NewStackService(happyConfig, paramStoreBackend, workspaceRepo)

	exists, err := checkImageExists(bootstrapConfig, happyConfig, tag)
	if err != nil {
		return err
	}
	if !exists {
		return errors.Errorf("image tag does not exist or cannot be verified: %s", tag)
	}

	fmt.Printf("Updating %s\n", stackName)

	stacks, err := stackService.GetStacks()
	if err != nil {
		return err
	}
	stack, ok := stacks[stackName]
	if !ok {
		return errors.Errorf("stack %s not found", stackName)
	}

	stackTags := map[string]string{}
	if len(sliceName) > 0 {
		stackTags, tag, err = buildSlice(happyConfig, sliceName, sliceDefaultTag)
		if err != nil {
			return err
		}
	}

	if tag == "" {
		tag, err = util.GenerateTag(happyConfig)
		if err != nil {
			return err
		}

		// invoke push cmd
		fmt.Printf("Pushing images with tags %s...\n", tag)
		err := runPush(tag)
		if err != nil {
			return err
		}
	}

	stackMeta, err := stack.Meta()
	if err != nil {
		return err
	}

	// reset the configsecret if it has changed
	secretArn := happyConfig.GetSecretArn()

	configSecret := map[string]string{"happy/meta/configsecret": secretArn}
	err = stackMeta.Load(configSecret)
	if err != nil {
		return err
	}

	err = stackMeta.Update(tag, stackTags, sliceDefaultTag, stackService)
	if err != nil {
		return err
	}

	err = stack.Apply(getWaitOptions(happyConfig, stackName))
	if err != nil {
		return errors.Wrap(err, "apply failed, skipping migrations")
	}

	stack.PrintOutputs()
	return nil
}
