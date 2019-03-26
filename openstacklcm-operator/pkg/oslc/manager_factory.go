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
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
	lcmif "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type managerFactory struct {
	kubeClient client.Client
}

// NewManagerFactory returns a new factory.
func NewManagerFactory(mgr manager.Manager) lcmif.OslcManagerFactory {
	return &managerFactory{kubeClient: mgr.GetClient()}
}

// NewOslcManager returns a new manager capable of controlling Oslc phase of the service lifecyle
func (f managerFactory) NewOslcManager(r *av1.Oslc) lcmif.OslcManager {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}

	return &oslcmanager{
		basemanager: basemanager{
			kubeClient: f.kubeClient,
			renderer:   NewOwnerRefRenderer(ownerRefs),
			namespace:  r.GetNamespace()},

		spec:   r.Spec,
		status: &r.Status,
	}
}
