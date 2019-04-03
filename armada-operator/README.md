# Kubernetes Operator for Armada and Helm

# Introduction

This README is mainly used a log / wiki of results and challenges encounted during the POC

## operator-sdk vs kubebuilder.

Things to clean up:
Did not had time to sort how to completly migrate from operator-sdh to kubebuilder.
Most of the scaffolding is done using the operator-sdk but once the default files are created,
the build process mainly relies on kubebuilder

## armada-operator code directory structure

###  cmd

Contains the main.go for the armada operator

###  pkg/apis/

Contains the golang definition of the CRDs. `make generate` will recreate the yaml definitions
of the CRDs that have to be provided to kubectl in order to deploy the new CRDs.
This current version of the operator uses "act" for shortname of ArmadaChart,
"acg" as shortname for ArmadaChartGroup and "amf" as shortname for ArmadaManifest.

The first version of the golang code has generated using tool such as "schema-generate" from
the schema definition provided with airship-armada project.

###  pkg/services

Contains the bulk of the interfaces used by the armada controller.

###  pkg/helm, pkg/helmv2 and pkg/helmv3

To get the process doing, some of the code for ArmadaChart handling is coming from the 
operator-sdk helm-operator. That code is relying on tiller component which is gone for Helm3.
Hence the three directory helm (Interface and Common code), helmv2 (tiller) and helmv3.

The golang package structure is different between helmv2 and helmv3. The Armada Operator will
most likely ultimatly have support two branches. In order to delay that milestone, the golang
code has been instrumentated with "v2" and "v3" tags which allows to compile either the
helm v3 version of the operator or the helm v3 version.

###  pkg/armada directory

Mainly contain the code for ArmadaChartGroup, ArmadaManifest as will ArmadaBackupLocation handling

###  pkg/controller directory

Contains the controller and the "Reconcile" functions for ArmadaChart, ArmadaChartGroup and ArmadaManifest.
There are currently three controllers (act-controller, acg-controller and amf-controller).

# Code changes.

## Adjusting the ArmadaOperator CRDs

Upon change of the CRD golang definition, the yaml files have to be regenerated

Note 1: Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

Note 2: The generation tool seems to comply with some of OpenAPI specs. The "validation" schema added
to in the CRD yaml definition does not contain fields using underscore. 
Most of those fields containing underscore where defined such a way in the original airship-armada.

```bash
make generate
```

## Compiling the armada-operator

To keep the directory tree ligthweight, the vendor directory is not checked in in the current repo.
TODO: Since the operator is only using one git branch, the developer has to comment out the helmv2
and add the helmv3 in Gopkg.toml if he wants to build the helmv3 version of the operator. This is still WIP.

```bash
dep ensure
```

To build the v2 version
```bash
make docker-build-v2
```

To build the v3 version
```bash
make docker-build-v3
```

## Run unit test and preintegration tests.

If you installed kubebuilder on your system, you will have access
to a standalone apiserver, etcd server and kubectl.

Because of a lack of time, the current makefile test statement,
will attempt to stop your kubelet and associated container in your local
kubernetes cluster, before starting apiserver, etcdserver.
TODO: We still need to figure out if it necessary

In order to run the unit tests and the e2e integration tests:
```bash
make unittest
```

# Deploying the operator.

Note the current deployment of the operator relies itself on helm.

To install the helm v2 version
```bash
make install-v2
```

To install the helm v3 version
```bash
make install-v3
```

# Testing the armada-controller

##  helm-charts/testchart directory

For testing purpose the current Docker file includes a dummy chart deliverd under armada-charts.
This removes the needs to access external chart repository which is also an aspect of helm changing from 2 to 3.

## examples/armada

In that directory, the ArmadaChart are enabled by default and the charts
are installed as soon as the ArmadaChart CRD are created.

### Deployment

Upon creation of the custom resource, the controller will
- Deploy the Armada Manifest described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/armada
kubectl describe amf/simple-armada
kubectl get amf
kubectl get acg
kubectl get act
```

### Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl delete deployment.apps/blog-2-testchart
kubectl get all
kubectl describe act blog-2
```

### Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding Armada Manifest should be uninstalled.

```bash
kubectl delete -f simple/armada
```

## examples/stepbystep

This directory contains invidual act,acg and amf files which allow the "step" by "step" testing of the deployment.

## examples/argo

In that directory, the ArmadaChart (act) are disabled by default and the charts not installed
automatically when the act CR are created.
This example assumes that the argo controller has been installed. When the "argo worflow"
CR is created, the "argo controller" is waked up and it orchestrates the enablement of the ArmadaChart
according to the worflow.

Note: You need to have argo installed in your cluster.

```bash
kubectl apply -f example/argo

armadachart.armada.airshipit.org/blog-1 created
armadachart.armada.airshipit.org/blog-2 created
workflow.argoproj.io/wf-blog-group created
```

The first ArmadaChart is installed:

```bash
kubectl get all

pod/armada-operator-cbbc7d7f7-zxj5n     1/1     Running             0          60s
pod/blog-1-testchart-5dd8c474f4-26574   0/1     ContainerCreating   0          6s
pod/wf-blog-group-1193326311            0/1     Completed           0          8s
pod/wf-blog-group-2026876860            0/1     ContainerCreating   0          1s
pod/wf-blog-group-2432013970            0/1     Completed           0          4s

NAME                       TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
service/armada-operator    ClusterIP   10.98.2.253    <none>        8383/TCP   57s
service/blog-1-testchart   ClusterIP   10.104.240.7   <none>        80/TCP     6s
service/kubernetes         ClusterIP   10.96.0.1      <none>        443/TCP    7m43s

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/armada-operator    1/1     1            1           60s
deployment.apps/blog-1-testchart   0/1     1            0           6s

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/armada-operator-cbbc7d7f7     1         1         1       60s
replicaset.apps/blog-1-testchart-5dd8c474f4   1         1         0       6s
```

Later both charts have been installed

```bash
NAME                                    READY   STATUS      RESTARTS   AGE
pod/armada-operator-cbbc7d7f7-zxj5n     1/1     Running     0          44m
pod/blog-1-testchart-5dd8c474f4-26574   1/1     Running     0          43m
pod/blog-2-testchart-57f86dd9c5-xmhd7   1/1     Running     0          43m
pod/wf-blog-group-1193326311            0/1     Completed   0          43m
pod/wf-blog-group-2026876860            0/1     Completed   0          43m
pod/wf-blog-group-2234690393            0/1     Completed   0          43m
pod/wf-blog-group-2432013970            0/1     Completed   0          43m

NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
service/armada-operator    ClusterIP   10.98.2.253      <none>        8383/TCP   44m
service/blog-1-testchart   ClusterIP   10.104.240.7     <none>        80/TCP     43m
service/blog-2-testchart   ClusterIP   10.110.160.242   <none>        80/TCP     43m
service/kubernetes         ClusterIP   10.96.0.1        <none>        443/TCP    51m

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/armada-operator    1/1     1            1           44m
deployment.apps/blog-1-testchart   1/1     1            1           43m
deployment.apps/blog-2-testchart   1/1     1            1           43m

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/armada-operator-cbbc7d7f7     1         1         1       44m
replicaset.apps/blog-1-testchart-5dd8c474f4   1         1         1       43m
replicaset.apps/blog-2-testchart-57f86dd9c5   1         1         1       43m
```

argo is tracing the steps take to sequence the deployment of the charts:

```bash
argo get wf-blog-group

Name:                wf-blog-group
Namespace:           default
ServiceAccount:      armada-argo-sa
Status:              Succeeded
Created:             Fri Mar 22 10:27:28 -0500 (26 seconds ago)
Started:             Fri Mar 22 10:27:28 -0500 (26 seconds ago)
Finished:            Fri Mar 22 10:27:42 -0500 (12 seconds ago)
Duration:            14 seconds

STEP                  PODNAME                   DURATION  MESSAGE
 ✔ wf-blog-group
 ├---✔ enable-blog-1  wf-blog-group-1193326311  2s
 ├---✔ blog-1-ready   wf-blog-group-2432013970  3s
 ├---✔ enable-blog-2  wf-blog-group-2026876860  3s
 └---✔ blog-2-ready   wf-blog-group-2234690393  3s
 ```

We can check the state of the ArmadaChart

```bash
kubectl get act

NAME     STATE      TARGET STATE   SATISFIED
blog-1   deployed   deployed       true
blog-2   deployed   deployed       true
```

Run the cleanup
```bash

kubectl delete -f examples/argo
```

## examples/sequenced

In that directory, the ArmadaChart (act) are disabled by default and the charts not installed
automatically when the act CR are created.
When the "ArmadaChartGroup" CR is created, the "chartgroup controller" receives an event and it
orchestrate the order of deployment/enablement of the ArmadaChart. The ArmadaChartGroup also
becomes owner of the ArmadaChart. 

This is basically the same sequencing that above except that it is implemented using an
ArmadaChartGroup and an ArmadaManifest

# Testing the oslc-controller

## Change lifecycle manually at first

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
```

Check argo
```bash
NAME                  STATUS      AGE   DURATION
wf-bootstrap          Running     57s   57s
wf-keystone-install   Succeeded   57s   50s
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
kubectl get all

NAME                                               READY   STATUS      RESTARTS   AGE
pod/keystone-api-777d75b84f-2j6pt                  1/1     Running     0          3m53s
pod/keystone-armada-operator-5cfccc74fb-dzfmz      1/1     Running     0          11m
pod/keystone-bootstrap-d9rpw                       0/1     Init:0/1    0          3m53s
pod/keystone-credential-setup-z8b4n                0/1     Completed   1          3m53s
pod/keystone-db-init-kjl7f                         0/1     Completed   0          3m53s
pod/keystone-db-sync-jqjvr                         0/1     Completed   0          3m53s
pod/keystone-domain-manage-kjf5z                   0/1     Completed   0          3m53s
pod/keystone-fernet-setup-kzwxk                    0/1     Completed   1          3m53s
pod/keystone-oslc-operator-747bb84698-ccgxq        1/1     Running     0          11m
pod/keystone-rabbit-init-pf6dq                     0/1     Completed   0          3m53s
pod/mariadb-ingress-7bb7bd65ff-lq799               1/1     Running     0          11m
pod/mariadb-ingress-error-pages-869655cc86-d2xgv   1/1     Running     0          11m
pod/mariadb-server-0                               1/1     Running     0          11m
pod/memcached-memcached-579f87f568-947zm           1/1     Running     0          11m
pod/rabbitmq-cluster-wait-xndv4                    0/1     Completed   0          10m
pod/rabbitmq-rabbitmq-0                            1/1     Running     0          10m
pod/wf-bootstrap-2256133693                        0/2     Error       0          3m
pod/wf-bootstrap-2275099864                        0/2     Completed   0          3m14s
pod/wf-bootstrap-458835584                         0/2     Completed   0          3m7s
pod/wf-bootstrap-636750498                         0/2     Completed   0          3m53s
pod/wf-keystone-install-1483869462                 0/1     Completed   0          3m22s
pod/wf-keystone-install-1938236639                 0/1     Completed   0          3m37s
pod/wf-keystone-install-24327275                   0/2     Completed   0          3m33s
pod/wf-keystone-install-2710739935                 0/2     Completed   0          3m30s
pod/wf-keystone-install-2782100822                 0/1     Completed   0          3m22s
pod/wf-keystone-install-3134557299                 0/1     Completed   0          3m39s
pod/wf-keystone-install-33467967                   0/2     Completed   0          3m53s
pod/wf-keystone-install-3464818                    0/2     Completed   0          3m53s
pod/wf-keystone-install-3672310916                 0/2     Completed   0          3m53s
pod/wf-keystone-install-376633117                  0/2     Completed   0          3m53s
pod/wf-keystone-install-724228482                  0/2     Completed   0          3m53s
pod/wf-keystone-install-823103181                  0/2     Completed   0          3m16s

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                        AGE
service/keystone                      ClusterIP   10.100.51.176    <none>        80/TCP,443/TCP                 3m53s
service/keystone-api                  ClusterIP   10.96.68.68      <none>        5000/TCP                       3m53s
service/keystone-armada-operator      ClusterIP   10.104.191.57    <none>        8383/TCP                       11m
service/keystone-oslc-operator        ClusterIP   10.99.143.77     <none>        8383/TCP                       11m
service/kubernetes                    ClusterIP   10.96.0.1        <none>        443/TCP                        57m
service/mariadb                       ClusterIP   10.100.6.122     <none>        3306/TCP                       11m
service/mariadb-discovery             ClusterIP   None             <none>        3306/TCP,4567/TCP              11m
service/mariadb-ingress-error-pages   ClusterIP   None             <none>        80/TCP                         11m
service/mariadb-server                ClusterIP   10.108.166.206   <none>        3306/TCP                       11m
service/memcached                     ClusterIP   10.104.245.239   <none>        11211/TCP                      11m
service/rabbitmq                      ClusterIP   10.108.133.228   <none>        5672/TCP,25672/TCP,15672/TCP   10m
service/rabbitmq-dsv-7b1733           ClusterIP   None             <none>        5672/TCP,25672/TCP,15672/TCP   10m
service/rabbitmq-mgr-7b1733           ClusterIP   10.101.43.52     <none>        80/TCP,443/TCP                 10m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/keystone-api                  1/1     1            1           3m53s
deployment.apps/keystone-armada-operator      1/1     1            1           11m
deployment.apps/keystone-oslc-operator        1/1     1            1           11m
deployment.apps/mariadb-ingress               1/1     1            1           11m
deployment.apps/mariadb-ingress-error-pages   1/1     1            1           11m
deployment.apps/memcached-memcached           1/1     1            1           11m

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/keystone-api-777d75b84f                  1         1         1       3m53s
replicaset.apps/keystone-armada-operator-5cfccc74fb      1         1         1       11m
replicaset.apps/keystone-oslc-operator-747bb84698        1         1         1       11m
replicaset.apps/mariadb-ingress-7bb7bd65ff               1         1         1       11m
replicaset.apps/mariadb-ingress-error-pages-869655cc86   1         1         1       11m
replicaset.apps/memcached-memcached-579f87f568           1         1         1       11m

NAME                                 READY   AGE
statefulset.apps/mariadb-server      1/1     11m
statefulset.apps/rabbitmq-rabbitmq   1/1     10m

NAME                                  COMPLETIONS   DURATION   AGE
job.batch/keystone-bootstrap          0/1           3m53s      3m53s
job.batch/keystone-credential-setup   1/1           38s        3m53s
job.batch/keystone-db-init            1/1           17s        3m53s
job.batch/keystone-db-sync            1/1           45s        3m53s
job.batch/keystone-domain-manage      1/1           42s        3m53s
job.batch/keystone-fernet-setup       1/1           38s        3m53s
job.batch/keystone-rabbit-init        1/1           16s        3m53s
job.batch/rabbitmq-cluster-wait       1/1           23s        10m

NAME                                       SCHEDULE       SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/keystone-credential-rotate   0 0 1 * *      False     0        <none>          3m53s
cronjob.batch/keystone-fernet-rotate       0 */12 * * *   False     0        <none>          3m53s
```

Check the error in argo

```bash
argo logs wf-bootstrap-2256133693

+ openstack role create --or-show member
An unexpected error prevented the server from fulfilling your request. (HTTP 500) (Request-ID: req-492c9d66-96a1-46f6-99b1-4439fd203362)
```

# Appendix

[POCs](./pocs/README.md) contains additional notes regarding successful and failed attempts.