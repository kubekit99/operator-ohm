# THIS REPOSITORY IS OBSOLETE. CONTENT HAS BEEN MIGRATED ONTO [Keleustes](https://github.com/keleustes/) 

# Keystone operator

Armada-Operator and OpenstackServiceLifeCycle-Operator deployment of keystone

## Deploy the Keystone-Operator using Helm

### Description.

That simple goal of this test is to check that helm can still be used to deploy keystone
without Argo or new CRD even if some new files have added to the chart. Basically backward
compatibility check.

### Logs

- [Helm](./docs/classichelm/helm.md)

## Use Argo within the Keystone Chart

### Description

- Argo is used within the chart deployment to sequence the database creation, the seeding of that data.

### Logs

- [First](./docs/argokeystone/first.md)
- [Second](./docs/argokeystone/second.md)
- [Third](./docs/argokeystone/third.md)
- [Purge](./docs/argokeystone/delete1.md)

## Use Argo for LifeCycle of an OpenstackService Chart

### Description

- Argo is used to sequence the lifecycle of an openstackservice. For instance:
  1. Install Service
  2. Test Service
  3. Rollout Traffic

### Logs

- [Install Flow](./docs/oslc/installflow.md)

## Deploy Keystone using ArmadaChart CRD and Argo for sequencing

### Description

- The ArmadaChart mariadb, rabbitmq, memcached and keystone are created as "disabled".
- An Argo Worflow CRD enables those charts following the sequence specified in the ArmadaChartGroup.

### Logs

- [Installation](./docs/argomanifest/argo.md)
- [Purge](./docs/argomanifest/argodelete.md)

## Deploy Keystone using ArmadaChart CRD and ArmadaManifest for sequencing

### Description

- The ArmadaChart mariadb, rabbitmq, memcached and keystone are created as "disabled".
- An ArmadaManifest CRD enables those charts following the sequence specified in the ArmadaChartGroup.

### Logs

- [Sequencing](./docs/armadamanifest/sequenced.md)
