package controller

import (
	"github.com/kubekit99/operator-ohm/armada-operator/pkg/controller/armada"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, armada.AddArmadaManifestController)
	AddToManagerFuncs = append(AddToManagerFuncs, armada.AddArmadaChartGroupController)
	AddToManagerFuncs = append(AddToManagerFuncs, armada.AddArmadaChartController)
}
