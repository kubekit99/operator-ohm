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

// JEB: This file has been at first generated using the kubebuilder
// and following the proposal described in:
// https://github.com/thomastaylor312/helm-3-crd/
// This file will be deleted once we figure out what we really want
// to put in our CRDs.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LifecycleSpec defines the desired state of Lifecycle
type LifecycleSpec struct {
	// The eventKey will be the value of the `helm-lifecycle-event` label for the
	// LifecycleEvent. The controller that handles this should filter events based
	// on this key
	EventKey string `json:"eventKey,omitempty"`
}

// LifecycleStatus defines the observed state of Lifecycle
type LifecycleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make generate" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Lifecycle is the Schema for the lifecycles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=lifecyles,shortName=life
type Lifecycle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LifecycleSpec   `json:"spec,omitempty"`
	Status LifecycleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LifecycleList contains a list of Lifecycle
type LifecycleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Lifecycle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Lifecycle{}, &LifecycleList{})
}
