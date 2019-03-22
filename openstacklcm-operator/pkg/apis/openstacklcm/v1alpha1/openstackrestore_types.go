package v1alpha1

//JEB: Inspired from ETCD backup and adapt to Openstack/Airship

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OpenstackRestoreSpec defines the desired state of OpenstackRestore
type OpenstackRestoreSpec struct {
	// BackupStorageType is the type of the backup storage which is used as RestoreSource.
	BackupStorageType BackupStorageType `json:"backupStorageType"`
	// RestoreSource tells the where to get the backup and restore from.
	RestoreSource `json:",inline"`
}

// OpenstackRestoreStatus defines the observed state of OpenstackRestore
type OpenstackRestoreStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

type RestoreSource struct {
	// Offsite tells where on Offsite the backup is saved and how to fetch the backup.
	Offsite *OffsiteRestoreSource `json:"offsite,omitempty"`

	// Ceph tells where on Ceph the backup is saved and how to fetch the backup.
	Ceph *CephRestoreSource `json:"ceph,omitempty"`
}

type OffsiteRestoreSource struct {
	// Path is the full offsite path where the backup is saved.
	// The format of the path must be: "<offsite-bucket-name>/<path-to-backup-file>"
	// e.g: "mybucket/openstack.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Offsite credential and config files.
	// The file name of the credential MUST be 'credentials'.
	// The file name of the config MUST be 'config'.
	// The profile to use in both files will be 'default'.
	//
	// OffsiteSecret overwrites the default openstack operator wide Offsite credential and config.
	OffsiteSecret string `json:"offsiteSecret"`

	// Endpoint if blank points to offsite. If specified, can point to offsite compatible object
	// stores.
	Endpoint string `json:"endpoint"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an offsite compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

type CephRestoreSource struct {
	// Path is the full Ceph path where the backup is saved.
	// The format of the path must be: "<ceph-bucket-name>/<path-to-backup-file>"
	// e.g: "mycephbucket/openstack.backup"
	Path string `json:"path"`

	// The name of the secret object that stores the Google storage credential
	// containing at most ONE of the following:
	// An access token with file name of 'access-token'.
	// JSON credentials with file name of 'credentials.json'.
	//
	// If omitted, client will use the default application credentials.
	CephSecret string `json:"cephSecret,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackRestore is the Schema for the openstackrestores API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=openstackrestores,shortName=orst
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type OpenstackRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackRestoreSpec   `json:"spec,omitempty"`
	Status OpenstackRestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackRestoreList contains a list of OpenstackRestore
type OpenstackRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackRestore{}, &OpenstackRestoreList{})
}
