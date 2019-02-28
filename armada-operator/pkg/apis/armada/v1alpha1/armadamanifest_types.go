package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Labels
type ArmadaLabels struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// Native
type ArmadaNative struct {
	Enabled bool `json:"enabled,omitempty"`
}

// ResourcesItems
type ArmadaResourcesItems struct {
	Labels *ArmadaLabels `json:"labels,omitempty"`
	//JEB MinReady interface{}   `json:"min_ready,omitempty"`
	Type string `json:"type"`
}

// Wait
type ArmadaWait struct {
	Labels    *ArmadaLabels           `json:"labels,omitempty"`
	Native    *ArmadaNative           `json:"native,omitempty"`
	Resources []*ArmadaResourcesItems `json:"resources,omitempty"`
	Timeout   int                     `json:"timeout,omitempty"`
}

// HookActionItems
type ArmadaHookActionItems struct {
	Labels *ArmadaLabels `json:"labels,omitempty"`
	Name   string        `json:"name,omitempty"`
	Type   string        `json:"type"`
}

// Options
type ArmadaOptions struct {
	Force        bool `json:"force,omitempty"`
	RecreatePods bool `json:"recreate_pods,omitempty"`
}

// Install
type ArmadaInstall struct {
}

// Delete
type ArmadaDelete struct {
	Timeout int `json:"timeout,omitempty"`
}

// Upgrade
type ArmadaUpgrade struct {
	NoHooks bool           `json:"no_hooks"`
	Options *ArmadaOptions `json:"options,omitempty"`
	Post    *ArmadaPost    `json:"post,omitempty"`
	Pre     *ArmadaPre     `json:"pre,omitempty"`
}

// Pre
type ArmadaPre struct {
	Create []*ArmadaHookActionItems `json:"create,omitempty"`
	Delete []*ArmadaHookActionItems `json:"delete,omitempty"`
	Update []*ArmadaHookActionItems `json:"update,omitempty"`
}

// Post
type ArmadaPost struct {
	Create []*ArmadaHookActionItems `json:"create,omitempty"`
}

// Protected
type ArmadaProtected struct {
	ContinueProcessing bool `json:"continue_processing,omitempty"`
}

// Values
type ArmadaValues struct {
}

// Source
type ArmadaSource struct {
	AuthMethod  string `json:"auth_method,omitempty"`
	Location    string `json:"location"`
	ProxyServer string `json:"proxy_server,omitempty"`
	Reference   string `json:"reference,omitempty"`
	Subpath     string `json:"subpath"`
	Type        string `json:"type"`
}

// Root
type ArmadaChart struct {
	ChartName    string        `json:"chart_name"`
	Namespace    string        `json:"namespace"`
	Release      string        `json:"release"`
	Source       *ArmadaSource `json:"source"`
	Dependencies []string      `json:"dependencies"`

	Install *ArmadaInstall `json:"install,omitempty"`
	Delete  *ArmadaDelete  `json:"delete,omitempty"`
	Upgrade *ArmadaUpgrade `json:"upgrade,omitempty"`
	Values  *ArmadaValues  `json:"values,omitempty"`

	Protected *ArmadaProtected `json:"protected,omitempty"`
	//JEB Test         interface{} `json:"test,omitempty"`
	Timeout int         `json:"timeout,omitempty"`
	Wait    *ArmadaWait `json:"wait,omitempty"`
}

type ArmadaChartGroup struct {
	ChartGroup  []string `json:"chart_group"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Sequenced   bool     `json:"sequenced,omitempty"`
	TestCharts  bool     `json:"test_charts,omitempty"`
}

// ArmadaManifestSpec defines the desired state of ArmadaManifest
type ArmadaManifestSpec struct {
	ChartGroups   []string `json:"chart_groups"`
	ReleasePrefix string   `json:"release_prefix"`
}

// ArmadaManifestStatus defines the observed state of ArmadaManifest
type ArmadaManifestStatus struct {
	// Succeeded indicates if the release is in the expected state
	Succeeded bool `json:"succeeded"`
	// Reason indicates the reason for any related failures.
	Reason string `json:"Reason,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaManifest is the Schema for the armadamanifests API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=armadamanifests,shortName=amf
// +kubebuilder:printcolumn:name="succeeded",type="boolean",JSONPath=".status.succeeded",description="success",priority=1
type ArmadaManifest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArmadaManifestSpec   `json:"spec,omitempty"`
	Status ArmadaManifestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArmadaManifestList contains a list of ArmadaManifest
type ArmadaManifestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArmadaManifest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArmadaManifest{}, &ArmadaManifestList{})
}
