# Kubernetes Operator for Openstack HELM

# Creation of armada-operator

## Initialising the armada-operator

```bash
operator-sdk new armada-operator --skip-git-init
```

## Coding the armada-operator

```bash
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=OpenstackChart
git add deploy/crds/
git add pkg/apis/armada/
git add pkg/apis/addtoscheme_armada_v1alpha1.go
git add deploy/role.yaml
```

```bash
vi pkg/apis/armada/v1alpha1/*_types.go
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=armada.airshipit.org/v1alpha1 --kind=OpenstackChart
```
## Adjusting crds

Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

```bash
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/
operator-sdk generate k8s
```

or

```bash
make generate
```

## Compiling the armada-operator

```
make docker-build
```

# Deploying

## Deployment of operator using helm

```bash
helm install --name osh-operator chart 
```

## Openstack Databases Backup CRD and Controller

### Trigger a backup

Upon creation of the custom resource, the controller will
- Create a new workflow owned by the customer resource.
- Add events to the custom resources.
- The workflow creation will get argo to react and run the workflow, and create news pods.


```bash
kubectl create -f examples/openstackchart-testchart.yaml
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackchart-testchart.yaml
```
