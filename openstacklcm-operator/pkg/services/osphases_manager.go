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

package services

import (
	"context"
	av1 "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/apis/openstacklcm/v1alpha1"
)

// PlanningPhaseManager manages the Planning Phase of an OpenstackServiceLifeCycle
type PlanningPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.PlanningPhase, error)
	UpdateResource(context.Context) (*av1.PlanningPhase, *av1.PlanningPhase, error)
	ReconcileResource(context.Context) (*av1.PlanningPhase, error)
	UninstallResource(context.Context) (*av1.PlanningPhase, error)
}

// InstallPhaseManager manages the Install Phase of an OpenstackServiceLifeCycle
type InstallPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.InstallPhase, error)
	UpdateResource(context.Context) (*av1.InstallPhase, *av1.InstallPhase, error)
	ReconcileResource(context.Context) (*av1.InstallPhase, error)
	UninstallResource(context.Context) (*av1.InstallPhase, error)
}

// TestPhaseManager manages the Test Phase of an OpenstackServiceLifeCycle
type TestPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.TestPhase, error)
	UpdateResource(context.Context) (*av1.TestPhase, *av1.TestPhase, error)
	ReconcileResource(context.Context) (*av1.TestPhase, error)
	UninstallResource(context.Context) (*av1.TestPhase, error)
}

// TrafficRolloutPhaseManager manages the TrafficRollout Phase of an OpenstackServiceLifeCycle
type TrafficRolloutPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.TrafficRolloutPhase, error)
	UpdateResource(context.Context) (*av1.TrafficRolloutPhase, *av1.TrafficRolloutPhase, error)
	ReconcileResource(context.Context) (*av1.TrafficRolloutPhase, error)
	UninstallResource(context.Context) (*av1.TrafficRolloutPhase, error)
}

// OperationalPhaseManager manages the Operational Phase of an OpenstackServiceLifeCycle
type OperationalPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.OperationalPhase, error)
	UpdateResource(context.Context) (*av1.OperationalPhase, *av1.OperationalPhase, error)
	ReconcileResource(context.Context) (*av1.OperationalPhase, error)
	UninstallResource(context.Context) (*av1.OperationalPhase, error)
}

// TrafficDrainPhaseManager manages the TrafficDrain Phase of an OpenstackServiceLifeCycle
type TrafficDrainPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.TrafficDrainPhase, error)
	UpdateResource(context.Context) (*av1.TrafficDrainPhase, *av1.TrafficDrainPhase, error)
	ReconcileResource(context.Context) (*av1.TrafficDrainPhase, error)
	UninstallResource(context.Context) (*av1.TrafficDrainPhase, error)
}

// UpgradePhaseManager manages the Upgrade Phase of an OpenstackServiceLifeCycle
type UpgradePhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.UpgradePhase, error)
	UpdateResource(context.Context) (*av1.UpgradePhase, *av1.UpgradePhase, error)
	ReconcileResource(context.Context) (*av1.UpgradePhase, error)
	UninstallResource(context.Context) (*av1.UpgradePhase, error)
}

// RollbackPhaseManager manages the Rollback Phase of an OpenstackServiceLifeCycle
type RollbackPhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.RollbackPhase, error)
	UpdateResource(context.Context) (*av1.RollbackPhase, *av1.RollbackPhase, error)
	ReconcileResource(context.Context) (*av1.RollbackPhase, error)
	UninstallResource(context.Context) (*av1.RollbackPhase, error)
}

// DeletePhaseManager manages the Delete Phase of an OpenstackServiceLifeCycle
type DeletePhaseManager interface {
	ResourceName() string
	IsInstalled() bool
	IsUpdateRequired() bool
	Sync(context.Context) error
	InstallResource(context.Context) (*av1.DeletePhase, error)
	UpdateResource(context.Context) (*av1.DeletePhase, *av1.DeletePhase, error)
	ReconcileResource(context.Context) (*av1.DeletePhase, error)
	UninstallResource(context.Context) (*av1.DeletePhase, error)
}
