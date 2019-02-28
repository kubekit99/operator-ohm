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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type HelmReleaseConditionType string
type ConditionStatus string
type HelmReleaseConditionReason string

type HelmReleaseCondition struct {
	Type    HelmReleaseConditionType   `json:"type"`
	Status  ConditionStatus            `json:"status"`
	Reason  HelmReleaseConditionReason `json:"reason,omitempty"`
	Message string                     `json:"message,omitempty"`
	// Release     *rpb.Release                  `json:"release,omitempty"`
	ReleaseName    string `json:"releaseName,omitempty"`
	ReleaseVersion int32  `json:"releaseVersion,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

const (
	ConditionInitialized    HelmReleaseConditionType = "Initialized"
	ConditionDeployed       HelmReleaseConditionType = "Deployed"
	ConditionReleaseFailed  HelmReleaseConditionType = "ReleaseFailed"
	ConditionIrreconcilable HelmReleaseConditionType = "Irreconcilable"
	ConditionBackedUp       HelmReleaseConditionType = "BackedUp"
	ConditionRestored       HelmReleaseConditionType = "Restored"
	ConditionUpgraded       HelmReleaseConditionType = "Upgraded"
	ConditionRolledBack     HelmReleaseConditionType = "RolledBack"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   HelmReleaseConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    HelmReleaseConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful HelmReleaseConditionReason = "UninstallSuccessful"
	ReasonInstallError        HelmReleaseConditionReason = "InstallError"
	ReasonUpdateError         HelmReleaseConditionReason = "UpdateError"
	ReasonReconcileError      HelmReleaseConditionReason = "ReconcileError"
	ReasonUninstallError      HelmReleaseConditionReason = "UninstallError"
	ReasonBackupError         HelmReleaseConditionReason = "BackupError"
	ReasonRestoreError        HelmReleaseConditionReason = "RestoreError"
	ReasonUpgradeError        HelmReleaseConditionReason = "UpgradeError"
	ReasonRollbackError       HelmReleaseConditionReason = "RollbackError"
)

// HelmReleaseSpec defines the desired state of HelmRelease
type HelmReleaseSpec struct {
	// Helm Chart releate information
	ReleaseName                 string `json:"releaseName,omitempty"`
	ChartDir                    string `json:"chartDir,omitempty"`
	WatchHelmDependentResources bool   `json:"watchHelmDependentResources"`

	// ReleaseDesc is the chart that was released.
	ReleaseDesc *HelmReleaseDesc `json:"releaseDesc,omitempty"`
}

// HelmReleaseStatus defines the observed state of HelmRelease
type HelmReleaseStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
	// List of conditions and states related to the release
	Conditions []HelmReleaseCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmRelease is the Schema for the armadas API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=helmreleases,shortName=hrel
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type HelmRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmReleaseSpec   `json:"spec,omitempty"`
	Status HelmReleaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelmReleaseList contains a list of HelmRelease
type HelmReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmRelease `json:"items"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *HelmReleaseStatus) SetCondition(condition HelmReleaseCondition) *HelmReleaseStatus {
	now := metav1.Now()
	for i := range s.Conditions {
		if s.Conditions[i].Type == condition.Type {
			if s.Conditions[i].Status != condition.Status {
				condition.LastTransitionTime = now
			} else {
				condition.LastTransitionTime = s.Conditions[i].LastTransitionTime
			}
			s.Conditions[i] = condition
			return s
		}
	}

	// If the condition does not exist,
	// initialize the lastTransitionTime
	condition.LastTransitionTime = now
	s.Conditions = append(s.Conditions, condition)
	return s
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *HelmReleaseStatus) RemoveCondition(conditionType HelmReleaseConditionType) *HelmReleaseStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return s
		}
	}
	return s
}

// Returns a GKV for HelmRelease
func NewHelmReleaseVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("armada.airshipit.org/v1alpha1")
	u.SetKind("HelmRelease")
	return u
}

// StatusFor safely returns a typed status block from a custom resource.
func StatusFor(cr *HelmRelease) *HelmReleaseStatus {
	res := &cr.Status

	if res.Conditions == nil {
		res.Conditions = make([]HelmReleaseCondition, 0)
	}

	return res
}

func init() {
	SchemeBuilder.Register(&HelmRelease{}, &HelmReleaseList{})
}
