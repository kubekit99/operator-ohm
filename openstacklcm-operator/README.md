# Kubernetes Operator for Openstack LCM

# Creation of openstacklcm-operator

## Initialising the openstacklcm-operator

```bash
operator-sdk new openstacklcm-operator --skip-git-init
```

## Coding the openstacklcm-operator

```bash
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add api --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackDeployment
git add deploy/crds/
git add pkg/apis/openstacklcm/
git add pkg/apis/addtoscheme_openstacklcm_v1alpha1.go
git add deploy/role.yaml
```

```bash
vi pkg/apis/openstacklcm/v1alpha1/*_types.go
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackBackup
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRestore
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackUpgrade
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackRollback
operator-sdk add controller --api-version=openstacklcm.airshipit.org/v1alpha1 --kind=OpenstackDeployment
```
## Adjusting crds

Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

```bash
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/
operator-sdk generate k8s
```

or

```bash
make generate
```

## Compiling the openstacklcm-operator

```
make docker-build
```

# Deploying

## Deployment of operator using helm

```bash
helm install --name lcm-operator chart 
```

## Openstack Service Phase CRD testing


```bash
kubectl apply -f examples/upgrade/
kubectl describe osupg
```

```bash
kubectl apply -f examples/rollback
kubectl describe osrbck
```

```bash
kubectl apply -f examples/trafficrollout
kubectl describe osroll
```

```bash
kubectl apply -f examples/trafficdrain
kubectl describe osdrain
```

```bash
kubectl apply -f examples/test
kubectl describe ostest
```

```bash
kubectl apply -f examples/operational
kubectl describe osplan
```

```bash
kubectl apply -f examples/install
kubectl describe osins
```

```bash
kubectl apply -f examples/delete
kubectl describe osdlt
```

# Simple sanity tests

```bash
for i in `cat phaselist.txt`; do kubectl apply -f examples/$i; done

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
for i in `cat phaselist.txt`; do kubectl delete -f examples/$i; done
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
