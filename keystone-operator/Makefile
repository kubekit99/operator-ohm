
# Image URL to use all building/pushing image targets
COMPONENT_A      ?= keystone-armada-operator
COMPONENT_O      ?= keystone-oslc-operator
VERSION          ?= poc
DHUBREPO_A       ?= kubekit99/${COMPONENT_A}
DHUBREPO_O       ?= kubekit99/${COMPONENT_O}
DOCKER_NAMESPACE ?= kubekit99
IMG_A            ?= ${DHUBREPO_A}:${VERSION}
IMG_O            ?= ${DHUBREPO_O}:${VERSION}

all: docker-build

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif

clean: setup
	rm -fr helm-charts/mariadb/
	rm -fr helm-charts/memcached/
	rm -fr helm-charts/rabbitmq/
	rm -fr helm-charts/helm-toolkit/

# Build the docker image
# Temporary kludge to add the depending charts to the current image
docker-build: setup
	tar -C helm-charts/ -xvf helm-charts/keystone/charts/helm-toolkit-0.1.0.tgz
	rsync -avz ../mariadb-operator/helm-charts/mariadb helm-charts
	rsync -avz ../memcached-operator/helm-charts/memcached helm-charts
	rsync -avz ../rabbitmq-operator/helm-charts/rabbitmq helm-charts
	docker build -t ${IMG_A} -f build/Dockerfile.armada-operator .
	docker tag ${IMG_A} ${DHUBREPO_A}:latest
	docker build -t ${IMG_O} -f build/Dockerfile.oslc-operator .
	docker tag ${IMG_O} ${DHUBREPO_O}:latest

# Push the docker image
docker-push:
	docker push ${IMG_A}
	docker push ${DHUBREPO_A}:latest
	docker push ${IMG_O}
	docker push ${DHUBREPO_O}:latest

# Run against the configured Kubernetes cluster in ~/.kube/config
install: docker-build install-oslc

install-armada:
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl apply -f ../armada-operator/chart/templates/role.yaml
	kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
	kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
	kubectl create -f deploy/armada-operator.yaml

install-oslc:
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/role_binding.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/roles.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/service_account.yaml
	kubectl apply -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml
	kubectl create -f deploy/oslc-operator.yaml

purge: setup purge-oslc

purge-armada:
	kubectl delete -f deploy/armada-operator.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl delete -f ../armada-operator/chart/templates/role.yaml
	kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
	kubectl delete -f ../armada-operator/chart/templates/argo_armada_role.yaml

purge-oslc:
	kubectl delete -f deploy/oslc-operator.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_deletephase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_installphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_operationalphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_planningphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_rollbackphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_testphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficdrainphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_trafficrolloutphase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_upgradephase.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/openstacklcm_v1alpha1_oslc.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/role_binding.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/roles.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/service_account.yaml
	kubectl delete -f ../openstacklcm-operator/chart/templates/argo_openstacklcm_role.yaml


installmanifest:
	kubectl label nodes airship openstack-control-plane=enabled --overwrite
	kubectl apply -f examples/keystone/simple_local.yaml
	kubectl describe amf/armada-manifest

describemanifest:
	kubectl describe amf/armada-manifest
	kubectl describe acg/keystone-infra-services
	kubectl describe acg/openstack-keystone
	kubectl describe act/helm-toolkit
	kubectl describe act/keystone
	kubectl describe act/mariadb
	kubectl describe act/memcached
	kubectl describe act/rabbitmq

deletemanifest:
	kubectl delete -f examples/keystone/simple_local.yaml
	kubectl delete configmap mariadb-mariadb-mariadb-ingress
	kubectl delete configmap mariadb-mariadb-state
	kubectl describe amf/armada-manifest 

getcrds:
	kubectl get workflows.argoproj.io

	kubectl get armadacharts.armada.airshipit.org
	kubectl get armadachartgroups.armada.airshipit.org
	kubectl get armadamanifests.armada.airshipit.org

	kubectl get oslcs.openstacklcm.airshipit.org

	kubectl get installphases.openstacklcm.airshipit.org
	kubectl get rollbackphases.openstacklcm.airshipit.org
	kubectl get testphases.openstacklcm.airshipit.org
	kubectl get trafficdrainphases.openstacklcm.airshipit.org
	kubectl get trafficrolloutphases.openstacklcm.airshipit.org
	kubectl get upgradephases.openstacklcm.airshipit.org
	kubectl get deletephases.openstacklcm.airshipit.org
	kubectl get planningphases.openstacklcm.airshipit.org
	kubectl get operationalphases.openstacklcm.airshipit.org
