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

type upgradephasemanager struct {
	phasemanager

	spec   av1.UpgradePhaseSpec
	status *av1.UpgradePhaseStatus

	deployedResource *av1.UpgradePhase
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this UpgradePhase CR
func (m *upgradephasemanager) Sync(ctx context.Context) error {
	m.deployedResource = &av1.UpgradePhase{}
	m.isInstalled = true
	m.isUpdateRequired = false

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this UpgradePhase CR
func (m upgradephasemanager) InstallResource(ctx context.Context) (*av1.UpgradePhase, error) {
	return &av1.UpgradePhase{}, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this UpgradePhase CR
func (m upgradephasemanager) UpdateResource(ctx context.Context) (*av1.UpgradePhase, *av1.UpgradePhase, error) {
	return m.deployedResource, &av1.UpgradePhase{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this UpgradePhase CR
func (m upgradephasemanager) ReconcileResource(ctx context.Context) (*av1.UpgradePhase, error) {
	return m.deployedResource, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this UpgradePhase CR
func (m upgradephasemanager) UninstallResource(ctx context.Context) (*av1.UpgradePhase, error) {
	return &av1.UpgradePhase{}, nil
}
