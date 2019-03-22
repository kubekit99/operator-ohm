
# Image URL to use all building/pushing image targets
COMPONENT        ?= openstacklcm-operator
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
	echo -e 'docker stop $$(docker ps -qa)'
	echo -e 'export PATH=$${PATH}:/usr/local/kubebuilder/bin'
	mkdir -p config/crds
	cp chart/templates/*v1alpha1* config/crds/
	go test ./pkg/... ./cmd/... -coverprofile cover.out

# Run go fmt against code
fmt: setup
	go fmt ./pkg/... ./cmd/...

# Run go vet against code
vet: setup
	go vet -composites=false -tags=v2 ./pkg/... ./cmd/...

# Generate code
generate: setup
	# JEB: go generate ./pkg/... ./cmd/...
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/ --domain airshipit.org
	operator-sdk generate k8s

# Build the docker image
docker-build: fmt vet
	# JEB: operator-sdk build ${IMG}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/_output/bin/openstacklcm-operator -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} -tags=v2 ./cmd/...
	docker build . -f build/Dockerfile -t ${IMG}
	docker tag ${IMG} ${DHUBREPO}:latest

# Push the docker image
docker-push:
	docker push ${IMG}

# Run against the configured Kubernetes cluster in ~/.kube/config
install: install

purge: setup
	helm delete --purge openstacklcm-operator

install: docker-build
	helm install --name openstacklcm-operator chart --set images.tags.operator=${IMG}

# Deploy and purge procedure which do not rely on helm itself
install-kubectl: docker-build
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstackbackup.yaml
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstackdeployment.yaml
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstackrestore.yaml
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstackrollback.yaml
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstackupgrade.yaml
	kubectl apply -f ./chart/templates/openstacklcm_v1alpha1_openstacksvc.yaml
	kubectl apply -f ./chart/templates/role_binding.yaml
	kubectl apply -f ./chart/templates/role.yaml
	kubectl apply -f ./chart/templates/service_account.yaml
	kubectl apply -f ./chart/templates/argo_openstacklcm_role.yaml
	kubectl create -f deploy/operator.yaml

purge-kubectl: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstackbackup.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstackdeployment.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstackrestore.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstackrollback.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstackupgrade.yaml
	kubectl delete -f ./chart/templates/openstacklcm_v1alpha1_openstacksvc.yaml
	kubectl delete -f ./chart/templates/role_binding.yaml
	kubectl delete -f ./chart/templates/role.yaml
	kubectl delete -f ./chart/templates/service_account.yaml
	kubectl delete -f ./chart/templates/argo_openstacklcm_role.yaml
