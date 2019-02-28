package v1alpha1

import (
// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
// rpb "k8s.io/helm/pkg/proto/hapi/release"
)

// Delete
type ArmadaDelete struct {
	Timeout int `json:"timeout,omitempty"`
}

// HookActionItems
type ArmadaHookActionItems struct {
	Labels *ArmadaLabels `json:"labels,omitempty"`
	Name   string        `json:"name,omitempty"`
	Type   string        `json:"type"`
}

// Install
type ArmadaInstall struct {
}

// Labels
type ArmadaLabels struct {
	AdditionalProperties map[string]string `json:"-,omitempty"`
}

// Native
type ArmadaNative struct {
	Enabled bool `json:"enabled,omitempty"`
}

// Options
type ArmadaOptions struct {
	Force        bool `json:"force,omitempty"`
	RecreatePods bool `json:"recreate_pods,omitempty"`
}

// Post
type ArmadaPost struct {
	Create []*ArmadaHookActionItems `json:"create,omitempty"`
}

// Pre
type ArmadaPre struct {
	Create []*ArmadaHookActionItems `json:"create,omitempty"`
	Delete []*ArmadaHookActionItems `json:"delete,omitempty"`
	Update []*ArmadaHookActionItems `json:"update,omitempty"`
}

// Protected
type ArmadaProtected struct {
	ContinueProcessing bool `json:"continue_processing,omitempty"`
}

// ResourcesItems
type ArmadaResourcesItems struct {
	Labels *ArmadaLabels `json:"labels,omitempty"`
	//JEB MinReady interface{}   `json:"min_ready,omitempty"`
	Type string `json:"type"`
}

// Root
type ArmadaArmadaChart struct {
	ChartName    string           `json:"chart_name"`
	Delete       *ArmadaDelete    `json:"delete,omitempty"`
	Dependencies []string         `json:"dependencies"`
	Install      *ArmadaInstall   `json:"install,omitempty"`
	Namespace    string           `json:"namespace"`
	Protected    *ArmadaProtected `json:"protected,omitempty"`
	Release      string           `json:"release"`
	Source       *ArmadaSource    `json:"source"`
	//JEB Test         interface{} `json:"test,omitempty"`
	Timeout int            `json:"timeout,omitempty"`
	Upgrade *ArmadaUpgrade `json:"upgrade,omitempty"`
	Values  *ArmadaValues  `json:"values,omitempty"`
	Wait    *ArmadaWait    `json:"wait,omitempty"`
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

// Upgrade
type ArmadaUpgrade struct {
	NoHooks bool           `json:"no_hooks"`
	Options *ArmadaOptions `json:"options,omitempty"`
	Post    *ArmadaPost    `json:"post,omitempty"`
	Pre     *ArmadaPre     `json:"pre,omitempty"`
}

// Values
type ArmadaValues struct {
}

// Wait
type ArmadaWait struct {
	Labels    *ArmadaLabels           `json:"labels,omitempty"`
	Native    *ArmadaNative           `json:"native,omitempty"`
	Resources []*ArmadaResourcesItems `json:"resources,omitempty"`
	Timeout   int                     `json:"timeout,omitempty"`
}

type ArmadaChartGroup struct {
	ChartGroup  []string `json:"chart_group"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Sequenced   bool     `json:"sequenced,omitempty"`
	TestCharts  bool     `json:"test_charts,omitempty"`
}

type ArmadaManifest struct {
	ChartGroups   []string `json:"chart_groups"`
	ReleasePrefix string   `json:"release_prefix"`
}
