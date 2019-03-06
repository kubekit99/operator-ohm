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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type chartgroupmanager struct {
	kubeClient       client.Client
	resourceName     string
	namespace        string
	spec             *av1.ArmadaChartGroupSpec
	status           *av1.ArmadaChartGroupStatus
	deployedResource *av1.ArmadaChartList
	isInstalled      bool
	isUpdateRequired bool
}

// ResourceName returns the name of the release.
func (m chartgroupmanager) ResourceName() string {
	return m.resourceName
}

func (m chartgroupmanager) IsInstalled() bool {
	return m.isInstalled
}

func (m chartgroupmanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync detects which ArmadaChart are already present for that ArmadaChartGroup
func (m *chartgroupmanager) Sync(ctx context.Context) error {
	m.deployedResource = &av1.ArmadaChartList{Items: make([]av1.ArmadaChart, 0)}
	errs := make([]error, 0)
	targetResourceList := m.newResourceForCR()
	for _, existingResource := range (*targetResourceList).Items {
		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.Name, Namespace: existingResource.Namespace}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(errs[0]) {
				log.Error(err, "Can't not Sync ArmadaChart")
			}
			errs = append(errs, err)
		} else {
			m.deployedResource.Items = append(m.deployedResource.Items, existingResource)
		}
	}

	// Let's check if some of the ArmaChart are already present.
	// If yes, let's consider the ArmadaChartGroup as installed and we will update it.
	if len(m.deployedResource.Items) == 0 {
		m.isInstalled = false
		return nil
	} else {
		m.isInstalled = true
	}

	if len(targetResourceList.Items) != len(m.deployedResource.Items) {
		m.isUpdateRequired = true
	} else {
		m.isUpdateRequired = false
	}

	return nil
}

func (m chartgroupmanager) InstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	errs := make([]error, 0)
	toInstallList := m.newResourceForCR()
	for _, toInstall := range (*toInstallList).Items {
		err := m.kubeClient.Create(context.TODO(), &toInstall)
		if err != nil {
			log.Error(err, "Can't not Create ArmadaChart")
			errs = append(errs, err)
		}
	}
	if len(errs) != 0 {
		return nil, errs[0]
	}
	return toInstallList.FromArmadaChartList(), nil
}

// UpdateResource performs a Helm release update.
func (m chartgroupmanager) UpdateResource(ctx context.Context) (*unstructured.Unstructured, *unstructured.Unstructured, error) {
	errs := make([]error, 0)
	toUpdateList := m.newResourceForCR()
	for _, toUpdate := range (*toUpdateList).Items {
		err := m.kubeClient.Update(context.TODO(), &toUpdate)
		if err != nil {
			log.Error(err, "Can't not Update ArmadaChart")
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
	return m.deployedResource.FromArmadaChartList(), toUpdateList.FromArmadaChartList(), nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartgroupmanager) ReconcileResource(ctx context.Context) (*unstructured.Unstructured, error) {
	toReconcile := m.newResourceForCR()
	return toReconcile.FromArmadaChartList(), nil
}

// UninstallResource performs a Helm release uninstall.
func (m chartgroupmanager) UninstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	errs := make([]error, 0)
	toDeleteList := m.newResourceForCR()
	for _, toDelete := range (*toDeleteList).Items {
		err := m.kubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			log.Error(err, "Can't not Delete ArmadaChart")
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
	return toDeleteList.FromArmadaChartList(), nil
}

// newResourceForCR returns a dummy ArmadaChart the same name/namespace as the cr
func (m chartgroupmanager) newResourceForCR() *av1.ArmadaChartList {
	labels := map[string]string{
		"app": m.resourceName,
	}

	var res = av1.ArmadaChartList{
		Items: make([]av1.ArmadaChart, 0),
	}

	for _, chartname := range m.spec.Charts {
		res.Items = append(res.Items, av1.ArmadaChart{
			ObjectMeta: metav1.ObjectMeta{
				Name:      chartname,
				Namespace: m.namespace,
				Labels:    labels,
			},
			Spec: av1.ArmadaChartSpec{
				ChartName: chartname,
				Release:   m.resourceName + "-release",
				Namespace: m.namespace,
				Upgrade: &av1.ArmadaUpgrade{
					NoHooks: false,
				},
				Source: &av1.ArmadaChartSource{
					Type:      "local",
					Location:  "/opt/armada/armada-charts/tiller-testchart/helm",
					Subpath:   ".",
					Reference: "master",
				},
				Dependencies: make([]string, 0),
				TargetState:  av1.StateInitialized,
			},
		})
	}

	return &res
}
