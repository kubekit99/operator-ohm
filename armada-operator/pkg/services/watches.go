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
	"reflect"
	"sync"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/source"

	crthandler "sigs.k8s.io/controller-runtime/pkg/handler"
	crtpredicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

type DependentResourceWatchUpdater func([]unstructured.Unstructured) error

// BuildDependentResourcesWatchUpdater builds a function that adds watches for resources in released Helm charts.
func BuildDependentResourceWatchUpdater(mgr manager.Manager, owner *unstructured.Unstructured, c controller.Controller) DependentResourceWatchUpdater {

	dependentPredicate := crtpredicate.Funcs{
		// We don't need to reconcile dependent resource creation events
		// because dependent resources are only ever created during
		// reconciliation. Another reconcile would be redundant.
		CreateFunc: func(e event.CreateEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			log.Info("CreateEvent. Skipping", "owner", owner.GetName(), "resource", o.GetName(), "namespace", o.GetNamespace(), "apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return false
		},

		// Reconcile when a dependent resource is deleted so that it can be
		// recreated.
		DeleteFunc: func(e event.DeleteEvent) bool {
			o := e.Object.(*unstructured.Unstructured)
			log.Info("DeleteEvent. Reconciling", "owner", owner.GetName(), "resource", o.GetName(), "namespace", o.GetNamespace(), "apiVersion", o.GroupVersionKind().GroupVersion(), "kind", o.GroupVersionKind().Kind)
			return true
		},

		// Reconcile when a dependent resource is updated, so that it can
		// be patched back to the resource managed by the Helm release, if
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
				return false
			}
			log.Info("UpdateEvent. Reconciling", "owner", owner.GetName(), "resource", new.GetName(), "namespace", new.GetNamespace(), "apiVersion", new.GroupVersionKind().GroupVersion(), "kind", new.GroupVersionKind().Kind)
			return true
		},
	}

	var m sync.RWMutex
	watches := map[schema.GroupVersionKind]struct{}{}
	watchUpdater := func(dependent []unstructured.Unstructured) error {
		for _, u := range dependent {
			gvk := u.GroupVersionKind()
			wlog := log.WithValues("owner", owner.GetName(), "resourceType", gvk.GroupVersion(), "resourceKind", gvk.Kind)
			m.RLock()
			_, ok := watches[gvk]
			m.RUnlock()
			if ok {
				continue
			}

			restMapper := mgr.GetRESTMapper()
			depMapping, err := restMapper.RESTMapping(gvk.GroupKind(), gvk.Version)
			if err != nil {
				wlog.Error(err, "GetRESTMapper")
				return err
			}
			ownerMapping, err := restMapper.RESTMapping(owner.GroupVersionKind().GroupKind(), owner.GroupVersionKind().Version)
			if err != nil {
				wlog.Error(err, "Build RESTMapping")
				return err
			}

			depClusterScoped := depMapping.Scope.Name() == meta.RESTScopeNameRoot
			ownerClusterScoped := ownerMapping.Scope.Name() == meta.RESTScopeNameRoot

			if !ownerClusterScoped && depClusterScoped {
				m.Lock()
				watches[gvk] = struct{}{}
				m.Unlock()
				wlog.Info("Cannot watch cluster-scoped")
				continue
			}

			err = c.Watch(&source.Kind{Type: &u}, &crthandler.EnqueueRequestForOwner{OwnerType: owner}, dependentPredicate)
			if err != nil {
				wlog.Error(err, "Add Watch to Controller")
				return err
			}

			m.Lock()
			watches[gvk] = struct{}{}
			m.Unlock()
			wlog.Info("Watching dependent resource")
		}

		return nil
	}

	return watchUpdater
}
