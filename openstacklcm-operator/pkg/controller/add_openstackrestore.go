package controller

import (
	"github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/controller/openstackrestore"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, openstackrestore.Add)
}
