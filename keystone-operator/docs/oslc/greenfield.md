# ./oslc/greenfield.md

```bash
prompt$ clear
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          20m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          20m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   20m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   20m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    32m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           20m
deployment.apps/keystone-oslc-operator     1/1     1            1           20m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       20m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       20m
```
 
```bash
prompt$ kubectl apply -f examples/
argo/             brownfield/       git/              greenfield/       keystone/         lifecycle-events/ rollback/         sequenced/        uninstall/
```
 
```bash
prompt$ kubectl apply -f examples/keystone/infra.yaml
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS     RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running    0          21m
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running    0          21m
pod/mariadb-ingress-5cfc68b86f-g5qb5              0/1     Running    0          28s
pod/mariadb-ingress-error-pages-67c88c67d-tsdqd   1/1     Running    0          28s
pod/mariadb-server-0                              0/1     Running    0          28s
pod/memcached-memcached-75447bffcf-6mf27          1/1     Running    0          27s
pod/rabbitmq-cluster-wait-p59mk                   0/1     Init:0/1   0          26s
pod/rabbitmq-rabbitmq-0                           0/1     Running    0          26s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       21m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       20m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        33m
service/mariadb                       ClusterIP   10.109.119.23    <none>        3306/TCP                       29s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              29s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         29s
service/mariadb-server                ClusterIP   10.104.200.51    <none>        3306/TCP                       29s
service/memcached                     ClusterIP   10.99.74.136     <none>        11211/TCP                      27s
service/rabbitmq                      ClusterIP   10.100.154.236   <none>        5672/TCP,25672/TCP,15672/TCP   26s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   26s
service/rabbitmq-mgr-7b1733           ClusterIP   10.111.241.140   <none>        80/TCP,443/TCP                 26s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator      1/1     1            1           21m
deployment.apps/keystone-oslc-operator        1/1     1            1           21m
deployment.apps/mariadb-ingress               0/1     1            0           29s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           29s
deployment.apps/memcached-memcached           1/1     1            1           27s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       21m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       21m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         0       28s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       29s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       27s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      0/1     28s
statefulset.apps/rabbitmq-rabbitmq   0/1     26s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   0/1           26s        26s
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
prompt$ argo list
NAME   STATUS   AGE   DURATION
```
 
```bash
prompt$ kubectl apply -f examples/
argo/             brownfield/       git/              greenfield/       keystone/         lifecycle-events/ rollback/         sequenced/        uninstall/
```
 
```bash
prompt$ kubectl apply -f examples/greenfield/
oslc.openstacklcm.airshipit.org/keystone-install-flow created
```
 
```bash
prompt$ argo list
NAME                    STATUS    AGE   DURATION
keystone-install-flow   Running   4s    4s
```
 
```bash
prompt$
```
# ./oslc/greenfield.md
```bash
prompt$ clear
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          20m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          20m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   20m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   20m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    32m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           20m
deployment.apps/keystone-oslc-operator     1/1     1            1           20m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       20m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       20m
```
 
```bash
prompt$ kubectl apply -f examples/
argo/             brownfield/       git/              greenfield/       keystone/         lifecycle-events/ rollback/         sequenced/        uninstall/
```
 
```bash
prompt$ kubectl apply -f examples/keystone/infra.yaml
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS     RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running    0          21m
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running    0          21m
pod/mariadb-ingress-5cfc68b86f-g5qb5              0/1     Running    0          28s
pod/mariadb-ingress-error-pages-67c88c67d-tsdqd   1/1     Running    0          28s
pod/mariadb-server-0                              0/1     Running    0          28s
pod/memcached-memcached-75447bffcf-6mf27          1/1     Running    0          27s
pod/rabbitmq-cluster-wait-p59mk                   0/1     Init:0/1   0          26s
pod/rabbitmq-rabbitmq-0                           0/1     Running    0          26s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       21m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       20m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        33m
service/mariadb                       ClusterIP   10.109.119.23    <none>        3306/TCP                       29s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              29s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         29s
service/mariadb-server                ClusterIP   10.104.200.51    <none>        3306/TCP                       29s
service/memcached                     ClusterIP   10.99.74.136     <none>        11211/TCP                      27s
service/rabbitmq                      ClusterIP   10.100.154.236   <none>        5672/TCP,25672/TCP,15672/TCP   26s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   26s
service/rabbitmq-mgr-7b1733           ClusterIP   10.111.241.140   <none>        80/TCP,443/TCP                 26s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator      1/1     1            1           21m
deployment.apps/keystone-oslc-operator        1/1     1            1           21m
deployment.apps/mariadb-ingress               0/1     1            0           29s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           29s
deployment.apps/memcached-memcached           1/1     1            1           27s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       21m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       21m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         0       28s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       29s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       27s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      0/1     28s
statefulset.apps/rabbitmq-rabbitmq   0/1     26s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   0/1           26s        26s
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
prompt$ argo list
NAME   STATUS   AGE   DURATION
```
 
```bash
prompt$ kubectl apply -f examples/
argo/             brownfield/       git/              greenfield/       keystone/         lifecycle-events/ rollback/         sequenced/        uninstall/
```
 
```bash
prompt$ kubectl apply -f examples/greenfield/
oslc.openstacklcm.airshipit.org/keystone-install-flow created
```
 
```bash
prompt$ argo list
NAME                    STATUS    AGE   DURATION
keystone-install-flow   Running   4s    4s
```
 
```bash
prompt$
```
