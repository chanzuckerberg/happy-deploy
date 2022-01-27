package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/chanzuckerberg/happy/pkg/backend"
	"github.com/chanzuckerberg/happy/pkg/config"
	stack_service "github.com/chanzuckerberg/happy/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/pkg/util"
	"github.com/chanzuckerberg/happy/pkg/workspace_repo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&sliceName, "slice", "s", "", "If you only need to test a slice of the app, specify it here")
	updateCmd.Flags().StringVar(&sliceDefaultTag, "slice-default-tag", "", "For stacks using slices, override the default tag for any images that aren't being built & pushed by the slice")
}

var updateCmd = &cobra.Command{
	Use:   "update STACK_NAME",
	Short: "update stack",
	Long:  "Update stack mathcing STACK_NAME",
	RunE:  runUpdate,
}

func runUpdate(cmd *cobra.Command, args []string) error {

	env := "rdev"

	if len(args) != 1 {
		return errors.New("incorrect number of arguments")
	}

	stackName := args[0]

	happyConfigPath, ok := os.LookupEnv("HAPPY_CONFIG_PATH")
	if !ok {
		return errors.New("please set env var HAPPY_CONFIG_PATH")
	}

	_, ok = os.LookupEnv("HAPPY_PROJECT_ROOT")
	if !ok {
		return errors.New("please set env var HAPPY_PROJECT_ROOT")
	}

	happyConfig, err := config.NewHappyConfig(happyConfigPath, env)
	if err != nil {
		return err
	}

	// taskRunner := backend.GetAwsEcs(happyConfig)

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

	fmt.Printf("Updating %s\n", stackName)

	stacks, err := stackService.GetStacks()
	if err != nil {
		return err
	}
	stack, ok := stacks[stackName]
	if !ok {
		return fmt.Errorf("stack %s not found", stackName)
	}

	// TODO pass tag as arg
	var tag string = ""
	var stackTags map[string]string
	if len(sliceName) > 0 {
		stackTags, tag, err = buildSlice(happyConfig, sliceName, sliceDefaultTag)
		if err != nil {
			return err
		}
	} else {
		stackTags = make(map[string]string)
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
	if err != nil {
		return err
	}
	configSecret := map[string]string{"happy/meta/configsecret": secretArn}
	stackMeta.Load(configSecret)
	stackMeta.Update(tag, stackTags, sliceDefaultTag, stackService)

	wait := true
	skipMigrations := true
	shouldWait := (wait || !skipMigrations)
	err = stack.Apply(shouldWait)
	if err != nil {
		return errors.New("apply failed, skipping migrations")
	}

	// TODO implement logic for shouldAutoMigrate
	shouldAutoMigrate := false
	if !skipMigrations && shouldAutoMigrate {
		// TODO run migrate cmd
		fmt.Println("We are migrating!")
	}
	stack.PrintOutputs()

	return nil
}
