package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChartGroupSpec defines the desired state of ChartGroup
type ChartGroupSpec struct {
	// ReleaseName indicates the name of the release
	ReleaseName string `json:"releaseName,omitempty"`
	// Directory for the configuration
	ChartDir string `json:"chartDir,omitempty"`
}

// ChartGroupStatus defines the observed state of ChartGroup
type ChartGroupStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChartGroup is the Schema for the chartgroups API
// +k8s:openapi-gen=true
type ChartGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChartGroupSpec   `json:"spec,omitempty"`
	Status ChartGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChartGroupList contains a list of ChartGroup
type ChartGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ChartGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ChartGroup{}, &ChartGroupList{})
}
