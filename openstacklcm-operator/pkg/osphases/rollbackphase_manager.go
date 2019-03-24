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

type rollbackphasemanager struct {
	phasemanager

	spec   av1.RollbackPhaseSpec
	status *av1.RollbackPhaseStatus

	deployedResource *av1.RollbackPhase
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this RollbackPhase CR
func (m *rollbackphasemanager) Sync(ctx context.Context) error {
	m.deployedResource = &av1.RollbackPhase{}
	m.isInstalled = true
	m.isUpdateRequired = false

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this RollbackPhase CR
func (m rollbackphasemanager) InstallResource(ctx context.Context) (*av1.RollbackPhase, error) {
	return &av1.RollbackPhase{}, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this RollbackPhase CR
func (m rollbackphasemanager) UpdateResource(ctx context.Context) (*av1.RollbackPhase, *av1.RollbackPhase, error) {
	return m.deployedResource, &av1.RollbackPhase{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this RollbackPhase CR
func (m rollbackphasemanager) ReconcileResource(ctx context.Context) (*av1.RollbackPhase, error) {
	return m.deployedResource, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this RollbackPhase CR
func (m rollbackphasemanager) UninstallResource(ctx context.Context) (*av1.RollbackPhase, error) {
	return &av1.RollbackPhase{}, nil
}
