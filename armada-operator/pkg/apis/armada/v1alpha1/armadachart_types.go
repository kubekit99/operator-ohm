package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ArmadaChartSpec defines the desired state of ArmadaChart
type ArmadaChartSpec struct {
	ChartName    string        `json:"chart_name"`
	Namespace    string        `json:"namespace"`
	Release      string        `json:"release"`
	Source       *ArmadaSource `json:"source"`
	Dependencies []string      `json:"dependencies"`

	Install *ArmadaInstall `json:"install,omitempty"`
	Delete  *ArmadaDelete  `json:"delete,omitempty"`
	Upgrade *ArmadaUpgrade `json:"upgrade,omitempty"`
	Values  *ArmadaValues  `json:"values,omitempty"`

	Protected *ArmadaProtected          `json:"protected,omitempty"`
	Test      unstructured.Unstructured `json:"test,omitempty"`
	Timeout   int                       `json:"timeout,omitempty"`
	Wait      *ArmadaWait               `json:"wait,omitempty"`

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
