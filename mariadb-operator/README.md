# MariaDB operator

Operator-SDK Helm-Operator deployment of MariaDB

## MariaDB

### Initialization

The basic structure was created using the following command

```bash
operator-sdk new mariadb-operator --api-version=openstackhelm.openstack.org/v1alpha1 --kind=Mariadb --type=helm --skip-git-init
```

### Implementation

First did ugly copy paste of mariadb openstackhelm chart and helmtoolkit into helm-charts

The created operator helm operator image. Helm charts ends up beeing delivered as part of the docker image.

```
docker build -t mariadb-operator:poc -f build/Dockerfile .
```

Note that the role had to be modified because the operator invokes the helm chart which in turn create roles, rolebinding and serviceaccounts.
Could not go in production with 


### Deployment

Then deployed the operator and the CR. Number of mariadb server is supposed to matched number in _cr.yaml

```
kubectl create namespace operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_mariadb_crd.yaml -n operatorpoc
kubectl create -f deploy/service_account.yaml -n operatorpoc
kubectl create -f deploy/role.yaml -n operatorpoc
kubectl create -f deploy/role_binding.yaml -n operatorpoc
kubectl create -f deploy/operator.yaml -n operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_mariadb_cr.yaml -n operatorpoc
```

Check deployment

```
kubectl get all -n operatorpoc

NAME                                               READY   STATUS    RESTARTS   AGE
pod/mariadb-ingress-6ccc4dbf76-ptfk7               0/1     Running   0          31s
pod/mariadb-ingress-error-pages-85c9996d68-45jpm   1/1     Running   0          31s
pod/mariadb-operator-6786859b89-kg8lx              1/1     Running   0          34s
pod/mariadb-server-0                               0/1     Running   0          31s
pod/mariadb-server-1                               0/1     Running   0          31s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
service/mariadb                       ClusterIP   10.105.88.65     <none>        3306/TCP            31s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP   31s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP              31s
service/mariadb-server                ClusterIP   10.103.136.223   <none>        3306/TCP            31s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               0/1     1            0           31s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           31s
deployment.apps/mariadb-operator              1/1     1            1           34s

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-6ccc4dbf76               1         1         0       31s
replicaset.apps/mariadb-ingress-error-pages-85c9996d68   1         1         1       31s
replicaset.apps/mariadb-operator-6786859b89              1         1         1       34s

NAME                              READY   AGE
statefulset.apps/mariadb-server   0/2     31s
```

### Conclusions

#### Note 1

Does not seem to work right. May be coming from the helm chart itself. 

```
vi deploy/crds/openstackhelm_v1alpha1_mariadb_cr.yaml
kubectl create -f deploy/crds/openstackhelm_v1alpha1_mariadb_cr.yaml -n operatorpoc
```

#### Note 2

Had to add roles, rolebinding in the roles.yaml to let operator run the helm chart which are creating roles and serviceaccount

#### Note 3

The Dockerfile should do some kind of wget or helm fetch in order to avoid ugly duplication of the chart and toolkit

#### Note 4

Note sure how far such as solution will be from helm v3....if every chart comes with its own CR and CRD.

#### Note 5

"helm ls" does not see the deployed chart, I guess the helm-operator is not accessing tiller.
