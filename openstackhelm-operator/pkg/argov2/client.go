package argov2

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// NewFromManager returns a Kubernetes client that can be used with
// argo CRD
func NewFromManager(mgr manager.Manager) (*client.Client, error) {
	clt := mgr.GetClient()
	return &clt, nil
}
