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

package utils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewPodGroupVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("v1")
	u.SetKind("Pod")
	return u

}

func NewPodForCR(name string, namespace string) *unstructured.Unstructured {

	reqLogger := log.WithValues("Namespace", namespace, "Name", name)

	labels := map[string]string{
		"app": name,
	}

	// Using a unstructured object.
	jsonfmt := []byte(`{ "apiVersion": "v1", "kind": "Pod", "spec": { "containers": [ { "command": [ "sleep", "3600" ], "image": "busybox", "name": "busybox" } ] } }`)
	u := &unstructured.Unstructured{}
	err := u.UnmarshalJSON(jsonfmt)
	if err != nil {
		reqLogger.Error(err, "something is wrong during scanning of json object", jsonfmt)
	}
	u.SetName(name + "-pod")
	u.SetNamespace(namespace)
	u.SetLabels(labels)

	return u
}
