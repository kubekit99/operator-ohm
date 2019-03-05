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

type manifestmanager struct {
	kubeClient       client.Client
	resourceName     string
	namespace        string
	spec             *av1.ArmadaManifestSpec
	status           *av1.ArmadaManifestStatus
	deployedResource *av1.ArmadaChartGroup
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

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *manifestmanager) Sync(ctx context.Context) error {
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

func (m manifestmanager) InstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	newResource := m.newResourceForCR()
	err := m.kubeClient.Create(context.TODO(), newResource)
	if err != nil {
		log.Error(err, "Can't not Create Resource")
		return nil, err
	}
	return newResource.FromArmadaChartGroup(), nil
}

// UpdateResource performs a Helm release update.
func (m manifestmanager) UpdateResource(ctx context.Context) (*unstructured.Unstructured, *unstructured.Unstructured, error) {
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
	return m.deployedResource.FromArmadaChartGroup(), toUpdate.FromArmadaChartGroup(), nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m manifestmanager) ReconcileResource(ctx context.Context) (*unstructured.Unstructured, error) {
	toReconcile := m.newResourceForCR()
	return toReconcile.FromArmadaChartGroup(), nil
}

// UninstallResource performs a Helm release uninstall.
func (m manifestmanager) UninstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
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
	return toDelete.FromArmadaChartGroup(), nil
}

// newResourceForCR returns a dummy ArmadaChartGroup the same name/namespace as the cr
func (m manifestmanager) newResourceForCR() *av1.ArmadaChartGroup {
	labels := map[string]string{
		"app": m.resourceName,
	}

	var charts = make([]string, 0)
	charts = append(charts, m.resourceName+"-chart")

	return &av1.ArmadaChartGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.resourceName + "-acg",
			Namespace: m.namespace,
			Labels:    labels,
		},
		Spec: av1.ArmadaChartGroupSpec{
			Charts:      charts,
			Description: "Created by " + m.resourceName,
			Name:        m.resourceName + "-acg",
			Sequenced:   false,
			TestCharts:  false,
			TargetState: av1.StateInitialized,
		},
	}
}
