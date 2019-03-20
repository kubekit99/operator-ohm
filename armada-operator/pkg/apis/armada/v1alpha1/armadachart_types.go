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

// ======= ArmadaChartSpec Definition =======
// ArmadaChartSpec defines the desired state of ArmadaChart
type ArmadaChartSpec struct {
	// name for the chart
	ChartName string `json:"chart_name"`
	// namespace of your chart
	Namespace string `json:"namespace"`
	// name of the release (Armada will prepend with ``release-prefix`` during processing)
	Release string `json:"release"`
	// provide a path to a ``git repo``, ``local dir``, or ``tarball url`` chart
	Source *ArmadaChartSource `json:"source"`
	// reference any chart dependencies before install
	Dependencies []string `json:"dependencies"`

	// JEB: install the chart into your Kubernetes cluster
	// JEB: Install *ArmadaInstall `json:"install,omitempty"`

	// override any default values in the charts
	Values *ArmadaChartValues `json:"values,omitempty"`
	// See Delete_.
	Delete *ArmadaDelete `json:"delete,omitempty"`
	// upgrade the chart managed by the armada yaml
	Upgrade *ArmadaUpgrade `json:"upgrade,omitempty"`

	// do not delete FAILED releases when encountered from previous run (provide the
	// 'continue_processing' bool to continue or halt execution (default: halt))
	Protected *ArmadaProtectedRelease `json:"protected,omitempty"`
	// See Test_.
	Test *ArmadaTest `json:"test,omitempty"`
	// time (in seconds) allotted for chart to deploy when 'wait' flag is set (DEPRECATED)
	Timeout int `json:"timeout,omitempty"`
	// See `ArmwadaWait`.
	Wait *ArmadaWait `json:"wait,omitempty"`

	// Administrative State of the resource. Is the reconcilation of the CRD by its controller enabled
	AdminState ArmadaAdminState `json:"admin_state"`
	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"target_state"`

	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the ArmadaChart's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// ArmadaChartSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
}

// ======= ArmadaChartStatus Definition =======
// ArmadaChartStatus defines the observed state of ArmadaChart
type ArmadaChartStatus struct {
	ArmadaStatus
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *ArmadaChartStatus) SetCondition(cond HelmResourceCondition, tgt HelmResourceState) {

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
func (s *ArmadaChartStatus) RemoveCondition(conditionType HelmResourceConditionType) *ArmadaChartStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.RemoveCondition(conditionType)
	return s
}

// ======= ArmadaChartList Definition =======
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaChart is the Schema for the armadacharts API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadacharts,shortName=act
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success"
type ArmadaChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaChartSpec   `json:"spec,omitempty"`
	Status ArmadaChartStatus `json:"status,omitempty"`
}

// Return the list of dependent resources to watch
func (obj *ArmadaChart) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed ArmadaChart
func ToArmadaChart(u *unstructured.Unstructured) *ArmadaChart {
	var obj *ArmadaChart
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaChart{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed ArmadaChart into an unstructured.Unstructured
func (obj *ArmadaChart) FromArmadaChart() *unstructured.Unstructured {
	u := NewArmadaChartVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArmadaChart) Equivalent(other *ArmadaChart) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Spec, other.Spec)
}

// IsDeleted returns true if the chart has been deleted
func (obj *ArmadaChart) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the chart if managed by the reconcilier
func (obj *ArmadaChart) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the chart is not managed by the reconcilier
func (obj *ArmadaChart) IsDisabled() bool {
	return !obj.IsEnabled()
}

// Returns a GKV for ArmadaChart
func NewArmadaChartVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaChart")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// ======= ArmadaChartList Definition =======
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaChartList contains a list of ArmadaChart
type ArmadaChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaChart `json:"items"`
}

// Convert an unstructured.Unstructured into a typed ArmadaChartList
func ToArmadaChartList(u *unstructured.Unstructured) *ArmadaChartList {
	var obj *ArmadaChartList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaChartList{}
	}
	return obj
}

// Convert a typed ArmadaChartList into an unstructured.Unstructured
func (obj *ArmadaChartList) FromArmadaChartList() *unstructured.Unstructured {
	u := NewArmadaChartListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *ArmadaChartList) Equivalent(other *ArmadaChartList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for ArmadaChartList
func NewArmadaChartListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaChartList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// ======= ArmadaCharts Definition =======
// ArmadaCharts is a wrapper around ArmadaChartList used for interface definitions
type ArmadaCharts struct {
	List *ArmadaChartList
	Name string
}

// Instantiate new ArmadaCharts
func NewArmadaCharts(name string) *ArmadaCharts {
	var emptyList = &ArmadaChartList{
		Items: make([]ArmadaChart, 0),
	}
	var res = ArmadaCharts{
		Name: name,
		List: emptyList,
	}

	return &res
}

// Convert the Name of an ArmadaCharts
func (obj *ArmadaCharts) GetName() string {
	return obj.Name
}

// ======= Schema Registration =======
func init() {
	SchemeBuilder.Register(&ArmadaChart{}, &ArmadaChartList{})
}
