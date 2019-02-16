package utils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NewWorkflowGroupVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("v1")
	u.SetKind("Pod")
	return u

}

func NewWorkflowForCR(name string, namespace string) *unstructured.Unstructured {

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
