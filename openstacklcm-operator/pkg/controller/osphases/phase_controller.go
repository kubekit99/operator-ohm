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

import (
	"reflect"
	"time"

	services "github.com/kubekit99/operator-ohm/openstacklcm-operator/pkg/services"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	crtpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var phaselog = logf.Log.WithName("phase-controller")

// PhaseReconciler reconciles custom resources as Workflow, Jobs....
type PhaseReconciler struct {
	client                  client.Client
	scheme                  *runtime.Scheme
	recorder                record.EventRecorder
	managerFactory          services.PhaseManagerFactory
	reconcilePeriod         time.Duration
	depResourceWatchUpdater services.DependentResourceWatchUpdater
}

func (r *PhaseReconciler) contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// buildDependentPredicate create the predicates used by subresources watches
func (r *PhaseReconciler) BuildDependentPredicate() *crtpredicate.Funcs {

	dependentPredicate := crtpredicate.Funcs{
		// We don't need to reconcile dependent resource creation events
		// because dependent resources are only ever created during
		// reconciliation. Another reconcile would be redundant.
		CreateFunc: func(e event.CreateEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			phaselog.Info("CreateEvent. Filtering", "resource", o.GetName(), "namespace", o.GetNamespace(),
				"apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return false
		},

		// Reconcile when a dependent resource is deleted so that it can be
		// recreated.
		DeleteFunc: func(e event.DeleteEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			phaselog.Info("DeleteEvent. Triggering", "resource", o.GetName(), "namespace", o.GetNamespace(),
				"apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return true
		},

		// Reconcile when a dependent resource is updated, so that it can
		// be patched back to the resource managed by the Argo workflow, if
		// necessary. Ignore updates that only change the status and
		// resourceVersion.
		UpdateFunc: func(e event.UpdateEvent) bool {
			old := e.ObjectOld.(*unstructured.Unstructured).DeepCopy()
			new := e.ObjectNew.(*unstructured.Unstructured).DeepCopy()

			delete(old.Object, "status")
			delete(new.Object, "status")
			old.SetResourceVersion("")
			new.SetResourceVersion("")

			if reflect.DeepEqual(old.Object, new.Object) {
				phaselog.Info("UpdateEvent. Filtering", "resource", new.GetName(), "namespace", new.GetNamespace(),
					"apiVersion", new.GroupVersionKind().GroupVersion(), "kind", new.GroupVersionKind().Kind)
				return false
			} else {
				phaselog.Info("UpdateEvent. Triggering", "resource", new.GetName(), "namespace", new.GetNamespace(),
					"apiVersion", new.GroupVersionKind().GroupVersion(), "kind", new.GroupVersionKind().Kind)
				return true
			}
		},
	}

	return &dependentPredicate
}
