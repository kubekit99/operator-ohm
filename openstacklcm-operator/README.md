# Kubernetes Operator for Openstack LCM

# Introduction

This README is mainly used a log / wiki of results and challenges encounted during the POC

## operator-sdk vs kubebuilder.

Things to clean up:
Did not had time to sort how to completly migrate from operator-sdh to kubebuilder.
Most of the scaffolding is done using the operator-sdk but once the default files are created,
the build process mainly relies on kubebuilder

## openstacklcm-operator code directory structure

###  cmd

Contains the main.go for the openstacklcm operator

###  pkg/apis/

Contains the golang definition of the CRDs. `make generate` will recreate the yaml definitions
of the CRDs that have to be provided to kubectl in order to deploy the new CRDs.

The first version of the golang code has generated using tool such as "schema-generate" from
the schema definition provided with airship-armada project.

###  pkg/services

Contains the bulk of the interfaces used by the openstacklcm controller.

###  pkg/osphases

Mainly contain the code for xxxPhase handling

###  pkg/oslc directory

Mainly contain the code for Oslc (OpenstackServiceLifeCycle) handling

###  pkg/controller directory

Contains the controller and the "Reconcile" functions for Oslc and xxxPhase.

# Code changes.

## Adjusting the OpenstackServiceLifeCycle CRDs

Upon change of the CRD golang definition, the yaml files have to be regenerated

Note 1: Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

Note 2: The generation tool seems to comply with some of OpenAPI specs. The "validation" schema added
to in the CRD yaml definition does not contain fields using underscore. 
Most of those fields containing underscore where defined such a way in the original airship-armada.

```bash
make generate
```

## Compiling the openstacklcm-operator

```bash
dep ensure
```

To build the version
```bash
make docker-build
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

To install the version
```bash
make install
```

# Openstack Service Invidual Phase CRD testing

For testing purpose the current Docker file includes a dummy chart deliverd under armada-charts.

## Invidual Phase CRD sanity tests

```bash
kubectl apply -f examples/phases/

deletephase.openstacklcm.airshipit.org/delete created
ginstallphase.openstacklcm.airshipit.org/install created
operationalphase.openstacklcm.airshipit.org/operational created
planningphase.openstacklcm.airshipit.org/planning created
rollbackphase.openstacklcm.airshipit.org/rollback created
testphase.openstacklcm.airshipit.org/test created
trafficdrainphase.openstacklcm.airshipit.org/trafficdrain created
trafficrolloutphase.openstacklcm.airshipit.org/trafficrollout created
upgradephase.openstacklcm.airshipit.org/upgrade created
```

```bash
kubectl describe osupg
kubectl describe osrbck
kubectl describe osroll
kubectl describe osdrain
kubectl describe ostest
kubectl describe osplan
kubectl describe osins
kubectl describe osdlt
```


```bash
for i in `cat phaselist.txt`; do kubectl logs pod/${i}-wf main; done
 ____________________
< workflow 1: delete >
 --------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 _____________________
< workflow 1: install >
 ---------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 _________________________
< workflow 1: operational >
 -------------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 ______________________
< workflow 1: planning >
 ----------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 ______________________
< workflow 1: rollback >
 ----------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 __________________
< workflow 1: test >
 ------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 __________________________
< workflow 1: trafficdrain >
 --------------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 ____________________________
< workflow 1: trafficrollout >
 ----------------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
 _____________________
< workflow 1: upgrade >
 ---------------------
    \
     \
      \
                    ##        .
              ## ## ##       ==
           ## ## ## ##      ===
       /""""""""""""""""___/ ===
  ~~~ {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~
       \______ o          __/
        \    \        __/
          \____\______/
```

```bash
kubectl get all

NAME                                         READY   STATUS      RESTARTS   AGE
pod/delete-wf                                0/2     Completed   0          18s
pod/install-wf                               0/2     Completed   0          18s
pod/openstacklcm-operator-6745fc85f4-b5f76   1/1     Running     0          4m46s
pod/operational-wf                           0/2     Completed   0          18s
pod/planning-wf                              0/2     Completed   0          17s
pod/rollback-wf                              0/2     Completed   0          16s
pod/test-wf                                  0/2     Completed   0          15s
pod/trafficdrain-wf                          0/2     Completed   0          14s
pod/trafficrollout-wf                        0/2     Completed   0          14s
pod/upgrade-wf                               0/2     Completed   0          14s

NAME                            TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE
service/kubernetes              ClusterIP   10.96.0.1       <none>        443/TCP    34m
service/openstacklcm-operator   ClusterIP   10.101.236.32   <none>        8383/TCP   33m

NAME                                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/openstacklcm-operator   1/1     1            1           4m46s

NAME                                               DESIRED   CURRENT   READY   AGE
replicaset.apps/openstacklcm-operator-6745fc85f4   1         1         1       4m46s
```

```bash
kubectl delete -f examples/phases/

deletephase.openstacklcm.airshipit.org "delete" deleted
installphase.openstacklcm.airshipit.org "install" deleted
operationalphase.openstacklcm.airshipit.org "operational" deleted
planningphase.openstacklcm.airshipit.org "planning" deleted
rollbackphase.openstacklcm.airshipit.org "rollback" deleted
testphase.openstacklcm.airshipit.org "test" deleted
trafficdrainphase.openstacklcm.airshipit.org "trafficdrain" deleted
trafficrolloutphase.openstacklcm.airshipit.org "trafficrollout" deleted
upgradephase.openstacklcm.airshipit.org "upgrade" deleted
```
