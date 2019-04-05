# ./classichelm/helm.md

```bash
prompt$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   128m
```
 
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                           READY   STATUS    RESTARTS   AGE
kube-system   pod/calico-etcd-k7h9l                          1/1     Running   0          128m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9   1/1     Running   0          128m
kube-system   pod/calico-node-qjfj8                          1/1     Running   1          128m
kube-system   pod/coredns-fb8b8dccf-kvdg6                    1/1     Running   0          128m
kube-system   pod/coredns-fb8b8dccf-p9zhw                    1/1     Running   0          128m
kube-system   pod/etcd-airship                               1/1     Running   0          127m
kube-system   pod/kube-apiserver-airship                     1/1     Running   0          127m
kube-system   pod/kube-controller-manager-airship            1/1     Running   0          127m
kube-system   pod/kube-proxy-dv5ls                           1/1     Running   0          128m
kube-system   pod/kube-scheduler-airship                     1/1     Running   0          127m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z             1/1     Running   0          128m

NAMESPACE     NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
default       service/kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP                  128m
kube-system   service/calico-etcd     ClusterIP   10.96.232.136   <none>        6666/TCP                 128m
kube-system   service/kube-dns        ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   128m
kube-system   service/tiller-deploy   ClusterIP   10.102.66.38    <none>        44134/TCP                128m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   128m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       128m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            128m

NAMESPACE     NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/calico-kube-controllers   1/1     1            1           128m
kube-system   deployment.apps/coredns                   2/2     2            2           128m
kube-system   deployment.apps/tiller-deploy             1/1     1            1           128m

NAMESPACE     NAME                                                 DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/calico-kube-controllers-79bd896977   1         1         1       128m
kube-system   replicaset.apps/coredns-fb8b8dccf                    2         2         2       128m
kube-system   replicaset.apps/tiller-deploy-8458f6c667             1         1         1       128m
```
 
```bash
prompt$ cd helm-charts/
```
 
```bash
prompt/helm-charts$ ls
helm-toolkit  keystone  mariadb  memcached  official  rabbitmq
```
 
```bash
prompt/helm-charts$ cd memcached/
```
 
```bash
prompt/helm-charts/memcached$ helm install --name memcached .
NAME:   memcached
LAST DEPLOYED: Fri Apr  5 09:49:14 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                     DATA  AGE
memcached-memcached-bin  1     1s

==> v1/Deployment
NAME                 READY  UP-TO-DATE  AVAILABLE  AGE
memcached-memcached  0/1    1           0          1s

==> v1/Pod(related)
NAME                                  READY  STATUS    RESTARTS  AGE
memcached-memcached-75447bffcf-w7sz2  0/1    Init:0/1  0         1s

==> v1/Service
NAME       TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)    AGE
memcached  ClusterIP  10.100.34.194  <none>       11211/TCP  1s

==> v1/ServiceAccount
NAME                 SECRETS  AGE
memcached-memcached  1        1s


```
 
```bash
prompt/helm-charts/memcached$ cd ..
```
 
```bash
prompt/helm-charts$ cd rabbitmq/
```
 
```bash
prompt/helm-charts/rabbitmq$ helm install --name rabbitmq .
NAME:   rabbitmq
LAST DEPLOYED: Fri Apr  5 09:49:28 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                   DATA  AGE
rabbitmq-rabbitmq-bin  7     0s
rabbitmq-rabbitmq-etc  2     0s

==> v1/Job
NAME                   COMPLETIONS  DURATION  AGE
rabbitmq-cluster-wait  0/1          0s        0s

==> v1/Pod(related)
NAME                         READY  STATUS    RESTARTS  AGE
rabbitmq-cluster-wait-mgzfq  0/1    Init:0/1  0         0s
rabbitmq-rabbitmq-0          0/1    Init:0/4  0         0s

==> v1/Secret
NAME                    TYPE    DATA  AGE
rabbitmq-admin-user     Opaque  2     0s
rabbitmq-erlang-cookie  Opaque  1     0s

==> v1/Service
NAME                 TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)                       AGE
rabbitmq             ClusterIP  10.103.214.106  <none>       5672/TCP,25672/TCP,15672/TCP  0s
rabbitmq-dsv-7b1733  ClusterIP  None            <none>       5672/TCP,25672/TCP,15672/TCP  0s
rabbitmq-mgr-7b1733  ClusterIP  10.100.37.18    <none>       80/TCP,443/TCP                0s

==> v1/ServiceAccount
NAME                   SECRETS  AGE
rabbitmq-cluster-wait  1        0s
rabbitmq-rabbitmq      1        0s
rabbitmq-test          1        0s

==> v1/StatefulSet
NAME               READY  AGE
rabbitmq-rabbitmq  0/1    0s

==> v1beta1/Ingress
NAME                 HOSTS                                                                                          ADDRESS  PORTS  AGE
rabbitmq-mgr-7b1733  rabbitmq-mgr-7b1733,rabbitmq-mgr-7b1733.default,rabbitmq-mgr-7b1733.default.svc.cluster.local  80       0s

==> v1beta1/Role
NAME                                    AGE
rabbitmq-default-rabbitmq-cluster-wait  0s
rabbitmq-default-rabbitmq-test          0s
rabbitmq-rabbitmq                       0s

==> v1beta1/RoleBinding
NAME                            AGE
rabbitmq-rabbitmq               0s
rabbitmq-rabbitmq-cluster-wait  0s
rabbitmq-rabbitmq-test          0s


```
 
```bash
prompt/helm-charts/rabbitmq$ cd ..
```
 
```bash
prompt/helm-charts$ ls
helm-toolkit  keystone  mariadb  memcached  official  rabbitmq
```
 
```bash
prompt/helm-charts$ cd mariadb/
```
 
```bash
prompt/helm-charts/mariadb$ helm install --name mariadb .
NAME:   mariadb
LAST DEPLOYED: Fri Apr  5 09:49:42 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                  DATA  AGE
mariadb-bin           7     1s
mariadb-etc           5     1s
mariadb-services-tcp  1     1s

==> v1/Deployment
NAME                         READY  UP-TO-DATE  AVAILABLE  AGE
mariadb-ingress              0/1    1           0          1s
mariadb-ingress-error-pages  0/1    1           0          1s

==> v1/Pod(related)
NAME                                         READY  STATUS    RESTARTS  AGE
mariadb-ingress-5cfc68b86f-4vb9q             0/1    Init:0/1  0         1s
mariadb-ingress-error-pages-67c88c67d-qsl9x  0/1    Init:0/1  0         1s
mariadb-server-0                             0/1    Pending   0         1s

==> v1/Secret
NAME                      TYPE    DATA  AGE
mariadb-dbadmin-password  Opaque  1     1s
mariadb-dbsst-password    Opaque  1     1s
mariadb-secrets           Opaque  1     1s

==> v1/Service
NAME                         TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)            AGE
mariadb                      ClusterIP  10.108.63.190  <none>       3306/TCP           1s
mariadb-discovery            ClusterIP  None           <none>       3306/TCP,4567/TCP  1s
mariadb-ingress-error-pages  ClusterIP  None           <none>       80/TCP             1s
mariadb-server               ClusterIP  10.104.239.55  <none>       3306/TCP           1s

==> v1/ServiceAccount
NAME                         SECRETS  AGE
mariadb-ingress              1        1s
mariadb-ingress-error-pages  1        1s
mariadb-mariadb              1        1s

==> v1/StatefulSet
NAME            READY  AGE
mariadb-server  0/1    1s

==> v1beta1/PodDisruptionBudget
NAME            MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
mariadb-server  0              N/A              0                    1s

==> v1beta1/Role
NAME                             AGE
mariadb-default-mariadb-ingress  1s
mariadb-ingress                  1s
mariadb-mariadb                  1s

==> v1beta1/RoleBinding
NAME                     AGE
mariadb-ingress          1s
mariadb-mariadb          1s
mariadb-mariadb-ingress  1s


```
 
```bash
prompt/helm-charts/mariadb$ helm ls
NAME            REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
mariadb         1               Fri Apr  5 09:49:42 2019        DEPLOYED        mariadb-0.1.0                   default
memcached       1               Fri Apr  5 09:49:14 2019        DEPLOYED        memcached-0.1.0                 default
rabbitmq        1               Fri Apr  5 09:49:28 2019        DEPLOYED        rabbitmq-0.1.0                  default
```
 
```bash
prompt/helm-charts/mariadb$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/mariadb-ingress-5cfc68b86f-4vb9q              0/1     Running     0          51s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          51s
pod/mariadb-server-0                              1/1     Running     0          51s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          80s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          66s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          66s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        130m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       51s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              51s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         51s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       51s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      80s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   66s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   66s
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 66s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               0/1     1            0           51s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           51s
deployment.apps/memcached-memcached           1/1     1            1           80s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         0       51s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       51s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       80s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     51s
statefulset.apps/rabbitmq-rabbitmq   1/1     66s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           28s        66s
```
 
```bash
prompt/helm-charts/mariadb$ cd ..
```
 
```bash
prompt/helm-charts$ cd keystone/
```
 
```bash
prompt/helm-charts/keystone$ ls
charts  Chart.yaml  no-lifecycle-values.yaml  requirements.yaml  templates  values.yaml
```
 
```bash
prompt/helm-charts/keystone$ cp no-lifecycle-values.yaml values.yaml
```
 
```bash
prompt/helm-charts/keystone$ kubectl get crds
NAME                    CREATED AT
workflows.argoproj.io   2019-04-05T12:41:07Z
```
 
```bash
prompt/helm-charts/keystone$ kubectl delete crds workflows.argoproj.io
customresourcedefinition.apiextensions.k8s.io "workflows.argoproj.io" deleted
```
 
```bash
prompt/helm-charts/keystone$ kubectl get crds
No resources found.
```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/mariadb-ingress-5cfc68b86f-4vb9q              1/1     Running     0          105s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          105s
pod/mariadb-server-0                              1/1     Running     0          105s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          2m14s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          2m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          2m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        131m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       105s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              105s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         105s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       105s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      2m14s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   2m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 2m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               1/1     1            1           105s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           105s
deployment.apps/memcached-memcached           1/1     1            1           2m14s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       105s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       105s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       2m14s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     105s
statefulset.apps/rabbitmq-rabbitmq   1/1     2m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           28s        2m
```
 
```bash
prompt/helm-charts/keystone$ argo list
2019/04/05 09:51:31 the server could not find the requested resource (get workflows.argoproj.io)
```
 
```bash
prompt/helm-charts/keystone$ make getcrds
make: *** No rule to make target 'getcrds'.  Stop.
```
 
```bash
prompt/helm-charts/keystone$ cd ../..
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
error: the server doesn't have a resource type "workflows"
Makefile:131: recipe for target 'getcrds' failed
make: *** [getcrds] Error 1
```
 
```bash
prompt$ cd helm-charts/keystone/
```
 
```bash
prompt/helm-charts/keystone$ helm install --name keystone .
NAME:   keystone
LAST DEPLOYED: Fri Apr  5 09:52:05 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME          DATA  AGE
keystone-bin  14    1s

==> v1/Deployment
NAME          READY  UP-TO-DATE  AVAILABLE  AGE
keystone-api  0/1    1           0          1s

==> v1/Job
NAME                       COMPLETIONS  DURATION  AGE
keystone-bootstrap         0/1          1s        1s
keystone-credential-setup  0/1          1s        1s
keystone-db-init           0/1          0s        1s
keystone-db-sync           0/1          0s        1s
keystone-domain-manage     0/1          0s        0s
keystone-fernet-setup      0/1          0s        0s
keystone-rabbit-init       0/1          0s        0s

==> v1/Pod(related)
NAME                             READY  STATUS    RESTARTS  AGE
keystone-api-5867d7c957-2rfdw    0/1    Init:0/1  0         1s
keystone-bootstrap-npmvx         0/1    Pending   0         1s
keystone-credential-setup-7bk9t  0/1    Pending   0         1s
keystone-db-init-wk48d           0/1    Pending   0         1s
keystone-db-sync-62b6d           0/1    Pending   0         0s
keystone-domain-manage-gzsfq     0/1    Pending   0         0s
keystone-fernet-setup-827ts      0/1    Pending   0         0s
keystone-rabbit-init-x5fs8       0/1    Pending   0         0s

==> v1/Secret
NAME                      TYPE    DATA  AGE
keystone-credential-keys  Opaque  0     1s
keystone-db-admin         Opaque  2     1s
keystone-db-user          Opaque  2     1s
keystone-etc              Opaque  9     1s
keystone-fernet-keys      Opaque  0     1s
keystone-keystone-admin   Opaque  8     1s
keystone-keystone-test    Opaque  8     1s
keystone-rabbitmq-admin   Opaque  1     1s
keystone-rabbitmq-user    Opaque  1     1s

==> v1/Service
NAME          TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)         AGE
keystone      ClusterIP  10.104.113.43  <none>       80/TCP,443/TCP  1s
keystone-api  ClusterIP  10.97.119.160  <none>       5000/TCP        1s

==> v1/ServiceAccount
NAME                        SECRETS  AGE
keystone-api                1        1s
keystone-bootstrap          1        1s
keystone-credential-rotate  1        1s
keystone-credential-setup   1        1s
keystone-db-init            1        1s
keystone-db-sync            1        1s
keystone-domain-manage      1        1s
keystone-fernet-rotate      1        1s
keystone-fernet-setup       1        1s
keystone-rabbit-init        1        1s
keystone-test               1        1s

==> v1beta1/CronJob
NAME                        SCHEDULE      SUSPEND  ACTIVE  LAST SCHEDULE  AGE
keystone-credential-rotate  0 0 1 * *     False    0       <none>         0s
keystone-fernet-rotate      0 */12 * * *  False    0       <none>         0s

==> v1beta1/Ingress
NAME      HOSTS                                                         ADDRESS  PORTS  AGE
keystone  keystone,keystone.default,keystone.default.svc.cluster.local  80       0s

==> v1beta1/PodDisruptionBudget
NAME          MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
keystone-api  0              N/A              0                    1s

==> v1beta1/Role
NAME                                         AGE
keystone-credential-rotate                   1s
keystone-credential-setup                    1s
keystone-default-keystone-api                1s
keystone-default-keystone-bootstrap          1s
keystone-default-keystone-credential-rotate  1s
keystone-default-keystone-db-init            1s
keystone-default-keystone-db-sync            1s
keystone-default-keystone-domain-manage      1s
keystone-default-keystone-fernet-rotate      1s
keystone-default-keystone-rabbit-init        1s
keystone-default-keystone-test               1s
keystone-fernet-rotate                       1s
keystone-fernet-setup                        1s

==> v1beta1/RoleBinding
NAME                                 AGE
keystone-credential-rotate           1s
keystone-credential-setup            1s
keystone-fernet-rotate               1s
keystone-fernet-setup                1s
keystone-keystone-api                1s
keystone-keystone-bootstrap          1s
keystone-keystone-credential-rotate  1s
keystone-keystone-db-init            1s
keystone-keystone-db-sync            1s
keystone-keystone-domain-manage      1s
keystone-keystone-fernet-rotate      1s
keystone-keystone-rabbit-init        1s
keystone-keystone-test               1s


```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-5867d7c957-2rfdw                 1/1     Running     0          116s
pod/keystone-bootstrap-npmvx                      0/1     Completed   0          116s
pod/keystone-credential-setup-7bk9t               0/1     Completed   0          116s
pod/keystone-db-init-wk48d                        0/1     Completed   0          116s
pod/keystone-db-sync-62b6d                        0/1     Completed   0          115s
pod/keystone-domain-manage-gzsfq                  0/1     Completed   0          115s
pod/keystone-fernet-setup-827ts                   0/1     Completed   0          115s
pod/keystone-rabbit-init-x5fs8                    0/1     Completed   0          115s
pod/mariadb-ingress-5cfc68b86f-4vb9q              1/1     Running     0          4m19s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          4m19s
pod/mariadb-server-0                              1/1     Running     0          4m19s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          4m48s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          4m34s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          4m34s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.104.113.43    <none>        80/TCP,443/TCP                 116s
service/keystone-api                  ClusterIP   10.97.119.160    <none>        5000/TCP                       116s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        133m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       4m19s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              4m19s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         4m19s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       4m19s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      4m48s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   4m34s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   4m34s
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 4m34s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           116s
deployment.apps/mariadb-ingress               1/1     1            1           4m19s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           4m19s
deployment.apps/memcached-memcached           1/1     1            1           4m48s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-5867d7c957                 1         1         1       116s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       4m19s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       4m19s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       4m48s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     4m19s
statefulset.apps/rabbitmq-rabbitmq   1/1     4m34s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          1/1           90s        116s
job.batch/keystone-credential-setup   1/1           13s        116s
job.batch/keystone-db-init            1/1           10s        116s
job.batch/keystone-db-sync            1/1           33s        116s
job.batch/keystone-domain-manage      1/1           77s        115s
job.batch/keystone-fernet-setup       1/1           13s        115s
job.batch/keystone-rabbit-init        1/1           11s        115s
job.batch/rabbitmq-cluster-wait       1/1           28s        4m34s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          115s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          115s
```
 
```bash
prompt/helm-charts/keystone$ helm ls
NAME            REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
keystone        1               Fri Apr  5 09:52:05 2019        DEPLOYED        keystone-0.1.0                  default
mariadb         1               Fri Apr  5 09:49:42 2019        DEPLOYED        mariadb-0.1.0                   default
memcached       1               Fri Apr  5 09:49:14 2019        DEPLOYED        memcached-0.1.0                 default
rabbitmq        1               Fri Apr  5 09:49:28 2019        DEPLOYED        rabbitmq-0.1.0                  default
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge keystone
release "keystone" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge mariadb
release "mariadb" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge rabbitmq
release "rabbitmq" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge memcached
release "memcached" deleted
```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   135m
```
 
```bash
prompt/helm-charts/keystone$ helm ls
```
 
```bash
prompt/helm-charts/keystone$ argo list
2019/04/05 09:56:21 the server could not find the requested resource (get workflows.argoproj.io)
```
 
```bash
prompt/helm-charts/keystone$
```
# ./classichelm/helm.md
```bash
prompt$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   128m
```
 
```bash
prompt$ kubectl get all --all-namespaces
NAMESPACE     NAME                                           READY   STATUS    RESTARTS   AGE
kube-system   pod/calico-etcd-k7h9l                          1/1     Running   0          128m
kube-system   pod/calico-kube-controllers-79bd896977-gw9q9   1/1     Running   0          128m
kube-system   pod/calico-node-qjfj8                          1/1     Running   1          128m
kube-system   pod/coredns-fb8b8dccf-kvdg6                    1/1     Running   0          128m
kube-system   pod/coredns-fb8b8dccf-p9zhw                    1/1     Running   0          128m
kube-system   pod/etcd-airship                               1/1     Running   0          127m
kube-system   pod/kube-apiserver-airship                     1/1     Running   0          127m
kube-system   pod/kube-controller-manager-airship            1/1     Running   0          127m
kube-system   pod/kube-proxy-dv5ls                           1/1     Running   0          128m
kube-system   pod/kube-scheduler-airship                     1/1     Running   0          127m
kube-system   pod/tiller-deploy-8458f6c667-sgz5z             1/1     Running   0          128m

NAMESPACE     NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                  AGE
default       service/kubernetes      ClusterIP   10.96.0.1       <none>        443/TCP                  128m
kube-system   service/calico-etcd     ClusterIP   10.96.232.136   <none>        6666/TCP                 128m
kube-system   service/kube-dns        ClusterIP   10.96.0.10      <none>        53/UDP,53/TCP,9153/TCP   128m
kube-system   service/tiller-deploy   ClusterIP   10.102.66.38    <none>        44134/TCP                128m

NAMESPACE     NAME                         DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                     AGE
kube-system   daemonset.apps/calico-etcd   1         1         1       1            1           node-role.kubernetes.io/master=   128m
kube-system   daemonset.apps/calico-node   1         1         1       1            1           beta.kubernetes.io/os=linux       128m
kube-system   daemonset.apps/kube-proxy    1         1         1       1            1           <none>                            128m

NAMESPACE     NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
kube-system   deployment.apps/calico-kube-controllers   1/1     1            1           128m
kube-system   deployment.apps/coredns                   2/2     2            2           128m
kube-system   deployment.apps/tiller-deploy             1/1     1            1           128m

NAMESPACE     NAME                                                 DESIRED   CURRENT   READY   AGE
kube-system   replicaset.apps/calico-kube-controllers-79bd896977   1         1         1       128m
kube-system   replicaset.apps/coredns-fb8b8dccf                    2         2         2       128m
kube-system   replicaset.apps/tiller-deploy-8458f6c667             1         1         1       128m
```
 
```bash
prompt$ cd helm-charts/
```
 
```bash
prompt/helm-charts$ ls
helm-toolkit  keystone  mariadb  memcached  official  rabbitmq
```
 
```bash
prompt/helm-charts$ cd memcached/
```
 
```bash
prompt/helm-charts/memcached$ helm install --name memcached .
NAME:   memcached
LAST DEPLOYED: Fri Apr  5 09:49:14 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                     DATA  AGE
memcached-memcached-bin  1     1s

==> v1/Deployment
NAME                 READY  UP-TO-DATE  AVAILABLE  AGE
memcached-memcached  0/1    1           0          1s

==> v1/Pod(related)
NAME                                  READY  STATUS    RESTARTS  AGE
memcached-memcached-75447bffcf-w7sz2  0/1    Init:0/1  0         1s

==> v1/Service
NAME       TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)    AGE
memcached  ClusterIP  10.100.34.194  <none>       11211/TCP  1s

==> v1/ServiceAccount
NAME                 SECRETS  AGE
memcached-memcached  1        1s


```
 
```bash
prompt/helm-charts/memcached$ cd ..
```
 
```bash
prompt/helm-charts$ cd rabbitmq/
```
 
```bash
prompt/helm-charts/rabbitmq$ helm install --name rabbitmq .
NAME:   rabbitmq
LAST DEPLOYED: Fri Apr  5 09:49:28 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                   DATA  AGE
rabbitmq-rabbitmq-bin  7     0s
rabbitmq-rabbitmq-etc  2     0s

==> v1/Job
NAME                   COMPLETIONS  DURATION  AGE
rabbitmq-cluster-wait  0/1          0s        0s

==> v1/Pod(related)
NAME                         READY  STATUS    RESTARTS  AGE
rabbitmq-cluster-wait-mgzfq  0/1    Init:0/1  0         0s
rabbitmq-rabbitmq-0          0/1    Init:0/4  0         0s

==> v1/Secret
NAME                    TYPE    DATA  AGE
rabbitmq-admin-user     Opaque  2     0s
rabbitmq-erlang-cookie  Opaque  1     0s

==> v1/Service
NAME                 TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)                       AGE
rabbitmq             ClusterIP  10.103.214.106  <none>       5672/TCP,25672/TCP,15672/TCP  0s
rabbitmq-dsv-7b1733  ClusterIP  None            <none>       5672/TCP,25672/TCP,15672/TCP  0s
rabbitmq-mgr-7b1733  ClusterIP  10.100.37.18    <none>       80/TCP,443/TCP                0s

==> v1/ServiceAccount
NAME                   SECRETS  AGE
rabbitmq-cluster-wait  1        0s
rabbitmq-rabbitmq      1        0s
rabbitmq-test          1        0s

==> v1/StatefulSet
NAME               READY  AGE
rabbitmq-rabbitmq  0/1    0s

==> v1beta1/Ingress
NAME                 HOSTS                                                                                          ADDRESS  PORTS  AGE
rabbitmq-mgr-7b1733  rabbitmq-mgr-7b1733,rabbitmq-mgr-7b1733.default,rabbitmq-mgr-7b1733.default.svc.cluster.local  80       0s

==> v1beta1/Role
NAME                                    AGE
rabbitmq-default-rabbitmq-cluster-wait  0s
rabbitmq-default-rabbitmq-test          0s
rabbitmq-rabbitmq                       0s

==> v1beta1/RoleBinding
NAME                            AGE
rabbitmq-rabbitmq               0s
rabbitmq-rabbitmq-cluster-wait  0s
rabbitmq-rabbitmq-test          0s


```
 
```bash
prompt/helm-charts/rabbitmq$ cd ..
```
 
```bash
prompt/helm-charts$ ls
helm-toolkit  keystone  mariadb  memcached  official  rabbitmq
```
 
```bash
prompt/helm-charts$ cd mariadb/
```
 
```bash
prompt/helm-charts/mariadb$ helm install --name mariadb .
NAME:   mariadb
LAST DEPLOYED: Fri Apr  5 09:49:42 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME                  DATA  AGE
mariadb-bin           7     1s
mariadb-etc           5     1s
mariadb-services-tcp  1     1s

==> v1/Deployment
NAME                         READY  UP-TO-DATE  AVAILABLE  AGE
mariadb-ingress              0/1    1           0          1s
mariadb-ingress-error-pages  0/1    1           0          1s

==> v1/Pod(related)
NAME                                         READY  STATUS    RESTARTS  AGE
mariadb-ingress-5cfc68b86f-4vb9q             0/1    Init:0/1  0         1s
mariadb-ingress-error-pages-67c88c67d-qsl9x  0/1    Init:0/1  0         1s
mariadb-server-0                             0/1    Pending   0         1s

==> v1/Secret
NAME                      TYPE    DATA  AGE
mariadb-dbadmin-password  Opaque  1     1s
mariadb-dbsst-password    Opaque  1     1s
mariadb-secrets           Opaque  1     1s

==> v1/Service
NAME                         TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)            AGE
mariadb                      ClusterIP  10.108.63.190  <none>       3306/TCP           1s
mariadb-discovery            ClusterIP  None           <none>       3306/TCP,4567/TCP  1s
mariadb-ingress-error-pages  ClusterIP  None           <none>       80/TCP             1s
mariadb-server               ClusterIP  10.104.239.55  <none>       3306/TCP           1s

==> v1/ServiceAccount
NAME                         SECRETS  AGE
mariadb-ingress              1        1s
mariadb-ingress-error-pages  1        1s
mariadb-mariadb              1        1s

==> v1/StatefulSet
NAME            READY  AGE
mariadb-server  0/1    1s

==> v1beta1/PodDisruptionBudget
NAME            MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
mariadb-server  0              N/A              0                    1s

==> v1beta1/Role
NAME                             AGE
mariadb-default-mariadb-ingress  1s
mariadb-ingress                  1s
mariadb-mariadb                  1s

==> v1beta1/RoleBinding
NAME                     AGE
mariadb-ingress          1s
mariadb-mariadb          1s
mariadb-mariadb-ingress  1s


```
 
```bash
prompt/helm-charts/mariadb$ helm ls
NAME            REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
mariadb         1               Fri Apr  5 09:49:42 2019        DEPLOYED        mariadb-0.1.0                   default
memcached       1               Fri Apr  5 09:49:14 2019        DEPLOYED        memcached-0.1.0                 default
rabbitmq        1               Fri Apr  5 09:49:28 2019        DEPLOYED        rabbitmq-0.1.0                  default
```
 
```bash
prompt/helm-charts/mariadb$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/mariadb-ingress-5cfc68b86f-4vb9q              0/1     Running     0          51s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          51s
pod/mariadb-server-0                              1/1     Running     0          51s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          80s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          66s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          66s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        130m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       51s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              51s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         51s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       51s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      80s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   66s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   66s
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 66s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               0/1     1            0           51s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           51s
deployment.apps/memcached-memcached           1/1     1            1           80s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         0       51s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       51s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       80s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     51s
statefulset.apps/rabbitmq-rabbitmq   1/1     66s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           28s        66s
```
 
```bash
prompt/helm-charts/mariadb$ cd ..
```
 
```bash
prompt/helm-charts$ cd keystone/
```
 
```bash
prompt/helm-charts/keystone$ ls
charts  Chart.yaml  no-lifecycle-values.yaml  requirements.yaml  templates  values.yaml
```
 
```bash
prompt/helm-charts/keystone$ cp no-lifecycle-values.yaml values.yaml
```
 
```bash
prompt/helm-charts/keystone$ kubectl get crds
NAME                    CREATED AT
workflows.argoproj.io   2019-04-05T12:41:07Z
```
 
```bash
prompt/helm-charts/keystone$ kubectl delete crds workflows.argoproj.io
customresourcedefinition.apiextensions.k8s.io "workflows.argoproj.io" deleted
```
 
```bash
prompt/helm-charts/keystone$ kubectl get crds
No resources found.
```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/mariadb-ingress-5cfc68b86f-4vb9q              1/1     Running     0          105s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          105s
pod/mariadb-server-0                              1/1     Running     0          105s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          2m14s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          2m
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          2m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        131m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       105s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              105s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         105s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       105s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      2m14s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   2m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   2m
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 2m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/mariadb-ingress               1/1     1            1           105s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           105s
deployment.apps/memcached-memcached           1/1     1            1           2m14s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       105s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       105s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       2m14s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     105s
statefulset.apps/rabbitmq-rabbitmq   1/1     2m

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           28s        2m
```
 
```bash
prompt/helm-charts/keystone$ argo list
2019/04/05 09:51:31 the server could not find the requested resource (get workflows.argoproj.io)
```
 
```bash
prompt/helm-charts/keystone$ make getcrds
make: *** No rule to make target 'getcrds'.  Stop.
```
 
```bash
prompt/helm-charts/keystone$ cd ../..
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
error: the server doesn't have a resource type "workflows"
Makefile:131: recipe for target 'getcrds' failed
make: *** [getcrds] Error 1
```
 
```bash
prompt$ cd helm-charts/keystone/
```
 
```bash
prompt/helm-charts/keystone$ helm install --name keystone .
NAME:   keystone
LAST DEPLOYED: Fri Apr  5 09:52:05 2019
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/ConfigMap
NAME          DATA  AGE
keystone-bin  14    1s

==> v1/Deployment
NAME          READY  UP-TO-DATE  AVAILABLE  AGE
keystone-api  0/1    1           0          1s

==> v1/Job
NAME                       COMPLETIONS  DURATION  AGE
keystone-bootstrap         0/1          1s        1s
keystone-credential-setup  0/1          1s        1s
keystone-db-init           0/1          0s        1s
keystone-db-sync           0/1          0s        1s
keystone-domain-manage     0/1          0s        0s
keystone-fernet-setup      0/1          0s        0s
keystone-rabbit-init       0/1          0s        0s

==> v1/Pod(related)
NAME                             READY  STATUS    RESTARTS  AGE
keystone-api-5867d7c957-2rfdw    0/1    Init:0/1  0         1s
keystone-bootstrap-npmvx         0/1    Pending   0         1s
keystone-credential-setup-7bk9t  0/1    Pending   0         1s
keystone-db-init-wk48d           0/1    Pending   0         1s
keystone-db-sync-62b6d           0/1    Pending   0         0s
keystone-domain-manage-gzsfq     0/1    Pending   0         0s
keystone-fernet-setup-827ts      0/1    Pending   0         0s
keystone-rabbit-init-x5fs8       0/1    Pending   0         0s

==> v1/Secret
NAME                      TYPE    DATA  AGE
keystone-credential-keys  Opaque  0     1s
keystone-db-admin         Opaque  2     1s
keystone-db-user          Opaque  2     1s
keystone-etc              Opaque  9     1s
keystone-fernet-keys      Opaque  0     1s
keystone-keystone-admin   Opaque  8     1s
keystone-keystone-test    Opaque  8     1s
keystone-rabbitmq-admin   Opaque  1     1s
keystone-rabbitmq-user    Opaque  1     1s

==> v1/Service
NAME          TYPE       CLUSTER-IP     EXTERNAL-IP  PORT(S)         AGE
keystone      ClusterIP  10.104.113.43  <none>       80/TCP,443/TCP  1s
keystone-api  ClusterIP  10.97.119.160  <none>       5000/TCP        1s

==> v1/ServiceAccount
NAME                        SECRETS  AGE
keystone-api                1        1s
keystone-bootstrap          1        1s
keystone-credential-rotate  1        1s
keystone-credential-setup   1        1s
keystone-db-init            1        1s
keystone-db-sync            1        1s
keystone-domain-manage      1        1s
keystone-fernet-rotate      1        1s
keystone-fernet-setup       1        1s
keystone-rabbit-init        1        1s
keystone-test               1        1s

==> v1beta1/CronJob
NAME                        SCHEDULE      SUSPEND  ACTIVE  LAST SCHEDULE  AGE
keystone-credential-rotate  0 0 1 * *     False    0       <none>         0s
keystone-fernet-rotate      0 */12 * * *  False    0       <none>         0s

==> v1beta1/Ingress
NAME      HOSTS                                                         ADDRESS  PORTS  AGE
keystone  keystone,keystone.default,keystone.default.svc.cluster.local  80       0s

==> v1beta1/PodDisruptionBudget
NAME          MIN AVAILABLE  MAX UNAVAILABLE  ALLOWED DISRUPTIONS  AGE
keystone-api  0              N/A              0                    1s

==> v1beta1/Role
NAME                                         AGE
keystone-credential-rotate                   1s
keystone-credential-setup                    1s
keystone-default-keystone-api                1s
keystone-default-keystone-bootstrap          1s
keystone-default-keystone-credential-rotate  1s
keystone-default-keystone-db-init            1s
keystone-default-keystone-db-sync            1s
keystone-default-keystone-domain-manage      1s
keystone-default-keystone-fernet-rotate      1s
keystone-default-keystone-rabbit-init        1s
keystone-default-keystone-test               1s
keystone-fernet-rotate                       1s
keystone-fernet-setup                        1s

==> v1beta1/RoleBinding
NAME                                 AGE
keystone-credential-rotate           1s
keystone-credential-setup            1s
keystone-fernet-rotate               1s
keystone-fernet-setup                1s
keystone-keystone-api                1s
keystone-keystone-bootstrap          1s
keystone-keystone-credential-rotate  1s
keystone-keystone-db-init            1s
keystone-keystone-db-sync            1s
keystone-keystone-domain-manage      1s
keystone-keystone-fernet-rotate      1s
keystone-keystone-rabbit-init        1s
keystone-keystone-test               1s


```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                                              READY   STATUS      RESTARTS   AGE
pod/keystone-api-5867d7c957-2rfdw                 1/1     Running     0          116s
pod/keystone-bootstrap-npmvx                      0/1     Completed   0          116s
pod/keystone-credential-setup-7bk9t               0/1     Completed   0          116s
pod/keystone-db-init-wk48d                        0/1     Completed   0          116s
pod/keystone-db-sync-62b6d                        0/1     Completed   0          115s
pod/keystone-domain-manage-gzsfq                  0/1     Completed   0          115s
pod/keystone-fernet-setup-827ts                   0/1     Completed   0          115s
pod/keystone-rabbit-init-x5fs8                    0/1     Completed   0          115s
pod/mariadb-ingress-5cfc68b86f-4vb9q              1/1     Running     0          4m19s
pod/mariadb-ingress-error-pages-67c88c67d-qsl9x   1/1     Running     0          4m19s
pod/mariadb-server-0                              1/1     Running     0          4m19s
pod/memcached-memcached-75447bffcf-w7sz2          1/1     Running     0          4m48s
pod/rabbitmq-cluster-wait-mgzfq                   0/1     Completed   0          4m34s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          4m34s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.104.113.43    <none>        80/TCP,443/TCP                 116s
service/keystone-api                  ClusterIP   10.97.119.160    <none>        5000/TCP                       116s
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        133m
service/mariadb                       ClusterIP   10.108.63.190    <none>        3306/TCP                       4m19s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              4m19s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         4m19s
service/mariadb-server                ClusterIP   10.104.239.55    <none>        3306/TCP                       4m19s
service/memcached                     ClusterIP   10.100.34.194    <none>        11211/TCP                      4m48s
service/rabbitmq                      ClusterIP   10.103.214.106   <none>        5672/TCP,25672/TCP,15672/TCP   4m34s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   4m34s
service/rabbitmq-mgr-7b1733           ClusterIP   10.100.37.18     <none>        80/TCP,443/TCP                 4m34s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           116s
deployment.apps/mariadb-ingress               1/1     1            1           4m19s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           4m19s
deployment.apps/memcached-memcached           1/1     1            1           4m48s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-5867d7c957                 1         1         1       116s
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       4m19s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       4m19s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       4m48s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     4m19s
statefulset.apps/rabbitmq-rabbitmq   1/1     4m34s

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          1/1           90s        116s
job.batch/keystone-credential-setup   1/1           13s        116s
job.batch/keystone-db-init            1/1           10s        116s
job.batch/keystone-db-sync            1/1           33s        116s
job.batch/keystone-domain-manage      1/1           77s        115s
job.batch/keystone-fernet-setup       1/1           13s        115s
job.batch/keystone-rabbit-init        1/1           11s        115s
job.batch/rabbitmq-cluster-wait       1/1           28s        4m34s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          115s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          115s
```
 
```bash
prompt/helm-charts/keystone$ helm ls
NAME            REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
keystone        1               Fri Apr  5 09:52:05 2019        DEPLOYED        keystone-0.1.0                  default
mariadb         1               Fri Apr  5 09:49:42 2019        DEPLOYED        mariadb-0.1.0                   default
memcached       1               Fri Apr  5 09:49:14 2019        DEPLOYED        memcached-0.1.0                 default
rabbitmq        1               Fri Apr  5 09:49:28 2019        DEPLOYED        rabbitmq-0.1.0                  default
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge keystone
release "keystone" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge mariadb
release "mariadb" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge rabbitmq
release "rabbitmq" deleted
```
 
```bash
prompt/helm-charts/keystone$ helm delete --purge memcached
release "memcached" deleted
```
 
```bash
prompt/helm-charts/keystone$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   135m
```
 
```bash
prompt/helm-charts/keystone$ helm ls
```
 
```bash
prompt/helm-charts/keystone$ argo list
2019/04/05 09:56:21 the server could not find the requested resource (get workflows.argoproj.io)
```
 
```bash
prompt/helm-charts/keystone$
```
