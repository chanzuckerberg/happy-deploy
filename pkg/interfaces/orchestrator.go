package interfaces

import "context"

type OrchestratorInterface interface {
	GetEvents(ctx context.Context, stack string, services []string) error
	IsDryRun() bool
}
