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
	"context"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
)

type trafficdrainmanager struct {
	phasemanager

	spec   av1.TrafficDrainPhaseSpec
	status *av1.TrafficDrainPhaseStatus
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this TrafficDrainPhase CR
func (m *trafficdrainmanager) Sync(ctx context.Context) error {

	m.deployedSubResourceList = av1.NewSubResourceList(m.phaseNamespace, m.phaseName)

	rendered, deployed, err := m.sync(ctx)
	if err != nil {
		return err
	}

	m.deployedSubResourceList = deployed
	if len(rendered.Items) != len(deployed.Items) {
		m.isInstalled = false
		m.isUpdateRequired = false
	} else {
		m.isInstalled = true
		m.isUpdateRequired = false
	}

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this TrafficDrainPhase CR
func (m trafficdrainmanager) InstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.installResource(ctx)
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this TrafficDrainPhase CR
func (m trafficdrainmanager) UpdateResource(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	return m.updateResource(ctx)
}

// ReconcileResource creates or patches resources as necessary to match this TrafficDrainPhase CR
func (m trafficdrainmanager) ReconcileResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.reconcileResource(ctx)
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this TrafficDrainPhase CR
func (m trafficdrainmanager) UninstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.uninstallResource(ctx)
}
