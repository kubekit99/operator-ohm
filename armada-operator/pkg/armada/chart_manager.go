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
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type chartmanager struct {
	kubeClient       client.Client
	resourceName     string
	namespace        string
	spec             *av1.ArmadaChartSpec
	status           *av1.ArmadaChartStatus
	deployedResource *corev1.Pod
	isInstalled      bool
	isUpdateRequired bool
}

// ResourceName returns the name of the release.
func (m chartmanager) ResourceName() string {
	return m.resourceName
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
	if targetResource.Spec.Containers[0].Image != m.deployedResource.Spec.Containers[0].Image {
		m.isUpdateRequired = true
	} else {
		m.isUpdateRequired = false
	}

	return nil
}

func (m chartmanager) InstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	newResource := m.newResourceForCR()
	err := m.kubeClient.Create(context.TODO(), newResource)
	if err != nil {
		log.Error(err, "Can't not Create Resource")
		return nil, err
	}
	return FromPod(newResource), nil
}

// UpdateResource performs a Helm release update.
func (m chartmanager) UpdateResource(ctx context.Context) (*unstructured.Unstructured, *unstructured.Unstructured, error) {
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
	return FromPod(m.deployedResource), FromPod(toUpdate), nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartmanager) ReconcileResource(ctx context.Context) (*unstructured.Unstructured, error) {
	toReconcile := m.newResourceForCR()
	return FromPod(toReconcile), nil
}

// UninstallResource performs a Helm release uninstall.
func (m chartmanager) UninstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
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
	return FromPod(toDelete), nil
}

// newResourceForCR returns a busybox pod with the same name/namespace as the cr
func (m chartmanager) newResourceForCR() *corev1.Pod {
	labels := map[string]string{
		"app": m.resourceName,
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      m.resourceName + "-pod",
			Namespace: m.namespace,
			Labels:    labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
