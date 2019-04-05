
# Image URL to use all building/pushing image targets
COMPONENT        ?= memcached-operator
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
	kubectl apply -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl apply -f ../armada-operator/chart/templates/role.yaml
	kubectl apply -f ../armada-operator/chart/templates/service_account.yaml
	kubectl apply -f ../armada-operator/chart/templates/argo_armada_role.yaml
	kubectl create -f deploy/operator.yaml

purge: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl delete -f ../armada-operator/chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl delete -f ../armada-operator/chart/templates/role_binding.yaml
	kubectl delete -f ../armada-operator/chart/templates/role.yaml
	kubectl delete -f ../armada-operator/chart/templates/service_account.yaml
	kubectl delete -f ../armada-operator/chart/templates/argo_armada_role.yaml

installmanifest:
	kubectl label nodes airship openstack-control-plane=enabled --overwrite
	kubectl apply -f examples/memcached/simple.yaml
	kubectl describe act/memcached

describemanifest:
	kubectl describe act/memcached

deletemanifest:
	kubectl delete -f examples/memcached/simple.yaml
	kubectl describe act/memcached

getcrds:
	kubectl get armadacharts.armada.airshipit.org
	kubectl get armadachartgroups.armada.airshipit.org
	kubectl get armadamanifests.armada.airshipit.org

	kubectl get workflows.argoproj.io
