# THIS REPOSITORY IS OBSOLETE. CONTENT HAS BEEN MIGRATED ONTO [Keleustes](https://github.com/keleustes/) 

# POCs parking lot

Before they get deleted from git, some of the POCs components
have been moved in that `pocs` subdirectory.

Some of the notes taken during the genesis of this operator have also been saved in this
present README.


# POC helmrelease Create a CRD to controller tiller

The bulk of the code was coming for the operator-sdk/helm-operator.
Was developed first and abandonned when the ArmadaChart was created.

## HelmRelease CRDs.

The HelmRelease represent temporalily used until the HelmV3 Release CRD is available.

## Trigger helm chart deployment

Upon creation of the custom resource, the controller will
- Deploy the Helm Chart described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/tiller/helm-testchart.yaml
kubectl describe hrel/testchart
```

## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe hrel/testchart
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding helm chart should be uninstalled.

```bash
kubectl delete -f examples/tiller/helm-testchart.yaml
```


# POC requests: Trigger Armada and Helm Request using CRD.

In this POC, CRD such as ArmadaApply or HelmInstall CRD are being
created to trigger the actions from Armada and Helm.

## HelmRequest CRD

The helm-v3 proposal appendix describes in Appendix the creation of the HelmRequest CRD to
work with a controller and the Helm code provided as a standalone library.
- Note sure the interest of the "Get"....operations since the same result can be achieve
using "kubectl get releases"

# POC helm3crd: Simple Helm3CRD Release and Controller

## Helm3CRD directories

### Release, Values, Chart... CRDs.

Those CRD and Controller have been build using the helm-3-crd discussion and repository.

### Notes

- JEB: Can't figure out the impact of having thousands of individual LifecycleEvent CR on the performance
and usability of the kubectl.

## Trigger helm chart deployment

Upon creation of the custom resource, the controller will
- Deploy the Helm Chart described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/helm3crd/release.yaml
kubectl describe rel/my-release
```

## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe rel/my-release
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding helm chart should be uninstalled.

```bash
kubectl delete -f examples/helm3crd/release.yaml
```

# Appendix A: Helm hooks for CRDs

Example of annotations hook for a CRD

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
  annotations:
    "helm.sh/hook": crd-install
spec:
  group: stable.example.com
  version: v1
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
```

# Appendix B: Genesis of the armada operator

```bash
operator-sdk new armada-operator --skip-git-init
```

## Coding the armada-operator

```bash
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaManifest
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaChart
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaChartGroup
git add pkg/apis/armada/
git add pkg/apis/addtoscheme_armada_v1alpha1.go
git add deploy/role.yaml
```

```bash
vi pkg/apis/armada/v1alpha1/*_types.go
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=armada.airshipit.org/v1alpha1 --kind=Tiller
operator-sdk add controller --api-version=armada.airshipit.org/v1alpha1 --kind=Armada
operator-sdk add controller --api-version=helm3crd.airshipit.org/v1beta1 --kind=Helm3CRD
```
