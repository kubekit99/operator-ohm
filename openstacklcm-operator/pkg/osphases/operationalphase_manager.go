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

type operationalphasemanager struct {
	phasemanager

	spec   av1.OperationalPhaseSpec
	status *av1.OperationalPhaseStatus

	deployedResource *av1.OperationalPhase
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this OperationalPhase CR
func (m *operationalphasemanager) Sync(ctx context.Context) error {
	m.deployedResource = &av1.OperationalPhase{}
	m.isInstalled = true
	m.isUpdateRequired = false

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this OperationalPhase CR
func (m operationalphasemanager) InstallResource(ctx context.Context) (*av1.OperationalPhase, error) {
	return &av1.OperationalPhase{}, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this OperationalPhase CR
func (m operationalphasemanager) UpdateResource(ctx context.Context) (*av1.OperationalPhase, *av1.OperationalPhase, error) {
	return m.deployedResource, &av1.OperationalPhase{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this OperationalPhase CR
func (m operationalphasemanager) ReconcileResource(ctx context.Context) (*av1.OperationalPhase, error) {
	return m.deployedResource, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this OperationalPhase CR
func (m operationalphasemanager) UninstallResource(ctx context.Context) (*av1.OperationalPhase, error) {
	return &av1.OperationalPhase{}, nil
}
