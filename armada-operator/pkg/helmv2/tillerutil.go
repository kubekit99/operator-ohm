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

// +build v2

package helmv2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	helmif "github.com/kubekit99/operator-ohm/armada-operator/pkg/services"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apitypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions/resource"
	"k8s.io/client-go/rest"
	"k8s.io/helm/pkg/chartutil"
	"k8s.io/helm/pkg/kube"
	"k8s.io/helm/pkg/storage"
	cpb "k8s.io/helm/pkg/proto/hapi/chart"
	rpb "k8s.io/helm/pkg/proto/hapi/release"
	svc "k8s.io/helm/pkg/proto/hapi/services"
	"k8s.io/helm/pkg/tiller"

	"github.com/mattbaird/jsonpatch"
)

func notFoundErr(err error) bool {
	return strings.Contains(err.Error(), "not found")
}

// processRequirements will process the requirements file
// It will disable/enable the charts based on condition in requirements file
// Also imports the specified chart values from child to parent.
func processRequirements(chart *cpb.Chart, values *cpb.Config) error {
	err := chartutil.ProcessRequirementsEnabled(chart, values)
	if err != nil {
		return err
	}
	err = chartutil.ProcessRequirementsImportValues(chart)
	if err != nil {
		return err
	}
	return nil
}

func installRelease(ctx context.Context, releaseServer *tiller.ReleaseServer, namespace, name string, chart *cpb.Chart, config *cpb.Config) (*rpb.Release, error) {
	installReq := &svc.InstallReleaseRequest{
		Namespace: namespace,
		Name:      name,
		Chart:     chart,
		Values:    config,
	}

	releaseResponse, err := releaseServer.InstallRelease(ctx, installReq)
	if err != nil {
		// Workaround for helm/helm#3338
		if releaseResponse.GetRelease() != nil {
			uninstallReq := &svc.UninstallReleaseRequest{
				Name:  releaseResponse.GetRelease().GetName(),
				Purge: true,
			}
			_, uninstallErr := releaseServer.UninstallRelease(ctx, uninstallReq)
			if uninstallErr != nil {
				return nil, fmt.Errorf("failed to roll back failed installation: %s: %s", uninstallErr, err)
			}
		}
		return nil, err
	}
	return releaseResponse.GetRelease(), nil
}

func updateRelease(ctx context.Context, releaseServer *tiller.ReleaseServer, name string, chart *cpb.Chart, config *cpb.Config) (*rpb.Release, error) {
	updateReq := &svc.UpdateReleaseRequest{
		Name:   name,
		Chart:  chart,
		Values: config,
	}

	releaseResponse, err := releaseServer.UpdateRelease(ctx, updateReq)
	if err != nil {
		// Workaround for helm/helm#3338
		if releaseResponse.GetRelease() != nil {
			rollbackReq := &svc.RollbackReleaseRequest{
				Name:  name,
				Force: true,
			}
			_, rollbackErr := releaseServer.RollbackRelease(ctx, rollbackReq)
			if rollbackErr != nil {
				return nil, fmt.Errorf("failed to roll back failed update: %s: %s", rollbackErr, err)
			}
		}
		return nil, err
	}
	return releaseResponse.GetRelease(), nil
}

func reconcileRelease(ctx context.Context, tillerKubeClient *kube.Client, namespace string, expectedManifest string) error {
	expectedInfos, err := tillerKubeClient.BuildUnstructured(namespace, bytes.NewBufferString(expectedManifest))
	if err != nil {
		return err
	}
	return expectedInfos.Visit(func(expected *resource.Info, err error) error {
		if err != nil {
			return err
		}

		expectedClient := resource.NewClientWithOptions(expected.Client, func(r *rest.Request) {
			*r = *r.Context(ctx)
		})
		helper := resource.NewHelper(expectedClient, expected.Mapping)

		existing, err := helper.Get(expected.Namespace, expected.Name, false)
		if apierrors.IsNotFound(err) {
			if _, err := helper.Create(expected.Namespace, true, expected.Object, &metav1.CreateOptions{}); err != nil {
				return fmt.Errorf("create error: %s", err)
			}
			return nil
		} else if err != nil {
			return err
		}

		patch, err := generatePatch(existing, expected.Object)
		if err != nil {
			return fmt.Errorf("failed to marshal JSON patch: %s", err)
		}

		if patch == nil {
			return nil
		}

		_, err = helper.Patch(expected.Namespace, expected.Name, apitypes.JSONPatchType, patch, &metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("patch error: %s", err)
		}
		return nil
	})
}

func generatePatch(existing, expected runtime.Object) ([]byte, error) {
	existingJSON, err := json.Marshal(existing)
	if err != nil {
		return nil, err
	}
	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		return nil, err
	}

	ops, err := jsonpatch.CreatePatch(existingJSON, expectedJSON)
	if err != nil {
		return nil, err
	}

	// We ignore the "remove" operations from the full patch because they are
	// fields added by Kubernetes or by the user after the existing release
	// resource has been applied. The goal for this patch is to make sure that
	// the fields managed by the Helm chart are applied.
	patchOps := make([]jsonpatch.JsonPatchOperation, 0)
	for _, op := range ops {
		if op.Operation != "remove" {
			patchOps = append(patchOps, op)
		}
	}

	// If there are no patch operations, return nil. Callers are expected
	// to check for a nil response and skip the patch operation to avoid
	// unnecessary chatter with the API server.
	if len(patchOps) == 0 {
		return nil, nil
	}

	return json.Marshal(patchOps)
}

func uninstallRelease(ctx context.Context, storageBackend *storage.Storage, releaseServer *tiller.ReleaseServer, releaseName string) (*rpb.Release, error) {
	// Get history of this release
	h, err := storageBackend.History(releaseName)
	if err != nil {
		return nil, fmt.Errorf("failed to get release history: %s", err)
	}

	// If there is no history, the release has already been uninstalled,
	// so return ErrNotFound.
	if len(h) == 0 {
		return nil, helmif.ErrNotFound
	}

	uninstallResponse, err := releaseServer.UninstallRelease(ctx, &svc.UninstallReleaseRequest{
		Name:  releaseName,
		Purge: true,
	})
	return uninstallResponse.GetRelease(), err
}
