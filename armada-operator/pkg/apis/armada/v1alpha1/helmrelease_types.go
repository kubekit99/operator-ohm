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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// HelmReleaseSpec defines the desired state of HelmRelease
type HelmReleaseSpec struct {
	// Helm Chart releate information
	ReleaseName string `json:"releaseName,omitempty"`
	// Path the chart in the container. Will be converted to a container later
	ChartDir string `json:"chartDir,omitempty"`
	// Set to true to add Watch to Kubernetes Resources created by the chart
	WatchHelmDependentResources bool `json:"watchHelmDependentResources"`
	// ReleaseDesc is the chart that was released.
	ReleaseDesc *HelmReleaseDesc `json:"releaseDesc,omitempty"`
	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"targetState"`
}

// HelmReleaseStatus defines the observed state of HelmRelease
type HelmReleaseStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"reason,omitempty"`
	// Actual state of the Helm Custom Resources
	ActualState HelmResourceState `json:"actual_state"`
	// List of conditions and states related to the resource. JEB: Feature kind of overlap with event recorder
	Conditions []HelmResourceCondition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRelease is the Schema for the armadas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=helmreleases,shortName=hrel
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type HelmRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmReleaseSpec   `json:"spec,omitempty"`
	Status HelmReleaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmReleaseList contains a list of HelmRelease
type HelmReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmRelease `json:"items"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *HelmReleaseStatus) SetCondition(condition HelmResourceCondition) *HelmReleaseStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.SetCondition(condition)
	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *HelmReleaseStatus) RemoveCondition(conditionType HelmResourceConditionType) *HelmReleaseStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.RemoveCondition(conditionType)
	return s
}

// Returns a GKV for HelmRelease
func NewHelmReleaseVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("HelmRelease")
	return u
}

func init() {
	SchemeBuilder.Register(&HelmRelease{}, &HelmReleaseList{})
}
