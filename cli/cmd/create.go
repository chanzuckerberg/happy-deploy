package cmd

import (
	"context"
	"fmt"
	"strings"

	happyCmd "github.com/chanzuckerberg/happy/cli/pkg/cmd"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/options"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/chanzuckerberg/happy/shared/workspace_repo"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	force        bool
	skipCheckTag bool
	createTag    bool
	tag          string
	dryRun       bool
)

func init() {
	rootCmd.AddCommand(createCmd)
	config.ConfigureCmdWithBootstrapConfig(createCmd)
	happyCmd.SupportUpdateSlices(createCmd, &sliceName, &sliceDefaultTag) // Should this function be renamed to something more generalized?
	happyCmd.SetMigrationFlags(createCmd)

	createCmd.Flags().StringVar(&tag, "tag", "", "Specify the tag for the docker images. If not specified we will generate a default tag.")
	createCmd.Flags().BoolVar(&createTag, "create-tag", true, "Will build, tag, and push images when set. Otherwise, assumes images already exist.")
	createCmd.Flags().BoolVar(&skipCheckTag, "skip-check-tag", false, "Skip checking that the specified tag exists (requires --tag)")
	createCmd.Flags().BoolVar(&force, "force", false, "Ignore the already-exists errors")
	createCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Plan all infrastructure changes, but do not apply them")
}

var createCmd = &cobra.Command{
	Use:          "create STACK_NAME",
	Short:        "Create new stack",
	Long:         "Create a new stack with a given tag.",
	SilenceUsage: true,
	PreRunE: happyCmd.Validate(
		happyCmd.IsTagUsedWithSkipTag,
		cobra.ExactArgs(1),
		happyCmd.IsStackNameDNSCharset,
		happyCmd.IsStackNameAlphaNumeric),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		checklist := util.NewValidationCheckList()
		return util.ValidateEnvironment(cmd.Context(),
			[]util.ValidationCallback{
				checklist.DockerEngineRunning,
				checklist.MinDockerComposeVersion,
				checklist.DockerInstalled,
				checklist.TerraformInstalled,
				checklist.AwsInstalled,
				checklist.AwsSessionManagerPluginInstalled,
			})
	},
	RunE: runCreate,
}

// keep in sync with happy-stack-eks terraform module
const terraformECRTargetPathTemplate = `module.stack.module.services["%s"].module.ecr`

func runCreate(
	cmd *cobra.Command,
	args []string,
) (err error) {
	stackName := args[0]
	happyClient, err := makeHappyClient(cmd, sliceName, stackName, []string{tag}, createTag)
	if err != nil {
		return errors.Wrap(err, "unable to initialize the happy client")
	}
	ctx := context.WithValue(cmd.Context(), options.DryRunKey, dryRun)
	message := workspace_repo.Message(fmt.Sprintf("Happy %s Create Stack [%s]", util.GetVersion().Version, stackName))
	err = validate(
		validateConfigurationIntegirty(ctx, happyClient),
		validateGitTree(happyClient.HappyConfig.GetProjectRoot()),
		validateTFEBackLog(ctx, happyClient.AWSBackend),
		validateStackNameAvailable(ctx, happyClient.StackService, stackName, force),
		validateStackExistsCreate(ctx, stackName, happyClient, message),
		validateECRExists(ctx, stackName, terraformECRTargetPathTemplate, happyClient, message),
		validateImageExists(ctx, createTag, skipCheckTag, happyClient.ArtifactBuilder),
	)
	if err != nil {
		return errors.Wrap(err, "failed one of the happy client validations")
	}

	// update the newly created stack
	stack, err := happyClient.StackService.GetStack(ctx, stackName)
	if err != nil {
		return errors.Wrapf(err, "stack %s doesn't exist; this should never happen", stackName)
	}

	err = updateStack(ctx, cmd, stack, force, happyClient)
	if err != nil {
		return errors.Wrapf(err, "unable to update the stack %s", stack.Name)
	}
	// if it was a dry run, we should remove the stack after we are done
	if dryRun {
		log.Debugf("cleaning up stack '%s'", stack.Name)
		return errors.Wrap(happyClient.StackService.Remove(ctx, stack.Name), "unable to remove stack")
	}
	return nil
}

func validateECRExists(ctx context.Context, stackName string, ecrTargetPathFormat string, happyClient *HappyClient, options ...workspace_repo.TFERunOption) validation {
	logrus.Debug("Scheduling validateECRExists()")
	return func() error {
		logrus.Debug("Running validateECRExists()")
		if !happyClient.HappyConfig.GetFeatures().EnableECRAutoCreation {
			return nil
		}

		stackECRS, err := happyClient.ArtifactBuilder.GetECRsForServices(ctx)
		if err != nil {
			return errors.Wrap(err, "unable to get ECRs for services; this shouldn't happen if the stack TF is configured correctly")
		}

		missingServiceECRs := []string{}
		for _, service := range happyClient.HappyConfig.GetServices() {
			if _, ok := stackECRS[service]; !ok {
				missingServiceECRs = append(missingServiceECRs, service)
			}
		}
		if len(missingServiceECRs) == 0 {
			return nil
		}

		log.Debugf("missing ECRs for the following services %s. making them now", strings.Join(missingServiceECRs, ","))
		targetAddrs := []string{}
		for _, service := range happyClient.HappyConfig.GetServices() {
			targetAddrs = append(targetAddrs, fmt.Sprintf(ecrTargetPathFormat, service))
		}
		stack, err := happyClient.StackService.GetStack(ctx, stackName)
		if err != nil {
			return errors.Wrapf(err, "stack %s doesn't exist; this should never happen", stackName)
		}
		stackMeta, err := updateStackMeta(ctx, stack.Name, happyClient)
		if err != nil {
			return errors.Wrap(err, "unable to update the stack's meta information")
		}

		// this has a strong coupling with the TF module version that we are using in happy-stack-eks,
		// so if the user isn't on it yet, this will fail or not do what you are expecting
		// TODO: maybe CDK
		// TODO: maybe we can peek at the version and fail if its not right or something?
		stack = stack.WithMeta(stackMeta)
		return stack.Apply(ctx, makeWaitOptions(stackName, happyClient.HappyConfig, happyClient.AWSBackend), append(options, workspace_repo.TargetAddrs(targetAddrs))...)
	}
}

func validateStackExistsCreate(ctx context.Context, stackName string, happyClient *HappyClient, options ...workspace_repo.TFERunOption) validation {
	logrus.Debug("Scheduling validateStackExistsCreate()")
	return func() error {
		logrus.Debug("Running validateStackExistsCreate()")
		// 1.) if the stack does not exist and force flag is used, call the create function first
		_, err := happyClient.StackService.GetStack(ctx, stackName)
		if err != nil {
			logrus.Debugf("Stack doesn't exist %s: %s\n", stackName, err.Error())
			_, err = happyClient.StackService.Add(ctx, stackName, options...)
			if err != nil {
				return errors.Wrap(err, "unable to create the stack")
			}
			logrus.Debugf("Stack added: %s", stackName)
		} else {
			logrus.Debugf("Stack exists: %s", stackName)
			if !force {
				return errors.Wrapf(err, "stack %s already exists", stackName)
			}
		}

		return nil
	}
}
