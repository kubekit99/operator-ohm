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

// ======= ArmadaChartGroupSpec Definition =======
// ArmadaChartGroupSpec defines the desired state of ArmadaChartGroup
type ArmadaChartGroupSpec struct {
	// reference to chart document
	Charts []string `json:"chart_group"`
	// description of chart set
	Description string `json:"description,omitempty"`
	// Name of the chartgroup
	Name string `json:"name,omitempty"`
	// enables sequenced chart deployment in a group
	Sequenced bool `json:"sequenced,omitempty"`
	// run pre-defined helm tests in a ChartGroup (DEPRECATED)
	TestCharts bool `json:"test_charts,omitempty"`

	// Administrative State of the resource. Is the reconcilation of the CRD by its controller enabled
	AdminState ArmadaAdminState `json:"admin_state"`
	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"target_state"`
	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the ArmadaChartGroup's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// ArmadaChartGroupSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
}

// ======= ArmadaChartGroupStatus Definition =======
// ArmadaChartGroupStatus defines the observed state of ArmadaChartGroup
type ArmadaChartGroupStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"reason,omitempty"`
	// Actual state of the Helm Custom Resources
	ActualState HelmResourceState `json:"actual_state"`
	// List of conditions and states related to the resource. JEB: Feature kind of overlap with event recorder
	Conditions []HelmResourceCondition `json:"conditions,omitempty"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *ArmadaChartGroupStatus) SetCondition(cond HelmResourceCondition, tgt HelmResourceState) *ArmadaChartGroupStatus {

	// Add the condition to the list
	chelper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = chelper.SetCondition(cond)

	// Recompute the state
	shelper := HelmResourceStatusHelper{
		Cond:             &cond,
		TargetState:      tgt,
		CurrentState:     s.ActualState,
		CurrentSucceeded: s.Succeeded,
		CurrentReason:    s.Reason}
	s.ActualState, s.Succeeded, s.Reason = shelper.ComputeActualState()
	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *ArmadaChartGroupStatus) RemoveCondition(conditionType HelmResourceConditionType) *ArmadaChartGroupStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.RemoveCondition(conditionType)
	return s
}

// ======= ArmadaChartGroup Definition =======
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaChartGroup is the Schema for the armadachartgroups API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadachartgroups,shortName=acg
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success"
type ArmadaChartGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaChartGroupSpec   `json:"spec,omitempty"`
	Status ArmadaChartGroupStatus `json:"status,omitempty"`
}

// Return the list of dependent resources to watch
func (obj *ArmadaChartGroup) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	for _, chartname := range obj.Spec.Charts {
		u := NewArmadaChartVersionKind(obj.GetNamespace(), chartname)
		res = append(res, *u)
	}
	return res
}

// Convert an unstructured.Unstructured into a typed ArmadaChartGroup
func ToArmadaChartGroup(u *unstructured.Unstructured) *ArmadaChartGroup {
	var obj *ArmadaChartGroup
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaChartGroup{}
	}
	return obj
}

// Convert a typed ArmadaChartGroup into an unstructured.Unstructured
func (obj *ArmadaChartGroup) FromArmadaChartGroup() *unstructured.Unstructured {
	u := NewArmadaChartGroupVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArmadaChartGroup) Equivalent(other *ArmadaChartGroup) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Spec.Charts, other.Spec.Charts)
}

// IsDeleted returns true if the chart group has been deleted
func (obj *ArmadaChartGroup) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the chart group if managed by the reconcilier
func (obj *ArmadaChartGroup) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the chart group is not managed by the reconcilier
func (obj *ArmadaChartGroup) IsDisabled() bool {
	return !obj.IsEnabled()
}

// Returns a GKV for ArmadaChartGroup
func NewArmadaChartGroupVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaChartGroup")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// ======= ArmadaChartGroupList Definition =======
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaChartGroupList contains a list of ArmadaChartGroup
type ArmadaChartGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaChartGroup `json:"items"`
}

// Convert an unstructured.Unstructured into a typed ArmadaChartGroupList
func ToArmadaChartGroupList(u *unstructured.Unstructured) *ArmadaChartGroupList {
	var obj *ArmadaChartGroupList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaChartGroupList{}
	}
	return obj
}

// Convert a typed ArmadaChartGroupList into an unstructured.Unstructured
func (obj *ArmadaChartGroupList) FromArmadaChartGroupList() *unstructured.Unstructured {
	u := NewArmadaChartGroupListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArmadaChartGroupList) Equivalent(other *ArmadaChartGroupList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for ArmadaChartGroupList
func NewArmadaChartGroupListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaChartGroupList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// ======= ArmadaChartGroups Definition =======
// ArmadaChartGroups is a wrapper around ArmadaChartGroupList used for interface definitions
type ArmadaChartGroups struct {
	Name string
	List *ArmadaChartGroupList
}

// Instantiate new ArmadaChartGroups
func NewArmadaChartGroups(name string) *ArmadaChartGroups {
	var emptyList = &ArmadaChartGroupList{
		Items: make([]ArmadaChartGroup, 0),
	}
	var res = ArmadaChartGroups{
		Name: name,
		List: emptyList,
	}

	return &res
}

// Convert the Name of an ArmadaChartGroupList
func (obj *ArmadaChartGroups) GetName() string {
	return obj.Name
}

// ======= Schema Registration =======
func init() {
	SchemeBuilder.Register(&ArmadaChartGroup{}, &ArmadaChartGroupList{})
}
