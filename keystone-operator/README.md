# Keystone operator

Armada-Operator deployment of keystone

# Prerequisite

Keystone-Operator still requires:
- Helm2 tiller to be deployed in your cluster
- Argo to be deployed in your cluster


# Deploy the Keystone-Operator


## Rebuild the operators.

This add the helm-charts to the operators containers and avoid the need
for git access

```bash
make docker-build
```

## Deploy the operators (armada and lifecycles) 

```bash
make install-armada && make install-oslc

kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount/armada-argo-sa created
role.rbac.authorization.k8s.io/armada-argo-role created
rolebinding.rbac.authorization.k8s.io/armada-argo-rolebinding created
kubectl create -f deploy/armada-operator.yaml

deployment.apps/keystone-armada-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io/deletephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io/installphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io/operationalphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io/planningphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io/rollbackphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io/testphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficdrainphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficrolloutphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io/upgradephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io/oslcs.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount/openstacklcm-argo-sa created
role.rbac.authorization.k8s.io/openstacklcm-argo-role created
rolebinding.rbac.authorization.k8s.io/openstacklcm-argo-rolebinding created
kubectl create -f deploy/oslc-operator.yaml
deployment.apps/keystone-oslc-operator created
```

## Check the installation of the CRD and operator

```bash
kubectl get all --all-namespaces

NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-cbsfv                     1/1     Running   0          4m20s
argo          pod/argo-workflow-controller-6d987f7fdc-vlrlf   1/1     Running   0          4m20s
default       pod/keystone-armada-operator-5cfccc74fb-kmhvp   1/1     Running   0          14s
default       pod/keystone-oslc-operator-747bb84698-sm5jg     1/1     Running   0          10s
kube-system   pod/calico-etcd-rf4nw                           1/1     Running   0          6m18s
kube-system   pod/calico-kube-controllers-79bd896977-q8hvr    1/1     Running   0          6m34s
kube-system   pod/calico-node-8fp6c                           1/1     Running   1          6m34s
kube-system   pod/coredns-fb8b8dccf-hxtbw                     1/1     Running   0          6m39s
kube-system   pod/coredns-fb8b8dccf-l4pzc                     1/1     Running   0          6m39s
kube-system   pod/etcd-airship                                1/1     Running   0          5m40s
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          5m44s
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          5m47s
kube-system   pod/kube-proxy-ww25b                            1/1     Running   0          6m40s
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          5m52s
kube-system   pod/tiller-deploy-8458f6c667-sd887              1/1     Running   0          6m30s

NAMESPACE     NAME                               TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui                    NodePort    10.99.142.95     <none>        80:32711/TCP             4m20s
default       service/keystone-armada-operator   ClusterIP   10.106.147.222   <none>        8383/TCP                 12s
default       service/keystone-oslc-operator     ClusterIP   10.96.251.99     <none>        8383/TCP                 8s
default       service/kubernetes                 ClusterIP   10.96.0.1        <none>        443/TCP                  6m55s
kube-system   service/calico-etcd                ClusterIP   10.96.232.136    <none>        6666/TCP                 6m35s
kube-system   service/kube-dns                   ClusterIP   10.96.0.10       <none>        53/UDP,53/TCP,9153/TCP   6m54s
kube-system   service/tiller-deploy              ClusterIP   10.99.3.180      <none>        44134/TCP                6m30s

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   6m35s
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       6m35s
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            6m54s

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           4m20s
argo          deployment.apps/argo-workflow-controller   1/1     1            1           4m20s
default       deployment.apps/keystone-armada-operator   1/1     1            1           14s
default       deployment.apps/keystone-oslc-operator     1/1     1            1           10s
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           6m35s
kube-system   deployment.apps/coredns                    2/2     2            2           6m54s
kube-system   deployment.apps/tiller-deploy              1/1     1            1           6m30s

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       4m20s
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       4m20s
default       replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       14s
default       replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       10s
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       6m35s
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       6m39s
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       6m30s
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
kubectl apply -f examples/keystone/
```

Result should be:

```bash
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
```

# Check the CRD and the deployment of the underlying helm chart

## Check the ArmachaChart Custom Resource

```bash
kubectl get act

NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
```

```bash
kubectl describe act/keystone

Name:         keystone
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChart","metadata":{"annotations":{},"name":"keystone","namespace":"default"},"...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChart
Metadata:
  Creation Timestamp:  2019-04-03T19:15:39Z
  Finalizers:
    uninstall-helm-release
  Generation:        2
  Resource Version:  5866
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadacharts/keystone
  UID:               de58271b-5644-11e9-8d63-0800272e6982
Spec:
  Chart Name:  keystone
  Dependencies:
    helm-toolkit
  Namespace:  openstack
  Release:    keystone
  Source:
    Location:    /opt/armada/helm-charts/keystone
    Reference:   master
    Subpath:     .
    Type:        local
  Target State:  deployed
  Test:
  Upgrade:
    No Hooks:  false
    Pre:
      Delete:
        Labels:
        Name:  keystone-bootstrap
        Type:  job
  Values:
  Wait:
    Labels:
    Timeout:  100
Status:
  Actual State:  deployed
  Conditions:
    Last Transition Time:  2019-04-03T19:15:44Z
    Status:                True
    Type:                  Initializing
    Last Transition Time:  2019-04-03T19:15:45Z
    Reason:                UpdateSuccessful
    Resource Name:         keystone
    Resource Version:      160
    Status:                True
    Type:                  Deployed
  Satisfied:               true
Events:
  Type    Reason    Age                    From          Message
  ----    ------    ----                   ----          -------
  Normal  Deployed  4m46s                  act-recorder  InstallSuccessful
  Normal  Deployed  4m1s (x24 over 4m42s)  act-recorder  UpdateSuccessful
```

## Check the Keystone Service deployment

```bash
kubectl get all

NAME                                               READY   STATUS      RESTARTS   AGE
pod/keystone-api-84f6b97b7d-9lh7b                  1/1     Running     0          2m43s
pod/keystone-armada-operator-5cfccc74fb-r7wlx      1/1     Running     0          3m10s
pod/keystone-bootstrap-bzc69                       0/1     Completed   0          2m43s
pod/keystone-credential-setup-4fm4k                0/1     Completed   0          2m43s
pod/keystone-db-init-cpcwv                         0/1     Completed   0          2m43s
pod/keystone-db-sync-2jjgm                         0/1     Completed   0          2m43s
pod/keystone-domain-manage-mn6l2                   0/1     Completed   0          2m43s
pod/keystone-fernet-setup-xrvcq                    0/1     Completed   0          2m43s
pod/keystone-oslc-operator-747bb84698-r8k8r        1/1     Running     0          3m6s
pod/keystone-rabbit-init-cpvc9                     0/1     Completed   0          2m43s
pod/mariadb-ingress-7bb7bd65ff-hpsc7               1/1     Running     0          2m48s
pod/mariadb-ingress-error-pages-869655cc86-cph5d   1/1     Running     0          2m48s
pod/mariadb-server-0                               1/1     Running     0          2m48s
pod/memcached-memcached-579f87f568-xvxb6           1/1     Running     0          2m47s
pod/rabbitmq-cluster-wait-wztz5                    0/1     Completed   0          2m45s
pod/rabbitmq-rabbitmq-0                            1/1     Running     0          2m45s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.97.70.16      <none>        80/TCP,443/TCP                 2m43s
service/keystone-api                  ClusterIP   10.100.217.228   <none>        5000/TCP                       2m43s
service/keystone-armada-operator      ClusterIP   10.103.102.104   <none>        8383/TCP                       3m8s
service/keystone-oslc-operator        ClusterIP   10.96.251.99     <none>        8383/TCP                       29m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        36m
service/mariadb                       ClusterIP   10.96.94.118     <none>        3306/TCP                       2m48s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              2m48s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         2m48s
service/mariadb-server                ClusterIP   10.105.147.246   <none>        3306/TCP                       2m48s
service/memcached                     ClusterIP   10.97.255.33     <none>        11211/TCP                      2m47s
service/rabbitmq                      ClusterIP   10.98.125.239    <none>        5672/TCP,25672/TCP,15672/TCP   2m45s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m45s
service/rabbitmq-mgr-7b1733           ClusterIP   10.102.235.132   <none>        80/TCP,443/TCP                 2m45s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           2m43s
deployment.apps/keystone-armada-operator      1/1     1            1           3m10s
deployment.apps/keystone-oslc-operator        1/1     1            1           3m6s
deployment.apps/mariadb-ingress               1/1     1            1           2m48s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           2m48s
deployment.apps/memcached-memcached           1/1     1            1           2m47s

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-84f6b97b7d                  1         1         1       2m43s
replicaset.apps/keystone-armada-operator-5cfccc74fb      1         1         1       3m10s
replicaset.apps/keystone-oslc-operator-747bb84698        1         1         1       3m6s
replicaset.apps/mariadb-ingress-7bb7bd65ff               1         1         1       2m48s
replicaset.apps/mariadb-ingress-error-pages-869655cc86   1         1         1       2m48s
replicaset.apps/memcached-memcached-579f87f568           1         1         1       2m47s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     2m48s
statefulset.apps/rabbitmq-rabbitmq   1/1     2m45s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          1/1           2m27s      2m43s
job.batch/keystone-credential-setup   1/1           15s        2m43s
job.batch/keystone-db-init            1/1           76s        2m43s
job.batch/keystone-db-sync            1/1           101s       2m43s
job.batch/keystone-domain-manage      1/1           2m15s      2m43s
job.batch/keystone-fernet-setup       1/1           15s        2m43s
job.batch/keystone-rabbit-init        1/1           37s        2m43s
job.batch/rabbitmq-cluster-wait       1/1           37s        2m45s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          2m43s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          2m43s
```

# Remove the CRD 

```bash
make deletemanifest
```

or

```bash
kubectl delete -f examples/keystone/

armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
```

Check the corresponding resources are being deleted

```bash
kubectl get all

NAME                                            READY   STATUS        RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-r7wlx   1/1     Running       0          7m1s
pod/keystone-oslc-operator-747bb84698-r8k8r     1/1     Running       0          6m57s
pod/mariadb-server-0                            0/1     Terminating   0          6m39s

NAME                               TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.103.102.104   <none>        8383/TCP   6m59s
service/keystone-oslc-operator     ClusterIP   10.96.251.99     <none>        8383/TCP   33m
service/kubernetes                 ClusterIP   10.96.0.1        <none>        443/TCP    39m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           7m1s
deployment.apps/keystone-oslc-operator     1/1     1            1           6m57s

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       7m1s
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       6m57s
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
make purge-akubectl delete -f deploy/armada-operator.yaml

deployment.apps "keystone-armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io "armadachartgroups.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io "armadacharts.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io "armadamanifests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount "armada-argo-sa" deleted
role.rbac.authorization.k8s.io "armada-argo-role" deleted
rolebinding.rbac.authorization.k8s.io "armada-argo-rolebinding" deleted
kubectl delete -f deploy/oslc-operator.yaml

deployment.apps "keystone-oslc-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io "deletephases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io "installphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io "operationalphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io "planningphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io "rollbackphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io "testphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io "trafficdrainphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io "trafficrolloutphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io "upgradephases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io "oslcs.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount "openstacklcm-argo-sa" deleted
role.rbac.authorization.k8s.io "openstacklcm-argo-role" deleted
rolebinding.rbac.authorization.k8s.io "openstacklcm-argo-rolebinding" deleted
```

```
kubectl get all

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   25m
```

Delete the configmaps related to keystone if needed
```
kubectl get configmaps

NAME                              DATA   AGE
mariadb-mariadb-mariadb-ingress   0      12m
mariadb-mariadb-state             5      12m
```

# Deployment using argo workflow

## Install the keystone operator

Ensure that argo is installed as well as helm

```bash
helm ls

NAME    REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
argo    1               Fri Mar 22 10:25:22 2019        DEPLOYED        argo-0.3.1                      argo
```

```bash
make install-armada && make install-olsc

kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount/armada-argo-sa created
role.rbac.authorization.k8s.io/armada-argo-role created
rolebinding.rbac.authorization.k8s.io/armada-argo-rolebinding created
kubectl create -f deploy/armada-operator.yaml
deployment.apps/keystone-armada-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io/deletephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io/installphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io/operationalphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io/planningphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io/rollbackphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io/testphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficdrainphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficrolloutphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io/upgradephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io/oslcs.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount/openstacklcm-argo-sa created
role.rbac.authorization.k8s.io/openstacklcm-argo-role created
rolebinding.rbac.authorization.k8s.io/openstacklcm-argo-rolebinding created
kubectl create -f deploy/oslc-operator.yaml
deployment.apps/keystone-oslc-operator created
```

## Create the ArmadaChart and ArgoWorkflow in K8s
```bash
kubectl apply -f examples/argo/

armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
workflow.argoproj.io/armada-manifest created
```

## Check the sequencement of the deployment

```bash
argo get armada-manifest

Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (20 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (20 seconds ago)
Duration:            20 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | └---◷ enable-memcached     armada-manifest-4290341812  3s        ContainerCreating
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (22 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (22 seconds ago)
Duration:            22 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | └---◷ memcached-ready      armada-manifest-13034731    0s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (25 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (25 seconds ago)
Duration:            25 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | └---● memcached-ready      armada-manifest-13034731    3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Mar 22 11:23:24 -0500 (33 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (33 seconds ago)
Duration:            33 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | ├---✔ memcached-ready      armada-manifest-13034731    4s
 | ├---✔ enable-rabbitmq      armada-manifest-1873087745  3s
 | └---● rabbitmq-ready       armada-manifest-3014489540  3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

```bash
argo get armada-manifest

Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Succeeded
Created:             Fri Mar 22 11:23:24 -0500 (35 seconds ago)
Started:             Fri Mar 22 11:23:24 -0500 (35 seconds ago)
Finished:            Fri Mar 22 11:23:58 -0500 (1 second ago)
Duration:            34 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ✔ armada-manifest
 ├-✔ keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   2s
 | ├---✔ mariadb-ready        armada-manifest-361794910   4s
 | ├---✔ enable-memcached     armada-manifest-4290341812  3s
 | ├---✔ memcached-ready      armada-manifest-13034731    4s
 | ├---✔ enable-rabbitmq      armada-manifest-1873087745  3s
 | └---✔ rabbitmq-ready       armada-manifest-3014489540  3s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   5s
```

## Check the state of the keystone service

Check that keystone is now up and running

```bash
NAME                                               READY   STATUS      RESTARTS   AGE
pod/armada-manifest-13034731                       0/1     Completed   0          5m6s
pod/armada-manifest-1873087745                     0/1     Completed   0          5m1s
pod/armada-manifest-2789817682                     0/1     Completed   0          5m20s
pod/armada-manifest-3014489540                     0/1     Completed   0          4m58s
pod/armada-manifest-361794910                      0/1     Completed   0          5m15s
pod/armada-manifest-4290341812                     0/1     Completed   0          5m9s
pod/armada-manifest-587324697                      0/1     Completed   0          5m15s
pod/armada-manifest-842022813                      0/1     Completed   0          5m20s
pod/keystone-api-84f6b97b7d-z7bxx                  1/1     Running     0          5m14s
pod/keystone-armada-operator-5cfccc74fb-kmhvp      1/1     Running     0          6m36s
pod/keystone-bootstrap-cj2q9                       0/1     Error       4          5m14s
pod/keystone-credential-setup-s5zxg                0/1     Completed   0          5m14s
pod/keystone-db-init-7xmth                         0/1     Completed   0          5m14s
pod/keystone-db-sync-nd2h5                         0/1     Completed   0          5m14s
pod/keystone-domain-manage-bskm7                   0/1     Completed   0          5m14s
pod/keystone-fernet-setup-wpxt4                    0/1     Completed   0          5m14s
pod/keystone-oslc-operator-747bb84698-sm5jg        1/1     Running     0          6m32s
pod/keystone-rabbit-init-2gg7q                     0/1     Completed   0          5m14s
pod/mariadb-ingress-7bb7bd65ff-bsztv               1/1     Running     0          5m16s
pod/mariadb-ingress-error-pages-869655cc86-nwqvx   1/1     Running     0          5m16s
pod/mariadb-server-0                               1/1     Running     0          5m16s
pod/memcached-memcached-579f87f568-cgswm           1/1     Running     0          5m5s
pod/rabbitmq-cluster-wait-hzq6s                    0/1     Completed   0          4m57s
pod/rabbitmq-rabbitmq-0                            1/1     Running     0          4m57s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.101.119.212   <none>        80/TCP,443/TCP                 5m14s
service/keystone-api                  ClusterIP   10.100.204.102   <none>        5000/TCP                       5m14s
service/keystone-armada-operator      ClusterIP   10.106.147.222   <none>        8383/TCP                       6m34s
service/keystone-oslc-operator        ClusterIP   10.96.251.99     <none>        8383/TCP                       6m30s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        13m
service/mariadb                       ClusterIP   10.101.28.26     <none>        3306/TCP                       5m16s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              5m16s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         5m16s
service/mariadb-server                ClusterIP   10.97.95.88      <none>        3306/TCP                       5m16s
service/memcached                     ClusterIP   10.110.229.196   <none>        11211/TCP                      5m5s
service/rabbitmq                      ClusterIP   10.99.22.59      <none>        5672/TCP,25672/TCP,15672/TCP   4m57s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   4m57s
service/rabbitmq-mgr-7b1733           ClusterIP   10.103.124.16    <none>        80/TCP,443/TCP                 4m57s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           5m14s
deployment.apps/keystone-armada-operator      1/1     1            1           6m36s
deployment.apps/keystone-oslc-operator        1/1     1            1           6m32s
deployment.apps/mariadb-ingress               1/1     1            1           5m16s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           5m16s
deployment.apps/memcached-memcached           1/1     1            1           5m5s

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-84f6b97b7d                  1         1         1       5m14s
replicaset.apps/keystone-armada-operator-5cfccc74fb      1         1         1       6m36s
replicaset.apps/keystone-oslc-operator-747bb84698        1         1         1       6m32s
replicaset.apps/mariadb-ingress-7bb7bd65ff               1         1         1       5m16s
replicaset.apps/mariadb-ingress-error-pages-869655cc86   1         1         1       5m16s
replicaset.apps/memcached-memcached-579f87f568           1         1         1       5m5s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     5m16s
statefulset.apps/rabbitmq-rabbitmq   1/1     4m57s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          0/1           5m14s      5m14s
job.batch/keystone-credential-setup   1/1           77s        5m14s
job.batch/keystone-db-init            1/1           2m21s      5m14s
job.batch/keystone-db-sync            1/1           2m41s      5m14s
job.batch/keystone-domain-manage      1/1           3m22s      5m14s
job.batch/keystone-fernet-setup       1/1           77s        5m14s
job.batch/keystone-rabbit-init        1/1           90s        5m14s
job.batch/rabbitmq-cluster-wait       1/1           75s        4m57s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          5m14s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          5m13s
```

## Delete the service

```bash
kubectl delete -f examples/argo

armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
workflow.argoproj.io "armada-manifest" deleted
```

Check in argo that the workflow is gone

```bash
argo list

NAME   STATUS   AGE   DURATION
```

Check in K8s that the keystone service is gone

```bash
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-kmhvp   1/1     Running   0          22m
pod/keystone-oslc-operator-747bb84698-sm5jg     1/1     Running   0          22m

NAME                               TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.106.147.222   <none>        8383/TCP   22m
service/keystone-oslc-operator     ClusterIP   10.96.251.99     <none>        8383/TCP   21m
service/kubernetes                 ClusterIP   10.96.0.1        <none>        443/TCP    28m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           22m
deployment.apps/keystone-oslc-operator     1/1     1            1           22m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       22m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       22m
```

## Delete the keystone operator

```bash
make purge-armada && make purge-oslc

kubectl delete -f deploy/armada-operator.yaml
deployment.apps "keystone-armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io "armadachartgroups.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io "armadacharts.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io "armadamanifests.armada.airshipit.org" deleted
kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount "armada-operator" deleted
kubectl delete -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount "armada-argo-sa" deleted
role.rbac.authorization.k8s.io "armada-argo-role" deleted
rolebinding.rbac.authorization.k8s.io "armada-argo-rolebinding" deleted
kubectl delete -f deploy/oslc-operator.yaml
deployment.apps "keystone-oslc-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io "deletephases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io "installphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io "operationalphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io "planningphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io "rollbackphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io "testphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io "trafficdrainphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io "trafficrolloutphases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io "upgradephases.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io "oslcs.openstacklcm.airshipit.org" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount "openstacklcm-operator" deleted
kubectl delete -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount "openstacklcm-argo-sa" deleted
role.rbac.authorization.k8s.io "openstacklcm-argo-role" deleted
rolebinding.rbac.authorization.k8s.io "openstacklcm-argo-rolebinding" deleted
```

# Testing the oslc-controller

## Check lifecycle manually at first

### Install the operators

```bash
make install-armada && make install-oslc

kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
customresourcedefinition.apiextensions.k8s.io/armadachartgroups.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
customresourcedefinition.apiextensions.k8s.io/armadacharts.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
customresourcedefinition.apiextensions.k8s.io/armadamanifests.armada.airshipit.org created
kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/role.yaml
role.rbac.authorization.k8s.io/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
serviceaccount/armada-operator created
kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
serviceaccount/armada-argo-sa created
role.rbac.authorization.k8s.io/armada-argo-role created
rolebinding.rbac.authorization.k8s.io/armada-argo-rolebinding created
kubectl create -f deploy/armada-operator.yaml
deployment.apps/keystone-armada-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
customresourcedefinition.apiextensions.k8s.io/deletephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
customresourcedefinition.apiextensions.k8s.io/installphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
customresourcedefinition.apiextensions.k8s.io/operationalphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
customresourcedefinition.apiextensions.k8s.io/planningphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
customresourcedefinition.apiextensions.k8s.io/rollbackphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
customresourcedefinition.apiextensions.k8s.io/testphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficdrainphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
customresourcedefinition.apiextensions.k8s.io/trafficrolloutphases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
customresourcedefinition.apiextensions.k8s.io/upgradephases.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
customresourcedefinition.apiextensions.k8s.io/oslcs.openstacklcm.airshipit.org created
kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
rolebinding.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
role.rbac.authorization.k8s.io/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
serviceaccount/openstacklcm-operator created
kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
serviceaccount/openstacklcm-argo-sa created
role.rbac.authorization.k8s.io/openstacklcm-argo-role created
rolebinding.rbac.authorization.k8s.io/openstacklcm-argo-rolebinding created
kubectl create -f deploy/oslc-operator.yaml
deployment.apps/keystone-oslc-operator created
```

Install the infra services

```bash
kubectl apply -f examples/keystone/infra.yaml

armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
```

```bash
kubectl get act

NAME        STATE      TARGET STATE   SATISFIED
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
```

```bash
kubectl get all

NAME                                               READY   STATUS      RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-dzfmz      1/1     Running     0          78s
pod/keystone-oslc-operator-747bb84698-ccgxq        1/1     Running     0          74s
pod/mariadb-ingress-7bb7bd65ff-lq799               0/1     Running     0          27s
pod/mariadb-ingress-error-pages-869655cc86-d2xgv   1/1     Running     0          27s
pod/mariadb-server-0                               0/1     Running     0          27s
pod/memcached-memcached-579f87f568-947zm           1/1     Running     0          26s
pod/rabbitmq-cluster-wait-xndv4                    0/1     Completed   0          24s
pod/rabbitmq-rabbitmq-0                            1/1     Running     0          24s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone-armada-operator      ClusterIP   10.104.191.57    <none>        8383/TCP                       76s
service/keystone-oslc-operator        ClusterIP   10.99.143.77     <none>        8383/TCP                       73s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        46m
service/mariadb                       ClusterIP   10.100.6.122     <none>        3306/TCP                       27s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              27s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         27s
service/mariadb-server                ClusterIP   10.108.166.206   <none>        3306/TCP                       27s
service/memcached                     ClusterIP   10.104.245.239   <none>        11211/TCP                      26s
service/rabbitmq                      ClusterIP   10.108.133.228   <none>        5672/TCP,25672/TCP,15672/TCP   24s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   24s
service/rabbitmq-mgr-7b1733           ClusterIP   10.101.43.52     <none>        80/TCP,443/TCP                 24s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator      1/1     1            1           78s
deployment.apps/keystone-oslc-operator        1/1     1            1           74s
deployment.apps/mariadb-ingress               0/1     1            0           27s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           27s
deployment.apps/memcached-memcached           1/1     1            1           26s

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb      1         1         1       78s
replicaset.apps/keystone-oslc-operator-747bb84698        1         1         1       74s
replicaset.apps/mariadb-ingress-7bb7bd65ff               1         1         0       27s
replicaset.apps/mariadb-ingress-error-pages-869655cc86   1         1         1       27s
replicaset.apps/memcached-memcached-579f87f568           1         1         1       26s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      0/1     27s
statefulset.apps/rabbitmq-rabbitmq   1/1     24s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           23s        24s
```

Install the chart using argo workflow instead of the kubernetesendpoint

```bash
cd helm-charts/keystone
helm install --name keystone --values ./lifecycle-values.yaml  .

NAME:   keystone
LAST DEPLOYED: Wed Apr  3 15:33:47 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME          DATA  AGE
keystone-bin  14    3s

==> v1/Deployment
NAME          READY  UP-TO-DATE  AVAILABLE  AGE
keystone-api  0/1    1           0          2s

==> v1/Job
NAME                       COMPLETIONS  DURATION  AGE
keystone-bootstrap         0/1          2s        2s
keystone-credential-setup  0/1          2s        2s
keystone-db-init           0/1          2s        2s
keystone-db-sync           0/1          2s        2s
keystone-domain-manage     0/1          2s        2s
keystone-fernet-setup      0/1          2s        2s
keystone-rabbit-init       0/1          2s        2s

==> v1/Pod(related)
NAME                             READY  STATUS    RESTARTS  AGE
keystone-api-7c45c747f-bwg8b     0/1    Init:0/1  0         2s
keystone-bootstrap-cdh26         0/1    Pending   0         2s
keystone-credential-setup-ht8rg  0/1    Pending   0         2s
keystone-db-init-jmsxk           0/1    Pending   0         2s
keystone-db-sync-pc87k           0/1    Pending   0         2s
keystone-domain-manage-dgkfj     0/1    Pending   0         2s
keystone-fernet-setup-vbplg      0/1    Pending   0         2s
keystone-rabbit-init-4pxjp       0/1    Pending   0         2s

==> v1/Role
NAME              AGE
wf-keystone-role  2s

==> v1/RoleBinding
NAME                     AGE
wf-keystone-rolebinding  2s

==> v1/Secret
NAME                      TYPE    DATA  AGE
keystone-credential-keys  Opaque  0     3s
keystone-db-admin         Opaque  2     3s
keystone-db-user          Opaque  2     3s
keystone-etc              Opaque  9     3s
keystone-fernet-keys      Opaque  0     3s
keystone-keystone-admin   Opaque  8     3s
keystone-keystone-test    Opaque  8     3s
keystone-rabbitmq-admin   Opaque  1     3s
keystone-rabbitmq-user    Opaque  1     3s

==> v1/Service
NAME          TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)         AGE
keystone      ClusterIP  10.97.109.34   <none>       80/TCP,443/TCP  2s
keystone-api  ClusterIP  10.109.216.81  <none>       5000/TCP        2s

==> v1/ServiceAccount
NAME                        SECRETS  AGE
keystone-api                1        3s
keystone-bootstrap          1        3s
keystone-credential-rotate  1        3s
keystone-credential-setup   1        2s
keystone-db-init            1        2s
keystone-db-sync            1        2s
keystone-domain-manage      1        2s
keystone-fernet-rotate      1        3s
keystone-fernet-setup       1        2s
keystone-rabbit-init        1        2s
keystone-test               1        2s
wf-keystone-sa              1        2s

==> v1alpha1/Workflow
NAME                 AGE
wf-bootstrap         2s
wf-keystone-install  2s

==> v1beta1/CronJob
NAME                        SCHEDULE      SUSPEND  ACTIVE  LAST SCHEDULE  AGE
keystone-credential-rotate  0 0 1 * *     False    0       <none>         2s
keystone-fernet-rotate      0 */12 * * *  False    0       <none>         2s

==> v1beta1/Ingress
NAME      HOSTS                                                         ADDRESS  PORTS  AGE
keystone  keystone,keystone.default,keystone.default.svc.cluster.local  80       2s

==> v1beta1/PodDisruptionBudget
NAME          MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
keystone-api  0              N/A              0                    3s

==> v1beta1/Role
NAME                                         AGE
keystone-credential-rotate                   2s
keystone-credential-setup                    2s
keystone-default-keystone-bootstrap          2s
keystone-default-keystone-credential-rotate  2s
keystone-default-keystone-domain-manage      2s
keystone-default-keystone-fernet-rotate      2s
keystone-default-keystone-test               2s
keystone-fernet-rotate                       2s
keystone-fernet-setup                        2s

==> v1beta1/RoleBinding
NAME                                 AGE
keystone-credential-rotate           2s
keystone-credential-setup            2s
keystone-fernet-rotate               2s
keystone-fernet-setup                2s
keystone-keystone-bootstrap          2s
keystone-keystone-credential-rotate  2s
keystone-keystone-domain-manage      2s
keystone-keystone-fernet-rotate      2s
keystone-keystone-test               2s
```

Check argo
```bash
argo list

NAME                  STATUS      AGE   DURATION
wf-bootstrap          Running     2m    2m
wf-keystone-install   Succeeded   2m    53s
```

```bash
argo get wf-keystone-install

Name:                wf-keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Succeeded
Created:             Wed Apr 03 14:35:38 -0500 (1 minute ago)
Started:             Wed Apr 03 14:35:38 -0500 (1 minute ago)
Finished:            Wed Apr 03 14:36:28 -0500 (45 seconds ago)
Duration:            50 seconds

STEP                                    PODNAME                         DURATION  MESSAGE
 ✔ wf-keystone-install
 ├-✔ svc-mariadb                        wf-keystone-install-3672310916  11s
 ├-✔ svc-memcached                      wf-keystone-install-33467967    13s
 ├-✔ svc-rabbitmq                       wf-keystone-install-724228482   13s
 ├-✔ task-keystone-credential-setup(0)  wf-keystone-install-3464818     22s
 ├-✔ task-keystone-fernet-setup(0)      wf-keystone-install-376633117   24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       wf-keystone-install-3134557299  5s
 | └---✔ task-keystone-db-init(0)       wf-keystone-install-24327275    7s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      wf-keystone-install-1938236639  5s
 | └---✔ task-keystone-rabbit-init(0)   wf-keystone-install-2710739935  6s
 └-✔ wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       wf-keystone-install-2782100822  4s
   | └-✔ svc-rabbitmq-just-in-time      wf-keystone-install-1483869462  4s
   └---✔ task-keystone-db-sync(0)       wf-keystone-install-823103181   12s
```

```bash
argo get wf-bootstrap

Name:                wf-bootstrap
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Failed
Created:             Wed Apr 03 14:35:38 -0500 (2 minutes ago)
Started:             Wed Apr 03 14:35:38 -0500 (2 minutes ago)
Finished:            Wed Apr 03 14:36:38 -0500 (1 minute ago)
Duration:            1 minute 0 seconds

STEP                              PODNAME                  DURATION  MESSAGE
 ✖ wf-bootstrap
 ├-✔ svc-keystone                 wf-bootstrap-636750498   37s
 ├-✔ wf-domain-manage
 | ├---✔ task-domain-manage-init  wf-bootstrap-2275099864  6s
 | └---✔ task-domain-manage       wf-bootstrap-458835584   5s
 └-✖ task-bootstrap               wf-bootstrap-2256133693  6s        failed with exit code 1
```

Let's check kubernetes
```bash

NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-7c45c747f-bwg8b                  0/1     Init:0/1    0          3m46s
pod/keystone-armada-operator-5cfccc74fb-6qw79     1/1     Running     0          6m36s
pod/keystone-bootstrap-cdh26                      0/1     Init:0/1    0          3m46s
pod/keystone-credential-setup-ht8rg               0/1     Completed   1          3m46s
pod/keystone-db-init-jmsxk                        0/1     Completed   0          3m46s
pod/keystone-db-sync-pc87k                        0/1     Completed   0          3m46s
pod/keystone-domain-manage-dgkfj                  0/1     Init:0/2    0          3m46s
pod/keystone-fernet-setup-vbplg                   0/1     Completed   1          3m46s
pod/keystone-oslc-operator-747bb84698-lh2sc       1/1     Running     0          6m32s
pod/keystone-rabbit-init-4pxjp                    0/1     Completed   0          3m46s
pod/mariadb-ingress-54fbf9dffb-m4tqc              1/1     Running     0          6m12s
pod/mariadb-ingress-error-pages-fd9b689f9-5z2xj   1/1     Running     0          6m12s
pod/mariadb-server-0                              1/1     Running     0          6m12s
pod/memcached-memcached-75447bffcf-8xtjx          1/1     Running     0          6m11s
pod/rabbitmq-cluster-wait-kb9ns                   0/1     Completed   0          6m10s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          6m10s
pod/wf-bootstrap-636750498                        2/2     Running     0          3m46s
pod/wf-keystone-install-1483869462                0/1     Completed   0          3m11s
pod/wf-keystone-install-1938236639                0/1     Completed   0          3m29s
pod/wf-keystone-install-24327275                  0/2     Completed   0          3m22s
pod/wf-keystone-install-2710739935                0/2     Completed   0          3m21s
pod/wf-keystone-install-2782100822                0/1     Completed   0          3m11s
pod/wf-keystone-install-3134557299                0/1     Completed   0          3m30s
pod/wf-keystone-install-33467967                  0/2     Completed   0          3m46s
pod/wf-keystone-install-3464818                   0/2     Completed   0          3m46s
pod/wf-keystone-install-3672310916                0/2     Completed   0          3m46s
pod/wf-keystone-install-376633117                 0/2     Completed   0          3m46s
pod/wf-keystone-install-724228482                 0/2     Completed   0          3m46s
pod/wf-keystone-install-823103181                 0/2     Completed   0          3m4s

NAME                                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.97.109.34    <none>        80/TCP,443/TCP                 3m46s
service/keystone-api                  ClusterIP   10.109.216.81   <none>        5000/TCP                       3m46s
service/keystone-armada-operator      ClusterIP   10.111.11.252   <none>        8383/TCP                       6m34s
service/keystone-oslc-operator        ClusterIP   10.99.191.7     <none>        8383/TCP                       6m30s
service/kubernetes                    ClusterIP   10.96.0.1       <none>        443/TCP                        10m
service/mariadb                       ClusterIP   10.98.8.153     <none>        3306/TCP                       6m12s
service/mariadb-discovery             ClusterIP   None            <none>        3306/TCP,4567/TCP              6m12s
service/mariadb-ingress-error-pages   ClusterIP   None            <none>        80/TCP                         6m12s
service/mariadb-server                ClusterIP   10.98.5.233     <none>        3306/TCP                       6m12s
service/memcached                     ClusterIP   10.103.97.4     <none>        11211/TCP                      6m11s
service/rabbitmq                      ClusterIP   10.101.6.23     <none>        5672/TCP,25672/TCP,15672/TCP   6m10s
service/rabbitmq-dsv-7b1733           ClusterIP   None            <none>        5672/TCP,25672/TCP,15672/TCP   6m10s
service/rabbitmq-mgr-7b1733           ClusterIP   10.101.77.163   <none>        80/TCP,443/TCP                 6m10s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  0/1     1            0           3m46s
deployment.apps/keystone-armada-operator      1/1     1            1           6m36s
deployment.apps/keystone-oslc-operator        1/1     1            1           6m32s
deployment.apps/mariadb-ingress               1/1     1            1           6m12s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           6m12s
deployment.apps/memcached-memcached           1/1     1            1           6m11s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-7c45c747f                  1         1         0       3m46s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       6m36s
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       6m32s
replicaset.apps/mariadb-ingress-54fbf9dffb              1         1         1       6m12s
replicaset.apps/mariadb-ingress-error-pages-fd9b689f9   1         1         1       6m12s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       6m11s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     6m12s
statefulset.apps/rabbitmq-rabbitmq   1/1     6m10s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          0/1           3m46s      3m46s
job.batch/keystone-credential-setup   1/1           43s        3m46s
job.batch/keystone-db-init            1/1           15s        3m46s
job.batch/keystone-db-sync            1/1           48s        3m46s
job.batch/keystone-domain-manage      0/1           3m46s      3m46s
job.batch/keystone-fernet-setup       1/1           43s        3m46s
job.batch/keystone-rabbit-init        1/1           18s        3m46s
job.batch/rabbitmq-cluster-wait       1/1           33s        6m10s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          3m46s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          3m46s
```

Check the error in argo

```bash
argo logs wf-bootstrap-2256133693

+ openstack role create --or-show member
An unexpected error prevented the server from fulfilling your request. (HTTP 500) (Request-ID: req-492c9d66-96a1-46f6-99b1-4439fd203362)
```

