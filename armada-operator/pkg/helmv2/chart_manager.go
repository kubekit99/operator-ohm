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

// +build v2

package helmv2

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	yaml "gopkg.in/yaml.v2"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/storage"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
	svc "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/tiller"
)

type chartmanager struct {
	storageBackend   *storage.Storage
	helmKubeClient   *kube.Client
	chartLocation    *av1.ArmadaChartSource

	releaseManager   *tiller.ReleaseServer
	releaseName string
	namespace   string

	spec   interface{}
	status *av1.ArmadaChartStatus

	isInstalled      bool
	isUpdateRequired bool
	deployedRelease  *helmif.HelmRelease
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
	// Replace this with sources from armada
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
	m.deployedRelease = &helmif.HelmRelease{deployedRelease}
	m.isInstalled = true

	// Get the next candidate release to determine if an update is necessary.
	candidateRelease, err := m.getCandidateRelease(ctx, m.releaseManager, m.releaseName, chart, config)
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
		// Still does not work right and cause fatal in the subsequent m.storageBackend.Create(release)
		// release = &rpb.Release{Name: condition.ResourceName, Version: condition.ResourceVersion}
		release = nil
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
	// so we need to reload it every time.
	chart, err := m.getChart()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load chart: %s", err)
	}

	cr, err := yaml.Marshal(m.spec)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse values: %s", err)
	}
	// JEB: Looks the Values field in the Chart with a bad structure
	// is messing the content in the "values.yaml" provided with the chart
	cr = make([]byte, 0)
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
	dryRunReq := &svc.UpdateReleaseRequest{
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
func (m chartmanager) InstallRelease(ctx context.Context) (*helmif.HelmRelease, error) {
	installedRelease, err := installRelease(ctx, m.releaseManager, m.namespace, m.releaseName, m.chart, m.config)
	return &helmif.HelmRelease{installedRelease}, err
}

// UpdateRelease performs a Helm release update.
func (m chartmanager) UpdateRelease(ctx context.Context) (*helmif.HelmRelease, *helmif.HelmRelease, error) {
	updatedRelease, err := updateRelease(ctx, m.releaseManager, m.releaseName, m.chart, m.config)
	return m.deployedRelease, &helmif.HelmRelease{updatedRelease}, err
}

// ReconcileRelease creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartmanager) ReconcileRelease(ctx context.Context) (*helmif.HelmRelease, error) {
	err := reconcileRelease(ctx, m.helmKubeClient, m.namespace, m.deployedRelease.GetManifest())
	return m.deployedRelease, err
}

// UninstallRelease performs a Helm release uninstall.
func (m chartmanager) UninstallRelease(ctx context.Context) (*helmif.HelmRelease, error) {
	uninstalledRelease, err := uninstallRelease(ctx, m.storageBackend, m.releaseManager, m.releaseName)
	return &helmif.HelmRelease{uninstalledRelease}, err
}

func (m chartmanager) getChart() (*cpb.Chart, error) {
	var pathToChart string
	var err error
	switch m.chartLocation.Type {
	case "git":
		pathToChart, err = m.gitClone()
	case "tar":
		pathToChart, err = m.getTarball()
	case "local":
		pathToChart = m.chartLocation.Location
	}

	if err != nil {
		return nil, err
	}

	chart, err := chartutil.LoadDir(pathToChart)
	if err != nil {
		return nil, err
	}

	err = sourceCleanup(pathToChart)
	if err != nil {
		// TODO: log a warning
	}

	return chart, nil
}

func (m *chartmanager) gitClone() (string, error) {
	// TODO(Ian): Finish this method
	repoURL := m.chartLocation.Location
	if repoURL == "" {
		return "", errors.New("Must provide a git url")
	}
	return "", nil
}

func (m *chartmanager) getTarball() (string, error) {
	tarballPath, err := m.downloadTarball(false)
	if err != nil {
		return "", err
	}
	return extractTarball(tarballPath)
}

// downloadTarball Downloads a tarball to /tmp and returns the path
func (m *chartmanager) downloadTarball(verify bool) (string, error) {
	file, err := ioutil.TempFile("", "armada")
	if err != nil {
		return "", err
	}
	response, err := http.Get(m.chartLocation.Location)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	file.Write(body)

	return file.Name(), nil
}

// extractTarball Extracts a tarball to /tmp and returns the path
func extractTarball(tarballPath string) (string, error) {
	if _, err := os.Stat(tarballPath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("%s does not exist", tarballPath)
		}
		return "", err
	}

	tempDir, err := ioutil.TempDir("", "armada")
	if err != nil {
		return "", err
	}

	fileContents, err := os.Open(tarballPath)
	if err != nil {
		return "", err
	}

	gzr, err := gzip.NewReader(fileContents)
	if err != nil {
		return "", err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	done := false
	for !done {
		if err := readFromArchive(tr, tempDir); err != nil {
			if err != io.EOF {
				return "", err
			}
			// io.EOF means there's no more data to be read
			done = true
		}
	}
	return tempDir, nil
}

// readFromArchive reads a an item from tr, saves it to dir, then move tr to the next item
func readFromArchive(tr *tar.Reader, dir string) error {
	header, err := tr.Next()
	if err != nil {
		// This catches EOF, which means that we're done
		return err
	}

	if header == nil {
		// if the header is nil, just skip it (not sure how this happens)
		return nil
	}

	target := filepath.Join(dir, header.Name)

	switch header.Typeflag {
	case tar.TypeDir:
		if _, err := os.Stat(target); err != nil {
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		}
	case tar.TypeReg:
		f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			return err
		}

		// copy over contents
		if _, err := io.Copy(f, tr); err != nil {
			return err
		}

		// manually close here after each file operation; defering would cause each file close
		// to wait until all operations have completed.
		f.Close()
	}
	return nil
}

func sourceCleanup(path string) error {
	// TODO(Ian): Finish this method
	return nil
}