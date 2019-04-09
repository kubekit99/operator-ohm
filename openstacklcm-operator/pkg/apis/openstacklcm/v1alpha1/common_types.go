// Copyright 2019 The OpenstackLcm Authors
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
	"reflect"

	yaml "gopkg.in/yaml.v2"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// LcmResourceState is the status of a release/chart/chartgroup/manifest
type LcmResourceState string

type LcmResourceConditionType string

// LcmResourceConditionStatus represents the current status of a Condition
type LcmResourceConditionStatus string

type LcmResourceConditionReason string

// String converts a LcmResourceState to a printable string
func (x LcmResourceState) String() string { return string(x) }

// String converts a LcmResourceConditionType to a printable string
func (x LcmResourceConditionType) String() string { return string(x) }

// String converts a LcmResourceConditionState to a printable string
func (x LcmResourceConditionStatus) String() string { return string(x) }

// String converts a LcmResourceConditionReason to a printable string
func (x LcmResourceConditionReason) String() string { return string(x) }

// Describe the status of a release
const (
	// StateUninitialied indicates that sub resource exists, but has not been acted upon
	StateUninitialized LcmResourceState = "uninitialized"
	// StateUnknown indicates that resource is in an uncertain state.
	StateUnknown LcmResourceState = "unknown"
	// StateInitialized indicates that resource is in an Kubernetes
	StateInitialized LcmResourceState = "initialized"
	// StateDeployed indicates that resource has been pushed to Kubernetes.
	StateDeployed LcmResourceState = "deployed"
	// StateUninstalled indicates the resource has been uninstalled from Kubermetes.
	StateUninstalled LcmResourceState = "uninstalled"
	// StateFailed indicates that resource was not successfully deployed.
	StateFailed LcmResourceState = "failed"
	// StatePending indicates that resource was xxx
	StatePending LcmResourceState = "pending"
	// StateRunning indicates that resource was xxx
	StateRunning LcmResourceState = "running"
	// StateSkipped indicates that resource was xxx
	StateSkipped LcmResourceState = "skipped"
	// StateError indicates that resource was xxx
	StateError LcmResourceState = "error"
)

// These represent acceptable values for a LcmResourceConditionStatus
const (
	ConditionStatusTrue    LcmResourceConditionStatus = "True"
	ConditionStatusFalse                              = "False"
	ConditionStatusUnknown                            = "Unknown"
)

// These represent acceptable values for a LcmResourceConditionType
const (
	ConditionIrreconcilable LcmResourceConditionType = "Irreconcilable"
	ConditionFailed                                  = "Failed"
	ConditionInitialized                             = "Initializing"
	ConditionDeployed                                = "Deployed"
)

// The following represent the more fine-grained reasons for a given condition
const (
	// Successful Conditions Reasons
	ReasonInstallSuccessful   LcmResourceConditionReason = "InstallSuccessful"
	ReasonReconcileSuccessful                            = "ReconcileSuccessful"
	ReasonUninstallSuccessful                            = "UninstallSuccessful"
	ReasonUpdateSuccessful                               = "UpdateSuccessful"

	// Error Condition Reasons
	ReasonInstallError   LcmResourceConditionReason = "InstallError"
	ReasonReconcileError                            = "ReconcileError"
	ReasonUninstallError                            = "UninstallError"
	ReasonUpdateError                               = "UpdateError"
)

// LcmResourceCondition represents one current condition of an Lcm resource
// A condition might not show up if it is not happening.
// For example, if a chart is not deploying, the Deploying condition would not show up.
// If a chart is deploying and encountered a problem that prevents the deployment,
// the Deploying condition's status will would be False and communicate the problem back.
type LcmResourceCondition struct {
	Type               LcmResourceConditionType   `json:"type"`
	Status             LcmResourceConditionStatus `json:"status"`
	Reason             LcmResourceConditionReason `json:"reason,omitempty"`
	Message            string                     `json:"message,omitempty"`
	ResourceName       string                     `json:"resourceName,omitempty"`
	ResourceVersion    int32                      `json:"resourceVersion,omitempty"`
	LastTransitionTime metav1.Time                `json:"lastTransitionTime,omitempty"`
}

type LcmResourceConditionListHelper struct {
	Items []LcmResourceCondition `json:"items"`
}

// OpenstackLcmStatus represents the common attributes shared amongst armada resources
type OpenstackLcmStatus struct {
	// Succeeded indicates if the release's ActualState satisfies its target state
	Succeeded bool `json:"satisfied"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"reason,omitempty"`
	// Actual state of the Lcm Custom Resources
	ActualState LcmResourceState `json:"actualState"`
	// List of conditions and states related to the resource. JEB: Feature kind of overlap with event recorder
	Conditions []LcmResourceCondition `json:"conditions,omitempty"`
}

// PhaseStatus represents the common attributes shared amongst armada resources
type PhaseStatus struct {
	OpenstackLcmStatus `json:",inline"`

	// OpenstackVersion is the version of the backup openstack server.
	ActualOpenstackServiceVersion string `json:"actualOpenstackServiceVersion,omitempty"`
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *OpenstackLcmStatus) SetCondition(cond LcmResourceCondition, tgt LcmResourceState) {

	// Add the condition to the list
	chelper := LcmResourceConditionListHelper{Items: s.Conditions}
	s.Conditions = chelper.SetCondition(cond)

	// Recompute the state
	s.ComputeActualState(cond, tgt)
}

// RemoveCondition removes the condition with the passed condition type from
// the status object. If the condition is not already present, the returned
// status object is returned unchanged. RemoveCondition does not update the
// resource in the cluster.
func (s *OpenstackLcmStatus) RemoveCondition(conditionType LcmResourceConditionType) {
	for i, cond := range s.Conditions {
		if cond.Type == conditionType {
			s.Conditions = append(s.Conditions[:i], s.Conditions[i+1:]...)
			return
		}
	}
}

// SetCondition sets a condition on the status object. If the condition already
// exists, it will be replaced. SetCondition does not update the resource in
// the cluster.
func (s *LcmResourceConditionListHelper) SetCondition(condition LcmResourceCondition) []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
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
func (s *LcmResourceConditionListHelper) RemoveCondition(conditionType LcmResourceConditionType) []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
	}

	for i := range s.Items {
		if s.Items[i].Type == conditionType {
			s.Items = append(s.Items[:i], s.Items[i+1:]...)
			return s.Items
		}
	}
	return s.Items
}

// Initialize the LcmResourceCondition list
func (s *LcmResourceConditionListHelper) InitIfEmpty() []LcmResourceCondition {

	// Initialize the Items array if needed
	if s.Items == nil {
		s.Items = make([]LcmResourceCondition, 0)
	}

	return s.Items
}

// Utility function to print an LcmResourceCondition list
func (s *LcmResourceConditionListHelper) PrettyPrint() string {
	// res, _ := json.MarshalIndent(s.Items, "", "\t")
	res, _ := yaml.Marshal(s.Items)
	return string(res)
}

// Utility function to find an LcmResourceCondition within the List
func (s *LcmResourceConditionListHelper) FindCondition(conditionType LcmResourceConditionType, conditionStatus LcmResourceConditionStatus) *LcmResourceCondition {
	var found *LcmResourceCondition
	for _, condition := range s.Items {
		if condition.Type == conditionType && condition.Status == conditionStatus {
			found = &condition
			break
		}
	}
	return found
}

func (s *OpenstackLcmStatus) ComputeActualState(cond LcmResourceCondition, target LcmResourceState) {
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
		} else if cond.Type == ConditionIrreconcilable {
			s.ActualState = StateError
			s.Succeeded = false
			s.Reason = cond.Reason.String()
		} else if cond.Type == ConditionFailed {
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
		} else {
			s.Succeeded = (s.ActualState == target)
			s.Reason = ""
		}
	}
}

// PhaseSource describe the location of the CR to create during a Phase of an
// Openstack Service Life Cycle.
type PhaseSource struct {
	// ``url`` or ``path`` to the chart's parent directory
	Location string `json:"location"`
	// source to build the chart: ``git``, ``local``, or ``tar``
	Type string `json:"type"`
}

// PhaseSpec defines the desired state of Phase
type PhaseSpec struct {
	// provide a path to a ``git repo``, ``local dir``, or ``tarball url`` chart
	Source *PhaseSource `json:"source"`
	// Openstack Service Name
	OpenstackServiceName string `json:"openstackServiceName"`
	// Openstack Service EndPoint
	OpenstackServiceEndPoint string `json:"openstackServiceEndPoint,omitempty"`

	// OpenstackServiceVersion is the version of the openstack service.
	TargetOpenstackServiceVersion string `json:"targetOpenstackServiceVersion,omitempty"`
	// Target state of the Lcm Custom Resources
	TargetState LcmResourceState `json:"targetState"`
}

// Backup/Restore related types
type BackupStorageType string

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

// SubResourceList represent the list of
type SubResourceList struct {
	Name      string
	Namespace string
	Notes     string
	Version   int32

	// Items is the list of Resources deployed in the K8s cluster
	Items [](unstructured.Unstructured)
}

// Returns the Name for the SubResourceList
func (obj *SubResourceList) GetName() string {
	return obj.Name
}

// Returns the Namespace for this SubResourceList
func (obj *SubResourceList) GetNamespace() string {
	return obj.Namespace
}

// Returns the Notes for this SubResourceList
func (obj *SubResourceList) GetNotes() string {
	return obj.Notes
}

// Returns the Version for this SubResourceList
func (obj *SubResourceList) GetVersion() int32 {
	return obj.Version
}

// Returns the DependentResource for this SubResourceList
func (obj *SubResourceList) GetDependentResources() []unstructured.Unstructured {
	return obj.Items
}

// JEB: Not sure yet if we really will need it
func (obj *SubResourceList) Equivalent(other *SubResourceList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Let's check the reference are setup properly.
func (obj *SubResourceList) CheckOwnerReference(refs []metav1.OwnerReference) bool {

	// Check that each sub resource is owned by the phase
	for _, item := range obj.Items {
		if !reflect.DeepEqual(item.GetOwnerReferences(), refs) {
			return false
		}
	}

	return true
}

// Check the state of a service
func (obj *SubResourceList) IsServiceReady(u unstructured.Unstructured) bool {
	endpoints := corev1.Endpoints{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &endpoints)
	if err1 != nil {
		return false
	}

	for _, subset := range endpoints.Subsets {
		if len(subset.Addresses) > 0 {
			return true
		}
	}
	return false
}

// Check the state of a container
func (obj *SubResourceList) isContainerReady(containerName string, u unstructured.Unstructured) bool {
	pod := corev1.Pod{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &pod)
	if err1 != nil {
		return false
	}

	containers := pod.Status.ContainerStatuses
	for _, container := range containers {
		if container.Name == containerName && container.Ready {
			return true
		}
	}
	return false
}

// Check the state of a job
func (obj *SubResourceList) isJobReady(u unstructured.Unstructured) bool {
	job := batchv1.Job{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &job)
	if err1 != nil {
		return false
	}

	if job.Status.Succeeded == 0 {
		return false
	}
	return true
}

// Check the state of a pod
func (obj *SubResourceList) isPodReady(u unstructured.Unstructured) bool {
	pod := corev1.Pod{}
	err1 := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &pod)
	if err1 != nil {
		return false
	}

	for _, condition := range pod.Status.Conditions {
		if condition.Type == corev1.PodReady && condition.Status == "True" {
			return true
		}
	}
	return false
}

// Returns a new SubResourceList
func NewSubResourceList(namespace string, name string) *SubResourceList {
	res := &SubResourceList{Namespace: namespace, Name: name}
	res.Items = make([]unstructured.Unstructured, 0)
	return res
}
