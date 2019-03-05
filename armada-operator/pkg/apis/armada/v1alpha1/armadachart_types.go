package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

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

	// Target state of the Helm Custom Resources
	TargetState HelmResourceState `json:"target_state"`
}

// ArmadaChartStatus defines the observed state of ArmadaChart
type ArmadaChartStatus struct {
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

// ArmadaChart is the Schema for the armadacharts API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadacharts,shortName=act
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type ArmadaChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaChartSpec   `json:"spec,omitempty"`
	Status ArmadaChartStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaChartList contains a list of ArmadaChart
type ArmadaChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaChart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArmadaChart{}, &ArmadaChartList{})
}

// Synthesis the actual state based on the conditions
func (s *ArmadaChartStatus) ComputeActualState(condition *HelmResourceCondition, targetState HelmResourceState) {
	s.ActualState = targetState
	s.Succeeded = (s.ActualState == targetState)
	s.Reason = ""
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *ArmadaChartStatus) SetCondition(condition HelmResourceCondition) *ArmadaChartStatus {

	helper := HelmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = helper.SetCondition(condition)
	return s
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

// Return the list of dependant resources to watch
func (obj *ArmadaChart) GetDependantResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed ArmadaChart
func ToArmadaChart(u *unstructured.Unstructured) *ArmadaChart {
	var obj *ArmadaChart
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &ArmadaChart{}
	}
	return obj
}

// Convert a typed ArmadaChart into an unstructured.Unstructured
func (obj *ArmadaChart) FromArmadaChart() *unstructured.Unstructured {
	u := NewArmadaChartVersionKind("", "")
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

// Returns a GKV for ArmadaChart
func NewArmadaChartVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("ArmadaChart")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
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
