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
	// "encoding/json"
	yaml "gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Administractive state of the reconcilation of a CRD by the corresponding controller
type ArmadaAdminState string

// Describe the Administrative State of the Chart
const (
	// StateUnknown indicates that a release/chart/chartgroup/manifest automatic reconcilation by the controller is enabled
	StateEnabled ArmadaAdminState = "enabled"
	// StateUnknown indicates that a release/chart/chartgroup/manifest automatic reconcilation by the controller is disabled
	StateDisabled ArmadaAdminState = "disabled"
)

// String converts a ArmadaAdminState to a printable string
func (x ArmadaAdminState) String() string { return string(x) }

// State is the status of a release/chart/chartgroup/manifest
type HelmResourceState string
type HelmResourceConditionType string
type HelmResourceConditionStatus string
type HelmResourceConditionReason string

// String converts a HelmResourceState to a printable string
func (x HelmResourceState) String() string { return string(x) }

// String converts a HelmResourceConditionType to a printable string
func (x HelmResourceConditionType) String() string { return string(x) }

// String converts a HelmResourceConditionState to a printable string
func (x HelmResourceConditionStatus) String() string { return string(x) }

// String converts a HelmResourceConditionReason to a printable string
func (x HelmResourceConditionReason) String() string { return string(x) }

// Describe the status of a release
const (
	// StateUnknown indicates that a release/chart/chartgroup/manifest is in an uncertain state.
	StateUnknown HelmResourceState = "unknown"
	// StateInitialized indicates that a release/chart/chartgroup/manifest is in an Kubernetes
	StateInitialized HelmResourceState = "initialized"
	// StateDeployed indicates that the release/chart/chartgroup/manifest has been downloaded from artifact repository
	StateDownloaded HelmResourceState = "downloaded"
	// StateDeployed indicates that the release/chart/chartgroup/manifest has been pushed to Kubernetes.
	StateDeployed HelmResourceState = "deployed"
	// StateUninstalled indicates that a release/chart/chartgroup/manifest has been uninstalled from Kubermetes.
	StateUninstalled HelmResourceState = "uninstalled"
	// StateSuperseded indicates that this release/chart/chartgroup/manifest object is outdated and a newer one exists.
	StateSuperseded HelmResourceState = "superseded"
	// StateFailed indicates that the release/chart/chartgroup/manifest was not successfully deployed.
	StateFailed HelmResourceState = "failed"
	// StateUninstalling indicates that a uninstall operation is underway.
	StateUninstalling HelmResourceState = "uninstalling"
	// StatePendingInstall indicates that an install operation is underway.
	StatePendingInstall HelmResourceState = "pending-install"
	// StatePendingUpgrade indicates that an upgrade operation is underway.
	StatePendingUpgrade HelmResourceState = "pending-upgrade"
	// StatePendingRollback indicates that an rollback operation is underway.
	StatePendingRollback HelmResourceState = "pending-rollback"
	// StatePendingBackup indicates that an data backup operation is underway.
	StatePendingBackup HelmResourceState = "pending-backup"
	// StatePendingRestore indicates that an data restore operation is underway.
	StatePendingRestore HelmResourceState = "pending-restore"
	// StatePendingRestore indicates that an data restore operation is underway.
	StatePendingInitialization HelmResourceState = "pending-initialization"
)

const (
	// XXX
	ConditionStatusTrue    HelmResourceConditionStatus = "True"
	ConditionStatusFalse   HelmResourceConditionStatus = "False"
	ConditionStatusUnknown HelmResourceConditionStatus = "Unknown"

	// ConditionType
	ConditionIrreconcilable HelmResourceConditionType = "Irreconcilable"
	ConditionFailed         HelmResourceConditionType = "Failed"
	ConditionInitialized    HelmResourceConditionType = "Initialized"
	ConditionEnabled        HelmResourceConditionType = "Enabled"
	ConditionDownloaded     HelmResourceConditionType = "Downloaded"
	ConditionDeployed       HelmResourceConditionType = "Deployed"

	// JEB: Not sure we will ever be able to use those conditions
	ConditionBackedUp   HelmResourceConditionType = "BackedUp"
	ConditionRestored   HelmResourceConditionType = "Restored"
	ConditionUpgraded   HelmResourceConditionType = "Upgraded"
	ConditionRolledBack HelmResourceConditionType = "RolledBack"

	// Successful Conditions Reasons
	ReasonInstallSuccessful   HelmResourceConditionReason = "InstallSuccessful"
	ReasonDownloadSuccessful  HelmResourceConditionReason = "DownloadSuccessful"
	ReasonReconcileSuccessful HelmResourceConditionReason = "ReconcileSuccessful"
	ReasonUninstallSuccessful HelmResourceConditionReason = "UninstallSuccessful"
	ReasonUpdateSuccessful    HelmResourceConditionReason = "UpdateSuccessful"

	// Finer grain successful reason of update
	ReasonBackupSuccessful   HelmResourceConditionReason = "BackupSuccessful"
	ReasonRestoreSuccessful  HelmResourceConditionReason = "RestoreSuccessful"
	ReasonUpgradeSuccessful  HelmResourceConditionReason = "UpgradeSuccessful"
	ReasonRollbackSuccessful HelmResourceConditionReason = "RollbackSuccessful"

	// Error Condition Reasons
	ReasonInstallError   HelmResourceConditionReason = "InstallError"
	ReasonDownloadError  HelmResourceConditionReason = "DownloadError"
	ReasonReconcileError HelmResourceConditionReason = "ReconcileError"
	ReasonUninstallError HelmResourceConditionReason = "UninstallError"
	ReasonUpdateError    HelmResourceConditionReason = "UpdateError"

	// Finer grain error reason of update
	ReasonBackupError   HelmResourceConditionReason = "BackupError"
	ReasonRestoreError  HelmResourceConditionReason = "RestoreError"
	ReasonUpgradeError  HelmResourceConditionReason = "UpgradeError"
	ReasonRollbackError HelmResourceConditionReason = "RollbackError"
)

type HelmResourceCondition struct {
	Type               HelmResourceConditionType   `json:"type"`
	Status             HelmResourceConditionStatus `json:"status"`
	Reason             HelmResourceConditionReason `json:"reason,omitempty"`
	Message            string                      `json:"message,omitempty"`
	ResourceName       string                      `json:"resourceName,omitempty"`
	ResourceVersion    int32                       `json:"resourceVersion,omitempty"`
	LastTransitionTime metav1.Time                 `json:"lastTransitionTime,omitempty"`
}

type HelmResourceConditionListHelper struct {
	Items []HelmResourceCondition `json:"items"`
}

type ArmadaStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"reason,omitempty"`
	// Actual state of the Helm Custom Resources
	ActualState HelmResourceState `json:"actual_state"`
	// List of conditions and states related to the resource. JEB: Feature kind of overlap with event recorder
	Conditions []HelmResourceCondition `json:"conditions,omitempty"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *HelmResourceConditionListHelper) SetCondition(condition HelmResourceCondition) []HelmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]HelmResourceCondition, 0)
	}

	now := metav1.Now()
	for i := range s.Items {
		if s.Items[i].Type == condition.Type {
			if s.Items[i].Status != condition.Status {
				condition.LastTransitionTime = now
			} else {
				condition.LastTransitionTime = s.Items[i].LastTransitionTime
			}
			s.Items[i] = condition
			return s.Items
		}
	}

	// If the condition does not exist,
	// initialize the lastTransitionTime
	condition.LastTransitionTime = now
	s.Items = append(s.Items, condition)
	return s.Items
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *HelmResourceConditionListHelper) RemoveCondition(conditionType HelmResourceConditionType) []HelmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]HelmResourceCondition, 0)
	}

	for i := range s.Items {
		if s.Items[i].Type == conditionType {
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return s.Items
		}
	}
	return s.Items
}

// Initialize the HelmResourceCondition list
func (s *HelmResourceConditionListHelper) InitIfEmpty() []HelmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]HelmResourceCondition, 0)
	}

	return s.Items
}

// Utility function to print an HelmResourceCondition list
func (s *HelmResourceConditionListHelper) PrettyPrint() string {
	// res, _ := json.MarshalIndent(s.Items, "", "\t")
	res, _ := yaml.Marshal(s.Items)
	return string(res)
}

// Utility function to find an HelmResourceCondition within the List
func (s *HelmResourceConditionListHelper) FindCondition(conditionType HelmResourceConditionType, conditionStatus HelmResourceConditionStatus) *HelmResourceCondition {
	var found *HelmResourceCondition
	for _, condition := range s.Items {
		if condition.Type == conditionType && condition.Status == conditionStatus {
			found = &condition
			break
		}
	}
	return found
}

func (s *ArmadaStatus) ComputeActualState(cond HelmResourceCondition, target HelmResourceState) {
	// TODO(Ian): finish this
	if cond.Status == ConditionStatusTrue {
		if cond.Type == ConditionInitialized {
			// Since that condition is set almost systematically
			// let's do not recompute the state.
			if (s.ActualState == "") || (s.ActualState == StateUnknown) {
				s.ActualState = StateInitialized
				s.Succeeded = (s.ActualState == target)
				s.Reason = ""
			}
		} else if cond.Type == ConditionDeployed {
			s.ActualState = StateDeployed
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		} else if cond.Type == ConditionEnabled {
			if (s.ActualState == "") || (s.ActualState == StateUnknown) {
				s.ActualState = StatePendingInitialization
				s.Succeeded = (s.ActualState == target)
				s.Reason = ""
			}
		} else if cond.Type == ConditionIrreconcilable {
			s.ActualState = StateFailed
			s.Succeeded = false
			s.Reason = cond.Reason.String()
		} else {
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		}
	} else {
		if cond.Type == ConditionDeployed {
			s.ActualState = StateUninstalled
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		} else if cond.Type == ConditionEnabled {
			s.ActualState = StateUnknown
			s.Succeeded = true
			s.Reason = "Disabled Resource is always successful"
		} else {
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		}
	}
}
