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

package helm

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	av1 "github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1"
	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
func NewManagerFactory(mgr manager.Manager) helmif.HelmManagerFactory {
	// Create Tiller's storage backend and kubernetes client
	storageBackend := storage.Init(driver.NewMemory())
	tillerKubeClient, err := NewFromManager(mgr)
	if err != nil {
		log.Error(err, "Failed to create new Tiller client.", storageBackend, tillerKubeClient)
		os.Exit(1)
	}

	return &managerFactory{storageBackend, tillerKubeClient}
}

func (f managerFactory) NewHelmReleaseTillerManager(r *av1.HelmRelease) helmif.HelmManager {
	return &helmreleasemanager{
		storageBackend:   f.storageBackend,
		tillerKubeClient: f.tillerKubeClient,
		chartDir:         r.Spec.ChartDir,

		tiller:      f.tillerRendererForHelmRelease(r),
		releaseName: r.Spec.ReleaseName,
		namespace:   r.GetNamespace(),

		spec:   r.Spec,
		status: &r.Status,
	}
}

func (f managerFactory) NewArmadaChartTillerManager(r *av1.ArmadaChart) helmif.HelmManager {
	return &chartmanager{
		storageBackend:   f.storageBackend,
		tillerKubeClient: f.tillerKubeClient,
		chartLocation:    r.Spec.Source,

		tiller:      f.tillerRendererForArmadaChart(r),
		releaseName: r.Spec.Release,
		namespace:   r.GetNamespace(),

		spec:   r.Spec,
		status: &r.Status,
	}
}

func (f managerFactory) NewHelm3Manager(r *av1.HelmRelease) helmif.HelmManager {
	return &helm3manager{}
}

// tillerRendererForCR creates a ReleaseServer configured with a rendering engine that adds ownerrefs to rendered assets
// based on the CR.
func (f managerFactory) tillerRendererForHelmRelease(r *av1.HelmRelease) *tiller.ReleaseServer {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}
	baseEngine := helmengine.New()
	e := helmif.NewOwnerRefEngine(baseEngine, ownerRefs)
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

// tillerRendererForCR creates a ReleaseServer configured with a rendering engine that adds ownerrefs to rendered assets
// based on the CR.
func (f managerFactory) tillerRendererForArmadaChart(r *av1.ArmadaChart) *tiller.ReleaseServer {
	controllerRef := metav1.NewControllerRef(r, r.GroupVersionKind())
	ownerRefs := []metav1.OwnerReference{
		*controllerRef,
	}
	baseEngine := helmengine.New()
	e := helmif.NewOwnerRefEngine(baseEngine, ownerRefs)
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
