// Copyright 2019 The Armada Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// JEB: This file has been created by from the ontent of
// https://github.com/helm/helm/blob/dev-v3/pkg/hapi/tiller.go
// This file will be deleted once we figure out what we really want
// to put in our CRDs.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// SortBy defines sort operations.
type HelmSortBy string

const (
	// SortByName requests releases sorted by name.
	SortByName HelmSortBy = "name"
	// SortByLastReleased requests releases sorted by last released.
	SortByLastReleased HelmSortBy = "last-released"
)

// SortOrder defines sort orders to augment sorting operations.
type HelmSortOrder string

const (
	//SortAsc defines ascending sorting.
	SortAsc HelmSortOrder = "ascending"
	//SortDesc defines descending sorting.
	SortDesc HelmSortOrder = "descending"
)

// ListReleasesRequest requests a list of releases.
//
// Releases can be retrieved in chunks by setting limit and offset.
//
// Releases can be sorted according to a few pre-determined sort stategies.
type HelmListReleasesRequest struct {
	// Limit is the maximum number of releases to be returned.
	Limit int64 `json:"limit,omitempty"`
	// Offset is the last release name that was seen. The next listing
	// operation will start with the name after this one.
	// Example: If list one returns albert, bernie, carl, and sets 'next: dennis'.
	// dennis is the offset. Supplying 'dennis' for the next request should
	// cause the next batch to return a set of results starting with 'dennis'.
	Offset string `json:"offset,omitempty"`
	// SortBy is the sort field that the ListReleases server should sort data before returning.
	SortBy HelmSortBy `json:"sort_by,omitempty"`
	// Filter is a regular expression used to filter which releases should be listed.
	//
	// Anything that matches the regexp will be included in the results.
	Filter string `json:"filter,omitempty"`
	// SortOrder is the ordering directive used for sorting.
	SortOrder   HelmSortOrder `json:"sort_order,omitempty"`
	StatusCodes []HelmStatus  `json:"status_codes,omitempty"`
}

// GetReleaseStatusRequest is a request to get the status of a release.
type HelmGetReleaseStatusRequest struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`
	// Version is the version of the release
	Version int `json:"version,omitempty"`
}

// GetReleaseStatusResponse is the response indicating the status of the named release.
type HelmGetReleaseStatusResponse struct {
	// Name is the name of the release.
	Name string `json:"name,omitempty"`
	// Info contains information about the release.
	Info *HelmInfo `json:"info,omitempty"`
	// Namespace the release was released into
	Namespace string `json:"namespace,omitempty"`
}

// GetReleaseContentRequest is a request to get the contents of a release.
type HelmGetReleaseContentRequest struct {
	// The name of the release
	Name string `json:"name,omitempty"`
	// Version is the version of the release
	Version int `json:"version,omitempty"`
}

// UpdateReleaseRequest updates a release.
type HelmUpdateReleaseRequest struct {
	// The name of the release
	Name string `json:"name,omitempty"`
	// Chart is the protobuf representation of a chart.
	Chart *HelmChart `json:"chart,omitempty"`
	// Values is a string containing (unparsed) YAML values.
	Values unstructured.Unstructured `json:"values,omitempty"`
	// dry_run, if true, will run through the release logic, but neither create
	DryRun bool `json:"dry_run,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the upgrade.
	DisableHooks bool `json:"disable_hooks,omitempty"`
	// Performs pods restart for resources if applicable
	Recreate bool `json:"recreate,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `json:"timeout,omitempty"`
	// ResetValues will cause Tiller to ignore stored values, resetting to default values.
	ResetValues bool `json:"reset_values,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `json:"wait,omitempty"`
	// ReuseValues will cause Tiller to reuse the values from the last release.
	// This is ignored if reset_values is set.
	ReuseValues bool `json:"reuse_values,omitempty"`
	// Force resource update through delete/recreate if needed.
	Force bool `json:"force,omitempty"`
	// Limit the maximum number of revisions saved per release.
	MaxHistory int `json:"max_history,omitempty"`
}

// RollbackReleaseRequest is the request for a release to be rolledback to a
// previous version.
type HelmRollbackReleaseRequest struct {
	// The name of the release
	Name string `json:"name,omitempty"`
	// dry_run, if true, will run through the release logic but no create
	DryRun bool `json:"dry_run,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the rollback
	DisableHooks bool `json:"disable_hooks,omitempty"`
	// Version is the version of the release to deploy.
	Version int `json:"version,omitempty"`
	// Performs pods restart for resources if applicable
	Recreate bool `json:"recreate,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `json:"timeout,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `json:"wait,omitempty"`
	// Force resource update through delete/recreate if needed.
	Force bool `json:"force,omitempty"`
}

// InstallReleaseRequest is the request for an installation of a chart.
type HelmInstallReleaseRequest struct {
	// Chart is the protobuf representation of a chart.
	Chart *HelmChart `json:"chart,omitempty"`
	// Values is a string containing (unparsed) YAML values.
	Values unstructured.Unstructured `json:"values,omitempty"`
	// DryRun, if true, will run through the release logic, but neither create
	// a release object nor deploy to Kubernetes. The release object returned
	// in the response will be fake.
	DryRun bool `json:"dry_run,omitempty"`
	// Name is the candidate release name. This must be unique to the
	// namespace, otherwise the server will return an error. If it is not
	// supplied, the server will autogenerate one.
	Name string `json:"name,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the install.
	DisableHooks bool `json:"disable_hooks,omitempty"`
	// Namepace is the kubernetes namespace of the release.
	Namespace string `json:"namespace,omitempty"`
	// ReuseName requests that Tiller re-uses a name, instead of erroring out.
	ReuseName bool `json:"reuse_name,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `json:"timeout,omitempty"`
	// wait, if true, will wait until all Pods, PVCs, and Services are in a ready state
	// before marking the release as successful. It will wait for as long as timeout
	Wait bool `json:"wait,omitempty"`
}

// UninstallReleaseRequest represents a request to uninstall a named release.
type HelmUninstallReleaseRequest struct {
	// Name is the name of the release to delete.
	Name string `json:"name,omitempty"`
	// DisableHooks causes the server to skip running any hooks for the uninstall.
	DisableHooks bool `json:"disable_hooks,omitempty"`
	// Purge removes the release from the store and make its name free for later use.
	Purge bool `json:"purge,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `json:"timeout,omitempty"`
}

// UninstallReleaseResponse represents a successful response to an uninstall request.
type UninstallReleaseResponse struct {
	// Release is the release that was marked deleted.
	Release *HelmReleaseDesc `json:"release,omitempty"`
	// Info is an uninstall message
	Info string `json:"info,omitempty"`
}

// GetHistoryRequest requests a release's history.
type HelmGetHistoryRequest struct {
	// The name of the release.
	Name string `json:"name,omitempty"`
	// The maximum number of releases to include.
	Max int `json:"max,omitempty"`
}

// TestReleaseRequest is a request to get the status of a release.
type HelmTestReleaseRequest struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`
	// timeout specifies the max amount of time any kubernetes client command can run.
	Timeout int64 `json:"timeout,omitempty"`
	// cleanup specifies whether or not to attempt pod deletion after test completes
	Cleanup bool `json:"cleanup,omitempty"`
}

// TestReleaseResponse represents a message from executing a test
type TestReleaseResponse struct {
	Msg    string            `json:"msg,omitempty"`
	Status HelmTestRunStatus `json:"status,omitempty"`
}

// HelmRequestSpec defines the parameters of an HelmRequest
type HelmRequestSpec struct {
	// Install the HelmManifest specified int the request
	Install *HelmInstallReleaseRequest `json:"apply,omitempty"`
	// Upgrade the HelmRelease specified in the request
	Upgrade *HelmUpdateReleaseRequest `json:"upgrade,omitempty"`
	// Rollback the HelmRelease specified in the request
	Rollback *HelmRollbackReleaseRequest `json:"rollback,omitempty"`
	// Delete the HelmRelease specified in the request
	Delete *HelmRollbackReleaseRequest `json:"delete,omitempty"`
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
