#!/bin/bash
kubectl create namespace operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_mariadb_crd.yaml -n operatorpoc
kubectl create -f deploy/service_account.yaml -n operatorpoc
kubectl create -f deploy/role.yaml -n operatorpoc
kubectl create -f deploy/role_binding.yaml -n operatorpoc
kubectl create -f deploy/operator.yaml -n operatorpoc
kubectl create -f deploy/crds/openstackhelm_v1alpha1_mariadb_cr.yaml -n operatorpoc
