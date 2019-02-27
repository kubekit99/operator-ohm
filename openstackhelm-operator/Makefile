
# Image URL to use all building/pushing image targets
COMPONENT        ?= openstackhelm-operator
VERSION          ?= 0.0.1
DHUBREPO         ?= kubekit99/${COMPONENT}-dev
DOCKER_NAMESPACE ?= kubekit99
IMG              ?= ${DHUBREPO}:v${VERSION}

all: docker-build

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	# dep ensure

# Run tests
unittest: setup generate fmt vet
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Run go fmt against code
fmt: setup
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet: setup
	go vet ./pkg/... ./cmd/...

# Generate code
generate: setup
	go generate ./pkg/... ./cmd/...
	# go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/
	# operator-sdk generate k8s

# Build the docker image
docker-build: unittest
	operator-sdk build ${IMG}
	docker tag ${IMG} ${DHUBREPO}:latest

# Push the docker image
docker-push:
	docker push ${IMG}

# Run against the configured Kubernetes cluster in ~/.kube/config
install: docker-build
	helm install --name osh-operator chart 

purge: setup
	helm delete --purge osh-operator

createchart:
	kubectl create -f examples/openstackchart-testchart.yaml

deletechart:
	kubectl delete -f examples/openstackchart-testchart.yaml
