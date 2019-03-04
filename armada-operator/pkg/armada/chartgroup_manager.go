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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type chartgroupmanager struct {
	resourceName     string
	namespace        string
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
	return nil
}

func (m chartgroupmanager) InstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	labels := map[string]string{
		"app": m.resourceName,
	}
	_ = &corev1.Pod{
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

	res := av1.NewArmadaChartGroupVersionKind()
	res.SetName(m.resourceName + "-ArmadaChartGroup")
	res.SetNamespace(m.namespace)
	return res, nil
}

// UpdateResource performs a Helm release update.
func (m chartgroupmanager) UpdateResource(ctx context.Context) (*unstructured.Unstructured, *unstructured.Unstructured, error) {
	oldValue := av1.NewArmadaChartGroupVersionKind()
	oldValue.SetName(m.resourceName + "-ArmadaChartGroup")
	oldValue.SetNamespace(m.namespace)
	newValue := av1.NewArmadaChartGroupVersionKind()
	newValue.SetName(m.resourceName + "-ArmadaChartGroup")
	newValue.SetNamespace(m.namespace)
	return oldValue, newValue, nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartgroupmanager) ReconcileResource(ctx context.Context) (*unstructured.Unstructured, error) {
	res := av1.NewArmadaChartGroupVersionKind()
	res.SetName(m.resourceName + "-ArmadaChartGroup")
	res.SetNamespace(m.namespace)
	return res, nil
}

// UninstallResource performs a Helm release uninstall.
func (m chartgroupmanager) UninstallResource(ctx context.Context) (*unstructured.Unstructured, error) {
	res := av1.NewArmadaChartGroupVersionKind()
	res.SetName(m.resourceName + "-ArmadaChartGroup")
	res.SetNamespace(m.namespace)
	return res, nil
}
