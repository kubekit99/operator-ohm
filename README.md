# operator-ohm
Operator POC

## Goal

Goal is to compare the behavior/usefullness of helm based operators vs Helm 3 based CRDs.

## Keystone operator 
```bash
operator-sdk new keystone-operator --api-version=openstackhelm.openstack.org/v1alpha1 --kind=Keystone --type=helm --skip-git-init
```

## MariaDB operator

```bash
operator-sdk new mariadb-operator --api-version=openstackhelm.openstack.org/v1alpha1 --kind=Mariadb --type=helm --skip-git-init
```
