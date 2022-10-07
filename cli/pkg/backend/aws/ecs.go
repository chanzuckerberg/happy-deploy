package aws

import (
	"context"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecstypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// TODO: Move this code into ecs_compute_backend

// checks if a happy service is a part of a happy stack and service combination. For now,
// we just check if both the happy stack and service names are includes in the ecs service name.
// TODO: make this more comprehensive in the future.
// I'm being a little loose with this matching to account for this convention and also
// give us some wiggle room to change this convention in the future.
func isStackECSService(happyServiceName, happyStackName string, ecsService ecstypes.Service) bool {
	if strings.Contains(*ecsService.ServiceName, happyServiceName) &&
		strings.Contains(*ecsService.ServiceName, happyStackName) {
		return true
	}
	return false
}

// GetECSServicesForStackService returns the ECS services that are associated with a happy stack and service.
// The filter is based on the name of the stack and the service name provided in the docker-compose file.
func (b *Backend) GetECSServicesForStackService(ctx context.Context, stackName, serviceName string) ([]ecstypes.Service, error) {
	clusterARN := b.integrationSecret.ClusterArn
	ls, err := b.ecsclient.ListServices(ctx, &ecs.ListServicesInput{
		Cluster: &clusterARN,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to list ECS services for a stack")
	}

	ds, err := b.ecsclient.DescribeServices(ctx, &ecs.DescribeServicesInput{
		Cluster:  &clusterARN,
		Services: ls.ServiceArns,
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to describe ECS services for stack")
	}

	// TODO: right now, happy has no control over what these ECS services are called
	// but a convention has started where the stack name is a part of the service name
	// and so is the docker-compose service name. Usually, its of the form <stackname>-<docker-compose-service-name>.
	stackServNames := []ecstypes.Service{}
	for _, s := range ds.Services {
		if isStackECSService(serviceName, stackName, s) {
			stackServNames = append(stackServNames, s)
		}
	}

	return stackServNames, nil
}

// GetECSTasksForStackService returns the task ARNs associated with a particular happy stack and service.
func (b *Backend) GetECSTasksForStackService(ctx context.Context, stackName, serviceName string) ([]string, error) {
	stackServNames, err := b.GetECSServicesForStackService(ctx, stackName, serviceName)
	if err != nil {
		return nil, err
	}

	clusterARN := b.integrationSecret.ClusterArn
	stackTaskARNs := []string{}
	for _, s := range stackServNames {
		lt, err := b.ecsclient.ListTasks(ctx, &ecs.ListTasksInput{
			Cluster:     &clusterARN,
			ServiceName: s.ServiceName,
		})

		if err != nil {
			return nil, errors.Wrapf(err, "unable to list ECS tasks for stack %s", *s.ServiceName)
		}

		stackTaskARNs = append(stackTaskARNs, lt.TaskArns...)
	}

	log.Debugf("found task ARNs associated with %s-%s: %+v", stackName, serviceName, stackTaskARNs)
	return stackTaskARNs, nil
}

// GetTaskDefinitions returns the task definitions for a slice of task ARNs.
func (b *Backend) GetTaskDefinitions(ctx context.Context, taskArns []string) ([]ecstypes.TaskDefinition, error) {
	tasks, err := b.GetTaskDetails(ctx, taskArns)
	if err != nil {
		return []ecstypes.TaskDefinition{}, errors.Wrap(err, "cannot describe ECS tasks")
	}
	taskDefinitions := []ecstypes.TaskDefinition{}
	for _, task := range tasks {
		taskDefResult, err := b.ecsclient.DescribeTaskDefinition(
			ctx,
			&ecs.DescribeTaskDefinitionInput{TaskDefinition: task.TaskDefinitionArn},
		)
		if err != nil {
			return []ecstypes.TaskDefinition{}, errors.Wrap(err, "cannot retrieve a task definition")
		}
		taskDefinitions = append(taskDefinitions, *taskDefResult.TaskDefinition)
	}
	return taskDefinitions, nil
}

// GetTaskDetails is a helper function to return the described task objects associated with a slice
// of task ARNs.
func (b *Backend) GetTaskDetails(ctx context.Context, taskArns []string) ([]ecstypes.Task, error) {
	clusterARN := b.integrationSecret.ClusterArn
	tasksResult, err := b.ecsclient.DescribeTasks(ctx, &ecs.DescribeTasksInput{
		Cluster: &clusterARN,
		Tasks:   taskArns,
	})
	if err != nil {
		return []ecstypes.Task{}, errors.Wrap(err, "could not describe tasks")
	}
	return tasksResult.Tasks, nil
}

func (ab *Backend) getTaskID(taskARN string) (string, error) {
	resourceArn, err := arn.Parse(taskARN)
	if err != nil {
		return "", errors.Wrapf(err, "unable to parse task ARN: '%s'", taskARN)
	}

	segments := strings.Split(resourceArn.Resource, "/")
	if len(segments) < 3 {
		return "", errors.Errorf("incomplete task ARN: '%s'", taskARN)
	}
	return segments[len(segments)-1], nil
}

type AWSLogConfiguration struct {
	GroupName   string
	StreamNames []string
}

const (
	AwsLogsGroup        = "awslogs-group"
	AwsLogsStreamPrefix = "awslogs-stream-prefix"
	AwsLogsRegion       = "awslogs-region"
)
