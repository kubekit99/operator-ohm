// Copyright 2018 The Operator-SDK Authors
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
	"github.com/ghodss/yaml"
	"io/ioutil"

	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// OwnerRefRenderer loads yaml file and adds ownerrefs to rendered assets
type OwnerRefRenderer struct {
	refs []metav1.OwnerReference
}

// Adds the ownerrefs to all the documents in a YAML file
func (o *OwnerRefRenderer) Render() (*av1.SubResourceList, error) {

	name := o.refs[0].Name
	namespace := "default" // o.refs[0].NameSpace
	fileName := "/root/argo-workflows/wf-" + name + ".yaml"

	ownedRenderedFiles := av1.NewSubResourceList(namespace, name)

	unstructured, err := o.fromYaml(fileName)
	if err != nil {
		log.Error(err, "Can not convert from yaml to unstructured", "filename", fileName)
		return ownedRenderedFiles, err
	}

	unstructured.SetName(name + "-wf")
	unstructured.SetNamespace(namespace)
	unstructured.SetOwnerReferences(o.refs)
	labels := map[string]string{
		"app": name,
	}

	unstructured.SetLabels(labels)

	ownedRenderedFiles.Items = append(ownedRenderedFiles.Items, unstructured)

	return ownedRenderedFiles, nil
}

// Reads a yaml file and converts into an Unstructured object
func (o *OwnerRefRenderer) fromYaml(filename string) (unstructured.Unstructured, error) {

	reqLogger := log.WithValues("filename", filename)

	yamlfmt, ferr := ioutil.ReadFile(filename)
	if ferr != nil {
		reqLogger.Error(ferr, "Can not read file")
		return unstructured.Unstructured{}, ferr
	}
	jsonfmt, ferr2 := yaml.YAMLToJSON(yamlfmt)
	if ferr2 != nil {
		reqLogger.Error(ferr2, "Can not convert from yaml to json")
		return unstructured.Unstructured{}, ferr2
	}

	// TODO(jeb): This following code is probably better than my JSON conversion
	// unst, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&manifestMap)
	// if err != nil {
	// 	return nil, err
	// }
	// unstructured := &unstructured.Unstructured{Object: unst}

	u := unstructured.Unstructured{}
	err := u.UnmarshalJSON(jsonfmt)
	if err != nil {
		reqLogger.Error(err, "something is wrong during scanning of json object")
		return unstructured.Unstructured{}, err
	}

	return u, nil
}

// NewOwnerRefRenderer creates a new OwnerRef engine with a set of metav1.OwnerReferences to be added to assets
func NewOwnerRefRenderer(refs []metav1.OwnerReference) *OwnerRefRenderer {
	return &OwnerRefRenderer{
		refs: refs,
	}
}
