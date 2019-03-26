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
	lcmif "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"
	"k8s.io/apimachinery/pkg/types"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type basemanager struct {
	kubeClient   client.Client
	renderer     *OwnerRefRenderer
	resourceName string
	namespace    string

	isInstalled       bool
	isUpdateRequired  bool
	deployedPhaseList *av1.PhaseList
}

// ResourceName returns the name of the release.
func (m basemanager) ResourceName() string {
	return m.resourceName
}

func (m basemanager) IsInstalled() bool {
	return m.isInstalled
}

func (m basemanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Attempts to compare the K8s object present with the rendered objects
func (m basemanager) sync(ctx context.Context) (*av1.PhaseList, *av1.PhaseList, error) {
	deployed := av1.NewPhaseList(m.namespace, m.resourceName)

	rendered, err := m.renderer.Render()
	if err != nil {
		return nil, deployed, err
	}

	errs := make([]error, 0)
	for _, renderedResource := range rendered.Items {
		// TODO(jeb): Don't undestand why need to code such a klduge
		existingResource := unstructured.Unstructured{}
		existingResource.SetAPIVersion(renderedResource.GetAPIVersion())
		existingResource.SetKind(renderedResource.GetKind())
		existingResource.SetName(renderedResource.GetName())
		existingResource.SetNamespace(renderedResource.GetNamespace())

		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.GetName(), Namespace: existingResource.GetNamespace()}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				// Don't want to trace is the error is not a NotFound.
				log.Error(err, "Can't not retrieve Resource")
			}
			errs = append(errs, err)
		} else {
			deployed.Items = append(deployed.Items, existingResource)
		}
	}

	return rendered, deployed, nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) installResource(ctx context.Context) (*av1.PhaseList, error) {
	rendered, err := m.renderer.Render()
	if err != nil {
		return m.deployedPhaseList, err
	}

	errs := make([]error, 0)
	for _, toCreate := range rendered.Items {
		err := m.kubeClient.Create(context.TODO(), &toCreate)
		if err != nil {
			log.Error(err, "Can't not Create Resource")
			errs = append(errs, err)
		} else {
			m.deployedPhaseList.Items = append(m.deployedPhaseList.Items, toCreate)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return m.deployedPhaseList, lcmif.ErrNotFound
		} else {
			return m.deployedPhaseList, errs[0]
		}
	}
	return m.deployedPhaseList, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) updateResource(ctx context.Context) (*av1.PhaseList, *av1.PhaseList, error) {
	return m.deployedPhaseList, &av1.PhaseList{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this Phase CR
func (m basemanager) reconcileResource(ctx context.Context) (*av1.PhaseList, error) {
	return m.deployedPhaseList, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) uninstallResource(ctx context.Context) (*av1.PhaseList, error) {
	errs := make([]error, 0)
	for _, toDelete := range m.deployedPhaseList.Items {
		err := m.kubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			log.Error(err, "Can't not Delete Resource")
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return nil, lcmif.ErrNotFound
		} else {
			return nil, errs[0]
		}
	}
	return m.deployedPhaseList, nil
}
