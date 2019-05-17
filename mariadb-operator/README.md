# THIS REPOSITORY IS OBSOLETE. CONTENT HAS BEEN MIGRATED ONTO [Keleustes](https://github.com/keleustes/) 

# MariaDB operator

Armada-Operator deployment of mariadb

# Prerequisite

MariaDB-Operator still requires:
- Helm2 tiller to be deployed in your cluster


# Deploy the MariaDB-Operator

## Rebuild the operator if needed and deploy it

```bash
make install

docker build -t kubekit99/mariadb-operator:poc -f build/Dockerfile .
Sending build context to Docker daemon  222.2kB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> 05ffe3152315
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> 74f4d8069d59
Successfully built 74f4d8069d59
Successfully tagged kubekit99/mariadb-operator:poc
docker tag kubekit99/mariadb-operator:poc kubekit99/mariadb-operator:latest
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
customresourcedefinition.apiextensions.k8s.io/armadarequests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
customresourcedefinition.apiextensions.k8s.io/helmreleases.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
customresourcedefinition.apiextensions.k8s.io/helmrequests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
customresourcedefinition.apiextensions.k8s.io/lifecyleevents.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
customresourcedefinition.apiextensions.k8s.io/lifecyles.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
customresourcedefinition.apiextensions.k8s.io/manifests.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
customresourcedefinition.apiextensions.k8s.io/releaseaudits.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
customresourcedefinition.apiextensions.k8s.io/releases.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
customresourcedefinition.apiextensions.k8s.io/values.helm3crd.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl create -f deploy/operator.yaml
deployment.apps/mariadb-operator created
```

## Check the installation of the CRD and operator

```bash
kubectl get all

NAME                                    READY   STATUS    RESTARTS   AGE
pod/mariadb-operator-7847f4c4dc-gtrq6   1/1     Running   0          61s

NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/kubernetes         ClusterIP   10.96.0.1        <none>        443/TCP    40m
service/mariadb-operator   ClusterIP   10.105.205.210   <none>        8383/TCP   59s

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-operator   1/1     1            1           61s

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-operator-7847f4c4dc   1         1         1       61s
```

```bash
kubectl get amf
No resources found.
```

```bash
kubectl get acg
No resources found.
```

```bash
kubectl get act
No resources found.
```


# Create the ArmadaChart CRD

```bash
make installmanifest
```

or

```bash
kubectl label nodes airship openstack-control-plane=enabled --overwrite
kubectl apply -f examples/mariadb/simple.yaml
kubectl describe act/mariadb
```

Result should be:

```bash
kubectl label nodes airship openstack-control-plane=enabled --overwrite
node/airship not labeled
kubectl apply -f examples/mariadb/simple.yaml
armadachart.armada.airshipit.org/mariadb created
```

# Check the CRD and the deployment of the underlying helm chart

## Check the ArmachaChart Custom Resource

```bash
kubectl describe act/mariadb

Name:         mariadb
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChart","metadata":{"annotations":{},"name":"mariadb","namespace":"default"},"s...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChart
Metadata:
  Creation Timestamp:  2019-03-07T16:35:44Z
  Finalizers:
    uninstall-helm-release
  Generation:        2
  Resource Version:  4133
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadacharts/mariadb
  UID:               0dd907db-40f7-11e9-a001-0800272e6982
Spec:
  Chart Name:  mariadb
  Dependencies:
  Namespace:  default
  Release:    mariadb
  Source:
    Location:    /opt/armada/helm-charts/mariadb
    Reference:   87aad18f7d8c6a1a08f3adc8866efd33bee6aa52
    Subpath:     .
    Type:        local
  Target State:  initialized
  Upgrade:
    No Hooks:  false
  Values:
Status:
  Actual State:  initialized
  Conditions:
    Last Transition Time:  2019-03-07T16:35:44Z
    Status:                True
    Type:                  Initialized
    Last Transition Time:  2019-03-07T16:35:45Z
    Reason:                UpdateSuccessful
    Resource Name:         mariadb
    Resource Version:      5
    Status:                True
    Type:                  Deployed
  Succeeded:               true
Events:
  Type    Reason    Age                From          Message
  ----    ------    ----               ----          -------
  Normal  Deployed  55s                act-recorder  InstallSuccessful
  Normal  Deployed  51s (x4 over 54s)  act-recorder  UpdateSuccessful
```

  
## Check the MariaDB service deployment

```bash
kubectl get all

NAME                                              READY   STATUS    RESTARTS   AGE
pod/mariadb-ingress-75cb44bc8c-tt2dx              1/1     Running   0          2m
pod/mariadb-ingress-error-pages-8f44b444b-lrcrx   1/1     Running   0          2m
pod/mariadb-operator-7847f4c4dc-gtrq6             1/1     Running   0          3m51s
pod/mariadb-server-0                              1/1     Running   0          2m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)             AGE
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP             42m
service/mariadb                       ClusterIP   10.108.196.84    <none>        3306/TCP            2m1s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP   2m1s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP              2m1s
service/mariadb-operator              ClusterIP   10.105.205.210   <none>        8383/TCP            3m49s
service/mariadb-server                ClusterIP   10.103.59.60     <none>        3306/TCP            2m1s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               1/1     1            1           2m
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           2m1s
deployment.apps/mariadb-operator              1/1     1            1           3m51s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-75cb44bc8c              1         1         1       2m
replicaset.apps/mariadb-ingress-error-pages-8f44b444b   1         1         1       2m1s
replicaset.apps/mariadb-operator-7847f4c4dc             1         1         1       3m51s

NAME                              READY   AGE
statefulset.apps/mariadb-server   1/1     2m
```

# Remove the CRD 

```bash
make deletemanifest
```

or

```bash
kubectl delete -f examples/mariadb/simple.yaml
```

Check the corresponding resources are being deleted

```bash
kubectl get all
NAME                                              READY   STATUS        RESTARTS   AGE
pod/mariadb-ingress-75cb44bc8c-tt2dx              0/1     Terminating   0          3m28s
pod/mariadb-ingress-error-pages-8f44b444b-lrcrx   0/1     Terminating   0          3m28s
pod/mariadb-operator-7847f4c4dc-gtrq6             1/1     Running       0          5m19s
pod/mariadb-server-0                              1/1     Terminating   0          3m28s

NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/kubernetes         ClusterIP   10.96.0.1        <none>        443/TCP    44m
service/mariadb-operator   ClusterIP   10.105.205.210   <none>        8383/TCP   5m17s

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-operator   1/1     1            1           5m19s

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-operator-7847f4c4dc   1         1         1       5m19
```

Check the ArmadaManifest, ChartGroup and Chart are no longer present
```bash
kubectl get amf
No resources found.
```

```bash
kubectl get acg
No resources found.
```

```bash
kubectl get act
No resources found.
```

# Purge artifacts from Kubernetes cluster

```bash
make purge

kubectl delete -f deploy/operator.yaml
deployment.apps "mariadb-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io "armadachartgroups.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io "armadacharts.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io "armadamanifests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
customresourcedefinition.apiextensions.k8s.io "armadarequests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
customresourcedefinition.apiextensions.k8s.io "helmreleases.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
customresourcedefinition.apiextensions.k8s.io "helmrequests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
customresourcedefinition.apiextensions.k8s.io "lifecyleevents.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
customresourcedefinition.apiextensions.k8s.io "lifecyles.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
customresourcedefinition.apiextensions.k8s.io "manifests.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
customresourcedefinition.apiextensions.k8s.io "releaseaudits.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
customresourcedefinition.apiextensions.k8s.io "releases.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
customresourcedefinition.apiextensions.k8s.io "values.helm3crd.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount "armada-operator" delete
```

```
kubectl get all

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   25m
```

Delete the configmaps related to mariadb if needed
```
kubectl get configmaps

NAME                              DATA   AGE
mariadb-mariadb-mariadb-ingress   0      5m3s
mariadb-mariadb-state             5      4m46s
```
