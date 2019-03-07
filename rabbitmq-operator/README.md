# RabbitMQ operator

Armada-Operator deployment of rabbitmq

# Prerequisite

RabbitMQ-Operator still requires:
- Helm2 tiller to be deployed in your cluster


# Deploy the RabbitMQ-Operator

## Rebuild the operator if needed and deploy it

```bash
make install

docker build -t kubekit99/rabbitmq-operator:poc -f build/Dockerfile .
Sending build context to Docker daemon  115.7kB
Step 1/2 : FROM kubekit99/armada-operator-dev:latest
 ---> 05ffe3152315
Step 2/2 : COPY helm-charts/ /opt/armada/helm-charts/
 ---> Using cache
 ---> f2deee4e5bb3
Successfully built f2deee4e5bb3
Successfully tagged kubekit99/rabbitmq-operator:poc
docker tag kubekit99/rabbitmq-operator:poc kubekit99/rabbitmq-operator:latest
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
deployment.apps/rabbitmq-operator created
```

## Check the installation of the CRD and operator

```bash
kubectl get all

NAME                                     READY   STATUS    RESTARTS   AGE
pod/rabbitmq-operator-778c9dcfd7-2zstw   1/1     Running   0          38s

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP    30m
service/rabbitmq-operator   ClusterIP   10.101.31.125   <none>        8383/TCP   37s

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/rabbitmq-operator   1/1     1            1           38s

NAME                                           DESIRED   CURRENT   READY   AGE
replicaset.apps/rabbitmq-operator-778c9dcfd7   1         1         1       38s
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
kubectl apply -f examples/rabbitmq/simple.yaml
kubectl describe act/rabbitmq
```

Result should be:

```bash
kubectl label nodes airship openstack-control-plane=enabled --overwrite
node/airship not labeled
kubectl apply -f examples/rabbitmq/simple.yaml
armadachart.armada.airshipit.org/rabbitmq created
```

# Check the CRD and the deployment of the underlying helm chart

## Check the ArmachaChart Custom Resource

```bash
kubectl describe act/rabbitmq

Name:         rabbitmq
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChart","metadata":{"annotations":{},"name":"rabbitmq","namespace":"default"},"...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChart
Metadata:
  Creation Timestamp:  2019-03-07T16:27:30Z
  Finalizers:
    uninstall-helm-release
  Generation:        2
  Resource Version:  3122
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadacharts/rabbitmq
  UID:               e785d146-40f5-11e9-a001-0800272e6982
Spec:
  Chart Name:  rabbitmq
  Dependencies:
  Namespace:  default
  Release:    rabbitmq
  Source:
    Location:    /opt/armada/helm-charts/rabbitmq
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
    Last Transition Time:  2019-03-07T16:27:30Z
    Status:                True
    Type:                  Initialized
    Last Transition Time:  2019-03-07T16:27:31Z
    Reason:                InstallSuccessful
    Resource Name:         rabbitmq
    Resource Version:      1
    Status:                True
    Type:                  Deployed
  Succeeded:               true
Events:
  Type    Reason    Age   From          Message
  ----    ------    ----  ----          -------
  Normal  Deployed  92s   act-recorder  InstallSuccessful
```

## Check RabbitMQ Service deployment

```bash
kubectl get all

NAME                                     READY   STATUS    RESTARTS   AGE
pod/rabbitmq-operator-778c9dcfd7-2zstw   1/1     Running   0          4m43s
pod/rabbitmq-rabbitmq-0                  1/1     Running   0          2m6s

NAME                          TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/kubernetes            ClusterIP   10.96.0.1        <none>        443/TCP                        34m
service/rabbitmq              ClusterIP   10.100.88.64     <none>        5672/TCP,25672/TCP,15672/TCP   2m7s
service/rabbitmq-dsv-7b1733   ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m7s
service/rabbitmq-mgr-7b1733   ClusterIP   10.100.147.217   <none>        80/TCP,443/TCP                 2m7s
service/rabbitmq-operator     ClusterIP   10.101.31.125    <none>        8383/TCP                       4m42s

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/rabbitmq-operator   1/1     1            1           4m43s

NAME                                           DESIRED   CURRENT   READY   AGE
replicaset.apps/rabbitmq-operator-778c9dcfd7   1         1         1       4m43s

NAME                                 READY   AGE
statefulset.apps/rabbitmq-rabbitmq   1/1     2m7s
```

# Remove the CRD 

```bash
make deletemanifest
```

or

```bash
kubectl delete -f examples/rabbitmq/simple.yaml
kubectl describe act/rabbitmq
```

Check the corresponding resources are being deleted

```bash
kubectl get all

NAME                                     READY   STATUS    RESTARTS   AGE
pod/rabbitmq-operator-778c9dcfd7-2zstw   1/1     Running   0          5m46s

NAME                        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/kubernetes          ClusterIP   10.96.0.1       <none>        443/TCP    35m
service/rabbitmq-operator   ClusterIP   10.101.31.125   <none>        8383/TCP   5m45s

NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/rabbitmq-operator   1/1     1            1           5m46s

NAME                                           DESIRED   CURRENT   READY   AGE
replicaset.apps/rabbitmq-operator-778c9dcfd7   1         1         1       5m46s
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
deployment.apps "rabbitmq-operator" deleted
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

