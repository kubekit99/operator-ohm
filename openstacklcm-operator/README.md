#operator-framework usage for openstack

## Initialization of openstacklcm-operator

```bash
operator-sdk new openstacklcm-operator --skip-git-init
```

## Coding the openstacklcm-operator

```bash
operator-sdk add api --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add api --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add api --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add api --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add api --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackDeployment
git add deploy/crds/
git add pkg/apis/openstackhelm/
git add pkg/apis/addtoscheme_openstackhelm_v1alpha1.go 
git add deploy/role.yaml 

```bash
vi pkg/apis/openstackhelm/v1alpha1/*_types.go 
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add controller --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add controller --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add controller --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add controller --api-version=openstackhelm.openstack.org/v1alpha1 --kind=OpenstackDeployment
```

## Compiling the openstacklcm-operator

```
dep ensure
./manualbuild.sh
```


