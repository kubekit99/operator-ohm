/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helm

import (
	"bytes"
	"io"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	yaml "gopkg.in/yaml.v2"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
)

// GetDependentResource extracts the list of dependent resources
// from the Helm Manifest in order to add Watch on those components.
func GetDependentResources(release *rpb.Release) []unstructured.Unstructured {

	var res = make([]unstructured.Unstructured, 0)
	dec := yaml.NewDecoder(bytes.NewBufferString(release.GetManifest()))
	for {
		var u unstructured.Unstructured
		err := dec.Decode(&u.Object)
		if err == io.EOF {
			return res
		}
		if err != nil {
			return nil
		}
		res = append(res, u)
	}
}
