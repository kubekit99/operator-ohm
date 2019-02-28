package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ArmadaApplyRequest struct {
	// ArmadaManifest CRD Name
	ManifestName string `json:"manifest,omitempty"`
}

type ArmadaRollbackRequest struct {
	// ArmadaRelease CRD Name
	ReleaseName string `json:"release,omitempty"`
}

type ArmadaTestRequest struct {
	// ArmadaRelease CRD Name
	ReleaseName string `json:"release,omitempty"`
}

type ArmadaValidateRequest struct {
	// ArmadaManifest CRD Name
	ManifestName string `json:"manifest,omitempty"`
}

// type ArmadaTillerRequest struct {
// }

// ArmadaRequestSpec defines the parameters of an ArmadaRequest
type ArmadaRequestSpec struct {
	// Apply the ArmadaManifest specified int the request
	Apply *ArmadaApplyRequest `json:"apply,omitempty"`
	// Rollback the ArmadaRelease specified in the request
	Rollback *ArmadaRollbackRequest `json:"rollback,omitempty"`
	// Test the ArmadaRelease specified in the request
	Test *ArmadaTestRequest `json:"test,omitempty"`
	// Validate the ArmadaManifest specified int the request
	Validate *ArmadaValidateRequest `json:"validate,omitempty"`
	// Invokes Tiller
	// Tiller   *ArmadaTillerRequest   `json:"tiller,omitempty"`
}

// ArmadaRequestStatus defines the current status of the ArmadaRequest
type ArmadaRequestStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaRequest is the Schema for the armadarequests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadarequests,shortName=areq
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type ArmadaRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaRequestSpec   `json:"spec,omitempty"`
	Status ArmadaRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaRequestList contains a list of ArmadaRequest
type ArmadaRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArmadaRequest{}, &ArmadaRequestList{})
}
