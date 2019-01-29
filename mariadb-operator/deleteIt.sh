#!/bin/bash
kubectl delete -f deploy/crds/openstackhelm_v1alpha1_mariadb_cr.yaml -n operatorpoc
kubectl delete -f deploy/operator.yaml -n operatorpoc
kubectl delete -f deploy/role_binding.yaml -n operatorpoc
kubectl delete -f deploy/role.yaml -n operatorpoc
kubectl delete -f deploy/service_account.yaml -n operatorpoc
kubectl delete -f deploy/crds/openstackhelm_v1alpha1_mariadb_crd.yaml -n operatorpoc
kubectl delete namespace operatorpoc
