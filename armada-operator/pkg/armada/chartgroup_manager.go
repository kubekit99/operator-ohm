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
	deployedResource *av1.ArmadaChart
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

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *chartgroupmanager) Sync(ctx context.Context) error {
	existingResource := m.newResourceForCR()
	err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.Name, Namespace: existingResource.Namespace}, existingResource)
	if err != nil {
		m.isInstalled = false
		m.deployedResource = nil
		if apierrors.IsNotFound(err) {
			return nil
		} else {
			log.Error(err, "Can't not Sync Resource")
			return err
		}
	}

	m.isInstalled = true
	m.deployedResource = existingResource

	targetResource := m.newResourceForCR()
	if !targetResource.Equivalent(m.deployedResource) {
		m.isUpdateRequired = true
	} else {
		m.isUpdateRequired = false
	}

	return nil
}

func (m chartgroupmanager) InstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	newResource := m.newResourceForCR()
	err := m.kubeClient.Create(context.TODO(), newResource)
	if err != nil {
		log.Error(err, "Can't not Create Resource")
		return nil, err
	}
	return newResource.FromArmadaChart(), nil
}

// UpdateResource performs a Helm release update.
func (m chartgroupmanager) UpdateResource(ctx context.Context) (*unstructured.Unstructured, *unstructured.Unstructured, error) {
	toUpdate := m.newResourceForCR()
	err := m.kubeClient.Update(context.TODO(), toUpdate)
	if err != nil {
		log.Error(err, "Can't not Update Resource")
		if apierrors.IsNotFound(err) {
			return nil, nil, armadaif.ErrNotFound
		} else {
			return nil, nil, err
		}
	}
	return m.deployedResource.FromArmadaChart(), toUpdate.FromArmadaChart(), nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartgroupmanager) ReconcileResource(ctx context.Context) (*unstructured.Unstructured, error) {
	toReconcile := m.newResourceForCR()
	return toReconcile.FromArmadaChart(), nil
}

// UninstallResource performs a Helm release uninstall.
func (m chartgroupmanager) UninstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	toDelete := m.newResourceForCR()
	err := m.kubeClient.Delete(context.TODO(), toDelete)
	if err != nil {
		log.Error(err, "Can't not Delete Resource")
		if apierrors.IsNotFound(err) {
			return nil, armadaif.ErrNotFound
		} else {
			return nil, err
		}
	}
	return toDelete.FromArmadaChart(), nil
}

// newResourceForCR returns a dummy ArmadaChart the same name/namespace as the cr
func (m chartgroupmanager) newResourceForCR() *av1.ArmadaChart {
	labels := map[string]string{
		"app": m.resourceName,
	}

	return &av1.ArmadaChart{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.resourceName + "-act",
			Namespace: m.namespace,
			Labels:    labels,
		},
		Spec: av1.ArmadaChartSpec{
			ChartName: m.resourceName + "-act",
			Release:   m.resourceName + "-release",
			Namespace: m.namespace,
			Upgrade: &av1.ArmadaUpgrade{
				NoHooks: false,
			},
			Source: &av1.ArmadaChartSource{
				Type:      "git",
				Location:  "https://github.com/gardlt/hello-world-chart",
				Subpath:   ".",
				Reference: "master",
			},
			Dependencies: make([]string, 0),
			TargetState:  av1.StateInitialized,
		},
	}
}
