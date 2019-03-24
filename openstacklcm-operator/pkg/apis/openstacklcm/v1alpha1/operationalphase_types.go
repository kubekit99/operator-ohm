package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// OperationalPhaseSpec defines the desired state of OperationalPhase
type OperationalPhaseSpec struct {
	// OpenstackVersion is the version of the backup openstack server.
	OpenstackVersion string `json:"openstackVersion,omitempty"`
	// OpenstackRevision is the revision of openstack's KV store where the backup is performed on.
	OpenstackRevision int32 `json:"openstackRevision,omitempty"`

	// Administrative State of the resource. Is the reconcilation of the CRD by its controller enabled
	AdminState OpenstackLcmAdminState `json:"admin_state"`
	// Target state of the Lcm Custom Resources
	TargetState LcmResourceState `json:"target_state"`
	// revisionHistoryLimit is the maximum number of revisions that will
	// be maintained in the OperationalPhase's revision history. The revision history
	// consists of all revisions not represented by a currently applied
	// OperationalPhaseSpec version. The default value is 10.
	RevisionHistoryLimit *int32 `json:"revisionHistoryLimit,omitempty"`
}

// OperationalPhaseStatus defines the observed state of OperationalPhase
type OperationalPhaseStatus struct {
	OpenstackLcmStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperationalPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=operationalphases,shortName=osvc
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type OperationalPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OperationalPhaseSpec   `json:"spec,omitempty"`
	Status OperationalPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an OperationalPhase. Namely, if the state has not been
// specified, it will be set
func (obj *OperationalPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *OperationalPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed OperationalPhase
func ToOperationalPhase(u *unstructured.Unstructured) *OperationalPhase {
	var obj *OperationalPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &OperationalPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed OperationalPhase into an unstructured.Unstructured
func (obj *OperationalPhase) FromOperationalPhase() *unstructured.Unstructured {
	u := NewOperationalPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *OperationalPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the chart group if managed by the reconcilier
func (obj *OperationalPhase) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the chart group is not managed by the reconcilier
func (obj *OperationalPhase) IsDisabled() bool {
	return !obj.IsEnabled()
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *OperationalPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *OperationalPhase) GetName() string {
	return obj.ObjectMeta.Name
}

func (obj *OperationalPhase) GetNotes() string {
	return "Notes"
}

func (obj *OperationalPhase) GetVersion() int32 {
	return obj.Spec.OpenstackRevision
}

// Returns a GKV for OperationalPhase
func NewOperationalPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("OperationalPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OperationalPhaseList contains a list of OperationalPhase
type OperationalPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OperationalPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed OperationalPhaseList
func ToOperationalPhaseList(u *unstructured.Unstructured) *OperationalPhaseList {
	var obj *OperationalPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &OperationalPhaseList{}
	}
	return obj
}

// Convert a typed OperationalPhaseList into an unstructured.Unstructured
func (obj *OperationalPhaseList) FromOperationalPhaseList() *unstructured.Unstructured {
	u := NewOperationalPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *OperationalPhaseList) Equivalent(other *OperationalPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for OperationalPhaseList
func NewOperationalPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("OperationalPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&OperationalPhase{}, &OperationalPhaseList{})
}
