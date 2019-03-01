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

// JEB: This file has been created by concatenation of the content of
// https://github.com/helm/helm/tree/dev-v3/pkg/hapi/release
// and https://github.com/helm/helm/tree/dev-v3/pkg/chart
// The content was then modified manually.
// This file will be deleted once we figure out what we really want
// to put in our CRDs.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Status is the status of a release
type HelmStatus string

// Describe the status of a release
const (
	// StatusUnknown indicates that a release is in an uncertain state.
	HelmStatusUnknown HelmStatus = "unknown"
	// StatusDeployed indicates that the release has been pushed to Kubernetes.
	HelmStatusDeployed HelmStatus = "deployed"
	// StatusUninstalled indicates that a release has been uninstalled from Kubermetes.
	HelmStatusUninstalled HelmStatus = "uninstalled"
	// StatusSuperseded indicates that this release object is outdated and a newer one exists.
	HelmStatusSuperseded HelmStatus = "superseded"
	// StatusFailed indicates that the release was not successfully deployed.
	HelmStatusFailed HelmStatus = "failed"
	// StatusUninstalling indicates that a uninstall operation is underway.
	HelmStatusUninstalling HelmStatus = "uninstalling"
	// StatusPendingInstall indicates that an install operation is underway.
	HelmStatusPendingInstall HelmStatus = "pending-install"
	// StatusPendingUpgrade indicates that an upgrade operation is underway.
	HelmStatusPendingUpgrade HelmStatus = "pending-upgrade"
	// StatusPendingRollback indicates that an rollback operation is underway.
	HelmStatusPendingRollback HelmStatus = "pending-rollback"
)

// Strng converts a status to a printable string
func (x HelmStatus) String() string { return string(x) }

// Maintainer describes a Chart maintainer.
type HelmMaintainer struct {
	// Name is a user name or organization name
	Name string `json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `json:"email,omitempty"`
	// URL is an optional URL to an address for the named maintainer
	URL string `json:"url,omitempty"`
}

// Metadata for a Chart file. This models the structure of a Chart.yaml file.
//
// Spec: https://k8s.io/helm/blob/master/docs/design/chart_format.md#the-chart-file
type HelmMetadata struct {
	// The name of the chart
	Name string `json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*HelmMaintainer `json:"maintainers,omitempty"`
	// The URL to an icon file.
	Icon string `json:"icon,omitempty"`
	// The API Version of this chart.
	APIVersion string `json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	Condition string `json:"condition,omitempty"`
	// The tags to check to enable chart
	Tags string `json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	AppVersion string `json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// Annotations are additional mappings uninterpreted by Helm,
	// made available for inspection by other applications.
	Annotations map[string]string `json:"annotations,omitempty"`
	// KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion,omitempty"`
	// Dependencies are a list of dependencies for a chart.
	Dependencies []*HelmDependency `json:"dependencies,omitempty"`
	// Specifies the chart type: application or library
	Type string `json:"type,omitempty"`
}

// File represents a file as a name/value pair.
//
// By convention, name is a relative path within the scope of the chart's
// base directory.
type HelmFile struct {
	// Name is the path-like name of the template.
	Name string
	// Data is the template as byte data.
	Data []byte
}

// Dependency describes a chart upon which another chart depends.
//
// Dependencies can be used to express developer intent, or to capture the state
// of a chart.
type HelmDependency struct {
	// Name is the name of the dependency.
	//
	// This must mach the name in the dependency's Chart.yaml.
	Name string `json:"name"`
	// Version is the version (range) of this chart.
	//
	// A lock file will always produce a single version, while a dependency
	// may contain a semantic version range.
	Version string `json:"version,omitempty"`
	// The URL to the repository.
	//
	// Appending `index.yaml` to this string should result in a URL that can be
	// used to fetch the repository index.
	Repository string `json:"repository"`
	// A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
	Condition string `json:"condition,omitempty"`
	// Tags can be used to group charts for enabling/disabling together
	Tags []string `json:"tags,omitempty"`
	// Enabled bool determines if chart should be loaded
	Enabled bool `json:"enabled,omitempty"`
	// ImportValues holds the mapping of source values to parent key to be imported. Each item can be a
	// string or pair of child/parent sublist items.
	ImportValues unstructured.Unstructured `json:"import-values,omitempty"`
	// Alias usable alias to be used for the chart
	Alias string `json:"alias,omitempty"`
}

// Lock is a lock file for dependencies.
//
// It represents the state that the dependencies should be in.
type HelmLock struct {
	// Genderated is the date the lock file was last generated.
	Generated metav1.Time `json:"generated"`
	// Digest is a hash of the dependencies in Chart.yaml.
	Digest string `json:"digest"`
	// Dependencies is the list of dependencies that this lock file has locked.
	Dependencies []*HelmDependency `json:"dependencies"`
}

type HelmChart struct {
	// Metadata is the contents of the Chartfile.
	Metadata *HelmMetadata
	// LocK is the contents of Chart.lock.
	Lock *HelmLock
	// Templates for this chart.
	Templates []*HelmFile
	// TODO Delete RawValues after unit tests for `create` are refactored.
	RawValues []byte
	// Values are default config for this template.
	Values unstructured.Unstructured
	// Files are miscellaneous files in a chart archive,
	// e.g. README, LICENSE, etc.
	Files []*HelmFile

	parent       *HelmChart
	dependencies []*HelmChart
}

// Info describes release information.
type HelmInfo struct {
	// FirstDeployed is when the release was first deployed.
	FirstDeployed metav1.Time `json:"first_deployed,omitempty"`
	// LastDeployed is when the release was last deployed.
	LastDeployed metav1.Time `json:"last_deployed,omitempty"`
	// Deleted tracks when this object was deleted.
	Deleted metav1.Time `json:"deleted,omitempty"`
	// Description is human-friendly "log entry" about this release.
	Description string `json:"Description,omitempty"`
	// Status is the current state of the release
	Status HelmStatus `json:"status,omitempty"`
	// Cluster resources as kubectl would print them.
	Resources string `json:"resources,omitempty"`
	// Contains the rendered templates/NOTES.txt if available
	Notes string `json:"notes,omitempty"`
	// LastTestSuiteRun provides results on the last test run on a release
	LastTestSuiteRun *HelmTestSuite `json:"last_test_suite_run,omitempty"`
}

// TestRunStatus is the status of a test run
type HelmTestRunStatus string

// Indicates the results of a test run
const (
	HelmTestRunUnknown HelmTestRunStatus = "unknown"
	HelmTestRunSuccess HelmTestRunStatus = "success"
	HelmTestRunFailure HelmTestRunStatus = "failure"
	HelmTestRunRunning HelmTestRunStatus = "running"
)

// Strng converts a test run status to a printable string
func (x HelmTestRunStatus) String() string { return string(x) }

// TestRun describes the run of a test
type HelmTestRun struct {
	Name        string            `json:"name,omitempty"`
	Status      HelmTestRunStatus `json:"status,omitempty"`
	Info        string            `json:"info,omitempty"`
	StartedAt   metav1.Time       `json:"started_at,omitempty"`
	CompletedAt metav1.Time       `json:"completed_at,omitempty"`
}

// TestSuite comprises of the last run of the pre-defined test suite of a release version
type HelmTestSuite struct {
	// StartedAt indicates the date/time this test suite was kicked off
	StartedAt metav1.Time `json:"started_at,omitempty"`
	// CompletedAt indicates the date/time this test suite was completed
	CompletedAt metav1.Time `json:"completed_at,omitempty"`
	// Results are the results of each segment of the test
	Results []*HelmTestRun `json:"results,omitempty"`
}

// HookEvent specifies the hook event
type HelmHookEvent string

// Hook event types
const (
	HelmHookPreInstall         HelmHookEvent = "pre-install"
	HelmHookPostInstall        HelmHookEvent = "post-install"
	HelmHookPreDelete          HelmHookEvent = "pre-delete"
	HelmHookPostDelete         HelmHookEvent = "post-delete"
	HelmHookPreUpgrade         HelmHookEvent = "pre-upgrade"
	HelmHookPostUpgrade        HelmHookEvent = "post-upgrade"
	HelmHookPreRollback        HelmHookEvent = "pre-rollback"
	HelmHookPostRollback       HelmHookEvent = "post-rollback"
	HelmHookReleaseTestSuccess HelmHookEvent = "release-test-success"
	HelmHookReleaseTestFailure HelmHookEvent = "release-test-failure"
)

func (x HelmHookEvent) String() string { return string(x) }

// HookDeletePolicy specifies the hook delete policy
type HelmHookDeletePolicy string

// Hook delete policy types
const (
	HelmHookSucceeded          HelmHookDeletePolicy = "succeeded"
	HelmHookFailed             HelmHookDeletePolicy = "failed"
	HelmHookBeforeHookCreation HelmHookDeletePolicy = "before-hook-creation"
)

func (x HelmHookDeletePolicy) String() string { return string(x) }

// Hook defines a hook object.
type HelmHook struct {
	Name string `json:"name,omitempty"`
	// Kind is the Kubernetes kind.
	Kind string `json:"kind,omitempty"`
	// Path is the chart-relative path to the template.
	Path string `json:"path,omitempty"`
	// Manifest is the manifest contents.
	Manifest string `json:"manifest,omitempty"`
	// Events are the events that this hook fires on.
	Events []HelmHookEvent `json:"events,omitempty"`
	// LastRun indicates the date/time this was last run.
	LastRun metav1.Time `json:"last_run,omitempty"`
	// Weight indicates the sort order for execution among similar Hook type
	Weight int `json:"weight,omitempty"`
	// DeletePolicies are the policies that indicate when to delete the hook
	DeletePolicies []HelmHookDeletePolicy `json:"delete_policies,omitempty"`
}

// Release describes a deployment of a chart, together with the chart
// and the variables used to deploy that chart.
type HelmReleaseDesc struct {
	// Name is the name of the release
	Name string `json:"name,omitempty"`
	// Info provides information about a release
	Info *HelmInfo `json:"info,omitempty"`
	// Chart is the chart that was released.
	Chart *HelmChart `json:"chart,omitempty"`
	// Config is the set of extra Values added to the chart.
	// These values override the default values inside of the chart.
	Config *unstructured.Unstructured `json:"config,omitempty"`
	// Manifest is the string representation of the rendered template.
	Manifest string `json:"manifest,omitempty"`
	// Hooks are all of the hooks declared for this release.
	Hooks []*HelmHook `json:"hooks,omitempty"`
	// Version is an int which represents the version of the release.
	Version int `json:"version,omitempty"`
	// Namespace is the kubernetes namespace of the release.
	Namespace string `json:"namespace,omitempty"`
}
