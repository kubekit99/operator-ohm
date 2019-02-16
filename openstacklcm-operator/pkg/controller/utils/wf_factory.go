package utils

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewWorkflowForCR(name string, namespace string) *unstructured.Unstructured {

	labels := map[string]string{
		"app": name,
	}

	// Using a unstructured object.
	u := &unstructured.Unstructured{}
	u.Object = map[string]interface{}{
		"metadata": map[string]interface{}{
			"name":      name + "-pod",
			"namespace": namespace,
			"labels":    labels,
		},
		"spec": map[string]interface{}{
			"containers": []map[string]interface{}{
				{
					"name":    "busybox",
					"image":   "busybox",
					"command": []string{"sleep", "3600"},
				},
			},
		},
	}
	u.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "apps",
		Kind:    "Pod",
		Version: "v1",
	})

	return u
}
