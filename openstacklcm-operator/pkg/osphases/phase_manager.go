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

package osphases

import (
	"context"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
	lcmif "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"
	"k8s.io/apimachinery/pkg/types"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type phasemanager struct {
	kubeClient     client.Client
	renderer       *OwnerRefRenderer
	phaseName      string
	phaseNamespace string
	source         *av1.PhaseSource

	isInstalled             bool
	isUpdateRequired        bool
	deployedSubResourceList *av1.SubResourceList
}

// ResourceName returns the name of the release.
func (m phasemanager) ResourceName() string {
	return m.phaseName
}

func (m phasemanager) IsInstalled() bool {
	return m.isInstalled
}

func (m phasemanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Render a chart or just a file
func (m phasemanager) render(ctx context.Context) (*av1.SubResourceList, error) {
	if m.source.Type == "tar" {
		return m.renderer.RenderChart(m.phaseName, m.phaseNamespace, m.source.Location)
	} else {
		return m.renderer.RenderFile(m.phaseName, m.phaseNamespace, m.source.Location)
	}
}

// Attempts to compare the K8s object present with the rendered objects
func (m phasemanager) sync(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	deployed := av1.NewSubResourceList(m.phaseNamespace, m.phaseName)

	rendered, err := m.render(ctx)
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
func (m phasemanager) installResource(ctx context.Context) (*av1.SubResourceList, error) {

	rendered, err := m.render(ctx)
	if err != nil {
		return m.deployedSubResourceList, err
	}

	errs := make([]error, 0)
	rendered.Items = lcmif.SortByInstallOrder(rendered.Items)
	for _, toCreate := range rendered.Items {
		err := m.kubeClient.Create(context.TODO(), &toCreate)
		if err != nil {
			log.Error(err, "Can't not Create Resource", "kind", toCreate.GetKind(), "name", toCreate.GetName())
			errs = append(errs, err)
		} else {
			log.Info("Created Resource", "kind", toCreate.GetKind(), "name", toCreate.GetName())
			m.deployedSubResourceList.Items = append(m.deployedSubResourceList.Items, toCreate)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return m.deployedSubResourceList, lcmif.ErrNotFound
		} else {
			return m.deployedSubResourceList, errs[0]
		}
	}
	return m.deployedSubResourceList, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m phasemanager) updateResource(ctx context.Context) (*av1.SubResourceList, *av1.SubResourceList, error) {
	return m.deployedSubResourceList, &av1.SubResourceList{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this Phase CR
func (m phasemanager) reconcileResource(ctx context.Context) (*av1.SubResourceList, error) {
	return m.deployedSubResourceList, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m phasemanager) uninstallResource(ctx context.Context) (*av1.SubResourceList, error) {
	errs := make([]error, 0)

	m.deployedSubResourceList.Items = lcmif.SortByUninstallOrder(m.deployedSubResourceList.Items)
	for _, toDelete := range m.deployedSubResourceList.Items {
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
	return m.deployedSubResourceList, nil
}
