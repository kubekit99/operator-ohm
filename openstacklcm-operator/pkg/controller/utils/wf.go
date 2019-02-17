package utils

import (
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
	jsonfmt := []byte(`{
	"apiVersion": "argoproj.io/v1alpha1",
	"kind": "Workflow",
	"metadata": {
	   "name": "openstackbackup-wf"
	},
	"spec": {
	   "entrypoint": "whalesay",
	   "templates": [
		  {
			 "name": "whalesay",
			 "container": {
				"image": "docker/whalesay:latest",
				"command": [
				   "cowsay"
				],
				"args": [
				   "openstackbackup"
				]
			 }
		  }
	   ]
	}
 }`)
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
