package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenstackDeploymentSpec defines the desired state of OpenstackDeployment
type OpenstackDeploymentSpec struct {
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

// OpenstackDeploymentStatus defines the observed state of OpenstackDeployment
type OpenstackDeploymentStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackDeployment is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=openstackdeployments,shortName=odpl
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type OpenstackDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackDeploymentSpec   `json:"spec,omitempty"`
	Status OpenstackDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackDeploymentList contains a list of OpenstackDeployment
type OpenstackDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackDeployment{}, &OpenstackDeploymentList{})
}
