package argoif

import (
	"context"
	"errors"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
)

var (
	// ErrNotFound indicates the release was not found.
	ErrNotFound = errors.New("release not found")
)

// Manager manages a Helm release. It can install, update, reconcile,
// and uninstall a release.
type Manager interface {
	WorkflowName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallWorkflow(context.Context) (*wfv1.Workflow, error)
	UpdateWorkflow(context.Context) (*wfv1.Workflow, *wfv1.Workflow, error)
	ReconcileWorkflow(context.Context) (*wfv1.Workflow, error)
	UninstallWorkflow(context.Context) (*wfv1.Workflow, error)
}
