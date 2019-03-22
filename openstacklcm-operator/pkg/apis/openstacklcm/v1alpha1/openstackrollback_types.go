package v1alpha1

//JEB: Inspired from ETCD backup and adapt to Openstack/Airship

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenstackRollbackSpec defines the desired state of OpenstackRollback
type OpenstackRollbackSpec struct {
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

// OpenstackRollbackStatus defines the observed state of OpenstackRollback
type OpenstackRollbackStatus struct {
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

// OpenstackRollback is the Schema for the openstackrollbacks API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=openstackrollbacks,shortName=orbck
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type OpenstackRollback struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackRollbackSpec   `json:"spec,omitempty"`
	Status OpenstackRollbackStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackRollbackList contains a list of OpenstackRollback
type OpenstackRollbackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackRollback `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackRollback{}, &OpenstackRollbackList{})
}
