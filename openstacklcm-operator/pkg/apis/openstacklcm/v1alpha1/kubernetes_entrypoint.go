// Copyright 2019 The OpenstackLcm Authors
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

package v1alpha1

import (
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type KubernetesDependency struct {
}

// Check the state of the Main workflow to figure out
// if the phase is still running
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsWorkflowReady(u *unstructured.Unstructured) bool {
	return obj.IsCustomResourceReady("status.phase", "Succeeded", u)
}

// Check the state of a custom resource
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsCustomResourceReady(key string, expectedValue string, u *unstructured.Unstructured) bool {
	if u == nil {
		return false
	}

	customResource := u.UnstructuredContent()

	for name, _ := range customResource {
		log.Info("HELLO", "name", name)
	}

	for i := strings.Index(key, "."); i != -1; i = strings.Index(key, ".") {
		first := key[:i]
		key = key[i+1:]
		if customResource[first] != nil {
			customResource = customResource[first].(map[string]interface{})
			for name, _ := range customResource {
				log.Info("HELLO", "name", name)
			}
		} else {
			return true
		}
	}

	if customResource != nil {
		value := customResource[key].(string)
		log.Info("HELLO", "value", value)
		return value == expectedValue
	} else {
		return true
	}

}

// Check the state of a service
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsServiceReady(u *unstructured.Unstructured) bool {
	if u == nil {
		return false
	}

	endpoints := corev1.Endpoints{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &endpoints)
	if err1 != nil {
		return false
	}

	for _, subset := range endpoints.Subsets {
		if len(subset.Addresses) > 0 {
			return true
		}
	}
	return false
}

// Check the state of a container
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsContainerReady(containerName string, u *unstructured.Unstructured) bool {
	if u == nil {
		return false
	}

	pod := corev1.Pod{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &pod)
	if err1 != nil {
		return false
	}

	containers := pod.Status.ContainerStatuses
	for _, container := range containers {
		if container.Name == containerName && container.Ready {
			return true
		}
	}
	return false
}

// Check the state of a job
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsJobReady(u *unstructured.Unstructured) bool {
	if u == nil {
		return false
	}

	job := batchv1.Job{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &job)
	if err1 != nil {
		return false
	}

	if job.Status.Succeeded == 0 {
		return false
	}
	return true
}

// Check the state of a pod
// This code is inspired from the kubernetes-entrypoint project
func (obj *KubernetesDependency) IsPodReady(u *unstructured.Unstructured) bool {
	if u == nil {
		return false
	}

	pod := corev1.Pod{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &pod)
	if err1 != nil {
		return false
	}

	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == "True" {
			return true
		}
	}
	return false
}
