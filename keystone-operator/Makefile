
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

# Build the docker image
docker-build: setup
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
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
	kubectl apply -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
	kubectl apply -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
	kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl apply -f ../armada-operator/chart/templates/role.yaml
	kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
	kubectl create -f deploy/operator.yaml

purge: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadarequest.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrelease.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_helmrequest.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycleevent.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_lifecycle.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_manifest.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_releaseaudit.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_release.yaml
	kubectl delete -f ../armada-operator/chart/templates/helm3crd_v1beta1_values.yaml
	kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl delete -f ../armada-operator/chart/templates/role.yaml
	kubectl delete -f ../armada-operator/chart/templates/service_account.yaml

installmanifest:
	kubectl label nodes airship openstack-control-plane=enabled --overwrite
	kubectl apply -f example/keystone/simple.yaml
	kubectl describe amf/keystone

describemanifest:
	kubectl describe act/keystone
	kubectl describe acg/keystone
	kubectl describe amf/keystone

deletemanifest:
	kubectl delete -f example/keystone/simple.yaml
	kubectl describe amf/keystone 
