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

package armada

import (
	"context"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	armadaif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var acglog = logf.Log.WithName("acg-manager")

type chartgroupmanager struct {
	kubeClient       client.Client
	resourceName     string
	namespace        string
	spec             *av1.ArmadaChartGroupSpec
	status           *av1.ArmadaChartGroupStatus
	deployedResource *av1.ArmadaCharts
	isUpdateRequired bool
}

// ResourceName returns the name of the release.
func (m chartgroupmanager) ResourceName() string {
	return m.resourceName
}

func (m chartgroupmanager) IsUpdateRequired() bool {
	return m.isUpdateRequired
}

// Sync detects which ArmadaCharts are already present for that ArmadaChartGroup
// to proceed. The ArmadaChartGroup should not proceed until all the Charts
// are present in the system
func (m *chartgroupmanager) Sync(ctx context.Context) error {
	m.deployedResource = av1.NewArmadaCharts(m.resourceName)
	errs := make([]error, 0)
	targetResourceList := m.expectedChartList()
	for _, existingResource := range (*targetResourceList).List.Items {
		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: existingResource.Name, Namespace: existingResource.Namespace}, &existingResource)
		if err != nil {
			if !apierrors.IsNotFound(err) {
				// Don't want to trace is the error is not a NotFound.
				acglog.Error(err, "Can't not retrieve ArmadaChart")
			}
			errs = append(errs, err)
		} else {
			m.deployedResource.List.Items = append(m.deployedResource.List.Items, existingResource)
		}
	}

	// The ChartGroup manager is not in charge of creating the ArmaChart since it
	// only contains the name of the charts.
	if len(errs) != 0 {
		// Regardless if the error is NotFound or something else,
		// we can't sync the ArmadaChartGroup with content of Kubernetes.
		m.isUpdateRequired = false
		return errs[0]
	}

	// TODO(jeb): We should check here the "admin_state" of the ArmadaChartGroup compared
	// it to the "admin_state" of the ArmadaCharts
	// TODO(jeb): We should check that the ArmadaChartGroup is still not the "owner" of
	// chartgroups which are not listed in its Spec anymore. In such as case we should put
	// the isUpdateRequired to true.
	m.isUpdateRequired = false
	return nil
}

// UpdateResource performs an update of an ArmadaChartGroup.
// Currently either the list of Charts or the Sequenced attribute may have change.
func (m chartgroupmanager) UpdateResource(ctx context.Context) (*av1.ArmadaCharts, *av1.ArmadaCharts, error) {
	errs := make([]error, 0)
	toUpdateList := m.expectedChartList()
	for _, toUpdate := range (*toUpdateList).List.Items {
		err := m.kubeClient.Update(context.TODO(), &toUpdate)
		if err != nil {
			acglog.Error(err, "Can't not Update ArmadaChart")
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return nil, nil, armadaif.ErrNotFound
		} else {
			return nil, nil, errs[0]
		}
	}
	return m.deployedResource, toUpdateList, nil
}

// ReconcileResource creates or patches resources as necessary to match the
// deployed release's manifest.
func (m chartgroupmanager) ReconcileResource(ctx context.Context) (*av1.ArmadaCharts, error) {

	nextToEnable := m.deployedResource.GetNextToEnable()
	if nextToEnable != nil {
		found := nextToEnable.FromArmadaChart()
		err := m.kubeClient.Get(context.TODO(), types.NamespacedName{Name: found.GetName(), Namespace: found.GetNamespace()}, nextToEnable)
		if err == nil {
			nextToEnable.Spec.AdminState = av1.StateEnabled
			if err2 := m.kubeClient.Update(context.TODO(), nextToEnable); err2 != nil {
				acglog.Error(err, "Can't get enable of ArmadaChart", "name", found.GetName())
				return m.deployedResource, err
			}
			acglog.Info("Enabled ArmadaChart", "name", found.GetName())
		} else {
			acglog.Error(err, "Can't enable ArmadaChart", "name", found.GetName())
			return m.deployedResource, err
		}
	}

	return m.deployedResource, nil
}

// UninstallResource currently delete Charts matching the ArmadaChartGroups.
// This is probably not the behavior we want to maitain in the long run.
func (m chartgroupmanager) UninstallResource(ctx context.Context) (*av1.ArmadaCharts, error) {
	errs := make([]error, 0)
	toDeleteList := m.expectedChartList()
	for _, toDelete := range (*toDeleteList).List.Items {
		err := m.kubeClient.Delete(context.TODO(), &toDelete)
		if err != nil {
			acglog.Error(err, "Can't not Delete ArmadaChart")
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		if apierrors.IsNotFound(errs[0]) {
			return nil, armadaif.ErrNotFound
		} else {
			return nil, errs[0]
		}
	}
	return toDeleteList, nil
}

// expectedChartList returns a dummy ArmadaChart the same name/namespace as the cr
func (m chartgroupmanager) expectedChartList() *av1.ArmadaCharts {
	labels := map[string]string{
		"app": m.resourceName,
	}

	var res = av1.NewArmadaCharts(m.resourceName)

	for _, chartname := range m.spec.Charts {
		res.List.Items = append(res.List.Items, av1.ArmadaChart{
			ObjectMeta: metav1.ObjectMeta{
				Name:      chartname,
				Namespace: m.namespace,
				Labels:    labels,
			},
			Spec: av1.ArmadaChartSpec{
				ChartName: chartname,
				Release:   m.resourceName + "-release",
				Namespace: m.namespace,
				Upgrade: &av1.ArmadaUpgrade{
					NoHooks: false,
				},
				Source: &av1.ArmadaChartSource{
					Type:      "local",
					Location:  "/opt/armada/helm-charts/testchart",
					Subpath:   ".",
					Reference: "master",
				},
				Dependencies: make([]string, 0),
				TargetState:  av1.StateInitialized,
				AdminState:   av1.StateDisabled,
			},
		})
	}

	return res
}
