package apis

import (
	"github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/helm3crd/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1beta1.SchemeBuilder.AddToScheme)
}
