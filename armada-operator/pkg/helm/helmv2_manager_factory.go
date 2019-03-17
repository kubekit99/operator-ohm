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

// +build v2

package helm

import (
	"sigs.k8s.io/controller-runtime/pkg/manager"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"
	helmv2 "github.com/kubekit99/operator-ohm/armada-operator/pkg/helmv2"
)

// NewManagerFactory returns a new Helm manager factory capable of installing and uninstalling releases.
func NewManagerFactory(mgr manager.Manager) helmif.HelmManagerFactory {
        return helmv2.NewManagerFactory(mgr)
}
