package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type HelmInstallRequest struct {
	// HelmRelease CRD Name
	ReleaseName string `json:"manifest,omitempty"`
}

type HelmUpgradeRequest struct {
	// HelmRelease CRD Name
	ReleaseName string `json:"release,omitempty"`
}

type HelmRollbackRequest struct {
	// HelmRelease CRD Name
	ReleaseName string `json:"release,omitempty"`
}

type HelmDeleteRequest struct {
	// HelmRelease CRD Name
	ReleaseName string `json:"release,omitempty"`
}

// HelmRequestSpec defines the parameters of an HelmRequest
type HelmRequestSpec struct {
	// Install the HelmManifest specified int the request
	Install *HelmInstallRequest `json:"apply,omitempty"`
	// Upgrade the HelmRelease specified in the request
	Upgrade *HelmUpgradeRequest `json:"upgrade,omitempty"`
	// Rollback the HelmRelease specified in the request
	Rollback *HelmRollbackRequest `json:"rollback,omitempty"`
	// Delete the HelmRelease specified in the request
	Delete *HelmRollbackRequest `json:"delete,omitempty"`
}

// HelmRequestStatus defines the current status of the HelmRequest
type HelmRequestStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRequest is the Schema for the helmrequests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=helmrequests,shortName=hreq
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
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
