package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ChartSpec defines the desired state of Chart
type ChartSpec struct {
	// ReleaseName indicates the name of the release
	ReleaseName string `json:"releaseName,omitempty"`
	// Directory for the configuration
	ChartDir string `json:"chartDir,omitempty"`
}

// ChartStatus defines the observed state of Chart
type ChartStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Chart is the Schema for the charts API
// +k8s:openapi-gen=true
type Chart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ChartSpec   `json:"spec,omitempty"`
	Status ChartStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ChartList contains a list of Chart
type ChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Chart `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Chart{}, &ChartList{})
}
