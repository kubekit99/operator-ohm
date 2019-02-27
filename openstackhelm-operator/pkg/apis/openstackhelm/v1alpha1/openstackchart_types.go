package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenstackChartSpec defines the desired state of OpenstackChart
type OpenstackChartSpec struct {
	// Helm Chart releate information
	ReleaseName                 string `json:"releaseName,omitempty"`
	ChartDir                    string `json:"chartDir,omitempty"`
	WatchHelmDependentResources bool   `json:"watchHelmDependentResources"`

	// DB Restore related workflows
	PreRestoreWorkflow    string `json:"preRestoreWorkflow,omitempty"`
	PostRestoreWorkflow   string `json:"postRestoreWorkflow,omitempty"`
	RestoreWorflowEnabled bool   `json:"restoreWorkflowEnabled"`

	// DB Backup related workflows
	PreBackupWorkflow     string `json:"preBackupWorkflow,omitempty"`
	PostBackupWorkflow    string `json:"postBackupWorkflow,omitempty"`
	BackupWorkflowEnabled bool   `json:"backupWorkflowEnabled"`

	// helm install related workflows
	PreDeploymentWorkflow  string `json:"preDeploymentWorkflow,omitempty"`
	PostDeploymentWorkflow string `json:"postDeploymentWorkflow,omitempty"`
	DeploymentEnabled      bool   `json:"deploymentWorflowEnabled"`

	// helm upgrade related workflows
	PreUpgradeWorkflow    string `json:"preUpgradeWorkflow,omitempty"`
	PostUpgradeWorkflow   string `json:"postUpgradeWorkflow,omitempty"`
	UpgradeWorflowEnabled bool   `json:"upgradeWorflowEnabled"`

	// helm rollback related workflows
	PreRollbackWorkflow    string `json:"preRollbackWorkflow,omitempty"`
	PostRollbackWorkflow   string `json:"postRollbackWorkflow,omitempty"`
	RollbackWorflowEnabled bool   `json:"rollbackWorkflowEnabled"`
}

type OpenstackChartConditionType string
type ConditionStatus string
type OpenstackChartConditionReason string

type OpenstackChartCondition struct {
	Type           OpenstackChartConditionType   `json:"type"`
	Status         ConditionStatus               `json:"status"`
	Reason         OpenstackChartConditionReason `json:"reason,omitempty"`
	Message        string                        `json:"message,omitempty"`
	Release        *rpb.Release                  `json:"release,omitempty"`
	ReleaseName    string                        `json:"releaseName,omitempty"`
	ReleaseVersion int32                         `json:"releaseVersion,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

const (
	ConditionInitialized    OpenstackChartConditionType = "Initialized"
	ConditionDeployed       OpenstackChartConditionType = "Deployed"
	ConditionReleaseFailed  OpenstackChartConditionType = "ReleaseFailed"
	ConditionIrreconcilable OpenstackChartConditionType = "Irreconcilable"
	ConditionBackedUp       OpenstackChartConditionType = "BackedUp"
	ConditionRestored       OpenstackChartConditionType = "Restored"
	ConditionUpgraded       OpenstackChartConditionType = "Upgraded"
	ConditionRolledBack     OpenstackChartConditionType = "RolledBack"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   OpenstackChartConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    OpenstackChartConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful OpenstackChartConditionReason = "UninstallSuccessful"
	ReasonInstallError        OpenstackChartConditionReason = "InstallError"
	ReasonUpdateError         OpenstackChartConditionReason = "UpdateError"
	ReasonReconcileError      OpenstackChartConditionReason = "ReconcileError"
	ReasonUninstallError      OpenstackChartConditionReason = "UninstallError"
	ReasonBackupError         OpenstackChartConditionReason = "BackupError"
	ReasonRestoreError        OpenstackChartConditionReason = "RestoreError"
	ReasonUpgradeError        OpenstackChartConditionReason = "UpgradeError"
	ReasonRollbackError       OpenstackChartConditionReason = "RollbackError"
)

// OpenstackChartStatus defines the observed state of OpenstackChart
type OpenstackChartStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`

	Conditions []OpenstackChartCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackChart is the Schema for the openstackhelms API
// +k8s:openapi-gen=true
type OpenstackChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackChartSpec   `json:"spec,omitempty"`
	Status OpenstackChartStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackChartList contains a list of OpenstackChart
type OpenstackChartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackChart `json:"items"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *OpenstackChartStatus) SetCondition(condition OpenstackChartCondition) *OpenstackChartStatus {
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
func (s *OpenstackChartStatus) RemoveCondition(conditionType OpenstackChartConditionType) *OpenstackChartStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return s
		}
	}
	return s
}

// Returns a GKV for OpenstackChart
func NewOpenstackChartVersionKind() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstackhelm.openstack.org/v1alpha1")
	u.SetKind("OpenstackChart")
	return u
}

// StatusFor safely returns a typed status block from a custom resource.
func StatusFor(cr *OpenstackChart) *OpenstackChartStatus {
	res := &cr.Status

	if res.Conditions == nil {
		res.Conditions = make([]OpenstackChartCondition, 0)
	}

	return res
}

func init() {
	SchemeBuilder.Register(&OpenstackChart{}, &OpenstackChartList{})
}
