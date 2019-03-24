// Copyright 2019 The Openstack-Service-Lifecyle Authors
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

package osphases

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
	lcmif "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"
)

type managerFactory struct {
}

// NewManagerFactory returns a new Helm manager factory capable of installing and uninstalling releases.
func NewManagerFactory(mgr manager.Manager) lcmif.PhaseManagerFactory {
	return &managerFactory{}
}

// NewPlanningPhaseManager returns a new manager capable of controlling PlanningPhase phase of the service lifecyle
func (f managerFactory) NewPlanningPhaseManager(r *av1.PlanningPhase) lcmif.PlanningPhaseManager {
	return &planningphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewInstallPhaseManager returns a new manager capable of controlling InstallPhase phase of the service lifecyle
func (f managerFactory) NewInstallPhaseManager(r *av1.InstallPhase) lcmif.InstallPhaseManager {
	return &installphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewTestPhaseManager returns a new manager capable of controlling TestPhase phase of the service lifecyle
func (f managerFactory) NewTestPhaseManager(r *av1.TestPhase) lcmif.TestPhaseManager {
	return &testphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewTrafficRolloutPhaseManager returns a new manager capable of controlling TrafficRolloutPhase phase of the service lifecyle
func (f managerFactory) NewTrafficRolloutPhaseManager(r *av1.TrafficRolloutPhase) lcmif.TrafficRolloutPhaseManager {
	return &trafficrolloutphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewOperationalPhaseManager returns a new manager capable of controlling OperationalPhase phase of the service lifecyle
func (f managerFactory) NewOperationalPhaseManager(r *av1.OperationalPhase) lcmif.OperationalPhaseManager {
	return &operationalphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewTrafficDrainPhaseManager returns a new manager capable of controlling TrafficDrainPhase phase of the service lifecyle
func (f managerFactory) NewTrafficDrainPhaseManager(r *av1.TrafficDrainPhase) lcmif.TrafficDrainPhaseManager {
	return &trafficdrainphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewUpgradePhaseManager returns a new manager capable of controlling UpgradePhase phase of the service lifecyle
func (f managerFactory) NewUpgradePhaseManager(r *av1.UpgradePhase) lcmif.UpgradePhaseManager {
	return &upgradephasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewRollbackPhaseManager returns a new manager capable of controlling RollbackPhase phase of the service lifecyle
func (f managerFactory) NewRollbackPhaseManager(r *av1.RollbackPhase) lcmif.RollbackPhaseManager {
	return &rollbackphasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}

// NewDeletePhaseManager returns a new manager capable of controlling DeletePhase phase of the service lifecyle
func (f managerFactory) NewDeletePhaseManager(r *av1.DeletePhase) lcmif.DeletePhaseManager {
	return &deletephasemanager{
		phasemanager: phasemanager{renderer: nil, namespace: r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}
