package v1alpha1

//JEB: Inspired from ETCD backup and adapt to Openstack/Airship

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// Offsite related consts
	BackupStorageTypeOffsite         BackupStorageType = "Offsite"
	OffsiteSecretCredentialsFileName                   = "credentials"
	OffsiteSecretConfigFileName                        = "config"

	// Ceph related consts
	BackupStorageTypeCeph BackupStorageType = "Ceph"
	CephAccessToken                         = "access-token"
	CephCredentialsJson                     = "credentials.json"
)

type BackupStorageType string

// OpenstackBackupSpec defines the desired state of OpenstackBackup
type OpenstackBackupSpec struct {
	// OpenstackEndpoints specifies the endpoints of an openstack cluster.
	// When multiple endpoints are given, the backup operator retrieves
	// the backup from the endpoint that has the most up-to-date state.
	// The given endpoints must belong to the same openstack cluster.
	OpenstackEndpoints []string `json:"openstackEndpoints,omitempty"`
	// StorageType is the openstack backup storage type.
	// We need this field because CRD doesn't support validation against invalid fields
	// and we cannot verify invalid backup storage source.
	StorageType BackupStorageType `json:"storageType"`
	// BackupPolicy configures the backup process.
	BackupPolicy *BackupPolicy `json:"backupPolicy,omitempty"`
	// BackupSource is the backup storage source.
	BackupSource `json:",inline"`
	// ClientTLSSecret is the secret containing the openstack TLS client certs and
	// must contain the following data items:
	// data:
	//    "openstack-client.crt": <pem-encoded-cert>
	//    "openstack-client.key": <pem-encoded-key>
	//    "openstack-client-ca.crt": <pem-encoded-ca-cert>
	ClientTLSSecret string `json:"clientTLSSecret,omitempty"`
}

// BackupSource contains the supported backup sources.
type BackupSource struct {
	// Offsite defines the Offsite backup source spec.
	Offsite *OffsiteBackupSource `json:"offsite,omitempty"`
	// Ceph defines the Ceph backup source spec.
	Ceph *CephBackupSource `json:"ceph,omitempty"`
}

// OffsiteBackupSource provides the spec how to store backups on Offsite.
type OffsiteBackupSource struct {
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
	Endpoint string `json:"endpoint,omitempty"`

	// ForcePathStyle forces to use path style over the default subdomain style.
	// This is useful when you have an offsite compatible endpoint that doesn't support
	// subdomain buckets.
	ForcePathStyle bool `json:"forcePathStyle"`
}

// CephBackupSource provides the spec how to store backups on Ceph.
type CephBackupSource struct {
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

// BackupPolicy defines backup policy.
type BackupPolicy struct {
	// TimeoutInSecond is the maximal allowed time in second of the entire backup process.
	TimeoutInSecond int64 `json:"timeoutInSecond,omitempty"`
}

// OpenstackBackupStatus defines the observed state of OpenstackBackup
type OpenstackBackupStatus struct {
	// Succeeded indicates if the backup has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"Reason,omitempty"`
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int64 `json:"openstackRevision,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackBackup is the Schema for the openstackbackups API
// +k8s:openapi-gen=true
type OpenstackBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackBackupSpec   `json:"spec,omitempty"`
	Status OpenstackBackupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackBackupList contains a list of OpenstackBackup
type OpenstackBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenstackBackup{}, &OpenstackBackupList{})
}
