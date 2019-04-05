# ./oslc/deleteflow.md

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
prompt$ vi traces/greenfield.txt
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Running
Created:             Fri Apr 05 08:14:27 -0500 (53 seconds ago)
Started:             Fri Apr 05 08:14:27 -0500 (53 seconds ago)
Duration:            53 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ● keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 └---◷ keystone-enable-trafficrollout  keystone-install-flow-3629395555  6s        ContainerCreating
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-665ddfbbbd-h2kq8                 1/1     Running     0          13m
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          35m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          11m
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          13m
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          11m
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          11m
pod/keystone-install-1608858220                   0/2     Completed   0          13m
pod/keystone-install-1630370687                   0/2     Completed   0          12m
pod/keystone-install-1659018021                   0/2     Completed   0          12m
pod/keystone-install-2012559777                   0/1     Completed   0          12m
pod/keystone-install-2401658373                   0/1     Completed   0          12m
pod/keystone-install-2716093231                   0/2     Completed   0          13m
pod/keystone-install-2958852122                   0/2     Completed   0          13m
pod/keystone-install-3754811077                   0/2     Completed   0          13m
pod/keystone-install-3843173896                   0/1     Completed   0          12m
pod/keystone-install-421213240                    0/2     Completed   0          13m
pod/keystone-install-4222398536                   0/1     Completed   0          12m
pod/keystone-install-494016497                    0/2     Completed   0          12m
pod/keystone-install-flow-1815575608              0/1     Completed   0          12m
pod/keystone-install-flow-2086467396              0/1     Completed   0          13m
pod/keystone-install-flow-2390570688              0/1     Completed   0          12m
pod/keystone-install-flow-2673199163              0/1     Completed   0          13m
pod/keystone-install-flow-2968515285              0/1     Completed   0          13m
pod/keystone-install-flow-3629395555              0/1     Completed   0          12m
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          35m
pod/keystone-test-1478702594                      0/2     Completed   0          12m
pod/keystone-test-2192614885                      0/2     Completed   0          12m
pod/keystone-test-2420766946                      0/2     Completed   0          12m
pod/keystone-test-3055163213                      0/2     Completed   0          12m
pod/keystone-test-test                            0/1     Completed   0          12m
pod/keystone-trafficrollout-1708264911            0/2     Completed   0          12m
pod/keystone-trafficrollout-3671107028            0/2     Completed   0          12m
pod/keystone-trafficrollout-4064757452            0/2     Completed   0          12m
pod/keystone-trafficrollout-597422575             0/2     Completed   0          12m
pod/mariadb-ingress-5cfc68b86f-g5qb5              1/1     Running     0          14m
pod/mariadb-ingress-error-pages-67c88c67d-tsdqd   1/1     Running     0          14m
pod/mariadb-server-0                              1/1     Running     0          14m
pod/memcached-memcached-75447bffcf-6mf27          1/1     Running     0          14m
pod/rabbitmq-cluster-wait-p59mk                   0/1     Completed   0          14m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          14m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.99.72.181     <none>        80/TCP,443/TCP                 13m
service/keystone-api                  ClusterIP   10.100.117.228   <none>        5000/TCP                       13m
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       35m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       35m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        47m
service/mariadb                       ClusterIP   10.109.119.23    <none>        3306/TCP                       14m
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              14m
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         14m
service/mariadb-server                ClusterIP   10.104.200.51    <none>        3306/TCP                       14m
service/memcached                     ClusterIP   10.99.74.136     <none>        11211/TCP                      14m
service/rabbitmq                      ClusterIP   10.100.154.236   <none>        5672/TCP,25672/TCP,15672/TCP   14m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   14m
service/rabbitmq-mgr-7b1733           ClusterIP   10.111.241.140   <none>        80/TCP,443/TCP                 14m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           13m
deployment.apps/keystone-armada-operator      1/1     1            1           35m
deployment.apps/keystone-oslc-operator        1/1     1            1           35m
deployment.apps/mariadb-ingress               1/1     1            1           14m
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           14m
deployment.apps/memcached-memcached           1/1     1            1           14m

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-665ddfbbbd                 1         1         1       13m
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       35m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       35m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       14m
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       14m
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       14m

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     14m
statefulset.apps/rabbitmq-rabbitmq   1/1     14m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           32s        14m

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          13m
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          13m
```
 
```bash
prompt$ argo list
NAME                      STATUS      AGE   DURATION
keystone-bootstrap        Succeeded   13m   2m
keystone-install          Succeeded   13m   1m
keystone-trafficrollout   Succeeded   12m   12s
keystone-install-flow     Succeeded   13m   1m
keystone-test             Succeeded   12m   17s
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:14:27 -0500 (14 minutes ago)
Started:             Fri Apr 05 08:14:27 -0500 (14 minutes ago)
Finished:            Fri Apr 05 08:15:33 -0500 (12 minutes ago)
Duration:            1 minute 6 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ✔ keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 ├---✔ keystone-enable-trafficrollout  keystone-install-flow-3629395555  7s
 └---✔ keystone-trafficrollout-ready   keystone-install-flow-1815575608  8s
```
 
```bash
prompt$ vi traces/installflow.txt
```
 
```bash
prompt$ kubectl apply -f examples/uninstall/
oslc.openstacklcm.airshipit.org/keystone-uninstall-flow created
```
 
```bash
prompt$ argo list
NAME                      STATUS      AGE   DURATION
keystone-delete           Succeeded   1m    9s
keystone-uninstall-flow   Succeeded   1m    22s
keystone-trafficdrain     Succeeded   1m    9s
keystone-bootstrap        Succeeded   16m   2m
keystone-install          Succeeded   16m   1m
keystone-trafficrollout   Succeeded   15m   12s
keystone-install-flow     Succeeded   16m   1m
keystone-test             Succeeded   15m   17s
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:14:27 -0500 (16 minutes ago)
Started:             Fri Apr 05 08:14:27 -0500 (16 minutes ago)
Finished:            Fri Apr 05 08:15:33 -0500 (15 minutes ago)
Duration:            1 minute 6 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ✔ keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 ├---✔ keystone-enable-trafficrollout  keystone-install-flow-3629395555  7s
 └---✔ keystone-trafficrollout-ready   keystone-install-flow-1815575608  8s
```
 
```bash
prompt$ argo get keystone-uninstall-flow
Name:                keystone-uninstall-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:29:31 -0500 (1 minute ago)
Started:             Fri Apr 05 08:29:31 -0500 (1 minute ago)
Finished:            Fri Apr 05 08:29:53 -0500 (1 minute ago)
Duration:            22 seconds

STEP                                 PODNAME                             DURATION  MESSAGE
 ✔ keystone-uninstall-flow
 ├---✔ keystone-enable-trafficdrain  keystone-uninstall-flow-3958365555  3s
 ├---✔ keystone-trafficdrain-ready   keystone-uninstall-flow-3432630062  6s
 ├---✔ keystone-enable-delete        keystone-uninstall-flow-3560870671  4s
 └---✔ keystone-delete-ready         keystone-uninstall-flow-2408050222  6s
```
 
```bash
prompt$ kubectl delete -f examples/uninstall/
oslc.openstacklcm.airshipit.org "keystone-uninstall-flow" deleted
```
 
```bash
prompt$ kubectl delete -f examples/greenfield/
oslc.openstacklcm.airshipit.org "keystone-install-flow" deleted
```
 
```bash
prompt$
```
# ./oslc/deleteflow.md
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
prompt$ vi traces/greenfield.txt
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Running
Created:             Fri Apr 05 08:14:27 -0500 (53 seconds ago)
Started:             Fri Apr 05 08:14:27 -0500 (53 seconds ago)
Duration:            53 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ● keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 └---◷ keystone-enable-trafficrollout  keystone-install-flow-3629395555  6s        ContainerCreating
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-665ddfbbbd-h2kq8                 1/1     Running     0          13m
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          35m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          11m
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          13m
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          11m
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          11m
pod/keystone-install-1608858220                   0/2     Completed   0          13m
pod/keystone-install-1630370687                   0/2     Completed   0          12m
pod/keystone-install-1659018021                   0/2     Completed   0          12m
pod/keystone-install-2012559777                   0/1     Completed   0          12m
pod/keystone-install-2401658373                   0/1     Completed   0          12m
pod/keystone-install-2716093231                   0/2     Completed   0          13m
pod/keystone-install-2958852122                   0/2     Completed   0          13m
pod/keystone-install-3754811077                   0/2     Completed   0          13m
pod/keystone-install-3843173896                   0/1     Completed   0          12m
pod/keystone-install-421213240                    0/2     Completed   0          13m
pod/keystone-install-4222398536                   0/1     Completed   0          12m
pod/keystone-install-494016497                    0/2     Completed   0          12m
pod/keystone-install-flow-1815575608              0/1     Completed   0          12m
pod/keystone-install-flow-2086467396              0/1     Completed   0          13m
pod/keystone-install-flow-2390570688              0/1     Completed   0          12m
pod/keystone-install-flow-2673199163              0/1     Completed   0          13m
pod/keystone-install-flow-2968515285              0/1     Completed   0          13m
pod/keystone-install-flow-3629395555              0/1     Completed   0          12m
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          35m
pod/keystone-test-1478702594                      0/2     Completed   0          12m
pod/keystone-test-2192614885                      0/2     Completed   0          12m
pod/keystone-test-2420766946                      0/2     Completed   0          12m
pod/keystone-test-3055163213                      0/2     Completed   0          12m
pod/keystone-test-test                            0/1     Completed   0          12m
pod/keystone-trafficrollout-1708264911            0/2     Completed   0          12m
pod/keystone-trafficrollout-3671107028            0/2     Completed   0          12m
pod/keystone-trafficrollout-4064757452            0/2     Completed   0          12m
pod/keystone-trafficrollout-597422575             0/2     Completed   0          12m
pod/mariadb-ingress-5cfc68b86f-g5qb5              1/1     Running     0          14m
pod/mariadb-ingress-error-pages-67c88c67d-tsdqd   1/1     Running     0          14m
pod/mariadb-server-0                              1/1     Running     0          14m
pod/memcached-memcached-75447bffcf-6mf27          1/1     Running     0          14m
pod/rabbitmq-cluster-wait-p59mk                   0/1     Completed   0          14m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          14m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.99.72.181     <none>        80/TCP,443/TCP                 13m
service/keystone-api                  ClusterIP   10.100.117.228   <none>        5000/TCP                       13m
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       35m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       35m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        47m
service/mariadb                       ClusterIP   10.109.119.23    <none>        3306/TCP                       14m
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              14m
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         14m
service/mariadb-server                ClusterIP   10.104.200.51    <none>        3306/TCP                       14m
service/memcached                     ClusterIP   10.99.74.136     <none>        11211/TCP                      14m
service/rabbitmq                      ClusterIP   10.100.154.236   <none>        5672/TCP,25672/TCP,15672/TCP   14m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   14m
service/rabbitmq-mgr-7b1733           ClusterIP   10.111.241.140   <none>        80/TCP,443/TCP                 14m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           13m
deployment.apps/keystone-armada-operator      1/1     1            1           35m
deployment.apps/keystone-oslc-operator        1/1     1            1           35m
deployment.apps/mariadb-ingress               1/1     1            1           14m
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           14m
deployment.apps/memcached-memcached           1/1     1            1           14m

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-665ddfbbbd                 1         1         1       13m
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       35m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       35m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       14m
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       14m
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       14m

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     14m
statefulset.apps/rabbitmq-rabbitmq   1/1     14m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           32s        14m

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          13m
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          13m
```
 
```bash
prompt$ argo list
NAME                      STATUS      AGE   DURATION
keystone-bootstrap        Succeeded   13m   2m
keystone-install          Succeeded   13m   1m
keystone-trafficrollout   Succeeded   12m   12s
keystone-install-flow     Succeeded   13m   1m
keystone-test             Succeeded   12m   17s
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:14:27 -0500 (14 minutes ago)
Started:             Fri Apr 05 08:14:27 -0500 (14 minutes ago)
Finished:            Fri Apr 05 08:15:33 -0500 (12 minutes ago)
Duration:            1 minute 6 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ✔ keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 ├---✔ keystone-enable-trafficrollout  keystone-install-flow-3629395555  7s
 └---✔ keystone-trafficrollout-ready   keystone-install-flow-1815575608  8s
```
 
```bash
prompt$ vi traces/installflow.txt
```
 
```bash
prompt$ kubectl apply -f examples/uninstall/
oslc.openstacklcm.airshipit.org/keystone-uninstall-flow created
```
 
```bash
prompt$ argo list
NAME                      STATUS      AGE   DURATION
keystone-delete           Succeeded   1m    9s
keystone-uninstall-flow   Succeeded   1m    22s
keystone-trafficdrain     Succeeded   1m    9s
keystone-bootstrap        Succeeded   16m   2m
keystone-install          Succeeded   16m   1m
keystone-trafficrollout   Succeeded   15m   12s
keystone-install-flow     Succeeded   16m   1m
keystone-test             Succeeded   15m   17s
```
 
```bash
prompt$ argo get keystone-install-flow
Name:                keystone-install-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:14:27 -0500 (16 minutes ago)
Started:             Fri Apr 05 08:14:27 -0500 (16 minutes ago)
Finished:            Fri Apr 05 08:15:33 -0500 (15 minutes ago)
Duration:            1 minute 6 seconds

STEP                                   PODNAME                           DURATION  MESSAGE
 ✔ keystone-install-flow
 ├---✔ keystone-enable-install         keystone-install-flow-2086467396  3s
 ├---✔ keystone-install-ready          keystone-install-flow-2968515285  24s
 ├---✔ keystone-enable-test            keystone-install-flow-2673199163  7s
 ├---✔ keystone-test-ready             keystone-install-flow-2390570688  7s
 ├---✔ keystone-enable-trafficrollout  keystone-install-flow-3629395555  7s
 └---✔ keystone-trafficrollout-ready   keystone-install-flow-1815575608  8s
```
 
```bash
prompt$ argo get keystone-uninstall-flow
Name:                keystone-uninstall-flow
Namespace:           default
ServiceAccount:      openstacklcm-argo-sa
Status:              Succeeded
Created:             Fri Apr 05 08:29:31 -0500 (1 minute ago)
Started:             Fri Apr 05 08:29:31 -0500 (1 minute ago)
Finished:            Fri Apr 05 08:29:53 -0500 (1 minute ago)
Duration:            22 seconds

STEP                                 PODNAME                             DURATION  MESSAGE
 ✔ keystone-uninstall-flow
 ├---✔ keystone-enable-trafficdrain  keystone-uninstall-flow-3958365555  3s
 ├---✔ keystone-trafficdrain-ready   keystone-uninstall-flow-3432630062  6s
 ├---✔ keystone-enable-delete        keystone-uninstall-flow-3560870671  4s
 └---✔ keystone-delete-ready         keystone-uninstall-flow-2408050222  6s
```
 
```bash
prompt$ kubectl delete -f examples/uninstall/
oslc.openstacklcm.airshipit.org "keystone-uninstall-flow" deleted
```
 
```bash
prompt$ kubectl delete -f examples/greenfield/
oslc.openstacklcm.airshipit.org "keystone-install-flow" deleted
```
 
```bash
prompt$
```
