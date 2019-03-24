// Copyright 2019 The Openstack-Service-Lifecyle Authors
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

package osphases

import ()

type phasemanager struct {
	renderer     interface{}
	resourceName string
	namespace    string

	isInstalled      bool
	isUpdateRequired bool
	config           *map[string]interface{}
}

// ResourceName returns the name of the release.
func (m phasemanager) ResourceName() string {
	return m.resourceName
}

func (m phasemanager) IsInstalled() bool {
	return m.isInstalled
}

func (m phasemanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}
