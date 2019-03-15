// Copyright 2018 The Operator-SDK Authors
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

package services

import (
	"context"
)

// Manager manages a Helm release. It can install, update, reconcile,
// and uninstall a release.
type HelmManager interface {
	ReleaseName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallRelease(context.Context) (*HelmRelease, error)
	UpdateRelease(context.Context) (*HelmRelease, *HelmRelease, error)
	ReconcileRelease(context.Context) (*HelmRelease, error)
	UninstallRelease(context.Context) (*HelmRelease, error)
}
