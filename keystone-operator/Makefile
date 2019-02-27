
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
	kubectl create namespace operatorpoc
	kubectl create -f deploy/crds/openstackhelm_v1alpha1_keystone_crd.yaml -n operatorpoc
	kubectl create -f deploy/service_account.yaml -n operatorpoc
	kubectl create -f deploy/role.yaml -n operatorpoc
	kubectl create -f deploy/role_binding.yaml -n operatorpoc
	kubectl create -f deploy/operator.yaml -n operatorpoc
	kubectl get all -n operatorpoc

purge: setup
	kubectl delete -f deploy/operator.yaml -n operatorpoc
	kubectl delete -f deploy/role_binding.yaml -n operatorpoc
	kubectl delete -f deploy/role.yaml -n operatorpoc
	kubectl delete -f deploy/service_account.yaml -n operatorpoc
	kubectl delete -f deploy/crds/openstackhelm_v1alpha1_keystone_crd.yaml -n operatorpoc
	kubectl get all -n operatorpoc
	kubectl delete namespace operatorpoc

createchart:
	kubectl label nodes airship openstack-control-plane=enabled --overwrite
	kubectl create -f deploy/crds/openstackhelm_v1alpha1_keystone_cr.yaml -n operatorpoc
	kubectl get keystones.openstackhelm.openstack.org -n operatorpoc
	kubectl get all -n operatorpoc

describechart:
	kubectl describe keystones/example-keystone -n operatorpoc

deletechart:
	kubectl delete -f deploy/crds/openstackhelm_v1alpha1_keystone_cr.yaml -n operatorpoc
	kubectl get keystones.openstackhelm.openstack.org -n operatorpoc
	kubectl get all -n operatorpoc
