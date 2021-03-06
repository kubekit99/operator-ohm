package controller

import (
	"github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/controller/openstackhelm"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, openstackhelm.Add)
}
