/*
Copyright 2019 The Openstack-Service-Lifecyle Authors.

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
	"github.com/ghodss/yaml"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewWorkflowGroupVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("argoproj.io/v1alpha1")
	u.SetKind("Workflow")
	return u

}

func NewWorkflowForCR(name string, namespace string) *unstructured.Unstructured {

	reqLogger := log.WithValues("Namespace", namespace, "Name", name)

	labels := map[string]string{
		"app": name,
	}

	// Using a unstructured object.
	// 	jsonfmtcst := []byte(`{
	// 	"apiVersion": "argoproj.io/v1alpha1",
	// 	"kind": "Workflow",
	// 	"metadata": {
	// 	   "name": "openstackbackup-wf"
	// 	},
	// 	"spec": {
	// 	   "entrypoint": "whalesay",
	// 	   "templates": [
	// 		  {
	// 			 "name": "whalesay",
	// 			 "container": {
	// 				"image": "docker/whalesay:latest",
	// 				"command": [
	// 				   "cowsay"
	// 				],
	// 				"args": [
	// 				   "openstackbackup"
	// 				]
	// 			 }
	// 		  }
	// 	   ]
	// 	}
	//  }`)

	filename := "/root/argo-workflows/wf-" + name + ".yaml"
	yamlfmt, ferr := ioutil.ReadFile(filename)
	if ferr != nil {
		reqLogger.Error(ferr, "Can not read file", filename)
		return nil
	}
	jsonfmt, ferr2 := yaml.YAMLToJSON(yamlfmt)
	if ferr2 != nil {
		reqLogger.Error(ferr2, "Can not convert from yaml to json")
		return nil
	}

	u := &unstructured.Unstructured{}
	err := u.UnmarshalJSON(jsonfmt)
	if err != nil {
		reqLogger.Error(err, "something is wrong during scanning of json object", jsonfmt)
	}
	u.SetName(name + "-wf")
	u.SetNamespace(namespace)
	u.SetLabels(labels)

	return u
}
