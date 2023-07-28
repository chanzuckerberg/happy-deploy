package cmd

import (
	"context"
	"fmt"
	"time"

	happyCmd "github.com/chanzuckerberg/happy/cli/pkg/cmd"
	"github.com/chanzuckerberg/happy/shared/backend/aws"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/options"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func init() {
	rootCmd.AddCommand(restartCmd)
	config.ConfigureCmdWithBootstrapConfig(restartCmd)
}

var restartCmd = &cobra.Command{
	Use:          "restart",
	Short:        "Restart a happy stack deployment, leaving everything else the same",
	SilenceUsage: true,
	PreRunE: happyCmd.Validate(
		cobra.ExactArgs(1),
		happyCmd.IsStackNameDNSCharset,
		happyCmd.IsStackNameAlphaNumeric,
		func(cmd *cobra.Command, args []string) error {
			checklist := util.NewValidationCheckList()
			return util.ValidateEnvironment(cmd.Context(), checklist.AwsInstalled)
		},
	),
	RunE: func(cmd *cobra.Command, args []string) error {
		stackName := args[0]
		happyClient, err := makeHappyClient(cmd, sliceName, stackName, []string{tag}, createTag)
		if err != nil {
			return errors.Wrap(err, "initializing the happy client")
		}

		ctx := context.WithValue(cmd.Context(), options.DryRunKey, dryRun)
		err = validate(
			validateConfigurationIntegirty(ctx, sliceName, happyClient),
			validateGitTree(happyClient.HappyConfig.GetProjectRoot()),
			validateStackNameAvailable(ctx, happyClient.StackService, stackName, force),
		)
		if err != nil {
			return errors.Wrap(err, "validating happy client")
		}

		backend, err := happyClient.AWSBackend.GetComputeBackend(cmd.Context())
		if err != nil {
			return err
		}

		k8s, ok := backend.(*aws.K8SComputeBackend)
		if !ok {
			return errors.New("not a k8s backend, nothing to do")
		}

		for _, service := range happyClient.HappyConfig.GetData().Services {
			deploymentName := k8s.GetDeploymentName(stackName, service)
			deploymentsClient := k8s.ClientSet.AppsV1().Deployments(k8s.KubeConfig.Namespace)
			logrus.Debugf("restarting deployment %s:%s", k8s.KubeConfig.Namespace, deploymentName)
			_, err = deploymentsClient.Patch(
				ctx,
				deploymentName,
				types.StrategicMergePatchType,
				[]byte(fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))),
				v1.PatchOptions{},
			)
			if err != nil {
				return errors.Wrapf(err, "patching deployment %s", deploymentName)
			}
		}
		return nil
	},
}
