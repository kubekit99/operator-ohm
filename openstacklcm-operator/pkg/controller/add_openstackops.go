package controller

import (
	"github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/controller/openstackops"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, openstackops.AddOpenstackBackup)
	AddToManagerFuncs = append(AddToManagerFuncs, openstackops.AddOpenstackDeployment)
	AddToManagerFuncs = append(AddToManagerFuncs, openstackops.AddOpenstackRestore)
	AddToManagerFuncs = append(AddToManagerFuncs, openstackops.AddOpenstackRollback)
	AddToManagerFuncs = append(AddToManagerFuncs, openstackops.AddOpenstackUpgrade)
}
