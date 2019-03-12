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

type LifecycleEventConfig struct {
	PodName string `json:"podName,omitempty"`
}

// LifecycleEventSpec defines the desired state of LifecycleEvent
type LifecycleEventSpec struct {
	Timeout int64                `json:"timeout,omitempty"`
	Config  LifecycleEventConfig `json:"config,omitempty"`
}

// LifecycleEventStatus defines the observed state of LifecycleEvent
type LifecycleEventStatus struct {
	// Will be set to true if there was an error with `message` containing more info
	Error bool `json:"error"`

	// Will be marked as "Running" if in progress, "Pending" if it hasn't been handled
	// yet, and "Error" if it is in an error state
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LifecycleEvent is the Schema for the lifecycleevents API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=lifecyleevents,shortName=le
type LifecycleEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LifecycleEventSpec   `json:"spec,omitempty"`
	Status LifecycleEventStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LifecycleEventList contains a list of LifecycleEvent
type LifecycleEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LifecycleEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LifecycleEvent{}, &LifecycleEventList{})
}
