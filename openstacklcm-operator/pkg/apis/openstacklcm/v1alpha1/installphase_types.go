package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

// InstallPhaseSpec defines the desired state of InstallPhase
type InstallPhaseSpec struct {
	PhaseSpec `json:",inline"`
}

// InstallPhaseStatus defines the observed state of InstallPhase
type InstallPhaseStatus struct {
	PhaseStatus `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstallPhase is the Schema for the openstackdeployments API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=installphases,shortName=osins
// +kubebuilder:printcolumn:name="Succeeded",type="boolean",JSONPath=".status.succeeded",description="Succeeded"
type InstallPhase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstallPhaseSpec   `json:"spec,omitempty"`
	Status InstallPhaseStatus `json:"status,omitempty"`
}

// Init is used to initialize an InstallPhase. Namely, if the state has not been
// specified, it will be set
func (obj *InstallPhase) Init() {
	if obj.Status.ActualState == "" {
		obj.Status.ActualState = StateUninitialied
	}
	if obj.Spec.TargetState == "" {
		obj.Spec.TargetState = StateUninitialied
	}
	obj.Status.Succeeded = (obj.Spec.TargetState == obj.Status.ActualState)
}

// Return the list of dependent resources to watch
func (obj *InstallPhase) GetDependentResources() []unstructured.Unstructured {
	var res = make([]unstructured.Unstructured, 0)
	return res
}

// Convert an unstructured.Unstructured into a typed InstallPhase
func ToInstallPhase(u *unstructured.Unstructured) *InstallPhase {
	var obj *InstallPhase
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &InstallPhase{
			ObjectMeta: metav1.ObjectMeta{
				Name:      u.GetName(),
				Namespace: u.GetNamespace(),
			},
		}
	}
	return obj
}

// Convert a typed InstallPhase into an unstructured.Unstructured
func (obj *InstallPhase) FromInstallPhase() *unstructured.Unstructured {
	u := NewInstallPhaseVersionKind(obj.ObjectMeta.Namespace, obj.ObjectMeta.Name)
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// IsDeleted returns true if the chart group has been deleted
func (obj *InstallPhase) IsDeleted() bool {
	return obj.GetDeletionTimestamp() != nil
}

// IsEnabled returns true if the chart group if managed by the reconcilier
func (obj *InstallPhase) IsEnabled() bool {
	return (obj.Spec.AdminState == "") || (obj.Spec.AdminState == StateEnabled)
}

// IsDisabled returns true if the chart group is not managed by the reconcilier
func (obj *InstallPhase) IsDisabled() bool {
	return !obj.IsEnabled()
}

// IsSatisfied returns true if the chart's actual state meets its target state
func (obj *InstallPhase) IsSatisfied() bool {
	return obj.Spec.TargetState == obj.Status.ActualState
}

func (obj *InstallPhase) GetName() string {
	return obj.ObjectMeta.Name
}

func (obj *InstallPhase) GetNotes() string {
	return "Notes"
}

func (obj *InstallPhase) GetVersion() int32 {
	return obj.Spec.OpenstackRevision
}

// Returns a GKV for InstallPhase
func NewInstallPhaseVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("InstallPhase")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// InstallPhaseList contains a list of InstallPhase
type InstallPhaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstallPhase `json:"items"`
}

// Convert an unstructured.Unstructured into a typed InstallPhaseList
func ToInstallPhaseList(u *unstructured.Unstructured) *InstallPhaseList {
	var obj *InstallPhaseList
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.UnstructuredContent(), &obj)
	if err != nil {
		return &InstallPhaseList{}
	}
	return obj
}

// Convert a typed InstallPhaseList into an unstructured.Unstructured
func (obj *InstallPhaseList) FromInstallPhaseList() *unstructured.Unstructured {
	u := NewInstallPhaseListVersionKind("", "")
	tmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(*obj)
	if err != nil {
		return u
	}
	u.SetUnstructuredContent(tmp)
	return u
}

// JEB: Not sure yet if we really will need it
func (obj *InstallPhaseList) Equivalent(other *InstallPhaseList) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(obj.Items, other.Items)
}

// Returns a GKV for InstallPhaseList
func NewInstallPhaseListVersionKind(namespace string, name string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetAPIVersion("openstacklcm.airshipit.org/v1alpha1")
	u.SetKind("InstallPhaseList")
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

func init() {
	SchemeBuilder.Register(&InstallPhase{}, &InstallPhaseList{})
}
