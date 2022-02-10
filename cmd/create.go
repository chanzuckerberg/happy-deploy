package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/chanzuckerberg/happy/pkg/artifact_builder"
	backend "github.com/chanzuckerberg/happy/pkg/backend/aws"
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/chanzuckerberg/happy/pkg/options"
	"github.com/chanzuckerberg/happy/pkg/orchestrator"
	stack_service "github.com/chanzuckerberg/happy/pkg/stack_mgr"
	"github.com/chanzuckerberg/happy/pkg/workspace_repo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	createTag       string
	force           bool
	sliceName       string
	sliceDefaultTag string
	skipCheckTag    bool
)

func init() {
	rootCmd.AddCommand(createCmd)
	config.ConfigureCmdWithBootstrapConfig(createCmd)

	createCmd.Flags().StringVar(&createTag, "tag", "", "Tag name for docker image. Leave empty to generate one")
	createCmd.Flags().BoolVar(&force, "force", false, "Ignore the already-exists errors")
	createCmd.Flags().StringVarP(&sliceName, "slice", "s", "", "If you only need to test a slice of the app, specify it here")
	createCmd.Flags().StringVar(&sliceDefaultTag, "slice-default-tag", "", "For stacks using slices, override the default tag for any images that aren't being built & pushed by the slice")
	createCmd.Flags().BoolVar(&skipCheckTag, "skip-check-tag", false, "Skip checking that the specified tag exists (requires --tag)")
}

var createCmd = &cobra.Command{
	Use:          "create STACK_NAME",
	Short:        "create new stack",
	Long:         "Create a new stack with a given tag.",
	SilenceUsage: true,
	PreRunE:      checkFlags,
	RunE:         runCreate,
	Args:         cobra.ExactArgs(1),
}

func checkFlags(cmd *cobra.Command, args []string) error {
	if cmd.Flags().Changed("skip-check-tag") && !cmd.Flags().Changed("tag") {
		return errors.New("--skip-check-tag can only be used when --tag is specified")
	}
	return nil
}

func runCreate(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	stackName := args[0]

	bootstrapConfig, err := config.NewBootstrapConfig()
	if err != nil {
		return err
	}
	happyConfig, err := config.NewHappyConfig(ctx, bootstrapConfig)
	if err != nil {
		return err
	}

	backend, err := backend.NewAWSBackend(ctx, happyConfig)
	if err != nil {
		return err
	}

	builderConfig := artifact_builder.NewBuilderConfig(bootstrapConfig, happyConfig)
	ab := artifact_builder.NewArtifactBuilder(builderConfig, backend)

	url := backend.Conf().GetTfeUrl()
	org := backend.Conf().GetTfeOrg()

	workspaceRepo, err := workspace_repo.NewWorkspaceRepo(url, org)
	if err != nil {
		return err
	}
	stackService := stack_service.NewStackService(backend, workspaceRepo)

	existingStacks, err := stackService.GetStacks(ctx)
	if err != nil {
		return err
	}
	if _, ok := existingStacks[stackName]; ok {
		if !force {
			return errors.Errorf("stack %s already exists", stackName)
		}
	}

	stackTags := map[string]string{}
	if len(sliceName) > 0 {
		stackTags, createTag, err = buildSlice(ctx, backend, sliceName, sliceDefaultTag)
		if err != nil {
			return err
		}
	}

	exists, err := checkImageExists(ctx, backend, ab, createTag)
	if err != nil {
		return err
	}
	if !exists {
		return errors.Errorf("image tag does not exist or cannot be verified: %s", createTag)
	}

	stackMeta := stackService.NewStackMeta(stackName)
	secretArn := happyConfig.GetSecretArn()
	if err != nil {
		return err
	}
	metaTag := map[string]string{"happy/meta/configsecret": secretArn}
	err = stackMeta.Load(metaTag)
	if err != nil {
		return err
	}

	if createTag == "" {
		createTag, err = backend.GenerateTag(ctx)
		if err != nil {
			return err
		}

		// invoke push cmd
		fmt.Printf("Pushing images with tags %s...\n", createTag)
		err := runPush(ctx, createTag)
		if err != nil {
			return errors.Errorf("failed to push image: %s", err)
		}
	}
	err = stackMeta.Update(ctx, createTag, stackTags, sliceName, stackService)
	if err != nil {
		return err
	}
	fmt.Printf("Creating %s\n", stackName)

	stack, err := stackService.Add(ctx, stackName)
	if err != nil {
		return err
	}
	fmt.Printf("setting stackMeta %v\n", stackMeta)
	stack.SetMeta(stackMeta)

	err = stack.Apply(getWaitOptions(backend, stackName))
	if err != nil {
		return err
	}

	autoRunMigration := happyConfig.AutoRunMigration()
	if autoRunMigration {
		err = runMigrate(ctx, stackName)
		if err != nil {
			return err
		}
	}

	stack.PrintOutputs()
	return nil
}

func checkImageExists(
	ctx context.Context,
	backend *backend.Backend,
	ab *artifact_builder.ArtifactBuilder,
	tag string,
) (bool, error) {
	if len(tag) == 0 || skipCheckTag {
		// TODO: maybe a bit misleading to say true here
		return true, nil
	}
	serviceRegistries := backend.Conf().GetServiceRegistries()
	return ab.CheckImageExists(serviceRegistries, tag)
}

func buildSlice(
	ctx context.Context,
	backend *backend.Backend,
	sliceName string,
	defaultSliceTag string) (stackTags map[string]string, defaultTag string, err error) {
	defaultTag = defaultSliceTag

	slices, err := backend.Conf().GetSlices()
	if err != nil {
		return stackTags, defaultTag, errors.Errorf("unable to retrieve slice configuration: %s", err.Error())
	}

	slice, ok := slices[sliceName]
	if !ok {
		validSlices := joinKeys(slices, ", ")
		return stackTags, defaultTag, errors.Errorf("slice %s is invalid - valid names: %s", sliceName, validSlices)
	}

	buildImages := slice.BuildImages
	sliceTag, err := backend.GenerateTag(ctx)
	if err != nil {
		return stackTags, defaultTag, err
	}

	err = runPushWithOptions(ctx, sliceTag, buildImages, "")
	if err != nil {
		return stackTags, defaultTag, errors.Wrap(err, "failed to push image")
	}

	if len(defaultTag) == 0 {
		defaultTag = backend.Conf().SliceDefaultTag()
	}

	stackTags = make(map[string]string)
	for _, sliceImg := range buildImages {
		stackTags[sliceImg] = sliceTag
	}

	return stackTags, defaultTag, nil
}

func joinKeys(m map[string]config.Slice, separator string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, separator)
}

func getWaitOptions(backend *backend.Backend, stackName string) options.WaitOptions {
	taskOrchestrator := orchestrator.NewOrchestrator(backend)
	waitOptions := options.WaitOptions{StackName: stackName, Orchestrator: taskOrchestrator, Services: backend.Conf().GetServices()}
	return waitOptions
}
