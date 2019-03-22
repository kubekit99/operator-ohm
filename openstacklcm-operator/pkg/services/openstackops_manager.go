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

package services

import (
	"context"
	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
)

// OpenstackBackupManager manages the backup of an OpenstackService
type OpenstackBackupManager interface {
	ResourceName() string
	IsInstallRequired() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OpenstackBackup, error)
	UpdateResource(context.Context) (*av1.OpenstackBackup, *av1.OpenstackBackup, error)
	ReconcileResource(context.Context) (*av1.OpenstackBackup, error)
	UninstallResource(context.Context) (*av1.OpenstackBackup, error)
}

// OpenstackRestoreManager manages the restore of an OpenstackService
type OpenstackRestoreManager interface {
	ResourceName() string
	IsInstallRequired() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OpenstackRestore, error)
	UpdateResource(context.Context) (*av1.OpenstackRestore, *av1.OpenstackRestore, error)
	ReconcileResource(context.Context) (*av1.OpenstackRestore, error)
	UninstallResource(context.Context) (*av1.OpenstackRestore, error)
}

// OpenstackDeploymentManager manages the deployment of an OpenstackService
type OpenstackDeploymentManager interface {
	ResourceName() string
	IsInstallRequired() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OpenstackDeployment, error)
	UpdateResource(context.Context) (*av1.OpenstackDeployment, *av1.OpenstackDeployment, error)
	ReconcileResource(context.Context) (*av1.OpenstackDeployment, error)
	UninstallResource(context.Context) (*av1.OpenstackDeployment, error)
}

// OpenstackRollbackManager manages the rollback of an OpenstackService
type OpenstackRollbackManager interface {
	ResourceName() string
	IsInstallRequired() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OpenstackRollback, error)
	UpdateResource(context.Context) (*av1.OpenstackRollback, *av1.OpenstackRollback, error)
	ReconcileResource(context.Context) (*av1.OpenstackRollback, error)
	UninstallResource(context.Context) (*av1.OpenstackRollback, error)
}

// OpenstackUpgradeManager manages the upgrade of an OpenstackService
type OpenstackUpgradeManager interface {
	ResourceName() string
	IsInstallRequired() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OpenstackUpgrade, error)
	UpdateResource(context.Context) (*av1.OpenstackUpgrade, *av1.OpenstackUpgrade, error)
	ReconcileResource(context.Context) (*av1.OpenstackUpgrade, error)
	UninstallResource(context.Context) (*av1.OpenstackUpgrade, error)
}
