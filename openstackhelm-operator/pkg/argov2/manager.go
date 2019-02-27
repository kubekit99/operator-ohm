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

package argov2

import (
	"context"
	"strings"

	wfv1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	oshv1 "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/apis/openstackhelm/v1alpha1"
	_ "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/argoif"

	_ "k8s.io/apimachinery/pkg/api/errors"
	_ "k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type argov2manager struct {
	argoKubeClient *client.Client
	chartDir       string

	workflowName string
	namespace    string

	spec   interface{}
	status *oshv1.OpenstackChartStatus

	isInstalled      bool
	isUpdateRequired bool
	deployedWorkflow *wfv1.Workflow
}

// WorkflowName returns the name of the release.
func (m argov2manager) WorkflowName() string {
	return m.workflowName
}

func (m argov2manager) IsInstalled() bool {
	return m.isInstalled
}

func (m argov2manager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync ensures the Helm storage backend is in sync with the status of the
// custom resource.
func (m *argov2manager) Sync(ctx context.Context) error {
	return nil
}

func (m argov2manager) syncWorkflowStatus(status oshv1.OpenstackChartStatus) error {
	var release *wfv1.Workflow

	if release == nil {
		return nil
	}

	return nil
}

func notFoundErr(err error) bool {
	return strings.Contains(err.Error(), "not found")
}

func (m argov2manager) loadChartAndConfig() error {
	return nil
}

func (m argov2manager) getDeployedWorkflow() (*wfv1.Workflow, error) {
	return nil, nil
}

// InstallWorkflow performs a Helm release install.
func (m argov2manager) InstallWorkflow(ctx context.Context) (*wfv1.Workflow, error) {
	return nil, nil
}

// UpdateWorkflow performs a Helm release update.
func (m argov2manager) UpdateWorkflow(ctx context.Context) (*wfv1.Workflow, *wfv1.Workflow, error) {
	return nil, nil, nil
}

// ReconcileWorkflow creates or patches resources as necessary to match the
// deployed release's manifest.
func (m argov2manager) ReconcileWorkflow(ctx context.Context) (*wfv1.Workflow, error) {
	return nil, nil
}

// UninstallWorkflow performs a Helm release uninstall.
func (m argov2manager) UninstallWorkflow(ctx context.Context) (*wfv1.Workflow, error) {
	// reqLogger := log.WithValues("Worflow.Namespace", m.namespace, "Worflow.Name", m.namespace)

	// found := argoif.NewWorkflowGroupVersionKind()
	// err := m.argoKubeClient.Get(context.TODO(), types.NamespacedName{Name: m.workflowName, Namespace: m.namespace}, found)
	// if err != nil && errors.IsNotFound(err) {
	// 	// Workflow was already deleted
	// 	return found, nil
	// } else if err != nil {
	// 	reqLogger.Error(err, "Failure to fetch workflow . Ignoring")
	// 	// Something else wrong with workflow
	// 	return found, nil
	// }

	// err = m.argoKubeClient.Delete(context.TODO(), found)
	// if err != nil {
	// 	reqLogger.Error(err, "Failure to delete workflow . Ignoring")
	// 	// Something wrong with workflow deletion
	// 	return found, nil
	// }

	// return found, nil

	return nil, nil
}
