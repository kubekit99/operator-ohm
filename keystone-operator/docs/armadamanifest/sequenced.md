# ./armadamanifest/sequenced.md

```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          105m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          105m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   105m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   105m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    118m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           105m
deployment.apps/keystone-oslc-operator     1/1     1            1           105m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       105m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       105m
```
 
```bash
prompt$ kubectl apply -f examples/sequenced/
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
armadachartgroup.armada.airshipit.org/keystone-infra-services created
armadachartgroup.armada.airshipit.org/openstack-keystone created
armadamanifest.armada.airshipit.org/armada-manifest created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   5s
keystone-install     5s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
NAME                      STATE                    TARGET STATE   SATISFIED
keystone-infra-services   pending-initialization   deployed       false
openstack-keystone        pending-initialization   deployed       false
kubectl get armadamanifests.armada.airshipit.org
NAME              STATE           TARGET STATE   SATISFIED
armada-manifest   uninitialized   deployed       false
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
prompt$ kubectl describe acg keystone-infra-services
Name:         keystone-infra-services
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChartGroup","metadata":{"annotations":{},"name":"keystone-infra-services","nam...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChartGroup
Metadata:
  Creation Timestamp:  2019-04-05T14:38:36Z
  Finalizers:
    uninstall-acg
  Generation:  2
  Owner References:
    API Version:           armada.airshipit.org/v1alpha1
    Block Owner Deletion:  true
    Controller:            true
    Kind:                  ArmadaManifest
    Name:                  armada-manifest
    UID:                   7f13f4da-57b0-11e9-bbfe-0800272e6982
  Resource Version:        8903
  Self Link:               /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadachartgroups/keystone-infra-services
  UID:                     7f0f7597-57b0-11e9-bbfe-0800272e6982
Spec:
  Chart Group:
    mariadb
    memcached
    rabbitmq
  Description:   Keystone Infra Services
  Sequenced:     true
  Target State:  deployed
Status:
  Actual State:  pending-initialization
  Conditions:
    Last Transition Time:  2019-04-05T14:38:38Z
    Reason:                ChartGroup is enabled
    Status:                True
    Type:                  Enabled
    Last Transition Time:  2019-04-05T14:38:38Z
    Status:                True
    Type:                  Initializing
  Satisfied:               false
Events:
  Type     Reason   Age                From          Message
  ----     ------   ----               ----          -------
  Warning  Enabled  69s (x2 over 69s)  acg-recorder  ChartGroup is disabled
  Normal   Enabled  65s (x5 over 69s)  acg-recorder  ChartGroup is enabled
```
 
```bash
prompt$ kubectl describe amf
Name:         armada-manifest
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaManifest","metadata":{"annotations":{},"name":"armada-manifest","namespace":"d...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaManifest
Metadata:
  Creation Timestamp:  2019-04-05T14:38:36Z
  Finalizers:
    uninstall-amf
  Generation:        1
  Resource Version:  8814
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadamanifests/armada-manifest
  UID:               7f13f4da-57b0-11e9-bbfe-0800272e6982
Spec:
  Chart Groups:
    keystone-infra-services
    openstack-keystone
  Release Prefix:  armada
  Target State:    deployed
Status:
  Actual State:  uninitialized
  Conditions:
    Last Transition Time:  2019-04-05T14:38:36Z
    Reason:                Manifest is enabled
    Status:                True
    Type:                  Enabled
    Last Transition Time:  2019-04-05T14:38:36Z
    Status:                True
    Type:                  Initializing
  Satisfied:               false
Events:
  Type    Reason   Age                From          Message
  ----    ------   ----               ----          -------
  Normal  Enabled  91s (x3 over 92s)  amf-recorder  Manifest is enabled
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Succeeded   3m    2m
keystone-install     Succeeded   3m    1m
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   3m55s
keystone-install     3m55s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
NAME                      STATE                    TARGET STATE   SATISFIED
keystone-infra-services   pending-initialization   deployed       false
openstack-keystone        pending-initialization   deployed       false
kubectl get armadamanifests.armada.airshipit.org
NAME              STATE           TARGET STATE   SATISFIED
armada-manifest   uninitialized   deployed       false
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
pod/keystone-api-6d7bd74f96-wdrjb                 1/1     Running     0          4m52s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          110m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          2m17s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          4m52s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          2m11s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          2m22s
pod/keystone-install-1608858220                   0/2     Completed   0          4m51s
pod/keystone-install-1630370687                   0/2     Completed   0          3m25s
pod/keystone-install-1659018021                   0/2     Completed   0          3m35s
pod/keystone-install-2012559777                   0/1     Completed   0          4m16s
pod/keystone-install-2401658373                   0/1     Completed   0          3m38s
pod/keystone-install-2716093231                   0/2     Completed   0          4m51s
pod/keystone-install-2958852122                   0/2     Completed   0          4m52s
pod/keystone-install-3754811077                   0/2     Completed   0          4m51s
pod/keystone-install-3843173896                   0/1     Completed   0          3m29s
pod/keystone-install-421213240                    0/2     Completed   0          4m52s
pod/keystone-install-4222398536                   0/1     Completed   0          3m29s
pod/keystone-install-494016497                    0/2     Completed   0          4m12s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          110m
pod/mariadb-ingress-5cfc68b86f-q6tz9              1/1     Running     0          4m53s
pod/mariadb-ingress-error-pages-67c88c67d-2vddk   1/1     Running     0          4m53s
pod/mariadb-server-0                              1/1     Running     0          4m53s
pod/memcached-memcached-75447bffcf-d9vtr          1/1     Running     0          4m49s
pod/rabbitmq-cluster-wait-klk2j                   0/1     Completed   0          4m49s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          4m49s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.109.236.218   <none>        80/TCP,443/TCP                 4m52s
service/keystone-api                  ClusterIP   10.101.48.61     <none>        5000/TCP                       4m52s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       110m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       110m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        123m
service/mariadb                       ClusterIP   10.97.251.223    <none>        3306/TCP                       4m53s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              4m53s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         4m53s
service/mariadb-server                ClusterIP   10.111.173.104   <none>        3306/TCP                       4m53s
service/memcached                     ClusterIP   10.110.125.14    <none>        11211/TCP                      4m49s
service/rabbitmq                      ClusterIP   10.107.157.11    <none>        5672/TCP,25672/TCP,15672/TCP   4m49s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   4m49s
service/rabbitmq-mgr-7b1733           ClusterIP   10.98.58.177     <none>        80/TCP,443/TCP                 4m49s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           4m52s
deployment.apps/keystone-armada-operator      1/1     1            1           110m
deployment.apps/keystone-oslc-operator        1/1     1            1           110m
deployment.apps/mariadb-ingress               1/1     1            1           4m53s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           4m53s
deployment.apps/memcached-memcached           1/1     1            1           4m49s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       4m52s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       110m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       110m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       4m53s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       4m53s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       4m49s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     4m53s
statefulset.apps/rabbitmq-rabbitmq   1/1     4m49s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           36s        4m49s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          4m52s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          4m52s
```
 
```bash
prompt$ kubectl delete -f examples/sequenced/
armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
armadachartgroup.armada.airshipit.org "keystone-infra-services" deleted
armadachartgroup.armada.airshipit.org "openstack-keystone" deleted
armadamanifest.armada.airshipit.org "armada-manifest" deleted
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          111m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          111m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   111m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   111m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    124m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           111m
deployment.apps/keystone-oslc-operator     1/1     1            1           111m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       111m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       111m
```
 
```bash
prompt$ make get crds
make: *** No rule to make target 'get'.  Stop.
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
prompt$ helm ls
NAME    REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
argo    1               Fri Apr  5 07:41:06 2019        DEPLOYED        argo-0.3.1                      argo
```
 
```bash
prompt$ make purge-oslc && make purge-armada
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
```
 
```bash
prompt$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   126m
```
 
```bash
prompt$
```
# ./armadamanifest/sequenced.md
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          105m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          105m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   105m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   105m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    118m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           105m
deployment.apps/keystone-oslc-operator     1/1     1            1           105m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       105m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       105m
```
 
```bash
prompt$ kubectl apply -f examples/sequenced/
armadachart.armada.airshipit.org/mariadb created
armadachart.armada.airshipit.org/memcached created
armadachart.armada.airshipit.org/rabbitmq created
armadachart.armada.airshipit.org/keystone created
armadachartgroup.armada.airshipit.org/keystone-infra-services created
armadachartgroup.armada.airshipit.org/openstack-keystone created
armadamanifest.armada.airshipit.org/armada-manifest created
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   5s
keystone-install     5s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
NAME                      STATE                    TARGET STATE   SATISFIED
keystone-infra-services   pending-initialization   deployed       false
openstack-keystone        pending-initialization   deployed       false
kubectl get armadamanifests.armada.airshipit.org
NAME              STATE           TARGET STATE   SATISFIED
armada-manifest   uninitialized   deployed       false
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
prompt$ kubectl describe acg keystone-infra-services
Name:         keystone-infra-services
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaChartGroup","metadata":{"annotations":{},"name":"keystone-infra-services","nam...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaChartGroup
Metadata:
  Creation Timestamp:  2019-04-05T14:38:36Z
  Finalizers:
    uninstall-acg
  Generation:  2
  Owner References:
    API Version:           armada.airshipit.org/v1alpha1
    Block Owner Deletion:  true
    Controller:            true
    Kind:                  ArmadaManifest
    Name:                  armada-manifest
    UID:                   7f13f4da-57b0-11e9-bbfe-0800272e6982
  Resource Version:        8903
  Self Link:               /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadachartgroups/keystone-infra-services
  UID:                     7f0f7597-57b0-11e9-bbfe-0800272e6982
Spec:
  Chart Group:
    mariadb
    memcached
    rabbitmq
  Description:   Keystone Infra Services
  Sequenced:     true
  Target State:  deployed
Status:
  Actual State:  pending-initialization
  Conditions:
    Last Transition Time:  2019-04-05T14:38:38Z
    Reason:                ChartGroup is enabled
    Status:                True
    Type:                  Enabled
    Last Transition Time:  2019-04-05T14:38:38Z
    Status:                True
    Type:                  Initializing
  Satisfied:               false
Events:
  Type     Reason   Age                From          Message
  ----     ------   ----               ----          -------
  Warning  Enabled  69s (x2 over 69s)  acg-recorder  ChartGroup is disabled
  Normal   Enabled  65s (x5 over 69s)  acg-recorder  ChartGroup is enabled
```
 
```bash
prompt$ kubectl describe amf
Name:         armada-manifest
Namespace:    default
Labels:       <none>
Annotations:  kubectl.kubernetes.io/last-applied-configuration:
                {"apiVersion":"armada.airshipit.org/v1alpha1","kind":"ArmadaManifest","metadata":{"annotations":{},"name":"armada-manifest","namespace":"d...
API Version:  armada.airshipit.org/v1alpha1
Kind:         ArmadaManifest
Metadata:
  Creation Timestamp:  2019-04-05T14:38:36Z
  Finalizers:
    uninstall-amf
  Generation:        1
  Resource Version:  8814
  Self Link:         /apis/armada.airshipit.org/v1alpha1/namespaces/default/armadamanifests/armada-manifest
  UID:               7f13f4da-57b0-11e9-bbfe-0800272e6982
Spec:
  Chart Groups:
    keystone-infra-services
    openstack-keystone
  Release Prefix:  armada
  Target State:    deployed
Status:
  Actual State:  uninitialized
  Conditions:
    Last Transition Time:  2019-04-05T14:38:36Z
    Reason:                Manifest is enabled
    Status:                True
    Type:                  Enabled
    Last Transition Time:  2019-04-05T14:38:36Z
    Status:                True
    Type:                  Initializing
  Satisfied:               false
Events:
  Type    Reason   Age                From          Message
  ----    ------   ----               ----          -------
  Normal  Enabled  91s (x3 over 92s)  amf-recorder  Manifest is enabled
```
 
```bash
prompt$ argo list
NAME                 STATUS      AGE   DURATION
keystone-bootstrap   Succeeded   3m    2m
keystone-install     Succeeded   3m    1m
```
 
```bash
prompt$ make getcrds
kubectl get workflows.argoproj.io
NAME                 AGE
keystone-bootstrap   3m55s
keystone-install     3m55s
kubectl get armadacharts.armada.airshipit.org
NAME        STATE      TARGET STATE   SATISFIED
keystone    deployed   deployed       true
mariadb     deployed   deployed       true
memcached   deployed   deployed       true
rabbitmq    deployed   deployed       true
kubectl get armadachartgroups.armada.airshipit.org
NAME                      STATE                    TARGET STATE   SATISFIED
keystone-infra-services   pending-initialization   deployed       false
openstack-keystone        pending-initialization   deployed       false
kubectl get armadamanifests.armada.airshipit.org
NAME              STATE           TARGET STATE   SATISFIED
armada-manifest   uninitialized   deployed       false
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
pod/keystone-api-6d7bd74f96-wdrjb                 1/1     Running     0          4m52s
pod/keystone-armada-operator-5cfccc74fb-zc7kq     1/1     Running     0          110m
pod/keystone-bootstrap-1692958623                 0/2     Completed   0          2m17s
pod/keystone-bootstrap-2032358733                 0/2     Completed   0          4m52s
pod/keystone-bootstrap-3637380442                 0/2     Completed   0          2m11s
pod/keystone-bootstrap-3824180533                 0/2     Completed   0          2m22s
pod/keystone-install-1608858220                   0/2     Completed   0          4m51s
pod/keystone-install-1630370687                   0/2     Completed   0          3m25s
pod/keystone-install-1659018021                   0/2     Completed   0          3m35s
pod/keystone-install-2012559777                   0/1     Completed   0          4m16s
pod/keystone-install-2401658373                   0/1     Completed   0          3m38s
pod/keystone-install-2716093231                   0/2     Completed   0          4m51s
pod/keystone-install-2958852122                   0/2     Completed   0          4m52s
pod/keystone-install-3754811077                   0/2     Completed   0          4m51s
pod/keystone-install-3843173896                   0/1     Completed   0          3m29s
pod/keystone-install-421213240                    0/2     Completed   0          4m52s
pod/keystone-install-4222398536                   0/1     Completed   0          3m29s
pod/keystone-install-494016497                    0/2     Completed   0          4m12s
pod/keystone-oslc-operator-747bb84698-cqvb8       1/1     Running     0          110m
pod/mariadb-ingress-5cfc68b86f-q6tz9              1/1     Running     0          4m53s
pod/mariadb-ingress-error-pages-67c88c67d-2vddk   1/1     Running     0          4m53s
pod/mariadb-server-0                              1/1     Running     0          4m53s
pod/memcached-memcached-75447bffcf-d9vtr          1/1     Running     0          4m49s
pod/rabbitmq-cluster-wait-klk2j                   0/1     Completed   0          4m49s
pod/rabbitmq-rabbitmq-0                           1/1     Running     0          4m49s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.109.236.218   <none>        80/TCP,443/TCP                 4m52s
service/keystone-api                  ClusterIP   10.101.48.61     <none>        5000/TCP                       4m52s
service/keystone-armada-operator      ClusterIP   10.102.35.39     <none>        8383/TCP                       110m
service/keystone-oslc-operator        ClusterIP   10.97.81.146     <none>        8383/TCP                       110m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        123m
service/mariadb                       ClusterIP   10.97.251.223    <none>        3306/TCP                       4m53s
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              4m53s
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         4m53s
service/mariadb-server                ClusterIP   10.111.173.104   <none>        3306/TCP                       4m53s
service/memcached                     ClusterIP   10.110.125.14    <none>        11211/TCP                      4m49s
service/rabbitmq                      ClusterIP   10.107.157.11    <none>        5672/TCP,25672/TCP,15672/TCP   4m49s
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   4m49s
service/rabbitmq-mgr-7b1733           ClusterIP   10.98.58.177     <none>        80/TCP,443/TCP                 4m49s

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           4m52s
deployment.apps/keystone-armada-operator      1/1     1            1           110m
deployment.apps/keystone-oslc-operator        1/1     1            1           110m
deployment.apps/mariadb-ingress               1/1     1            1           4m53s
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           4m53s
deployment.apps/memcached-memcached           1/1     1            1           4m49s

NAME                                                    DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-6d7bd74f96                 1         1         1       4m52s
replicaset.apps/keystone-armada-operator-5cfccc74fb     1         1         1       110m
replicaset.apps/keystone-oslc-operator-747bb84698       1         1         1       110m
replicaset.apps/mariadb-ingress-5cfc68b86f              1         1         1       4m53s
replicaset.apps/mariadb-ingress-error-pages-67c88c67d   1         1         1       4m53s
replicaset.apps/memcached-memcached-75447bffcf          1         1         1       4m49s

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     4m53s
statefulset.apps/rabbitmq-rabbitmq   1/1     4m49s

NAME                              COMPLETIONS   DURATION   AGE
job.batch/rabbitmq-cluster-wait   1/1           36s        4m49s

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          4m52s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          4m52s
```
 
```bash
prompt$ kubectl delete -f examples/sequenced/
armadachart.armada.airshipit.org "mariadb" deleted
armadachart.armada.airshipit.org "memcached" deleted
armadachart.armada.airshipit.org "rabbitmq" deleted
armadachart.armada.airshipit.org "keystone" deleted
armadachartgroup.armada.airshipit.org "keystone-infra-services" deleted
armadachartgroup.armada.airshipit.org "openstack-keystone" deleted
armadamanifest.armada.airshipit.org "armada-manifest" deleted
```
 
```bash
prompt$ kubectl get all
NAME                                            READY   STATUS    RESTARTS   AGE
pod/keystone-armada-operator-5cfccc74fb-zc7kq   1/1     Running   0          111m
pod/keystone-oslc-operator-747bb84698-cqvb8     1/1     Running   0          111m

NAME                               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/keystone-armada-operator   ClusterIP   10.102.35.39   <none>        8383/TCP   111m
service/keystone-oslc-operator     ClusterIP   10.97.81.146   <none>        8383/TCP   111m
service/kubernetes                 ClusterIP   10.96.0.1      <none>        443/TCP    124m

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-armada-operator   1/1     1            1           111m
deployment.apps/keystone-oslc-operator     1/1     1            1           111m

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-armada-operator-5cfccc74fb   1         1         1       111m
replicaset.apps/keystone-oslc-operator-747bb84698     1         1         1       111m
```
 
```bash
prompt$ make get crds
make: *** No rule to make target 'get'.  Stop.
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
prompt$ helm ls
NAME    REVISION        UPDATED                         STATUS          CHART           APP VERSION     NAMESPACE
argo    1               Fri Apr  5 07:41:06 2019        DEPLOYED        argo-0.3.1                      argo
```
 
```bash
prompt$ make purge-oslc && make purge-armada
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
```
 
```bash
prompt$ kubectl get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   126m
```
