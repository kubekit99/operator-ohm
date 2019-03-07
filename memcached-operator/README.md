# Memcached operator

Armada-Operator deployment of memcached

# Prerequisite

Memcached-Operator still requires:
- Helm2 tiller to be deployed in your cluster


# Deploy the Memcached-Operator

## Rebuild the operator if needed and deploy it

```bash
make install

docker build -t kubekit99/memcached-operator:poc -f build/Dockerfile .
Sending build context to Docker daemon  77.82kB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> 05ffe3152315
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> 8daf96e356dd
Successfully built 8daf96e356dd
Successfully tagged kubekit99/memcached-operator:poc
docker tag kubekit99/memcached-operator:poc kubekit99/memcached-operator:latest
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
deployment.apps/memcached-operator created
```

## Check the installation of the CRD and operator

```bash
kubectl get all

NAME                                      READY   STATUS    RESTARTS   AGE
pod/memcached-operator-59474c5bc5-wj824   1/1     Running   0          15m

NAME                         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/kubernetes           ClusterIP   10.96.0.1      <none>        443/TCP    21m
service/memcached-operator   ClusterIP   10.110.14.86   <none>        8383/TCP   15m

NAME                                 READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/memcached-operator   1/1     1            1           15m

NAME                                            DESIRED   CURRENT   READY   AGE
replicaset.apps/memcached-operator-59474c5bc5   1         1         1       15m
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
kubectl apply -f examples/memcached/simple.yaml
kubectl describe act/memcached
```

Result should be:

```bash
kubectl label nodes airship openstack-control-plane=enabled --overwrite
node/airship not labeled
kubectl apply -f examples/memcached/simple.yaml
armadachart.armada.airshipit.org/memcached created
```

# Check the CRD and the deployment of the underlying helm chart

## Check the ArmachaChart Custom Resource

```bash
kubectl describe act/memcached

Name:         memcached
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChart","metadata":{"annotations":{},"name":"memcached","namespace":"default"},...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChart
Metadata:
  Creation Timestamp:  2019-03-07T16:02:50Z
  Finalizers:
    uninstall-helm-release
  Generation:        2
  Resource Version:  1155
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadacharts/memcached
  UID:               754fb964-40f2-11e9-a001-0800272e6982
Spec:
  Chart Name:  memcached
  Dependencies:
  Namespace:  default
  Release:    memcached
  Source:
    Location:    /opt/armada/helm-charts/memcached
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
    Last Transition Time:  2019-03-07T16:02:50Z
    Status:                True
    Type:                  Initialized
    Last Transition Time:  2019-03-07T16:02:51Z
    Reason:                InstallSuccessful
    Resource Name:         memcached
    Resource Version:      1
    Status:                True
    Type:                  Deployed
  Succeeded:               true
Events:
  Type    Reason    Age   From          Message
  ----    ------    ----  ----          -------
  Normal  Deployed  108s  act-recorder  InstallSuccessful
```

  
## Check the Memcached Service deployment

```bash
kubectl get all

NAME                                       READY   STATUS    RESTARTS   AGE
pod/memcached-memcached-5bc79f976c-sh2hm   1/1     Running   0          3m52s
pod/memcached-operator-59474c5bc5-wj824    1/1     Running   0          5m39s

NAME                         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)     AGE
service/kubernetes           ClusterIP   10.96.0.1      <none>        443/TCP     11m
service/memcached            ClusterIP   10.96.171.54   <none>        11211/TCP   3m52s
service/memcached-operator   ClusterIP   10.110.14.86   <none>        8383/TCP    5m37s

NAME                                  READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/memcached-memcached   1/1     1            1           3m52s
deployment.apps/memcached-operator    1/1     1            1           5m39s

NAME                                             DESIRED   CURRENT   READY   AGE
replicaset.apps/memcached-memcached-5bc79f976c   1         1         1       3m52s
replicaset.apps/memcached-operator-59474c5bc5    1         1         1       5m39s
```

# Remove the CRD 

```bash
make deletemanifest
```

or

```bash
kubectl delete -f examples/memcached/simple.yaml
kubectl describe act/memcached
```

Check the corresponding resources are being deleted

```bash
kubectl get all

NAME                                       READY   STATUS        RESTARTS   AGE
pod/memcached-memcached-5bc79f976c-sh2hm   0/1     Terminating   0          7m15s
pod/memcached-operator-59474c5bc5-wj824    1/1     Running       0          9m2s

NAME                         TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/kubernetes           ClusterIP   10.96.0.1      <none>        443/TCP    15m
service/memcached-operator   ClusterIP   10.110.14.86   <none>        8383/TCP   9m

NAME                                 READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/memcached-operator   1/1     1            1           9m2s

NAME                                            DESIRED   CURRENT   READY   AGE
replicaset.apps/memcached-operator-59474c5bc5   1         1         1       9m2s
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
deployment.apps "memcached-operator" deleted
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
serviceaccount "armada-operator" deleted
```

```
kubectl get all

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   25m
```

