
# Image URL to use all building/pushing image targets
COMPONENT        ?= keystone-operator
VERSION          ?= poc
DHUBREPO         ?= kubekit99/${COMPONENT}
DOCKER_NAMESPACE ?= kubekit99
IMG              ?= ${DHUBREPO}:${VERSION}

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
	docker build -t ${IMG} -f build/Dockerfile .
	docker tag ${IMG} ${DHUBREPO}:latest

# Push the docker image
docker-push:
	docker push ${IMG}
	docker push ${DHUBREPO}:latest

# Run against the configured Kubernetes cluster in ~/.kube/config
install: docker-build
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl apply -f ../armada-operator/chart/templates/role.yaml
	kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
	kubectl create -f deploy/operator.yaml

purge: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl delete -f ../armada-operator/chart/templates/role.yaml
	kubectl delete -f ../armada-operator/chart/templates/service_account.yaml

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

