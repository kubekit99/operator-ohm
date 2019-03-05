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
	"fmt"
	"strings"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	yaml "gopkg.in/yaml.v2"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/kube"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/storage"
	"k8s.io/helm/pkg/tiller"
)

type chartmanager struct {
	storageBackend   *storage.Storage
	tillerKubeClient *kube.Client
	chartDir         string

	tiller      *tiller.ReleaseServer
	releaseName string
	namespace   string

	spec   interface{}
	status *av1.ArmadaChartStatus

	isInstalled      bool
	isUpdateRequired bool
	deployedRelease  *rpb.Release
	chart            *cpb.Chart
	config           *cpb.Config
}

// ReleaseName returns the name of the release.
func (m chartmanager) ReleaseName() string {
	return m.releaseName
}

func (m chartmanager) IsInstalled() bool {
	return m.isInstalled
}

func (m chartmanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *chartmanager) Sync(ctx context.Context) error {
	if err := m.syncReleaseStatus(*m.status); err != nil {
		return fmt.Errorf("failed to sync release status to storage backend: %s", err)
	}

	// Get release history for this release name
	releases, err := m.storageBackend.History(m.releaseName)
	if err != nil && !notFoundErr(err) {
		return fmt.Errorf("failed to retrieve release history: %s", err)
	}

	// Cleanup non-deployed release versions. If all release versions are
	// non-deployed, this will ensure that failed installations are correctly
	// retried.
	for _, rel := range releases {
		if rel.GetInfo().GetStatus().GetCode() != rpb.Status_DEPLOYED {
			_, err := m.storageBackend.Delete(rel.GetName(), rel.GetVersion())
			if err != nil && !notFoundErr(err) {
				return fmt.Errorf("failed to delete stale release version: %s", err)
			}
		}
	}

	// Load the chart and config based on the current state of the custom resource.
	chart, config, err := m.loadChartAndConfig()
	if err != nil {
		return fmt.Errorf("failed to load chart and config: %s", err)
	}
	m.chart = chart
	m.config = config

	// Load the most recently deployed release from the storage backend.
	deployedRelease, err := m.getDeployedRelease()
	if err == helmif.ErrNotFound {
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to get deployed release: %s", err)
	}
	m.deployedRelease = deployedRelease
	m.isInstalled = true

	// Get the next candidate release to determine if an update is necessary.
	candidateRelease, err := m.getCandidateRelease(ctx, m.tiller, m.releaseName, chart, config)
	if err != nil {
		return fmt.Errorf("failed to get candidate release: %s", err)
	}
	if deployedRelease.GetManifest() != candidateRelease.GetManifest() {
		m.isUpdateRequired = true
	}

	return nil
}

func (m chartmanager) syncReleaseStatus(status av1.ArmadaChartStatus) error {
	var release *rpb.Release
	helper := av1.HelmResourceConditionListHelper{Items: status.Conditions}
	condition := helper.FindCondition(av1.ConditionDeployed, av1.ConditionStatusTrue)
	if condition == nil {
		return nil
	} else {
		// JEB: Big issue here. Original code was saving the release in the Condition
		release = &rpb.Release{Name: condition.ResourceName, Version: condition.ResourceVersion}
	}
	if release == nil {
		return nil
	}

	name := release.GetName()
	version := release.GetVersion()
	_, err := m.storageBackend.Get(name, version)
	if err == nil {
		return nil
	}

	if !notFoundErr(err) {
		return err
	}
	return m.storageBackend.Create(release)
}

func (m chartmanager) loadChartAndConfig() (*cpb.Chart, *cpb.Config, error) {
	// chart is mutated by the call to processRequirements,
	// so we need to reload it from disk every time.
	chart, err := chartutil.LoadDir(m.chartDir)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load chart: %s", err)
	}

	cr, err := yaml.Marshal(m.spec)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse values: %s", err)
	}
	config := &cpb.Config{Raw: string(cr)}

	err = processRequirements(chart, config)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to process chart requirements: %s", err)
	}
	return chart, config, nil
}

func (m chartmanager) getDeployedRelease() (*rpb.Release, error) {
	deployedRelease, err := m.storageBackend.Deployed(m.releaseName)
	if err != nil {
		if strings.Contains(err.Error(), "has no deployed releases") {
			return nil, helmif.ErrNotFound
		}
		return nil, err
	}
	return deployedRelease, nil
}

func (m chartmanager) getCandidateRelease(ctx context.Context, tiller *tiller.ReleaseServer, name string, chart *cpb.Chart, config *cpb.Config) (*rpb.Release, error) {
	dryRunReq := &services.UpdateReleaseRequest{
		Name:   name,
		Chart:  chart,
		Values: config,
		DryRun: true,
	}
	dryRunResponse, err := tiller.UpdateRelease(ctx, dryRunReq)
	if err != nil {
		return nil, err
	}
	return dryRunResponse.GetRelease(), nil
}

// InstallRelease performs a Helm release install.
func (m chartmanager) InstallRelease(ctx context.Context) (*rpb.Release, error) {
	return installRelease(ctx, m.tiller, m.namespace, m.releaseName, m.chart, m.config)
}

// UpdateRelease performs a Helm release update.
func (m chartmanager) UpdateRelease(ctx context.Context) (*rpb.Release, *rpb.Release, error) {
	updatedRelease, err := updateRelease(ctx, m.tiller, m.releaseName, m.chart, m.config)
	return m.deployedRelease, updatedRelease, err
}

// ReconcileRelease creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartmanager) ReconcileRelease(ctx context.Context) (*rpb.Release, error) {
	err := reconcileRelease(ctx, m.tillerKubeClient, m.namespace, m.deployedRelease.GetManifest())
	return m.deployedRelease, err
}

// UninstallRelease performs a Helm release uninstall.
func (m chartmanager) UninstallRelease(ctx context.Context) (*rpb.Release, error) {
	return uninstallRelease(ctx, m.storageBackend, m.tiller, m.releaseName)
}
