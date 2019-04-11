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

package services

import (
	"bytes"
	"io"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	yaml "gopkg.in/yaml.v2"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
)

type HelmRelease struct {
	*rpb.Release
	cached []unstructured.Unstructured
}

func (r *HelmRelease) GetNotes() string {
	return r.GetInfo().GetStatus().GetNotes()
}

// GetDependentResource extracts the list of dependent resources
// from the Helm Manifest in order to add Watch on those components.
func (release *HelmRelease) GetDependentResources() []unstructured.Unstructured {

	if len(release.cached) != 0 {
		return release.cached
	}

	release.cached = make([]unstructured.Unstructured, 0)
	dec := yaml.NewDecoder(bytes.NewBufferString(release.GetManifest()))
	for {
		var u unstructured.Unstructured
		err := dec.Decode(&u.Object)
		if err == io.EOF {
			return release.cached
		}
		if err != nil {
			return nil
		}
		release.cached = append(release.cached, u)
	}
}

// Let's check the reference are setup properly.
func (release *HelmRelease) CheckOwnerReference(refs []metav1.OwnerReference) bool {

	// Check that each sub resource is owned by the phase
	items := release.GetDependentResources()
	for _, item := range items {
		if !reflect.DeepEqual(item.GetOwnerReferences(), refs) {
			return false
		}
	}

	return true
}

// Check the state of a service
func (release *HelmRelease) IsReady() bool {

	dep := &KubernetesDependency{}

	// Check that each sub resource is owned by the phase
	items := release.GetDependentResources()
	for _, item := range items {
		if !dep.IsUnstructuredReady(&item) {
			return false
		}
	}

	return true
}
