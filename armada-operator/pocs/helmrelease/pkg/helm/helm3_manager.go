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

package helm

import (
	"context"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"

	"k8s.io/helm/pkg/kube"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
)

type helm3manager struct {
	tillerKubeClient *kube.Client
	chartDir         string

	releaseName string
	namespace   string

	spec   interface{}
	status *av1.HelmReleaseStatus

	isInstalled      bool
	isUpdateRequired bool
	deployedRelease  *rpb.Release
	chart            *cpb.Chart
	config           *cpb.Config
}

// ReleaseName returns the name of the release.
func (m helm3manager) ReleaseName() string {
	return m.releaseName
}

func (m helm3manager) IsInstalled() bool {
	return m.isInstalled
}

func (m helm3manager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *helm3manager) Sync(ctx context.Context) error {
	return nil
}

// InstallRelease performs a Helm release install.
func (m helm3manager) InstallRelease(ctx context.Context) (*rpb.Release, error) {
	return nil, nil
}

// UpdateRelease performs a Helm release update.
func (m helm3manager) UpdateRelease(ctx context.Context) (*rpb.Release, *rpb.Release, error) {
	return nil, nil, nil
}

// ReconcileRelease creates or patches resources as necessary to match the
// deployed release's manifest.
func (m helm3manager) ReconcileRelease(ctx context.Context) (*rpb.Release, error) {
	return nil, nil
}

// UninstallRelease performs a Helm release uninstall.
func (m helm3manager) UninstallRelease(ctx context.Context) (*rpb.Release, error) {
	return nil, nil
}
