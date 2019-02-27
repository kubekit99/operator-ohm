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
	"fmt"
	"os"
	"strings"

	"github.com/martinlindhe/base36"
	"github.com/pborman/uuid"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	oshv1 "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/apis/openstackhelm/v1alpha1"
	argoif "github.com/kubekit99/operator-ohm/openstackhelm-operator/pkg/argoif"

	apitypes "k8s.io/apimachinery/pkg/types"
)

type managerFactory struct {
	argoKubeClient *client.Client
}

// NewManagerFactory returns a new Helm manager factory capable of installing and uninstalling releases.
func NewManagerFactory(mgr manager.Manager) argoif.ManagerFactory {
	// Create Argo KubeClient.
	// We currently do not use the argocli/client hence the client is
	// still the normal kubeclient. This implementation relies on CR and CRD to control argo
	// behavior
	argoKubeClient, err := NewFromManager(mgr)
	if err != nil {
		log.Error(err, "Failed to create new Kube client.", argoKubeClient)
		os.Exit(1)
	}

	return &managerFactory{argoKubeClient}
}

func (f managerFactory) NewManager(r *oshv1.OpenstackChart) argoif.Manager {
	return f.newManagerForCR(r)
}

func (f managerFactory) newManagerForCR(r *oshv1.OpenstackChart) argoif.Manager {
	return &argov2manager{
		argoKubeClient: f.argoKubeClient,
		chartDir:       getChartDir(r),

		workflowName: getWorkflowName(r),
		namespace:    r.GetNamespace(),

		spec:   r.Spec,
		status: oshv1.StatusFor(r),
	}
}

func getChartDir(r *oshv1.OpenstackChart) string {
	if r.Spec.ChartDir != "" {
		// JEB: We should check for duplicates here as well as syntax of ReleaseName
		return r.Spec.ChartDir
	} else {
		return fmt.Sprintf("%s-%s", r.GetName(), shortenUID(r.GetUID()))
	}
}

func getWorkflowName(r *oshv1.OpenstackChart) string {
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
