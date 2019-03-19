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
	"reflect"

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

	// Administrative State of the resource. Is the reconcilation of the CRD by its controller enabled
	AdminState ArmadaAdminState `json:"admin_state"`
	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"target_state"`
}

// ArmadaManifestStatus defines the observed state of ArmadaManifest
type ArmadaManifestStatus struct {
	ArmadaStatus
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaManifest is the Schema for the armadamanifests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadamanifests,shortName=amf
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success"
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

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *ArmadaManifestStatus) SetCondition(cond HelmResourceCondition, tgt HelmResourceState) {

	// Add the condition to the list
	chelper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = chelper.SetCondition(cond)

	// Recompute the state
	s.ComputeActualState(cond, tgt)
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

// Return the list of dependent resources to watch
func (obj *ArmadaManifest) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	for _, chartname := range obj.Spec.ChartGroups {
		u := NewArmadaChartGroupVersionKind(obj.GetNamespace(), chartname)
		res = append(res, *u)
	}
	return res
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

// JEB: Not sure yet if we really will need it
func (obj *ArmadaManifest) Equivalent(other *ArmadaManifest) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Spec.ChartGroups, other.Spec.ChartGroups)
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

// Convert an unstructured.Unstructured into a typed ArmadaManifestList
func ToArmadaManifestList(u *unstructured.Unstructured) *ArmadaManifestList {
	var obj *ArmadaManifestList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaManifestList{}
	}
	return obj
}

// Convert a typed ArmadaManifestList into an unstructured.Unstructured
func (obj *ArmadaManifestList) FromArmadaManifestList() *unstructured.Unstructured {
	u := NewArmadaManifestListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArmadaManifestList) Equivalent(other *ArmadaManifestList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for ArmadaManifestList
func NewArmadaManifestListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaManifestList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// IsDeleted returns true if the manifest has been deleted
func (obj *ArmadaManifest) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the manifest if managed by the reconcilier
func (obj *ArmadaManifest) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the manifest is not managed by the reconcilier
func (obj *ArmadaManifest) IsDisabled() bool {
	return !obj.IsEnabled()
}
