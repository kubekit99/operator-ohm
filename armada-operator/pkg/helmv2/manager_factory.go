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

package helmv2

import (
	"fmt"
	"os"
	"strings"

	"github.com/martinlindhe/base36"
	"github.com/pborman/uuid"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	oshv1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/helmif"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apitypes "k8s.io/apimachinery/pkg/types"
	clientset "k8s.io/client-go/kubernetes"
	helmengine "k8s.io/helm/pkg/engine"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/storage"
	"k8s.io/helm/pkg/storage/driver"
	"k8s.io/helm/pkg/tiller"
	"k8s.io/helm/pkg/tiller/environment"
)

type managerFactory struct {
	storageBackend   *storage.Storage
	tillerKubeClient *kube.Client
}

// NewManagerFactory returns a new Helm manager factory capable of installing and uninstalling releases.
func NewManagerFactory(mgr manager.Manager) helmif.ManagerFactory {
	// Create Tiller's storage backend and kubernetes client
	storageBackend := storage.Init(driver.NewMemory())
	tillerKubeClient, err := NewFromManager(mgr)
	if err != nil {
		log.Error(err, "Failed to create new Tiller client.", storageBackend, tillerKubeClient)
		os.Exit(1)
	}

	return &managerFactory{storageBackend, tillerKubeClient}
}

func (f managerFactory) NewManager(r *oshv1.HelmRelease) helmif.Manager {
	return f.newManagerForCR(r)
}

func (f managerFactory) newManagerForCR(r *oshv1.HelmRelease) helmif.Manager {
	return &helmv2manager{
		storageBackend:   f.storageBackend,
		tillerKubeClient: f.tillerKubeClient,
		chartDir:         getChartDir(r),

		tiller:      f.tillerRendererForCR(r),
		releaseName: getReleaseName(r),
		namespace:   r.GetNamespace(),

		spec:   r.Spec,
		status: &r.Status,
	}
}

// tillerRendererForCR creates a ReleaseServer configured with a rendering engine that adds ownerrefs to rendered assets
// based on the CR.
func (f managerFactory) tillerRendererForCR(r *oshv1.HelmRelease) *tiller.ReleaseServer {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}
	baseEngine := helmengine.New()
	e := NewOwnerRefEngine(baseEngine, ownerRefs)
	var ey environment.EngineYard = map[string]environment.Engine{
		environment.GoTplEngine: e,
	}
	env := &environment.Environment{
		EngineYard: ey,
		Releases:   f.storageBackend,
		KubeClient: f.tillerKubeClient,
	}
	kubeconfig, _ := f.tillerKubeClient.ToRESTConfig()
	cs := clientset.NewForConfigOrDie(kubeconfig)

	return tiller.NewReleaseServer(env, cs, false)
}

func getChartDir(r *oshv1.HelmRelease) string {
	if r.Spec.ChartDir != "" {
		// JEB: We should check for duplicates here as well as syntax of ReleaseName
		return r.Spec.ChartDir
	} else {
		return fmt.Sprintf("%s-%s", r.GetName(), shortenUID(r.GetUID()))
	}
}

func getReleaseName(r *oshv1.HelmRelease) string {
	if r.Spec.ReleaseName != "" {
		// JEB: We should check for duplicates here as well as syntax of ReleaseName
		return r.Spec.ReleaseName
	} else {
		return fmt.Sprintf("%s-%s", r.GetName(), shortenUID(r.GetUID()))
	}
}

func shortenUID(uid apitypes.UID) string {
	u := uuid.Parse(string(uid))
	uidBytes, err := u.MarshalBinary()
	if err != nil {
		return strings.Replace(string(uid), "-", "", -1)
	}
	return strings.ToLower(base36.EncodeBytes(uidBytes))
}
