# Kubernetes Operator for Openstack LCM

# Creation of openstacklcm-operator

## Initialising the openstacklcm-operator

```bash
operator-sdk new openstacklcm-operator --skip-git-init
```

## Coding the openstacklcm-operator

```bash
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackDeployment
git add deploy/crds/
git add pkg/apis/openstacklcm/
git add pkg/apis/addtoscheme_openstacklcm_v1alpha1.go
git add deploy/role.yaml
```

```bash
vi pkg/apis/openstacklcm/v1alpha1/*_types.go
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackDeployment
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

## Compiling the openstacklcm-operator

```
make docker-build
```

# Deploying

## Deployment of operator using helm

```bash
helm install --name lcm-operator chart 
```

## Openstack Databases Backup CRD and Controller

### Trigger a backup

Upon creation of the custom resource, the controller will
- Create a new workflow owned by the customer resource.
- Add events to the custom resources.
- The workflow creation will get argo to react and run the workflow, and create news pods.


```bash
kubectl apply -f examples/openstackbackup/backup-example.yaml
kubectl describe openstackbackups/openstackbackup
kubectl describe workflows/openstackbackup-wf
kubectl describe pod/openstackbackup-wf
kubectl logs pod/openstackbackup-wf main
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
kubectl delete workflows/openstackbackup-wf
kubectl logs pod/openstackbackup-wf main
kubectl describe openstackbackups/openstackbackup
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackbackup/backup-example.yaml
kubectl describe workflows/openstackbackup-wf
kubectl describe pod/openstackbackup-wf
```


## Openstack Databases Data Restore CRD and Controller

Upon creation of the custom resource, the controller will
- Create a new workflow owned by the customer resource.
- Add events to the custom resources.
- The workflow creation will get argo to react and run the workflow, and create news pods.

### Trigger a restore

```bash
kubectl apply -f examples/openstackrestore/restore-example.yaml
kubectl describe openstackrestores/openstackrestore
kubectl describe workflows/openstackrestore-wf
kubectl describe pod/openstackrestore-wf
kubectl logs pod/openstackrestore-wf main
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
kubectl delete workflows/openstackrestore-wf
kubectl logs pod/openstackrestore-wf main
kubectl describe openstackrestores/openstackrestore
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackrestore/restore-example.yaml
kubectl describe workflows/openstackrestore-wf
kubectl describe pod/openstackrestore-wf
```


## Openstack Upgrade CRD and Controller

### Trigger an upgrade

```bash
kubectl apply -f examples/openstackupgrade/upgrade-example.yaml
kubectl describe openstackupgrades/openstackupgrade
kubectl describe workflows/openstackupgrade-wf
kubectl describe pod/openstackupgrade-wf
kubectl logs pod/openstackupgrade-wf main
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
kubectl delete workflows/openstackupgrade-wf
kubectl logs pod/openstackupgrade-wf main
kubectl describe openstackupgrades/openstackupgrade
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackupgrade/upgrade-example.yaml
kubectl describe workflows/openstackupgrade-wf
kubectl describe pod/openstackupgrade-wf
```


## Openstack Rollback CRD and Controller

### Trigger a rollback

Upon creation of the custom resource, the controller will
- Create a new workflow owned by the customer resource.
- Add events to the custom resources.
- The workflow creation will get argo to react and run the workflow, and create news pods.

```bash
kubectl apply -f examples/openstackrollback/rollback-example.yaml
kubectl describe openstackrollbacks/openstackrollback
kubectl describe workflows/openstackrollback-wf
kubectl describe pod/openstackrollback-wf
kubectl logs pod/openstackrollback-wf main
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
kubectl delete workflows/openstackrollback-wf
kubectl logs pod/openstackrollback-wf main
kubectl describe openstackrollbacks/openstackrollback
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackrollback/rollback-example.yaml
kubectl describe workflows/openstackrollback-wf
kubectl describe pod/openstackrollback-wf
```


## Openstack Greenfield Deployment CRD and Controller

### Trigger a deployment

Upon creation of the custom resource, the controller will
- Create a new workflow owned by the customer resource.
- Add events to the custom resources.
- The workflow creation will get argo to react and run the workflow, and create news pods.

```bash
kubectl apply -f examples/openstackdeployment/deployment-example.yaml
kubectl describe openstackdeployments/openstackdeployment
kubectl describe workflows/openstackdeployment-wf
kubectl describe pod/openstackdeployment-wf
kubectl logs pod/openstackdeployment-wf main
```

### Test controller reconcilation logic (for depending workflows)

Upon deletion of its workflow, the controller will recreate it,
which will get argo to rerun the workflow.

```bash
kubectl delete workflows/openstackdeployment-wf
kubectl logs pod/openstackdeployment-wf main
kubectl describe openstackdeployments/openstackdeployment
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding workflow should be deleted.
Argo in turn, will delete the corresponding pods.

```bash
kubectl delete -f examples/openstackdeployment/deployment-example.yaml
kubectl describe workflows/openstackdeployment-wf
kubectl describe pod/openstackdeployment-wf
```


