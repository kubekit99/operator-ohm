
# Image URL to use all building/pushing image targets
COMPONENT        ?= armada-operator
VERSION          ?= 0.0.1
DHUBREPO         ?= kubekit99/${COMPONENT}-dev
DOCKER_NAMESPACE ?= kubekit99
IMG              ?= ${DHUBREPO}:v${VERSION}

all: docker-build

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	# JEB: dep ensure

clean:
	rm -fr vendor
	rm -fr cover.out
	rm -fr build/_output
	rm -fr config/crds

# Run tests
unittest: setup fmt vet
	echo "sudo systemctl stop kubelet"
	echo "docker stop $(docker ps -qa)"
	echo "export PATH=$PATH:/usr/local/kubebuilder/bin"
	mkdir -p config/crds
	cp chart/templates/*v1alpha1* config/crds/
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Run go fmt against code
fmt: setup
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet: setup
	go vet -composites=false ./pkg/... ./cmd/...

# Generate code
generate: setup
	# JEB: go generate ./pkg/... ./cmd/...
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/ --domain airshipit.org
	operator-sdk generate k8s

# Build the docker image
docker-build: fmt vet
	# JEB: operator-sdk build ${IMG}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/_output/bin/armada-operator -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} ./cmd/...
	docker build . -f build/Dockerfile -t ${IMG}
	docker tag ${IMG} ${DHUBREPO}:latest

# Push the docker image
docker-push:
	docker push ${IMG}

# Run against the configured Kubernetes cluster in ~/.kube/config
install: docker-build
	helm install --name armada-operator chart 

purge: setup
	helm delete --purge armada-operator

installhelmchart:
	kubectl create -f examples/tiller/helm-testchart.yaml

deletehelmchart:
	kubectl delete -f examples/tiller/helm-testchart.yaml

installhelm3crdrelease:
	kubectl create -f examples/helm3crd/release.yaml

deletehelm3crdrelease:
	kubectl delete -f examples/helm3crd/release.yaml
