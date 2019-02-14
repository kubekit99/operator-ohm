#operator-framework usage for openstack

# Creation of openstacklcm-operator

## Initialising the openstacklcm-operator

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
## Adjusting crds

Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

```bash
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./deploy/crds/
ls chart/templates/*_crd.yaml > filelist
for i in `cat filelist`; do NEWNAME=`echo $i | sed -e "s/_crd//g"`; mv $NEWNAME $i; done
rm filelist
```

## Compiling the openstacklcm-operator

```
dep ensure
./manualbuild.sh
```

# Deploying

## Deployment of operator manually

```bash
cd chart
helm install --name openstacklcm .
```

## Deployment of operator manually

```bash
kubectl apply -f deploy/crds/arpscan_v1alpha1_kubedge_cr.yaml
```


