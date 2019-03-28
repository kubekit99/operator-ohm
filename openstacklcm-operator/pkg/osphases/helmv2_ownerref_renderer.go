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

package osphases

import (
	helmv2 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/helmv2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OwnerRefRenderer struct {
	*helmv2.OwnerRefHelmv2Renderer
}

// NewOwnerRefRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewOwnerRefRenderer(refs []metav1.OwnerReference, suffix string,
	renderFiles []string, renderValues map[string]interface{}) *OwnerRefRenderer {
	return &OwnerRefRenderer{
		OwnerRefHelmv2Renderer: helmv2.NewOwnerRefHelmv2Renderer(refs, suffix, renderFiles, renderValues),
	}
}
