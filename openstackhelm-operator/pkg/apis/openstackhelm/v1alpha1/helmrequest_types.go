package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HelmRequestSpec defines the desired state of HelmRequest
type HelmRequestSpec struct {
	// ReleaseName indicates the name of the release
	ReleaseName string `json:"releaseName,omitempty"`
	// Directory for the configuration
	ChartDir string `json:"chartDir,omitempty"`
}

// HelmRequestStatus defines the observed state of HelmRequest
type HelmRequestStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRequest is the Schema for the helmrequests API
// +k8s:openapi-gen=true
type HelmRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmRequestSpec   `json:"spec,omitempty"`
	Status HelmRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRequestList contains a list of HelmRequest
type HelmRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelmRequest{}, &HelmRequestList{})
}
