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

// JEB: This file has been at first generated from the file presents in
// https://github.com/openstack/airship-armada/tree/master/armada/schemas
// and then through yaml2json tools followed by a call to schema-generate
// This file will be deleted once we figure out what we really want
// to put in our CRDs.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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
	Labels   *ArmadaLabels             `json:"labels,omitempty"`
	MinReady unstructured.Unstructured `json:"min_ready,omitempty"`
	Type     string                    `json:"type"`
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

	Protected *ArmadaProtected          `json:"protected,omitempty"`
	Test      unstructured.Unstructured `json:"test,omitempty"`
	Timeout   int                       `json:"timeout,omitempty"`
	Wait      *ArmadaWait               `json:"wait,omitempty"`
}

type ArmadaChartGroup struct {
	ChartGroup  []string `json:"chart_group"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Sequenced   bool     `json:"sequenced,omitempty"`
	TestCharts  bool     `json:"test_charts,omitempty"`
}
