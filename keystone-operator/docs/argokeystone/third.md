# ./argokeystone/third.md

```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          12m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          12m
default       pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          36s
default       pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          31s
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          12m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          12m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          12m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          12m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          12m
kube-system   pod/etcd-airship                                1/1     Running   0          12m
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          12m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          12m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          12m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          12m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          12m

NAMESPACE     NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui                    NodePort    10.107.6.139    <none>        80:32711/TCP             12m
default       service/keystone-armada-operator   ClusterIP   10.102.35.39    <none>        8383/TCP                 34s
default       service/keystone-oslc-operator     ClusterIP   10.97.81.146    <none>        8383/TCP                 29s
default       service/kubernetes                 ClusterIP   10.96.0.1       <none>        443/TCP                  13m
kube-system   service/calico-etcd                ClusterIP   10.96.232.136   <none>        6666/TCP                 12m
kube-system   service/kube-dns                   ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   13m
kube-system   service/tiller-deploy              ClusterIP   10.102.66.38    <none>        44134/TCP                12m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   12m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       12m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            13m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           12m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           12m
default       deployment.apps/keystone-armada-operator   1/1     1            1           37s
default       deployment.apps/keystone-oslc-operator     1/1     1            1           31s
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           12m
kube-system   deployment.apps/coredns                    2/2     2            2           13m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           12m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       12m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       12m
default       replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       36s
default       replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       31s
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       12m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       12m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       12m
```
 
```bash
prompt$ clear
```
 
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          12m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          12m
default       pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          54s
default       pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          49s
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          12m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          13m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          13m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          13m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          13m
kube-system   pod/etcd-airship                                1/1     Running   0          12m
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          12m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          12m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          13m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          12m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          13m

NAMESPACE     NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui                    NodePort    10.107.6.139    <none>        80:32711/TCP             12m
default       service/keystone-armada-operator   ClusterIP   10.102.35.39    <none>        8383/TCP                 52s
default       service/keystone-oslc-operator     ClusterIP   10.97.81.146    <none>        8383/TCP                 47s
default       service/kubernetes                 ClusterIP   10.96.0.1       <none>        443/TCP                  13m
kube-system   service/calico-etcd                ClusterIP   10.96.232.136   <none>        6666/TCP                 13m
kube-system   service/kube-dns                   ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   13m
kube-system   service/tiller-deploy              ClusterIP   10.102.66.38    <none>        44134/TCP                13m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   13m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       13m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            13m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           12m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           12m
default       deployment.apps/keystone-armada-operator   1/1     1            1           55s
default       deployment.apps/keystone-oslc-operator     1/1     1            1           49s
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           13m
kube-system   deployment.apps/coredns                    2/2     2            2           13m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           13m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       12m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       12m
default       replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       54s
default       replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       49s
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       13m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       13m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       13m
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
No resources found.
kubectl get armadacharts.armada.airshipit.org
No resources found.
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ kubectl apply -f examples/keystone/infra.yaml
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
No resources found.
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          3m51s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          3m46s
pod/mariadb-ingress-5cfc68b86f-g7sgm              1/1     Running     0          93s
pod/mariadb-ingress-error-pages-67c88c67d-s5r7s   1/1     Running     0          94s
pod/mariadb-server-0                              1/1     Running     0          93s
pod/memcached-memcached-75447bffcf-bgxft          1/1     Running     0          92s
pod/rabbitmq-cluster-wait-ztvqh                   0/1     Completed   0          90s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          90s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       3m49s
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       3m44s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        16m
service/mariadb                       ClusterIP   10.98.4.164      <none>        3306/TCP                       94s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              94s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         94s
service/mariadb-server                ClusterIP   10.105.137.86    <none>        3306/TCP                       94s
service/memcached                     ClusterIP   10.100.216.105   <none>        11211/TCP                      92s
service/rabbitmq                      ClusterIP   10.98.191.40     <none>        5672/TCP,25672/TCP,15672/TCP   90s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   90s
service/rabbitmq-mgr-7b1733           ClusterIP   10.97.20.244     <none>        80/TCP,443/TCP                 90s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator      1/1     1            1           3m52s
deployment.apps/keystone-oslc-operator        1/1     1            1           3m46s
deployment.apps/mariadb-ingress               1/1     1            1           94s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           94s
deployment.apps/memcached-memcached           1/1     1            1           92s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       3m51s
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       3m46s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       94s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       94s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       92s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     93s
statefulset.apps/rabbitmq-rabbitmq   1/1     90s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           25s        90s
```
 
```bash
prompt$ vi second.txt
```
 
```bash
prompt$ kubectl apply -f examples/keystone/keystone.yaml
armadachart.armada.airshipit.org/keystone created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   8s
keystone-install     8s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Running
Created:             Fri Apr 05 07:57:35 -0500 (30 seconds ago)
Started:             Fri Apr 05 07:57:35 -0500 (30 seconds ago)
Duration:            30 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ● keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 └-✔ wf-keystone-rabbit-init
   ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
   └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Running
Created:             Fri Apr 05 07:57:35 -0500 (38 seconds ago)
Started:             Fri Apr 05 07:57:35 -0500 (38 seconds ago)
Duration:            38 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ● keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
 └-● wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  4s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---◷ task-keystone-db-sync(0)       keystone-install-1630370687  2s        ContainerCreating
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Succeeded
Created:             Fri Apr 05 07:57:35 -0500 (1 minute ago)
Started:             Fri Apr 05 07:57:35 -0500 (1 minute ago)
Finished:            Fri Apr 05 07:58:34 -0500 (2 seconds ago)
Duration:            59 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ✔ keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
 └-✔ wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  4s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---✔ task-keystone-db-sync(0)       keystone-install-1630370687  21s
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     1m    1m
keystone-install     Succeeded   1m    59s
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-6d7bd74f96-6s8qk                 1/1     Running     0          2m36s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          7m21s
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          40s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          2m36s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          33s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          44s
pod/keystone-install-1608858220                   0/2     Completed   0          2m36s
pod/keystone-install-1630370687                   0/2     Completed   0          2m
pod/keystone-install-1659018021                   0/2     Completed   0          2m18s
pod/keystone-install-2012559777                   0/1     Completed   0          2m24s
pod/keystone-install-2401658373                   0/1     Completed   0          2m24s
pod/keystone-install-2716093231                   0/2     Completed   0          2m36s
pod/keystone-install-2958852122                   0/2     Completed   0          2m36s
pod/keystone-install-3754811077                   0/2     Completed   0          2m36s
pod/keystone-install-3843173896                   0/1     Completed   0          2m6s
pod/keystone-install-421213240                    0/2     Completed   0          2m36s
pod/keystone-install-4222398536                   0/1     Completed   0          2m6s
pod/keystone-install-494016497                    0/2     Completed   0          2m15s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          7m16s
pod/mariadb-ingress-5cfc68b86f-g7sgm              1/1     Running     0          5m3s
pod/mariadb-ingress-error-pages-67c88c67d-s5r7s   1/1     Running     0          5m4s
pod/mariadb-server-0                              1/1     Running     0          5m3s
pod/memcached-memcached-75447bffcf-bgxft          1/1     Running     0          5m2s
pod/rabbitmq-cluster-wait-ztvqh                   0/1     Completed   0          5m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          5m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.103.84.0      <none>        80/TCP,443/TCP                 2m36s
service/keystone-api                  ClusterIP   10.98.137.145    <none>        5000/TCP                       2m36s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       7m19s
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       7m14s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        19m
service/mariadb                       ClusterIP   10.98.4.164      <none>        3306/TCP                       5m4s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              5m4s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         5m4s
service/mariadb-server                ClusterIP   10.105.137.86    <none>        3306/TCP                       5m4s
service/memcached                     ClusterIP   10.100.216.105   <none>        11211/TCP                      5m2s
service/rabbitmq                      ClusterIP   10.98.191.40     <none>        5672/TCP,25672/TCP,15672/TCP   5m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   5m
service/rabbitmq-mgr-7b1733           ClusterIP   10.97.20.244     <none>        80/TCP,443/TCP                 5m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           2m36s
deployment.apps/keystone-armada-operator      1/1     1            1           7m22s
deployment.apps/keystone-oslc-operator        1/1     1            1           7m16s
deployment.apps/mariadb-ingress               1/1     1            1           5m4s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           5m4s
deployment.apps/memcached-memcached           1/1     1            1           5m2s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       2m36s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       7m21s
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       7m16s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       5m4s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       5m4s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       5m2s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     5m3s
statefulset.apps/rabbitmq-rabbitmq   1/1     5m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           25s        5m

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          2m36s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          2m36s
```
 
```bash
prompt$ argo logs keystone-install -w
svc-mariadb:    Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Invalid format:
svc-mariadb:    Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Resolving Service mariadb in namespace default
svc-mariadb:    Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Dependency Service mariadb in namespace default is resolved.
svc-mariadb:    done
svc-memcached:  Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Invalid format:
svc-memcached:  Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Resolving Service memcached in namespace default
svc-memcached:  Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Dependency Service memcached in namespace default is resolved.
svc-memcached:  done
svc-rabbitmq:   Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Invalid format:
svc-rabbitmq:   Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Resolving Service rabbitmq in namespace default
svc-rabbitmq:   Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Dependency Service rabbitmq in namespace default is resolved.
svc-rabbitmq:   done
task-keystone-credential-setup(0):      2019-04-05 12:57:47.445 - INFO - Executing 'keystone-manage credential_setup --keystone-user=keystone --keystone-group=keystone' command.
task-keystone-fernet-setup(0):  2019-04-05 12:57:48.681 - INFO - Executing 'keystone-manage fernet_setup --keystone-user=keystone --keystone-group=keystone' command.
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Creating a docker executor"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: mariadb\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"mariadb\"\n"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Creating a docker executor"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: rabbitmq\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"rabbitmq\"\n"
svc-rabbitmq-just-in-time:      kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:52Z" level=info msg=default/Service./mariadb
svc-mariadb-just-in-time:       time="2019-04-05T12:57:52Z" level=info msg="No output parameters"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:53Z" level=info msg=default/Service./rabbitmq
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:53Z" level=info msg="No output parameters"
task-keystone-credential-setup(0):      2019-04-05 12:57:57.902 8 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/credential-keys/
task-keystone-credential-setup(0):      2019-04-05 12:57:57.902 8 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/credential-keys/
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/credential-keys/0']
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/credential-keys/0']
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.909 8 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.909 8 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:58.011 - INFO - Updating data for 'keystone-credential-keys' secret.
task-keystone-credential-setup(0):      2019-04-05 12:57:58.057 - INFO - 2 fernet keys have been placed to secret 'keystone-credential-keys'
task-keystone-credential-setup(0):      2019-04-05 12:57:58.057 - INFO - Credential keys generation has been completed
task-keystone-db-init(0):       2019-04-05 12:57:58,718 - OpenStack-Helm DB Init - INFO - Got DB root connection
task-keystone-db-init(0):       2019-04-05 12:57:58,718 - OpenStack-Helm DB Init - INFO - Using /etc/keystone/keystone.conf as db config source
task-keystone-db-init(0):       2019-04-05 12:57:58,720 - OpenStack-Helm DB Init - INFO - Trying to load db config from database:connection
task-keystone-db-init(0):       2019-04-05 12:57:58,720 - OpenStack-Helm DB Init - INFO - Got config from /etc/keystone/keystone.conf
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.761 9 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/fernet-keys/
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.761 9 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/fernet-keys/
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.762 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.762 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/fernet-keys/0']
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/fernet-keys/0']
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.765 9 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.765 9 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.864 - INFO - Updating data for 'keystone-fernet-keys' secret.
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.906 - INFO - 2 fernet keys have been placed to secret 'keystone-fernet-keys'
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.906 - INFO - Fernet keys generation has been completed
task-keystone-db-init(0):       2019-04-05 12:57:58,946 - OpenStack-Helm DB Init - INFO - Tested connection to DB @ mariadb.default.svc.cluster.local:3306 as root
task-keystone-db-init(0):       2019-04-05 12:57:58,948 - OpenStack-Helm DB Init - INFO - Got user db config
task-keystone-db-init(0):       2019-04-05 12:57:58,981 - OpenStack-Helm DB Init - INFO - Created database keystone
task-keystone-db-init(0):       2019-04-05 12:57:58,991 - OpenStack-Helm DB Init - INFO - Created user keystone for keystone
task-keystone-db-init(0):       2019-04-05 12:57:59,062 - OpenStack-Helm DB Init - INFO - Tested connection to DB @ mariadb.default.svc.cluster.local:3306/keystone as keystone
task-keystone-db-init(0):       2019-04-05 12:57:59,063 - OpenStack-Helm DB Init - INFO - Finished DB Management
task-keystone-rabbit-init(0):   Managing: User: keystone
task-keystone-rabbit-init(0):   user declared
task-keystone-rabbit-init(0):   Managing: vHost: keystone
task-keystone-rabbit-init(0):   vhost declared
task-keystone-rabbit-init(0):   Managing: Permissions: keystone on keystone
task-keystone-rabbit-init(0):   permission declared
task-keystone-rabbit-init(0):   Applying additional configuration
task-keystone-rabbit-init(0):   Imported definitions for rabbitmq.default.svc.cluster.local from "/tmp/rmq_definitions.json"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Creating a docker executor"
svc-rabbitmq-just-in-time:      kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: rabbitmq\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"rabbitmq\"\n"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Creating a docker executor"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: mariadb\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"mariadb\"\n"
svc-mariadb-just-in-time:       kubectl get -f /tmp/manifest.yaml -o json
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:09Z" level=info msg=default/Service./rabbitmq
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:09Z" level=info msg="No output parameters"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:09Z" level=info msg=default/Service./mariadb
svc-mariadb-just-in-time:       time="2019-04-05T12:58:09Z" level=info msg="No output parameters"
task-keystone-db-sync(0):       + keystone-manage --config-file=/etc/keystone/keystone.conf db_sync
task-keystone-db-sync(0):       2019-04-05 12:58:22.006 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`inherited` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       2019-04-05 12:58:22.197 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`enabled` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       2019-04-05 12:58:23.309 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`impersonation` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       + keystone-manage --config-file=/etc/keystone/keystone.conf bootstrap --bootstrap-username admin --bootstrap-password password --bootstrap-project-name admin --bootstrap-admin-url http://keystone.default.svc.cluster.local:80/v3 --bootstrap-public-url http://keystone.default.svc.cluster.local:80/v3 --bootstrap-internal-url http://keystone-api.default.svc.cluster.local:5000/v3 --bootstrap-region-id RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:31.547 11 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/pycadf/identifier.py:60: UserWarning: Invalid uuid. To ensure interoperability, identifiers should be a valid uuid.
task-keystone-db-sync(0):         warnings.warn('Invalid uuid. To ensure interoperability, identifiers '
task-keystone-db-sync(0):       2019-04-05 12:58:31.806 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created domain default
task-keystone-db-sync(0):       2019-04-05 12:58:31.806 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created domain default
task-keystone-db-sync(0):       2019-04-05 12:58:31.851 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created project admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.851 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created project admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.852 11 WARNING keystone.identity.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Unable to locate domain config directory: /etc/keystonedomains
task-keystone-db-sync(0):       2019-04-05 12:58:31.852 11 WARNING keystone.identity.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Unable to locate domain config directory: /etc/keystonedomains
task-keystone-db-sync(0):       2019-04-05 12:58:31.937 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created user admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.937 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created user admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.955 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created role admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.955 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created role admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.970 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Granted admin on admin to user admin.
task-keystone-db-sync(0):       2019-04-05 12:58:31.970 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Granted admin on admin to user admin.
task-keystone-db-sync(0):       2019-04-05 12:58:32.004 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created region RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:32.004 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created region RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:32.049 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created admin endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.049 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created admin endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.064 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created internal endpoint http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.064 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created internal endpoint http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.083 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created public endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.083 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created public endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.089 11 INFO keystone.assignment.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Creating the default role 9fe2ff9ee4384b1894a90878d3e92bab because it does not exist.
task-keystone-db-sync(0):       2019-04-05 12:58:32.089 11 INFO keystone.assignment.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Creating the default role 9fe2ff9ee4384b1894a90878d3e92bab because it does not exist.
task-keystone-db-sync(0):       + exec python /tmp/endpoint-update.py
task-keystone-db-sync(0):       2019-04-05 12:58:32,433 - OpenStack-Helm Keystone Endpoint management - INFO - Using /etc/keystone/keystone.conf as db config source
task-keystone-db-sync(0):       2019-04-05 12:58:32,439 - OpenStack-Helm Keystone Endpoint management - INFO - Trying to load db config from database:connection
task-keystone-db-sync(0):       2019-04-05 12:58:32,440 - OpenStack-Helm Keystone Endpoint management - INFO - Got config from /etc/keystone/keystone.conf
task-keystone-db-sync(0):       2019-04-05 12:58:32,521 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (admin): http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (public): http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (internal): http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - Finished Endpoint Management
```
 
```bash
prompt$ argo logs keystone-bootstrap -w
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Invalid format:
svc-keystone:   Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Resolving Service keystone-api in namespace default
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 2.000 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 2.500 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:45 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 3.125 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:48 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 3.906 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:52 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 4.883 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:56 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 6.104 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:03 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 7.629 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:10 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 9.537 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:20 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 11.921 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:32 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 14.901 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:47 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 18.626 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:59:05 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 20.000 seconds
svc-keystone:   Entrypoint INFO: 2019/04/05 12:59:25 logger.go:32: Dependency Service keystone-api in namespace default is resolved.
svc-keystone:   done
task-bootstrap: + openstack role create --or-show member
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: | Field     | Value                            |
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: | domain_id | None                             |
task-bootstrap: | id        | ccbaa551824c4a519f2ec42501c493e4 |
task-bootstrap: | name      | member                           |
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: + openstack role add --user=admin --user-domain=default --project-domain=default --project=admin member
```
 
```bash
prompt$
```
# ./argokeystone/third.md
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          12m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          12m
default       pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          36s
default       pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          31s
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          12m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          12m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          12m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          12m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          12m
kube-system   pod/etcd-airship                                1/1     Running   0          12m
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          12m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          12m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          12m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          12m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          12m

NAMESPACE     NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui                    NodePort    10.107.6.139    <none>        80:32711/TCP             12m
default       service/keystone-armada-operator   ClusterIP   10.102.35.39    <none>        8383/TCP                 34s
default       service/keystone-oslc-operator     ClusterIP   10.97.81.146    <none>        8383/TCP                 29s
default       service/kubernetes                 ClusterIP   10.96.0.1       <none>        443/TCP                  13m
kube-system   service/calico-etcd                ClusterIP   10.96.232.136   <none>        6666/TCP                 12m
kube-system   service/kube-dns                   ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   13m
kube-system   service/tiller-deploy              ClusterIP   10.102.66.38    <none>        44134/TCP                12m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   12m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       12m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            13m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           12m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           12m
default       deployment.apps/keystone-armada-operator   1/1     1            1           37s
default       deployment.apps/keystone-oslc-operator     1/1     1            1           31s
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           12m
kube-system   deployment.apps/coredns                    2/2     2            2           13m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           12m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       12m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       12m
default       replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       36s
default       replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       31s
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       12m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       12m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       12m
```
 
```bash
prompt$ clear
```
 
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                            READY   STATUS    RESTARTS   AGE
argo          pod/argo-ui-f57bb579c-swz7h                     1/1     Running   0          12m
argo          pod/argo-workflow-controller-6d987f7fdc-l5cln   1/1     Running   0          12m
default       pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          54s
default       pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          49s
kube-system   pod/calico-etcd-k7h9l                           1/1     Running   0          12m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9    1/1     Running   0          13m
kube-system   pod/calico-node-qjfj8                           1/1     Running   1          13m
kube-system   pod/coredns-fb8b8dccf-kvdg6                     1/1     Running   0          13m
kube-system   pod/coredns-fb8b8dccf-p9zhw                     1/1     Running   0          13m
kube-system   pod/etcd-airship                                1/1     Running   0          12m
kube-system   pod/kube-apiserver-airship                      1/1     Running   0          12m
kube-system   pod/kube-controller-manager-airship             1/1     Running   0          12m
kube-system   pod/kube-proxy-dv5ls                            1/1     Running   0          13m
kube-system   pod/kube-scheduler-airship                      1/1     Running   0          12m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z              1/1     Running   0          13m

NAMESPACE     NAME                               TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
argo          service/argo-ui                    NodePort    10.107.6.139    <none>        80:32711/TCP             12m
default       service/keystone-armada-operator   ClusterIP   10.102.35.39    <none>        8383/TCP                 52s
default       service/keystone-oslc-operator     ClusterIP   10.97.81.146    <none>        8383/TCP                 47s
default       service/kubernetes                 ClusterIP   10.96.0.1       <none>        443/TCP                  13m
kube-system   service/calico-etcd                ClusterIP   10.96.232.136   <none>        6666/TCP                 13m
kube-system   service/kube-dns                   ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   13m
kube-system   service/tiller-deploy              ClusterIP   10.102.66.38    <none>        44134/TCP                13m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   13m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       13m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            13m

NAMESPACE     NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
argo          deployment.apps/argo-ui                    1/1     1            1           12m
argo          deployment.apps/argo-workflow-controller   1/1     1            1           12m
default       deployment.apps/keystone-armada-operator   1/1     1            1           55s
default       deployment.apps/keystone-oslc-operator     1/1     1            1           49s
kube-system   deployment.apps/calico-kube-controllers    1/1     1            1           13m
kube-system   deployment.apps/coredns                    2/2     2            2           13m
kube-system   deployment.apps/tiller-deploy              1/1     1            1           13m

NAMESPACE     NAME                                                  DESIRED   CURRENT   READY   AGE
argo          replicaset.apps/argo-ui-f57bb579c                     1         1         1       12m
argo          replicaset.apps/argo-workflow-controller-6d987f7fdc   1         1         1       12m
default       replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       54s
default       replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       49s
kube-system   replicaset.apps/calico-kube-controllers-79bd896977    1         1         1       13m
kube-system   replicaset.apps/coredns-fb8b8dccf                     2         2         2       13m
kube-system   replicaset.apps/tiller-deploy-8458f6c667              1         1         1       13m
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
No resources found.
kubectl get armadacharts.armada.airshipit.org
No resources found.
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ kubectl apply -f examples/keystone/infra.yaml
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
No resources found.
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          3m51s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          3m46s
pod/mariadb-ingress-5cfc68b86f-g7sgm              1/1     Running     0          93s
pod/mariadb-ingress-error-pages-67c88c67d-s5r7s   1/1     Running     0          94s
pod/mariadb-server-0                              1/1     Running     0          93s
pod/memcached-memcached-75447bffcf-bgxft          1/1     Running     0          92s
pod/rabbitmq-cluster-wait-ztvqh                   0/1     Completed   0          90s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          90s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       3m49s
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       3m44s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        16m
service/mariadb                       ClusterIP   10.98.4.164      <none>        3306/TCP                       94s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              94s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         94s
service/mariadb-server                ClusterIP   10.105.137.86    <none>        3306/TCP                       94s
service/memcached                     ClusterIP   10.100.216.105   <none>        11211/TCP                      92s
service/rabbitmq                      ClusterIP   10.98.191.40     <none>        5672/TCP,25672/TCP,15672/TCP   90s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   90s
service/rabbitmq-mgr-7b1733           ClusterIP   10.97.20.244     <none>        80/TCP,443/TCP                 90s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator      1/1     1            1           3m52s
deployment.apps/keystone-oslc-operator        1/1     1            1           3m46s
deployment.apps/mariadb-ingress               1/1     1            1           94s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           94s
deployment.apps/memcached-memcached           1/1     1            1           92s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       3m51s
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       3m46s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       94s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       94s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       92s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     93s
statefulset.apps/rabbitmq-rabbitmq   1/1     90s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           25s        90s
```
 
```bash
prompt$ vi second.txt
```
 
```bash
prompt$ kubectl apply -f examples/keystone/keystone.yaml
armadachart.armada.airshipit.org/keystone created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   8s
keystone-install     8s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
No resources found.
kubectl get armadamanifests.armada.airshipit.org
No resources found.
kubectl get oslcs.openstacklcm.airshipit.org
No resources found.
kubectl get installphases.openstacklcm.airshipit.org
No resources found.
kubectl get rollbackphases.openstacklcm.airshipit.org
No resources found.
kubectl get testphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficdrainphases.openstacklcm.airshipit.org
No resources found.
kubectl get trafficrolloutphases.openstacklcm.airshipit.org
No resources found.
kubectl get upgradephases.openstacklcm.airshipit.org
No resources found.
kubectl get deletephases.openstacklcm.airshipit.org
No resources found.
kubectl get planningphases.openstacklcm.airshipit.org
No resources found.
kubectl get operationalphases.openstacklcm.airshipit.org
No resources found.
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Running
Created:             Fri Apr 05 07:57:35 -0500 (30 seconds ago)
Started:             Fri Apr 05 07:57:35 -0500 (30 seconds ago)
Duration:            30 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ● keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 └-✔ wf-keystone-rabbit-init
   ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
   └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Running
Created:             Fri Apr 05 07:57:35 -0500 (38 seconds ago)
Started:             Fri Apr 05 07:57:35 -0500 (38 seconds ago)
Duration:            38 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ● keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
 └-● wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  4s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---◷ task-keystone-db-sync(0)       keystone-install-1630370687  2s        ContainerCreating
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Succeeded
Created:             Fri Apr 05 07:57:35 -0500 (1 minute ago)
Started:             Fri Apr 05 07:57:35 -0500 (1 minute ago)
Finished:            Fri Apr 05 07:58:34 -0500 (2 seconds ago)
Duration:            59 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ✔ keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  11s
 ├-✔ svc-memcached                      keystone-install-3754811077  10s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   11s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  23s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  24s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  5s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  6s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  7s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   8s
 └-✔ wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  4s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---✔ task-keystone-db-sync(0)       keystone-install-1630370687  21s
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     1m    1m
keystone-install     Succeeded   1m    59s
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-6d7bd74f96-6s8qk                 1/1     Running     0          2m36s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          7m21s
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          40s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          2m36s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          33s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          44s
pod/keystone-install-1608858220                   0/2     Completed   0          2m36s
pod/keystone-install-1630370687                   0/2     Completed   0          2m
pod/keystone-install-1659018021                   0/2     Completed   0          2m18s
pod/keystone-install-2012559777                   0/1     Completed   0          2m24s
pod/keystone-install-2401658373                   0/1     Completed   0          2m24s
pod/keystone-install-2716093231                   0/2     Completed   0          2m36s
pod/keystone-install-2958852122                   0/2     Completed   0          2m36s
pod/keystone-install-3754811077                   0/2     Completed   0          2m36s
pod/keystone-install-3843173896                   0/1     Completed   0          2m6s
pod/keystone-install-421213240                    0/2     Completed   0          2m36s
pod/keystone-install-4222398536                   0/1     Completed   0          2m6s
pod/keystone-install-494016497                    0/2     Completed   0          2m15s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          7m16s
pod/mariadb-ingress-5cfc68b86f-g7sgm              1/1     Running     0          5m3s
pod/mariadb-ingress-error-pages-67c88c67d-s5r7s   1/1     Running     0          5m4s
pod/mariadb-server-0                              1/1     Running     0          5m3s
pod/memcached-memcached-75447bffcf-bgxft          1/1     Running     0          5m2s
pod/rabbitmq-cluster-wait-ztvqh                   0/1     Completed   0          5m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          5m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.103.84.0      <none>        80/TCP,443/TCP                 2m36s
service/keystone-api                  ClusterIP   10.98.137.145    <none>        5000/TCP                       2m36s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       7m19s
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       7m14s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        19m
service/mariadb                       ClusterIP   10.98.4.164      <none>        3306/TCP                       5m4s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              5m4s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         5m4s
service/mariadb-server                ClusterIP   10.105.137.86    <none>        3306/TCP                       5m4s
service/memcached                     ClusterIP   10.100.216.105   <none>        11211/TCP                      5m2s
service/rabbitmq                      ClusterIP   10.98.191.40     <none>        5672/TCP,25672/TCP,15672/TCP   5m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   5m
service/rabbitmq-mgr-7b1733           ClusterIP   10.97.20.244     <none>        80/TCP,443/TCP                 5m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           2m36s
deployment.apps/keystone-armada-operator      1/1     1            1           7m22s
deployment.apps/keystone-oslc-operator        1/1     1            1           7m16s
deployment.apps/mariadb-ingress               1/1     1            1           5m4s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           5m4s
deployment.apps/memcached-memcached           1/1     1            1           5m2s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       2m36s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       7m21s
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       7m16s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       5m4s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       5m4s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       5m2s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     5m3s
statefulset.apps/rabbitmq-rabbitmq   1/1     5m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           25s        5m

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          2m36s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          2m36s
```
 
```bash
prompt$ argo logs keystone-install -w
svc-mariadb:    Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Invalid format:
svc-mariadb:    Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Resolving Service mariadb in namespace default
svc-mariadb:    Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Dependency Service mariadb in namespace default is resolved.
svc-mariadb:    done
svc-memcached:  Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Invalid format:
svc-memcached:  Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Resolving Service memcached in namespace default
svc-memcached:  Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Dependency Service memcached in namespace default is resolved.
svc-memcached:  done
svc-rabbitmq:   Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Invalid format:
svc-rabbitmq:   Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Resolving Service rabbitmq in namespace default
svc-rabbitmq:   Entrypoint INFO: 2019/04/05 12:57:42 logger.go:32: Dependency Service rabbitmq in namespace default is resolved.
svc-rabbitmq:   done
task-keystone-credential-setup(0):      2019-04-05 12:57:47.445 - INFO - Executing 'keystone-manage credential_setup --keystone-user=keystone --keystone-group=keystone' command.
task-keystone-fernet-setup(0):  2019-04-05 12:57:48.681 - INFO - Executing 'keystone-manage fernet_setup --keystone-user=keystone --keystone-group=keystone' command.
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Creating a docker executor"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: mariadb\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"mariadb\"\n"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:51Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Creating a docker executor"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: rabbitmq\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"rabbitmq\"\n"
svc-rabbitmq-just-in-time:      kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:52Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       time="2019-04-05T12:57:52Z" level=info msg=default/Service./mariadb
svc-mariadb-just-in-time:       time="2019-04-05T12:57:52Z" level=info msg="No output parameters"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:53Z" level=info msg=default/Service./rabbitmq
svc-rabbitmq-just-in-time:      time="2019-04-05T12:57:53Z" level=info msg="No output parameters"
task-keystone-credential-setup(0):      2019-04-05 12:57:57.902 8 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/credential-keys/
task-keystone-credential-setup(0):      2019-04-05 12:57:57.902 8 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/credential-keys/
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.907 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/credential-keys/0']
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/credential-keys/0']
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/credential-keys/0.tmp
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.908 8 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.909 8 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.909 8 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:57.910 8 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/credential-keys/0
task-keystone-credential-setup(0):      2019-04-05 12:57:58.011 - INFO - Updating data for 'keystone-credential-keys' secret.
task-keystone-credential-setup(0):      2019-04-05 12:57:58.057 - INFO - 2 fernet keys have been placed to secret 'keystone-credential-keys'
task-keystone-credential-setup(0):      2019-04-05 12:57:58.057 - INFO - Credential keys generation has been completed
task-keystone-db-init(0):       2019-04-05 12:57:58,718 - OpenStack-Helm DB Init - INFO - Got DB root connection
task-keystone-db-init(0):       2019-04-05 12:57:58,718 - OpenStack-Helm DB Init - INFO - Using /etc/keystone/keystone.conf as db config source
task-keystone-db-init(0):       2019-04-05 12:57:58,720 - OpenStack-Helm DB Init - INFO - Trying to load db config from database:connection
task-keystone-db-init(0):       2019-04-05 12:57:58,720 - OpenStack-Helm DB Init - INFO - Got config from /etc/keystone/keystone.conf
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.761 9 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/fernet-keys/
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.761 9 WARNING keystone.common.fernet_utils [-] key_repository is world readable: /etc/keystone/fernet-keys/
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.762 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.762 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/fernet-keys/0']
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.763 9 INFO keystone.common.fernet_utils [-] Starting key rotation with 1 key files: ['/etc/keystone/fernet-keys/0']
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Created a new temporary key: /etc/keystone/fernet-keys/0.tmp
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.764 9 INFO keystone.common.fernet_utils [-] Current primary key is: 0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.765 9 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.765 9 INFO keystone.common.fernet_utils [-] Next primary key will be: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Promoted key 0 to be the primary: 1
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.766 9 INFO keystone.common.fernet_utils [-] Become a valid new key: /etc/keystone/fernet-keys/0
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.864 - INFO - Updating data for 'keystone-fernet-keys' secret.
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.906 - INFO - 2 fernet keys have been placed to secret 'keystone-fernet-keys'
task-keystone-fernet-setup(0):  2019-04-05 12:57:58.906 - INFO - Fernet keys generation has been completed
task-keystone-db-init(0):       2019-04-05 12:57:58,946 - OpenStack-Helm DB Init - INFO - Tested connection to DB @ mariadb.default.svc.cluster.local:3306 as root
task-keystone-db-init(0):       2019-04-05 12:57:58,948 - OpenStack-Helm DB Init - INFO - Got user db config
task-keystone-db-init(0):       2019-04-05 12:57:58,981 - OpenStack-Helm DB Init - INFO - Created database keystone
task-keystone-db-init(0):       2019-04-05 12:57:58,991 - OpenStack-Helm DB Init - INFO - Created user keystone for keystone
task-keystone-db-init(0):       2019-04-05 12:57:59,062 - OpenStack-Helm DB Init - INFO - Tested connection to DB @ mariadb.default.svc.cluster.local:3306/keystone as keystone
task-keystone-db-init(0):       2019-04-05 12:57:59,063 - OpenStack-Helm DB Init - INFO - Finished DB Management
task-keystone-rabbit-init(0):   Managing: User: keystone
task-keystone-rabbit-init(0):   user declared
task-keystone-rabbit-init(0):   Managing: vHost: keystone
task-keystone-rabbit-init(0):   vhost declared
task-keystone-rabbit-init(0):   Managing: Permissions: keystone on keystone
task-keystone-rabbit-init(0):   permission declared
task-keystone-rabbit-init(0):   Applying additional configuration
task-keystone-rabbit-init(0):   Imported definitions for rabbitmq.default.svc.cluster.local from "/tmp/rmq_definitions.json"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Creating a docker executor"
svc-rabbitmq-just-in-time:      kubectl get -f /tmp/manifest.yaml -o json
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: rabbitmq\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"rabbitmq\"\n"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:08Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Creating a docker executor"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Executor (version: v2.3.0+9702c44.dirty, build_date: 2019-03-19T15:43:51Z) initialized with template:\narchiveLocation: {}\ninputs:\n  parameters:\n  - name: service\n    value: mariadb\nmetadata: {}\nname: svc-just-in-time\noutputs: {}\nresource:\n  action: get\n  manifest: |-\n    apiVersion: v1\n    kind: Service\n    metadata:\n      name: \"mariadb\"\n"
svc-mariadb-just-in-time:       kubectl get -f /tmp/manifest.yaml -o json
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="Loading manifest to /tmp/manifest.yaml"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:08Z" level=info msg="kubectl get -f /tmp/manifest.yaml -o json"
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:09Z" level=info msg=default/Service./rabbitmq
svc-rabbitmq-just-in-time:      time="2019-04-05T12:58:09Z" level=info msg="No output parameters"
svc-mariadb-just-in-time:       time="2019-04-05T12:58:09Z" level=info msg=default/Service./mariadb
svc-mariadb-just-in-time:       time="2019-04-05T12:58:09Z" level=info msg="No output parameters"
task-keystone-db-sync(0):       + keystone-manage --config-file=/etc/keystone/keystone.conf db_sync
task-keystone-db-sync(0):       2019-04-05 12:58:22.006 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`inherited` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       2019-04-05 12:58:22.197 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`enabled` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       2019-04-05 12:58:23.309 8 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/sqlalchemy/dialects/mysql/base.py:3016: SAWarning: Unknown schema content: u'  CONSTRAINT `CONSTRAINT_1` CHECK (`impersonation` in (0,1))'
task-keystone-db-sync(0):         util.warn("Unknown schema content: %r" % line)
task-keystone-db-sync(0):       + keystone-manage --config-file=/etc/keystone/keystone.conf bootstrap --bootstrap-username admin --bootstrap-password password --bootstrap-project-name admin --bootstrap-admin-url http://keystone.default.svc.cluster.local:80/v3 --bootstrap-public-url http://keystone.default.svc.cluster.local:80/v3 --bootstrap-internal-url http://keystone-api.default.svc.cluster.local:5000/v3 --bootstrap-region-id RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:31.547 11 WARNING py.warnings [-] /var/lib/openstack/local/lib/python2.7/site-packages/pycadf/identifier.py:60: UserWarning: Invalid uuid. To ensure interoperability, identifiers should be a valid uuid.
task-keystone-db-sync(0):         warnings.warn('Invalid uuid. To ensure interoperability, identifiers '
task-keystone-db-sync(0):       2019-04-05 12:58:31.806 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created domain default
task-keystone-db-sync(0):       2019-04-05 12:58:31.806 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created domain default
task-keystone-db-sync(0):       2019-04-05 12:58:31.851 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created project admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.851 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created project admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.852 11 WARNING keystone.identity.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Unable to locate domain config directory: /etc/keystonedomains
task-keystone-db-sync(0):       2019-04-05 12:58:31.852 11 WARNING keystone.identity.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Unable to locate domain config directory: /etc/keystonedomains
task-keystone-db-sync(0):       2019-04-05 12:58:31.937 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created user admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.937 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created user admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.955 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created role admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.955 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created role admin
task-keystone-db-sync(0):       2019-04-05 12:58:31.970 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Granted admin on admin to user admin.
task-keystone-db-sync(0):       2019-04-05 12:58:31.970 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Granted admin on admin to user admin.
task-keystone-db-sync(0):       2019-04-05 12:58:32.004 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created region RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:32.004 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created region RegionOne
task-keystone-db-sync(0):       2019-04-05 12:58:32.049 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created admin endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.049 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created admin endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.064 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created internal endpoint http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.064 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created internal endpoint http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.083 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created public endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.083 11 INFO keystone.cmd.cli [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Created public endpoint http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32.089 11 INFO keystone.assignment.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Creating the default role 9fe2ff9ee4384b1894a90878d3e92bab because it does not exist.
task-keystone-db-sync(0):       2019-04-05 12:58:32.089 11 INFO keystone.assignment.core [req-24b607b1-6586-46fe-89b4-b21e317d111f - - - - -] Creating the default role 9fe2ff9ee4384b1894a90878d3e92bab because it does not exist.
task-keystone-db-sync(0):       + exec python /tmp/endpoint-update.py
task-keystone-db-sync(0):       2019-04-05 12:58:32,433 - OpenStack-Helm Keystone Endpoint management - INFO - Using /etc/keystone/keystone.conf as db config source
task-keystone-db-sync(0):       2019-04-05 12:58:32,439 - OpenStack-Helm Keystone Endpoint management - INFO - Trying to load db config from database:connection
task-keystone-db-sync(0):       2019-04-05 12:58:32,440 - OpenStack-Helm Keystone Endpoint management - INFO - Got config from /etc/keystone/keystone.conf
task-keystone-db-sync(0):       2019-04-05 12:58:32,521 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (admin): http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (public): http://keystone.default.svc.cluster.local:80/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - endpoint (internal): http://keystone-api.default.svc.cluster.local:5000/v3
task-keystone-db-sync(0):       2019-04-05 12:58:32,522 - OpenStack-Helm Keystone Endpoint management - INFO - Finished Endpoint Management
```
 
```bash
prompt$ argo logs keystone-bootstrap -w
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Invalid format:
svc-keystone:   Entrypoint INFO: 2019/04/05 12:57:40 logger.go:32: Resolving Service keystone-api in namespace default
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:40 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 2.000 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:42 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 2.500 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:45 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 3.125 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:48 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 3.906 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:52 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 4.883 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:57:56 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 6.104 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:03 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 7.629 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:10 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 9.537 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:20 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 11.921 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:32 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 14.901 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:58:47 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 18.626 seconds
svc-keystone:   Entrypoint WARNING: 2019/04/05 12:59:05 logger.go:39: Resolving dependency Service keystone-api in namespace default failed: Service keystone-api has no endpoints - Trying again in 20.000 seconds
svc-keystone:   Entrypoint INFO: 2019/04/05 12:59:25 logger.go:32: Dependency Service keystone-api in namespace default is resolved.
svc-keystone:   done
task-bootstrap: + openstack role create --or-show member
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: | Field     | Value                            |
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: | domain_id | None                             |
task-bootstrap: | id        | ccbaa551824c4a519f2ec42501c493e4 |
task-bootstrap: | name      | member                           |
task-bootstrap: +-----------+----------------------------------+
task-bootstrap: + openstack role add --user=admin --user-domain=default --project-domain=default --project=admin member
```
 
```bash
prompt$
```
