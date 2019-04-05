# ./argomanifest/argo.md

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
prompt$
```
# ./argomanifest/argo.md
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
prompt$
```
