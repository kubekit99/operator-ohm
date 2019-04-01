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

package oslc

import (
	"context"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
)

type oslcmanager struct {
	basemanager

	spec   av1.OslcSpec
	status *av1.OslcStatus
}

// Sync retrieves from K8s the sub resources (Workflow, Job, ....) attached to this Oslc CR
func (m *oslcmanager) Sync(ctx context.Context) error {

	m.deployedLifecycleFlow = av1.NewLifecycleFlow(m.serviceNamespace, m.serviceName)

	rendered, deployed, err := m.sync(ctx)
	if err != nil {
		return err
	}

	m.deployedLifecycleFlow = deployed
	if len(rendered.Phases) != len(deployed.Phases) {
		m.isInstalled = false
		m.isUpdateRequired = false
	} else {
		m.isInstalled = true
		m.isUpdateRequired = false
	}

	return nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this Oslc CR
func (m oslcmanager) InstallResource(ctx context.Context) (*av1.LifecycleFlow, error) {
	return m.installResource(ctx)
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this Oslc CR
func (m oslcmanager) UpdateResource(ctx context.Context) (*av1.LifecycleFlow, *av1.LifecycleFlow, error) {
	return m.updateResource(ctx)
}

// ReconcileResource creates or patches resources as necessary to match this Oslc CR
func (m oslcmanager) ReconcileResource(ctx context.Context) (*av1.LifecycleFlow, error) {
	return m.reconcileResource(ctx)
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this Oslc CR
func (m oslcmanager) UninstallResource(ctx context.Context) (*av1.LifecycleFlow, error) {
	return m.uninstallResource(ctx)
}
