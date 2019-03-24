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
	renderer    interface{}
	releaseName string
	namespace   string

	spec   interface{}
	status *av1.OslcStatus

	deployedRelease  *av1.Oslc
	isInstalled      bool
	isUpdateRequired bool
	config           *map[string]interface{}
}

// ReleaseName returns the name of the release.
func (m oslcmanager) ReleaseName() string {
	return m.releaseName
}

func (m oslcmanager) IsInstalled() bool {
	return m.isInstalled
}

func (m oslcmanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *oslcmanager) Sync(ctx context.Context) error {
	return nil
}

func (m oslcmanager) InstallRelease(ctx context.Context) (*av1.Oslc, error) {
	return &av1.Oslc{}, nil
}

// UpdateRelease performs a Helm release update.
func (m oslcmanager) UpdateRelease(ctx context.Context) (*av1.Oslc, *av1.Oslc, error) {
	return m.deployedRelease, &av1.Oslc{}, nil
}

// ReconcileRelease creates or patches resources as necessary to match the
// deployed release's manifest.
func (m oslcmanager) ReconcileRelease(ctx context.Context) (*av1.Oslc, error) {
	return m.deployedRelease, nil
}

// UninstallRelease performs a Helm release uninstall.
func (m oslcmanager) UninstallRelease(ctx context.Context) (*av1.Oslc, error) {
	return &av1.Oslc{}, nil
}
