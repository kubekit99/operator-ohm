package v1alpha1

//JEB: Inspired from ETCD backup and adapt to Openstack/Airship

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenstackUpgradeSpec defines the desired state of OpenstackUpgrade
type OpenstackUpgradeSpec struct {
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

// OpenstackUpgradeStatus defines the observed state of OpenstackUpgrade
type OpenstackUpgradeStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackUpgrade is the Schema for the openstackupgrades API
// +k8s:openapi-gen=true
type OpenstackUpgrade struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackUpgradeSpec   `json:"spec,omitempty"`
	Status OpenstackUpgradeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackUpgradeList contains a list of OpenstackUpgrade
type OpenstackUpgradeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackUpgrade `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackUpgrade{}, &OpenstackUpgradeList{})
}
