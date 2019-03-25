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
	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
	lcmif "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"
)

type managerFactory struct {
}

// NewManagerFactory returns a new factory.
func NewManagerFactory(mgr manager.Manager) lcmif.OslcManagerFactory {
	return &managerFactory{}
}

// NewOslcManager returns a new OslcManagerr factory capable of managing an Openstack Service Lifecyle
func (f managerFactory) NewOslcManager(r *av1.Oslc) lcmif.OslcManager {
	return &oslcmanager{
		renderer:  nil,
		namespace: r.GetNamespace(),

		spec:   r.Spec,
		status: &r.Status,
	}
}
