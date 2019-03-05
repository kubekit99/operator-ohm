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
	"k8s.io/apimachinery/pkg/runtime"
)

// ArmadaManifestSpec defines the desired state of ArmadaManifest
type ArmadaManifestSpec struct {

	// References ChartGroup document of all groups
	ChartGroups []string `json:"chart_groups"`
	// Appends to the front of all charts released by the manifest in order to manage releases throughout their lifecycle
	ReleasePrefix string `json:"release_prefix"`

	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"target_state"`
}

// ArmadaManifestStatus defines the observed state of ArmadaManifest
type ArmadaManifestStatus struct {
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

// ArmadaManifest is the Schema for the armadamanifests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadamanifests,shortName=amf
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type ArmadaManifest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaManifestSpec   `json:"spec,omitempty"`
	Status ArmadaManifestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaManifestList contains a list of ArmadaManifest
type ArmadaManifestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaManifest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArmadaManifest{}, &ArmadaManifestList{})
}

// Synthesis the actual state based on the conditions
func (s *ArmadaManifestStatus) ComputeActualState(condition *HelmResourceCondition, targetState HelmResourceState) {
	s.ActualState = targetState
	s.Succeeded = (s.ActualState == targetState)
	s.Reason = ""
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *ArmadaManifestStatus) SetCondition(condition HelmResourceCondition) *ArmadaManifestStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.SetCondition(condition)
	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *ArmadaManifestStatus) RemoveCondition(conditionType HelmResourceConditionType) *ArmadaManifestStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.RemoveCondition(conditionType)
	return s
}

// Convert an unstructured.Unstructured into a typed ArmadaManifest
func ToArmadaManifest(u *unstructured.Unstructured) *ArmadaManifest {
	var obj *ArmadaManifest
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaManifest{}
	}
	return obj
}

// Convert a typed ArmadaManifest into an unstructured.Unstructured
func (obj *ArmadaManifest) FromArmadaManifest() *unstructured.Unstructured {
	u := NewArmadaManifestVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// Return the list of dependant resources to watch
func (obj *ArmadaManifest) GetDependantResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	for _, chartname := range obj.Spec.ChartGroups {
		u := NewArmadaChartGroupVersionKind(obj.GetNamespace(), chartname)
		res = append(res, *u)
	}
	return res
}

// Returns a GKV for ArmadaManifest
func NewArmadaManifestVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaManifest")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}
