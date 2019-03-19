// Copyright 2019 The Armada Authors
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

package armada

import (
	"context"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	armadaif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type manifestmanager struct {
	kubeClient       client.Client
	resourceName     string
	namespace        string
	spec             *av1.ArmadaManifestSpec
	status           *av1.ArmadaManifestStatus
	deployedResource *av1.ArmadaChartGroups
	isInstalled      bool
	isUpdateRequired bool
}

// ResourceName returns the name of the release.
func (m manifestmanager) ResourceName() string {
	return m.resourceName
}

func (m manifestmanager) IsInstalled() bool {
	return m.isInstalled
}

func (m manifestmanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync detects which ArmadaChartGroup listed this ArmadaManifest are already present in
// the K8s cluster.
func (m *manifestmanager) Sync(ctx context.Context) error {
	m.deployedResource = av1.NewArmadaChartGroups(m.resourceName)
	errs := make([]error, 0)
	targetResourceList := m.expectedChartGroupList()
	for _, existingResource := range (*targetResourceList).List.Items {
		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.Name, Namespace: existingResource.Namespace}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				log.Error(err, "Can't not Sync ArmadaChartGroup")
			}
			errs = append(errs, err)
		} else {
			m.deployedResource.List.Items = append(m.deployedResource.List.Items, existingResource)
		}
	}

	// Let's check if some of the ArmaChartGroup are already present.
	// If yes, let's consider the ArmadaManifest as installed and we will update it.
	if len(m.deployedResource.List.Items) == 0 {
		m.isInstalled = false
		return nil
	} else {
		m.isInstalled = true
	}

	if len(targetResourceList.List.Items) != len(m.deployedResource.List.Items) {
		m.isUpdateRequired = true
	} else {
		m.isUpdateRequired = false
	}

	return nil

}

// InstallResource currently create dummy ChartGroups in the K8s cluster if those are not found.
// This is probably not the behavior we want to maitain in the long run.
func (m manifestmanager) InstallResource(ctx context.Context) (*av1.ArmadaChartGroups, error) {
	errs := make([]error, 0)
	toInstallList := m.expectedChartGroupList()
	for _, toInstall := range (*toInstallList).List.Items {
		err := m.kubeClient.Create(context.TODO(), &toInstall)
		if err != nil {
			log.Error(err, "Can't not Create ArmadaChartGroup")
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		return nil, errs[0]
	}
	return toInstallList, nil
}

// UpdateResource performs an update of an ArmadaManifest.
// Currently either the list of ChartGroupss or the "prefix" attribute may have changed.
func (m manifestmanager) UpdateResource(ctx context.Context) (*av1.ArmadaChartGroups, *av1.ArmadaChartGroups, error) {
	errs := make([]error, 0)
	toUpdateList := m.expectedChartGroupList()
	for _, toUpdate := range (*toUpdateList).List.Items {
		err := m.kubeClient.Update(context.TODO(), &toUpdate)
		if err != nil {
			log.Error(err, "Can't not Update ArmadaChartGroup")
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return nil, nil, armadaif.ErrNotFound
		} else {
			return nil, nil, errs[0]
		}
	}
	return m.deployedResource, toUpdateList, nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m manifestmanager) ReconcileResource(ctx context.Context) (*av1.ArmadaChartGroups, error) {
	toReconcile := m.expectedChartGroupList()
	return toReconcile, nil
}

// UninstallResource currently delete ChartGroups matching the manifest.
// This is probably not the behavior we want to maitain in the long run.
func (m manifestmanager) UninstallResource(ctx context.Context) (*av1.ArmadaChartGroups, error) {
	errs := make([]error, 0)
	toDeleteList := m.expectedChartGroupList()
	for _, toDelete := range (*toDeleteList).List.Items {
		err := m.kubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			log.Error(err, "Can't not Delete ArmadaChartGroup")
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return nil, armadaif.ErrNotFound
		} else {
			return nil, errs[0]
		}
	}
	return toDeleteList, nil
}

// expectedChartGroupList returns a dummy list of ArmadaChartGroup the same name/namespace as the cr
func (m manifestmanager) expectedChartGroupList() *av1.ArmadaChartGroups {
	labels := map[string]string{
		"app": m.resourceName,
	}

	var res = av1.NewArmadaChartGroups(m.resourceName)

	for _, chartgroupname := range m.spec.ChartGroups {
		res.List.Items = append(res.List.Items,
			av1.ArmadaChartGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      chartgroupname,
					Namespace: m.namespace,
					Labels:    labels,
				},
				Spec: av1.ArmadaChartGroupSpec{
					Charts:      make([]string, 0),
					Description: "Created by " + m.resourceName,
					Name:        chartgroupname,
					Sequenced:   false,
					TestCharts:  false,
					TargetState: av1.StateInitialized,
					AdminState:  av1.StateDisabled,
				},
			},
		)
	}

	return res
}
