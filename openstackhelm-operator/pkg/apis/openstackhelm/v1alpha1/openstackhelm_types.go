package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenstackHelmSpec defines the desired state of OpenstackHelm
type OpenstackHelmSpec struct {
	// ReleaseName indicates the name of the release
	ReleaseName string `json:"releaseName,omitempty"`
	// Directory for the configuration
	ChartDir string `json:"chartDir,omitempty"`
}

// OpenstackHelmStatus defines the observed state of OpenstackHelm
type OpenstackHelmStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackHelm is the Schema for the openstackhelms API
// +k8s:openapi-gen=true
type OpenstackHelm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackHelmSpec   `json:"spec,omitempty"`
	Status OpenstackHelmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackHelmList contains a list of OpenstackHelm
type OpenstackHelmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackHelm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackHelm{}, &OpenstackHelmList{})
}
