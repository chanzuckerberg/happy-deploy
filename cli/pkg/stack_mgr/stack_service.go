package stack_mgr

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	backend "github.com/chanzuckerberg/happy/shared/backend/aws"
	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/chanzuckerberg/happy/shared/diagnostics"
	"github.com/chanzuckerberg/happy/shared/options"
	"github.com/chanzuckerberg/happy/shared/util"
	"github.com/chanzuckerberg/happy/shared/util/tf"
	workspacerepo "github.com/chanzuckerberg/happy/shared/workspace_repo"
	"github.com/hashicorp/go-getter"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/go-tfe"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/maps"
	"golang.org/x/sync/errgroup"
)

type StackServiceIface interface {
	NewStackMeta(stackName string) *StackMeta
	Add(ctx context.Context, stackName string, options ...workspacerepo.TFERunOption) (*Stack, error)
	Remove(ctx context.Context, stackName string, options ...workspacerepo.TFERunOption) error
	GetStacks(ctx context.Context) (map[string]*Stack, error)
	GetStackWorkspace(ctx context.Context, stackName string) (workspacerepo.Workspace, error)
	GetConfig() *config.HappyConfig
}

type StackService struct {
	// dependencies
	backend       *backend.Backend
	workspaceRepo workspacerepo.WorkspaceRepoIface
	dirProcessor  util.DirProcessor
	executor      util.Executor
	happyConfig   *config.HappyConfig

	// NOTE: creator Workspace is a workspace that creates dependent workspaces with
	// given default values and configuration
	// the derived workspace is then used to launch the actual happy infrastructure
	creatorWorkspaceName string
}

func NewStackService() *StackService {
	return &StackService{
		dirProcessor: util.NewLocalProcessor(),
		executor:     util.NewDefaultExecutor(),
	}
}

func (s *StackService) GetWritePath() string {
	return fmt.Sprintf("/happy/%s/stacklist", s.happyConfig.GetEnv())
}

func (s *StackService) GetNamespacedWritePath() string {
	return fmt.Sprintf("/happy/%s/%s/stacklist", s.happyConfig.App(), s.happyConfig.GetEnv())
}

func (s *StackService) WithBackend(backend *backend.Backend) *StackService {
	creatorWorkspaceName := fmt.Sprintf("env-%s", s.happyConfig.GetEnv())

	s.creatorWorkspaceName = creatorWorkspaceName
	s.backend = backend

	return s
}

func (s *StackService) WithHappyConfig(happyConfig *config.HappyConfig) *StackService {
	s.happyConfig = happyConfig
	return s
}

func (s *StackService) WithExecutor(executor util.Executor) *StackService {
	s.executor = executor
	return s
}

func (s *StackService) WithWorkspaceRepo(workspaceRepo workspacerepo.WorkspaceRepoIface) *StackService {
	s.workspaceRepo = workspaceRepo
	return s
}

func (s *StackService) GetConfig() *config.HappyConfig {
	return s.happyConfig
}

// Invoke a specific TFE workspace that creates/deletes TFE workspaces,
// with prepopulated variables for identifier tokens.
func (s *StackService) resync(ctx context.Context, wait bool, options ...workspacerepo.TFERunOption) error {
	log.Debug("resyncing new workspace...")
	log.Debugf("running creator workspace %s...", s.creatorWorkspaceName)
	creatorWorkspace, err := s.workspaceRepo.GetWorkspace(ctx, s.creatorWorkspaceName)
	if err != nil {
		return errors.Wrapf(err, "unable to get workspace %s", s.creatorWorkspaceName)
	}
	err = creatorWorkspace.Run(ctx, options...)
	if err != nil {
		return errors.Wrapf(err, "error running latest %s workspace version", s.creatorWorkspaceName)
	}
	if wait {
		return creatorWorkspace.Wait(ctx)
	}
	return nil
}

func (s *StackService) GetLatestDeployedTag(ctx context.Context, stackName string) (string, error) {
	stack, err := s.GetStack(ctx, stackName)
	if err != nil {
		return "", errors.Wrap(err, "unable to get the stack")
	}
	stackInfo, err := stack.GetStackInfo(ctx)
	if err != nil {
		return "", errors.Wrap(err, "unable to get the stack info")
	}
	return stackInfo.Tag, nil
}

func (s *StackService) Remove(ctx context.Context, stackName string, opts ...workspacerepo.TFERunOption) error {
	dryRun, ok := ctx.Value(options.DryRunKey).(bool)
	if !ok {
		dryRun = false
	}
	if dryRun {
		return nil
	}
	var err error
	if s.GetConfig().GetFeatures().EnableDynamoLocking {
		err = s.removeFromStacklistWithLock(ctx, stackName)
	} else {
		err = s.removeFromStacklist(ctx, stackName)
	}
	if err != nil {
		return err
	}

	return s.resync(ctx, false, opts...)
}

func (s *StackService) removeFromStacklistWithLock(ctx context.Context, stackName string) error {
	distributedLock, err := s.getDistributedLock()
	if err != nil {
		return err
	}
	defer distributedLock.Close(ctx)

	lockKey := s.GetNamespacedWritePath()
	lock, err := distributedLock.AcquireLock(ctx, lockKey)
	if err != nil {
		return err
	}

	// don't return if there was an error here, we still need to release the lock so we'll use multierror instead
	ret := s.removeFromStacklist(ctx, stackName)

	_, err = distributedLock.ReleaseLock(ctx, lock)
	if err != nil {
		ret = multierror.Append(ret, errors.Wrapf(err, "unable to release the lock on %s", lockKey))
	}

	return ret
}

func (s *StackService) removeFromStacklist(ctx context.Context, stackName string) error {
	log.WithField("stack_name", stackName).Debug("Removing stack...")

	stacks, err := s.GetStacks(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get a list of stacks")
	}
	stackNamesList := []string{}
	for name := range stacks {
		if name != stackName {
			stackNamesList = append(stackNamesList, name)
		}
	}

	return s.writeStacklist(ctx, stackNamesList)
}

func (s *StackService) Add(ctx context.Context, stackName string, opts ...workspacerepo.TFERunOption) (*Stack, error) {
	log.WithField("stack_name", stackName).Debug("Adding a new stack...")
	dryRun, ok := ctx.Value(options.DryRunKey).(bool)
	if !ok {
		dryRun = false
	}
	if dryRun {
		log.Debugf("temporarily creating a TFE workspace for stack '%s'", stackName)
	} else {
		log.Debugf("creating stack '%s'", stackName)
	}

	var err error
	if s.GetConfig().GetFeatures().EnableDynamoLocking {
		err = s.addToStacklistWithLock(ctx, stackName)
	} else {
		err = s.addToStacklist(ctx, stackName)
	}
	if err != nil {
		return nil, err
	}

	if !util.IsLocalstackMode() {
		// Create the workspace
		wait := true
		if err := s.resync(ctx, wait, opts...); err != nil {
			return nil, err
		}
	}

	_, err = s.GetStackWorkspace(ctx, stackName)
	if err != nil {
		return nil, err
	}
	return s.createStack(stackName), nil
}

func (s *StackService) addToStacklistWithLock(ctx context.Context, stackName string) error {
	log.WithField("stack_name", stackName).Debug("Adding new stack with a lock...")
	distributedLock, err := s.getDistributedLock()
	if err != nil {
		return err
	}
	defer distributedLock.Close(ctx)

	lockKey := s.GetNamespacedWritePath()
	lock, err := distributedLock.AcquireLock(ctx, lockKey)
	if err != nil {
		return err
	}

	// don't return if there was an error here, we still need to release the lock so we'll use multierror instead
	ret := s.addToStacklist(ctx, stackName)

	_, err = distributedLock.ReleaseLock(ctx, lock)
	if err != nil {
		ret = multierror.Append(ret, errors.Wrapf(err, "unable to release the lock on %s", lockKey))
	}

	return ret
}

func (s *StackService) addToStacklist(ctx context.Context, stackName string) error {
	log.WithField("stack_name", stackName).Debug("Adding new stack...")
	existStacks, err := s.GetStacks(ctx)
	if err != nil {
		return err
	}

	newStackNames := []string{}
	stackNameExists := false
	for name := range existStacks {
		newStackNames = append(newStackNames, name)
		if name == stackName {
			stackNameExists = true
		}
	}
	if !stackNameExists {
		newStackNames = append(newStackNames, stackName)
	}

	return s.writeStacklist(ctx, newStackNames)
}

func (s *StackService) writeStacklist(ctx context.Context, stackNames []string) error {
	sort.Strings(stackNames)

	stackNamesJson, err := json.Marshal(stackNames)
	if err != nil {
		return errors.Wrap(err, "unable to serialize stack list as json")
	}

	stackNamesStr := string(stackNamesJson)
	log.WithFields(log.Fields{"path": s.GetNamespacedWritePath(), "data": stackNamesStr}).Debug("Writing to paramstore...")
	if err := s.backend.ComputeBackend.WriteParam(ctx, s.GetNamespacedWritePath(), stackNamesStr); err != nil {
		return errors.Wrap(err, "unable to write a workspace param")
	}
	log.WithFields(log.Fields{"path": s.GetWritePath(), "data": stackNamesStr}).Debug("Writing to paramstore...")
	if err := s.backend.ComputeBackend.WriteParam(ctx, s.GetWritePath(), stackNamesStr); err != nil {
		return errors.Wrap(err, "unable to write a workspace param")
	}

	return nil
}

func (s *StackService) GetStacks(ctx context.Context) (map[string]*Stack, error) {
	defer diagnostics.AddProfilerRuntime(ctx, time.Now(), "GetStacks")
	log.WithField("path", s.GetNamespacedWritePath()).Debug("Reading stacks from paramstore at path...")
	paramOutput, err := s.backend.ComputeBackend.GetParam(ctx, s.GetNamespacedWritePath())
	if err != nil && strings.Contains(err.Error(), "ParameterNotFound") {
		log.WithField("path", s.GetWritePath()).Debug("Reading stacks from paramstore at path...")
		paramOutput, err = s.backend.ComputeBackend.GetParam(ctx, s.GetWritePath())
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get stacks")
	}

	log.WithField("output", paramOutput).Debug("read stacks info from param store")

	var stacklist []string
	err = json.Unmarshal([]byte(paramOutput), &stacklist)
	if err != nil {
		return nil, errors.Wrap(err, "could not parse json")
	}

	log.WithField("output", stacklist).Debug("marshalled json output to string slice")

	stacks := map[string]*Stack{}
	for _, stackName := range stacklist {
		stacks[stackName] = s.createStack(stackName)
	}

	return stacks, nil
}

func (s *StackService) CollectStackInfo(ctx context.Context, listAll bool, app string) ([]StackInfo, error) {
	g, ctx := errgroup.WithContext(ctx)
	stacks, err := s.GetStacks(ctx)
	if err != nil {
		return nil, err
	}
	// Iterate in order
	stackNames := maps.Keys(stacks)
	stackInfos := make([]*StackInfo, len(stackNames))
	sort.Strings(stackNames)
	for i, name := range stackNames {
		i, name := i, name // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			stackInfo, err := stacks[name].GetStackInfo(ctx)
			if err != nil {
				log.Warnf("unable to get stack info for %s: %s (likely means the deploy failed the first time)", name, err)
				if !diagnostics.IsInteractiveContext(ctx) {
					stackInfos[i] = &StackInfo{
						Name:    name,
						Status:  "error",
						Message: err.Error(),
					}
				}
				// we still want to show the other stacks if this errors
				return nil
			}

			// only show the stacks that belong to this app or they want to list all
			if listAll || (stackInfo != nil && stackInfo.App == app) {
				stackInfos[i] = stackInfo
			}

			return nil
		})
	}
	err = g.Wait()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get stack infos")
	}

	// remove empties
	nonEmptyStackInfos := []StackInfo{}
	for _, stackInfo := range stackInfos {
		if stackInfo == nil {
			continue
		}
		nonEmptyStackInfos = append(nonEmptyStackInfos, *stackInfo)
	}
	return nonEmptyStackInfos, g.Wait()
}

func (s *StackService) GetStack(ctx context.Context, stackName string) (*Stack, error) {
	existingStacks, err := s.GetStacks(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get stacks")
	}
	stack, ok := existingStacks[stackName]
	if !ok {
		return nil, errors.Errorf("stack %s doesn't exist", stackName)
	}

	return stack, nil
}

// pre-format stack name and call workspaceRepo's GetWorkspace method
func (s *StackService) GetStackWorkspace(ctx context.Context, stackName string) (workspacerepo.Workspace, error) {
	workspaceName := fmt.Sprintf("%s-%s", s.happyConfig.GetEnv(), stackName)

	ws, err := s.workspaceRepo.GetWorkspace(ctx, workspaceName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get workspace")
	}

	return ws, nil
}

func (s *StackService) createStack(stackName string) *Stack {
	return &Stack{
		stackService: s,
		Name:         stackName,
		dirProcessor: s.dirProcessor,
		executor:     s.executor,
	}
}

func (s *StackService) HasState(ctx context.Context, stackName string) (bool, error) {
	workspace, err := s.GetStackWorkspace(ctx, stackName)
	if err != nil {
		if errors.Is(err, tfe.ErrInvalidWorkspaceValue) || errors.Is(err, tfe.ErrResourceNotFound) {
			// Workspace doesn't exist, thus no state
			return false, nil
		}
		return true, errors.Wrap(err, "Cannot get the stack workspace")
	}
	return workspace.HasState(ctx)
}

func (s *StackService) getDistributedLock() (*backend.DistributedLock, error) {
	lockConfig := backend.DistributedLockConfig{DynamodbTableName: s.backend.Conf().GetDynamoLocktableName()}
	return backend.NewDistributedLock(&lockConfig, s.backend.GetDynamoDBClient())
}

func (s *StackService) Generate(ctx context.Context) error {
	stackConfig, _, err := s.GetConfig().GetStackConfig()
	if err != nil {
		return errors.Wrap(err, "Unable to get stack config")
	}
	moduleSource := "git@github.com:chanzuckerberg/happy//terraform/modules/happy-stack-%s?ref=main"
	if s.GetConfig().TaskLaunchType() == util.LaunchTypeK8S {
		moduleSource = fmt.Sprintf(moduleSource, "eks")
	} else {
		moduleSource = fmt.Sprintf(moduleSource, "ecs")
	}
	if stackConfig.Source != nil && len(*stackConfig.Source) > 0 {
		moduleSource = *stackConfig.Source
	}

	_, modulePath, _, err := tf.ParseModuleSource(moduleSource)
	if err != nil {
		return errors.Wrap(err, "Unable to parse module path out")
	}
	modulePathParts := strings.Split(modulePath, "/")
	moduleName := modulePathParts[len(modulePathParts)-1]

	tempDir, err := os.MkdirTemp("", moduleName)
	if err != nil {
		return errors.Wrap(err, "Unable to create temp directory")
	}
	defer os.RemoveAll(tempDir)

	// Download the module source
	err = getter.GetAny(tempDir, moduleSource)
	if err != nil {
		return errors.Wrap(err, "Unable to download module source")
	}

	parser := tf.NewTfParser()

	// Extract variable information from the module
	variables, err := parser.ParseVariables(tempDir)
	if err != nil {
		return errors.Wrap(err, "Unable to parse out variables from the module")
	}

	outputs, err := parser.ParseOutputs(tempDir)
	if err != nil {
		return errors.Wrap(err, "Unable to parse out variables from the module")
	}

	tfDirPath := s.GetConfig().TerraformDirectory()

	happyProjectRoot := s.GetConfig().GetProjectRoot()
	srcDir := filepath.Join(happyProjectRoot, tfDirPath)

	gen := tf.NewTfGenerator(s.GetConfig())

	if _, err := os.Stat(srcDir); os.IsNotExist(err) {
		err = os.MkdirAll(srcDir, 0777)
		if err != nil {
			return errors.Wrapf(err, "Unable to create terraform directory: %s", srcDir)
		}
	}

	logrus.Infof("Generating terraform files in %s", srcDir)

	err = gen.GenerateMain(srcDir, moduleSource, variables)
	if err != nil {
		return errors.Wrap(err, "Unable to generate main.tf")
	}

	err = gen.GenerateProviders(srcDir)
	if err != nil {
		return errors.Wrap(err, "Unable to generate providers.tf")
	}

	err = gen.GenerateVersions(srcDir)
	if err != nil {
		return errors.Wrap(err, "Unable to generate versions.tf")
	}

	err = gen.GenerateOutputs(srcDir, outputs)
	if err != nil {
		return errors.Wrap(err, "Unable to generate outputs.tf")
	}

	err = gen.GenerateVariables(srcDir)
	if err != nil {
		return errors.Wrap(err, "Unable to generate variables.tf")
	}

	return nil
}
