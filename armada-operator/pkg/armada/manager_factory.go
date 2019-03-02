// Copyright 2018 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package armada

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	armadaif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"
)

type managerFactory struct {
}

// NewManagerFactory returns a new Helm manager factory capable of installing and uninstalling releases.
func NewManagerFactory(mgr manager.Manager) armadaif.ArmadaManagerFactory {
	return &managerFactory{}
}

func (f managerFactory) NewArmadaChartManager(r *av1.ArmadaChart) armadaif.ArmadaManager {
	return &chartmanager{}
}

func (f managerFactory) NewArmadaChartGroupManager(r *av1.ArmadaChartGroup) armadaif.ArmadaManager {
	return &chartgroupmanager{}
}

func (f managerFactory) NewArmadaManifestManager(r *av1.ArmadaManifest) armadaif.ArmadaManager {
	return &manifestmanager{}
}
