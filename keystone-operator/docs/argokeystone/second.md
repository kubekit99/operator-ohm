# ./argokeystone/second.md

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
prompt$
```
# ./argokeystone/second.md
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
