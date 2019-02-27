# keystone operator

Operator-SDK Helm-Operator deployment of Keystone

## Keystone

### Initialization

The basic structure was created using the following command

```bash
operator-sdk new keystone-operator --api-version=openstackhelm.openstack.org/v1alpha1 --kind=Mariadb --type=helm --skip-git-init
```

### Implementation

First did ugly copy paste of keystone openstackhelm chart and helmtoolkit into helm-charts

The created operator helm operator image. Helm charts ends up beeing delivered as part of the docker image.

```
docker build -t keystone-operator:poc -f build/Dockerfile .
```

Note that the role had to be modified because the operator invokes the helm chart which in turn create roles, rolebinding and serviceaccounts.
Could not go in production with 


### Deployment

Then deployed the operator and the CR. Number of keystone server is supposed to matched number in _cr.yaml

```
kubectl create namespace operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_keystone_crd.yaml -n operatorpoc
kubectl create -f deploy/service_account.yaml -n operatorpoc
kubectl create -f deploy/role.yaml -n operatorpoc
kubectl create -f deploy/role_binding.yaml -n operatorpoc
kubectl create -f deploy/operator.yaml -n operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_keystone_cr.yaml -n operatorpoc
```

Check deployment

```
kubectl get all -n operatorpoc

```

### Conclusions


