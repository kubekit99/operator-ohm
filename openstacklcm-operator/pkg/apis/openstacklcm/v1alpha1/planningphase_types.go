package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// PlanningPhaseSpec defines the desired state of PlanningPhase
type PlanningPhaseSpec struct {
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int32 `json:"openstackRevision,omitempty"`

	// Administrative State of the resource. Is the reconcilation of the CRD by its controller enabled
	AdminState OpenstackLcmAdminState `json:"admin_state"`
	// Target state of the Lcm Custom Resources
	TargetState LcmResourceState `json:"target_state"`
	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the PlanningPhase's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// PlanningPhaseSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
}

// PlanningPhaseStatus defines the observed state of PlanningPhase
type PlanningPhaseStatus struct {
	OpenstackLcmStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanningPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=planningphases,shortName=osvc
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type PlanningPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanningPhaseSpec   `json:"spec,omitempty"`
	Status PlanningPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an PlanningPhase. Namely, if the state has not been
// specified, it will be set
func (obj *PlanningPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *PlanningPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed PlanningPhase
func ToPlanningPhase(u *unstructured.Unstructured) *PlanningPhase {
	var obj *PlanningPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &PlanningPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed PlanningPhase into an unstructured.Unstructured
func (obj *PlanningPhase) FromPlanningPhase() *unstructured.Unstructured {
	u := NewPlanningPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *PlanningPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the chart group if managed by the reconcilier
func (obj *PlanningPhase) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the chart group is not managed by the reconcilier
func (obj *PlanningPhase) IsDisabled() bool {
	return !obj.IsEnabled()
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *PlanningPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *PlanningPhase) GetName() string {
	return obj.ObjectMeta.Name
}

func (obj *PlanningPhase) GetNotes() string {
	return "Notes"
}

func (obj *PlanningPhase) GetVersion() int32 {
	return obj.Spec.OpenstackRevision
}

// Returns a GKV for PlanningPhase
func NewPlanningPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("PlanningPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PlanningPhaseList contains a list of PlanningPhase
type PlanningPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PlanningPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed PlanningPhaseList
func ToPlanningPhaseList(u *unstructured.Unstructured) *PlanningPhaseList {
	var obj *PlanningPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &PlanningPhaseList{}
	}
	return obj
}

// Convert a typed PlanningPhaseList into an unstructured.Unstructured
func (obj *PlanningPhaseList) FromPlanningPhaseList() *unstructured.Unstructured {
	u := NewPlanningPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *PlanningPhaseList) Equivalent(other *PlanningPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for PlanningPhaseList
func NewPlanningPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("PlanningPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&PlanningPhase{}, &PlanningPhaseList{})
}
