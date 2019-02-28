package controller

import (
	"github.com/kubekit99/operator-ohm/armada-operator/pkg/controller/helm3crd"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, helm3crd.Add)
}
