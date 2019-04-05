# ./argomanifest/argodelete.md

```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          14m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          14m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   14m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   14m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    27m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           14m
deployment.apps/keystone-oslc-operator     1/1     1            1           14m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       14m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       14m
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
prompt$ kubectl apply -f examples/argo/
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
workflow.argoproj.io/armada-manifest created
```
 
```bash
prompt$ argo list
NAME                 STATUS    AGE   DURATION
keystone-install     Running   6s    6s
keystone-bootstrap   Running   6s    6s
armada-manifest      Running   10s   10s
```
 
```bash
prompt$ argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Apr 05 08:07:48 -0500 (22 seconds ago)
Started:             Fri Apr 05 08:07:48 -0500 (22 seconds ago)
Duration:            22 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   3s
 | ├---✔ mariadb-ready        armada-manifest-361794910   7s
 | └---● enable-memcached     armada-manifest-4290341812  4s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   9s
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     1m    1m
keystone-install     Succeeded   1m    1m
armada-manifest      Succeeded   1m    37s
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
armada-manifest      2m5s
keystone-bootstrap   2m1s
keystone-install     2m1s
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
prompt$ vi traces/argo.txt
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     2m    2m
keystone-install     Succeeded   2m    1m
armada-manifest      Succeeded   2m    37s
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Succeeded
Created:             Fri Apr 05 08:07:52 -0500 (3 minutes ago)
Started:             Fri Apr 05 08:07:52 -0500 (3 minutes ago)
Finished:            Fri Apr 05 08:09:27 -0500 (1 minute ago)
Duration:            1 minute 35 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ✔ keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  57s
 ├-✔ svc-memcached                      keystone-install-3754811077  36s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   57s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  18s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  18s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  4s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  4s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  4s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   4s
 └-✔ wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  3s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---✔ task-keystone-db-sync(0)       keystone-install-1630370687  20s
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/armada-manifest-13034731                      0/1     Completed   0          2m48s
pod/armada-manifest-1873087745                    0/1     Completed   0          2m43s
pod/armada-manifest-2789817682                    0/1     Completed   0          3m11s
pod/armada-manifest-3014489540                    0/1     Completed   0          2m39s
pod/armada-manifest-361794910                     0/1     Completed   0          3m1s
pod/armada-manifest-4290341812                    0/1     Completed   0          2m53s
pod/armada-manifest-587324697                     0/1     Completed   0          3m7s
pod/armada-manifest-842022813                     0/1     Completed   0          3m11s
pod/keystone-api-6d7bd74f96-qjq8c                 1/1     Running     0          3m7s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          18m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          30s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          3m7s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          24s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          37s
pod/keystone-install-1608858220                   0/2     Completed   0          3m7s
pod/keystone-install-1630370687                   0/2     Completed   0          112s
pod/keystone-install-1659018021                   0/2     Completed   0          2m5s
pod/keystone-install-2012559777                   0/1     Completed   0          2m9s
pod/keystone-install-2401658373                   0/1     Completed   0          2m10s
pod/keystone-install-2716093231                   0/2     Completed   0          3m7s
pod/keystone-install-2958852122                   0/2     Completed   0          3m7s
pod/keystone-install-3754811077                   0/2     Completed   0          3m7s
pod/keystone-install-3843173896                   0/1     Completed   0          117s
pod/keystone-install-421213240                    0/2     Completed   0          3m7s
pod/keystone-install-4222398536                   0/1     Completed   0          117s
pod/keystone-install-494016497                    0/2     Completed   0          2m4s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          18m
pod/mariadb-ingress-5cfc68b86f-9qpss              1/1     Running     0          3m5s
pod/mariadb-ingress-error-pages-67c88c67d-v92gn   1/1     Running     0          3m5s
pod/mariadb-server-0                              1/1     Running     0          3m5s
pod/memcached-memcached-75447bffcf-2p87s          1/1     Running     0          2m48s
pod/rabbitmq-cluster-wait-vn4qs                   0/1     Completed   0          2m40s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          2m40s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.109.208.195   <none>        80/TCP,443/TCP                 3m7s
service/keystone-api                  ClusterIP   10.111.11.14     <none>        5000/TCP                       3m7s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       18m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       18m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        30m
service/mariadb                       ClusterIP   10.109.71.179    <none>        3306/TCP                       3m6s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              3m6s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         3m6s
service/mariadb-server                ClusterIP   10.106.67.9      <none>        3306/TCP                       3m6s
service/memcached                     ClusterIP   10.100.139.48    <none>        11211/TCP                      2m48s
service/rabbitmq                      ClusterIP   10.106.131.188   <none>        5672/TCP,25672/TCP,15672/TCP   2m40s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m40s
service/rabbitmq-mgr-7b1733           ClusterIP   10.98.27.66      <none>        80/TCP,443/TCP                 2m40s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           3m7s
deployment.apps/keystone-armada-operator      1/1     1            1           18m
deployment.apps/keystone-oslc-operator        1/1     1            1           18m
deployment.apps/mariadb-ingress               1/1     1            1           3m5s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           3m6s
deployment.apps/memcached-memcached           1/1     1            1           2m48s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       3m7s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       18m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       18m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       3m5s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       3m6s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       2m48s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     3m5s
statefulset.apps/rabbitmq-rabbitmq   1/1     2m40s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           27s        2m40s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          3m7s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          3m7s
```
 
```bash
prompt$ kubectl delete -f examples/argo/
armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
workflow.argoproj.io "armada-manifest" deleted
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
prompt$ argo list
NAME   STATUS   AGE   DURATION
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          19m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          19m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   19m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   18m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    31m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           19m
deployment.apps/keystone-oslc-operator     1/1     1            1           19m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       19m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       19m
```
 
```bash
prompt$
```
# ./argomanifest/argodelete.md
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          14m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          14m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   14m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   14m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    27m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           14m
deployment.apps/keystone-oslc-operator     1/1     1            1           14m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       14m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       14m
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
prompt$ kubectl apply -f examples/argo/
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
workflow.argoproj.io/armada-manifest created
```
 
```bash
prompt$ argo list
NAME                 STATUS    AGE   DURATION
keystone-install     Running   6s    6s
keystone-bootstrap   Running   6s    6s
armada-manifest      Running   10s   10s
```
 
```bash
prompt$ argo get armada-manifest
Name:                armada-manifest
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Running
Created:             Fri Apr 05 08:07:48 -0500 (22 seconds ago)
Started:             Fri Apr 05 08:07:48 -0500 (22 seconds ago)
Duration:            22 seconds

STEP                          PODNAME                     DURATION  MESSAGE
 ● armada-manifest
 ├-● keystone-infra-services
 | ├---✔ enable-mariadb       armada-manifest-842022813   3s
 | ├---✔ mariadb-ready        armada-manifest-361794910   7s
 | └---● enable-memcached     armada-manifest-4290341812  4s
 └-✔ openstack-keystone
   ├---✔ enable-keystone      armada-manifest-2789817682  3s
   └---✔ keystone-ready       armada-manifest-587324697   9s
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     1m    1m
keystone-install     Succeeded   1m    1m
armada-manifest      Succeeded   1m    37s
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
armada-manifest      2m5s
keystone-bootstrap   2m1s
keystone-install     2m1s
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
prompt$ vi traces/argo.txt
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Running     2m    2m
keystone-install     Succeeded   2m    1m
armada-manifest      Succeeded   2m    37s
```
 
```bash
prompt$ argo get keystone-install
Name:                keystone-install
Namespace:           default
ServiceAccount:      wf-keystone-sa
Status:              Succeeded
Created:             Fri Apr 05 08:07:52 -0500 (3 minutes ago)
Started:             Fri Apr 05 08:07:52 -0500 (3 minutes ago)
Finished:            Fri Apr 05 08:09:27 -0500 (1 minute ago)
Duration:            1 minute 35 seconds

STEP                                    PODNAME                      DURATION  MESSAGE
 ✔ keystone-install
 ├-✔ svc-mariadb                        keystone-install-2958852122  57s
 ├-✔ svc-memcached                      keystone-install-3754811077  36s
 ├-✔ svc-rabbitmq                       keystone-install-421213240   57s
 ├-✔ task-keystone-credential-setup(0)  keystone-install-1608858220  18s
 ├-✔ task-keystone-fernet-setup(0)      keystone-install-2716093231  18s
 ├-✔ wf-keystone-db-init
 | ├---✔ svc-mariadb-just-in-time       keystone-install-2401658373  4s
 | └---✔ task-keystone-db-init(0)       keystone-install-1659018021  4s
 ├-✔ wf-keystone-rabbit-init
 | ├---✔ svc-rabbitmq-just-in-time      keystone-install-2012559777  4s
 | └---✔ task-keystone-rabbit-init(0)   keystone-install-494016497   4s
 └-✔ wf-keystone-db-sync
   ├-·-✔ svc-mariadb-just-in-time       keystone-install-4222398536  3s
   | └-✔ svc-rabbitmq-just-in-time      keystone-install-3843173896  4s
   └---✔ task-keystone-db-sync(0)       keystone-install-1630370687  20s
```
 
```bash
prompt$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/armada-manifest-13034731                      0/1     Completed   0          2m48s
pod/armada-manifest-1873087745                    0/1     Completed   0          2m43s
pod/armada-manifest-2789817682                    0/1     Completed   0          3m11s
pod/armada-manifest-3014489540                    0/1     Completed   0          2m39s
pod/armada-manifest-361794910                     0/1     Completed   0          3m1s
pod/armada-manifest-4290341812                    0/1     Completed   0          2m53s
pod/armada-manifest-587324697                     0/1     Completed   0          3m7s
pod/armada-manifest-842022813                     0/1     Completed   0          3m11s
pod/keystone-api-6d7bd74f96-qjq8c                 1/1     Running     0          3m7s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          18m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          30s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          3m7s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          24s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          37s
pod/keystone-install-1608858220                   0/2     Completed   0          3m7s
pod/keystone-install-1630370687                   0/2     Completed   0          112s
pod/keystone-install-1659018021                   0/2     Completed   0          2m5s
pod/keystone-install-2012559777                   0/1     Completed   0          2m9s
pod/keystone-install-2401658373                   0/1     Completed   0          2m10s
pod/keystone-install-2716093231                   0/2     Completed   0          3m7s
pod/keystone-install-2958852122                   0/2     Completed   0          3m7s
pod/keystone-install-3754811077                   0/2     Completed   0          3m7s
pod/keystone-install-3843173896                   0/1     Completed   0          117s
pod/keystone-install-421213240                    0/2     Completed   0          3m7s
pod/keystone-install-4222398536                   0/1     Completed   0          117s
pod/keystone-install-494016497                    0/2     Completed   0          2m4s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          18m
pod/mariadb-ingress-5cfc68b86f-9qpss              1/1     Running     0          3m5s
pod/mariadb-ingress-error-pages-67c88c67d-v92gn   1/1     Running     0          3m5s
pod/mariadb-server-0                              1/1     Running     0          3m5s
pod/memcached-memcached-75447bffcf-2p87s          1/1     Running     0          2m48s
pod/rabbitmq-cluster-wait-vn4qs                   0/1     Completed   0          2m40s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          2m40s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.109.208.195   <none>        80/TCP,443/TCP                 3m7s
service/keystone-api                  ClusterIP   10.111.11.14     <none>        5000/TCP                       3m7s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       18m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       18m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        30m
service/mariadb                       ClusterIP   10.109.71.179    <none>        3306/TCP                       3m6s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              3m6s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         3m6s
service/mariadb-server                ClusterIP   10.106.67.9      <none>        3306/TCP                       3m6s
service/memcached                     ClusterIP   10.100.139.48    <none>        11211/TCP                      2m48s
service/rabbitmq                      ClusterIP   10.106.131.188   <none>        5672/TCP,25672/TCP,15672/TCP   2m40s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m40s
service/rabbitmq-mgr-7b1733           ClusterIP   10.98.27.66      <none>        80/TCP,443/TCP                 2m40s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           3m7s
deployment.apps/keystone-armada-operator      1/1     1            1           18m
deployment.apps/keystone-oslc-operator        1/1     1            1           18m
deployment.apps/mariadb-ingress               1/1     1            1           3m5s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           3m6s
deployment.apps/memcached-memcached           1/1     1            1           2m48s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       3m7s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       18m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       18m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       3m5s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       3m6s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       2m48s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     3m5s
statefulset.apps/rabbitmq-rabbitmq   1/1     2m40s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           27s        2m40s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          3m7s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          3m7s
```
 
```bash
prompt$ kubectl delete -f examples/argo/
armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
workflow.argoproj.io "armada-manifest" deleted
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
prompt$ argo list
NAME   STATUS   AGE   DURATION
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          19m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          19m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   19m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   18m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    31m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           19m
deployment.apps/keystone-oslc-operator     1/1     1            1           19m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       19m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       19m
```
 
```bash
prompt$
```
