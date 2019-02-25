package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenstackHelmSpec defines the desired state of OpenstackHelm
type OpenstackHelmSpec struct {
	// ReleaseName indicates the name of the release
	ReleaseName string `json:"releaseName,omitempty"`
	// Directory for the configuration
	ChartDir string `json:"chartDir,omitempty"`
}

type OpenstackHelmConditionType string
type ConditionStatus string
type OpenstackHelmConditionReason string

type OpenstackHelmCondition struct {
	Type           OpenstackHelmConditionType   `json:"type"`
	Status         ConditionStatus              `json:"status"`
	Reason         OpenstackHelmConditionReason `json:"reason,omitempty"`
	Message        string                       `json:"message,omitempty"`
	ReleaseName    string                       `json:"releaseName,omitempty"`
	ReleaseVersion int32                        `json:"releaseVersion,omitempty"`

	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

const (
	ConditionInitialized    OpenstackHelmConditionType = "Initialized"
	ConditionDeployed       OpenstackHelmConditionType = "Deployed"
	ConditionReleaseFailed  OpenstackHelmConditionType = "ReleaseFailed"
	ConditionIrreconcilable OpenstackHelmConditionType = "Irreconcilable"

	StatusTrue    ConditionStatus = "True"
	StatusFalse   ConditionStatus = "False"
	StatusUnknown ConditionStatus = "Unknown"

	ReasonInstallSuccessful   OpenstackHelmConditionReason = "InstallSuccessful"
	ReasonUpdateSuccessful    OpenstackHelmConditionReason = "UpdateSuccessful"
	ReasonUninstallSuccessful OpenstackHelmConditionReason = "UninstallSuccessful"
	ReasonInstallError        OpenstackHelmConditionReason = "InstallError"
	ReasonUpdateError         OpenstackHelmConditionReason = "UpdateError"
	ReasonReconcileError      OpenstackHelmConditionReason = "ReconcileError"
	ReasonUninstallError      OpenstackHelmConditionReason = "UninstallError"
)

// OpenstackHelmStatus defines the observed state of OpenstackHelm
type OpenstackHelmStatus struct {
	// Succeeded indicates if the operattion has Succeeded.
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`

	Conditions []OpenstackHelmCondition `json:"conditions"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackHelm is the Schema for the openstackhelms API
// +k8s:openapi-gen=true
type OpenstackHelm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenstackHelmSpec   `json:"spec,omitempty"`
	Status OpenstackHelmStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OpenstackHelmList contains a list of OpenstackHelm
type OpenstackHelmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenstackHelm `json:"items"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *OpenstackHelmStatus) SetCondition(condition OpenstackHelmCondition) *OpenstackHelmStatus {
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
func (s *OpenstackHelmStatus) RemoveCondition(conditionType OpenstackHelmConditionType) *OpenstackHelmStatus {
	for i := range s.Conditions {
		if s.Conditions[i].Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return s
		}
	}
	return s
}

// StatusFor safely returns a typed status block from a custom resource.
func StatusFor(cr *OpenstackHelm) *OpenstackHelmStatus {
	res := &cr.Status

	if res.Conditions == nil {
		res.Conditions = make([]OpenstackHelmCondition, 0)
	}

	return res
}

func init() {
	SchemeBuilder.Register(&OpenstackHelm{}, &OpenstackHelmList{})
}
