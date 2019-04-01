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
	kubeClient       client.Client
	renderer         *OwnerRefRenderer
	serviceName      string
	serviceNamespace string
	source           av1.FlowSource

	isInstalled           bool
	isUpdateRequired      bool
	deployedLifecycleFlow *av1.LifecycleFlow
}

// ResourceName returns the name of the release.
func (m basemanager) ResourceName() string {
	return m.serviceName
}

func (m basemanager) IsInstalled() bool {
	return m.isInstalled
}

func (m basemanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Render a chart or just a file
func (m basemanager) render(ctx context.Context) (*av1.LifecycleFlow, error) {
	var err error
	var subResourceList *av1.SubResourceList

	if m.source.Type == "tar" {
		subResourceList, err = m.renderer.RenderChart(m.serviceName, m.serviceNamespace, m.source.Location)
	} else {
		subResourceList, err = m.renderer.RenderFile(m.serviceName, m.serviceNamespace, m.source.Location)
	}

	phaseList := av1.NewLifecycleFlow(m.serviceNamespace, m.serviceName)
	if subResourceList != nil {
		for _, item := range subResourceList.Items {
			log.Info(item.GetAPIVersion())
			if item.GetAPIVersion() == "openstacklcm.airshipit.org/v1alpha1" {
				phaseList.Phases[item.GetKind()] = item
			} else {
				phaseList.Main = &item
			}
		}
	}
	return phaseList, err
}

// Attempts to compare the K8s object present with the rendered objects
func (m basemanager) sync(ctx context.Context) (*av1.LifecycleFlow, *av1.LifecycleFlow, error) {
	deployed := av1.NewLifecycleFlow(m.serviceNamespace, m.serviceName)

	rendered, err := m.render(ctx)
	if err != nil {
		return nil, deployed, err
	}

	errs := make([]error, 0)

	if rendered.Main != nil {
		existingResource := unstructured.Unstructured{}
		existingResource.SetAPIVersion(rendered.Main.GetAPIVersion())
		existingResource.SetKind(rendered.Main.GetKind())
		existingResource.SetName(rendered.Main.GetName())
		existingResource.SetNamespace(rendered.Main.GetNamespace())

		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.GetName(), Namespace: existingResource.GetNamespace()}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				// Don't want to trace is the error is not a NotFound.
				log.Error(err, "Can't not retrieve main workflow")
			}
			errs = append(errs, err)
		} else {
			deployed.Main = &existingResource
		}
	}

	for phaseName, renderedResource := range rendered.Phases {
		existingResource := unstructured.Unstructured{}
		existingResource.SetAPIVersion(renderedResource.GetAPIVersion())
		existingResource.SetKind(renderedResource.GetKind())
		existingResource.SetName(renderedResource.GetName())
		existingResource.SetNamespace(renderedResource.GetNamespace())

		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.GetName(), Namespace: existingResource.GetNamespace()}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				// Don't want to trace is the error is not a NotFound.
				log.Error(err, "Can't not retrieve phase")
			}
			errs = append(errs, err)
		} else {
			deployed.Phases[phaseName] = existingResource
		}
	}

	return rendered, deployed, nil
}

// InstallResource creates K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) installResource(ctx context.Context) (*av1.LifecycleFlow, error) {

	rendered, err := m.render(ctx)
	if err != nil {
		return m.deployedLifecycleFlow, err
	}

	errs := make([]error, 0)

	for phaseName, toCreate := range rendered.Phases {
		err := m.kubeClient.Create(context.TODO(), &toCreate)
		if err != nil {
			log.Error(err, "Can't not create phase")
			errs = append(errs, err)
		} else {
			m.deployedLifecycleFlow.Phases[phaseName] = toCreate
		}
	}

	if rendered.Main != nil {
		err := m.kubeClient.Create(context.TODO(), rendered.Main)
		if err != nil {
			log.Error(err, "Can't not create main flow")
			errs = append(errs, err)
		}

		m.deployedLifecycleFlow.Main = rendered.Main
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return m.deployedLifecycleFlow, lcmif.ErrNotFound
		} else {
			return m.deployedLifecycleFlow, errs[0]
		}
	}
	return m.deployedLifecycleFlow, nil
}

// InstallResource updates K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) updateResource(ctx context.Context) (*av1.LifecycleFlow, *av1.LifecycleFlow, error) {
	return m.deployedLifecycleFlow, &av1.LifecycleFlow{}, nil
}

// ReconcileResource creates or patches resources as necessary to match this Phase CR
func (m basemanager) reconcileResource(ctx context.Context) (*av1.LifecycleFlow, error) {
	return m.deployedLifecycleFlow, nil
}

// UninstallResource delete K8s sub resources (Workflow, Job, ....) attached to this Phase CR
func (m basemanager) uninstallResource(ctx context.Context) (*av1.LifecycleFlow, error) {
	errs := make([]error, 0)

	if m.deployedLifecycleFlow.Main != nil {
		err := m.kubeClient.Delete(context.TODO(), m.deployedLifecycleFlow.Main)
		if err != nil {
			log.Error(err, "Can't not delete main flow")
			errs = append(errs, err)
		}
	}

	for _, toDelete := range m.deployedLifecycleFlow.Phases {
		err := m.kubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			log.Error(err, "Can't not delete phase")
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
	return m.deployedLifecycleFlow, nil
}
