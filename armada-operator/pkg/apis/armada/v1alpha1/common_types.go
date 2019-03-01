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
)

// State is the status of a release
type HelmResourceState string
type HelmResourceConditionType string
type HelmResourceConditionStatus string
type HelmResourceConditionReason string

// Strng converts a HelmResourceState to a printable string
func (x HelmResourceState) String() string { return string(x) }

// Strng converts a HelmResourceConditionType to a printable string
func (x HelmResourceConditionType) String() string { return string(x) }

// Strng converts a HelmResourceConditionState to a printable string
func (x HelmResourceConditionStatus) String() string { return string(x) }

// Strng converts a HelmResourceConditionReason to a printable string
func (x HelmResourceConditionReason) String() string { return string(x) }

// Describe the status of a release
const (
	// StateUnknown indicates that a release/chart/chartgroup/manifest is in an uncertain state.
	StateUnknown HelmResourceState = "unknown"
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
)

const (
	// XXX
	ConditionStatusTrue    HelmResourceConditionStatus = "True"
	ConditionStatusFalse   HelmResourceConditionStatus = "False"
	ConditionStatusUnknown HelmResourceConditionStatus = "Unknown"

	// ConditionType:w
	ConditionIrreconcilable HelmResourceConditionType = "Irreconcilable"
	ConditionFailed         HelmResourceConditionType = "Failed"
	ConditionInitialized    HelmResourceConditionType = "Initialized"
	ConditionDownloaded     HelmResourceConditionType = "Downloaded"
	ConditionDeployed       HelmResourceConditionType = "Deployed"
	ConditionBackedUp       HelmResourceConditionType = "BackedUp"
	ConditionRestored       HelmResourceConditionType = "Restored"
	ConditionUpgraded       HelmResourceConditionType = "Upgraded"
	ConditionRolledBack     HelmResourceConditionType = "RolledBack"

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