
# Image URL to use all building/pushing image targets
COMPONENT        ?= armada-operator
VERSION_V2       ?= 2.13.0
VERSION_V3       ?= 3.0.0
DHUBREPO         ?= kubekit99/${COMPONENT}-dev
DOCKER_NAMESPACE ?= kubekit99
IMG_V2           ?= ${DHUBREPO}:v${VERSION_V2}
IMG_V3           ?= ${DHUBREPO}:v${VERSION_V3}

all: docker-build

setup:
ifndef GOPATH
	$(error GOPATH not defined, please define GOPATH. Run "go help gopath" to learn more about GOPATH)
endif
	# JEB: dep ensure
	# JEB: mkdir -p vendor/helm.sh
	# JEB: cd vendor/helm.sh && git clone -b dev-v3 https://github.com/helm/helm.git

clean:
	rm -fr vendor
	rm -fr cover.out
	rm -fr build/_output
	rm -fr config/crds

# Run tests
unittest: setup fmt vet-v2
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
vet-v2: setup
	go vet -composites=false -tags=v2 ./pkg/... ./cmd/...

vet-v3: setup
	go vet -composites=false -tags=v3 ./pkg/... ./cmd/...

# Generate code
generate: setup
	go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/ --domain airshipit.org --skip-map-validation=false
	go run vendor/k8s.io/code-generator/cmd/deepcopy-gen/main.go --input-dirs github.com/kubekit99/operator-ohm/armada-operator/pkg/apis/armada/v1alpha1 -O zz_generated.deepcopy --bounding-dirs github.com/kubekit99/operator-ohm/armada-operator/pkg/apis

# Build the docker image
docker-build: fmt docker-build-v2

docker-build-v2: vet-v2
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/_output/bin/armada-operator -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} -tags=v2 ./cmd/...
	docker build . -f build/Dockerfile -t ${IMG_V2}
	docker tag ${IMG_V2} ${DHUBREPO}:latest

docker-build-v3: vet-v3
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/_output/bin/armada-operator -gcflags all=-trimpath=${GOPATH} -asmflags all=-trimpath=${GOPATH} -tags=v3 ./cmd/...
	docker build . -f build/Dockerfile -t ${IMG_V3}
	docker tag ${IMG_V3} ${DHUBREPO}:latest


# Push the docker image
docker-push: docker-push-v2

docker-push-v2:
	docker push ${IMG_V2}

docker-push-v3:
	docker push ${IMG_V3}

# Run against the configured Kubernetes cluster in ~/.kube/config
install: install-v2

purge: setup
	kubectl delete act --all
	kubectl delete acg --all
	kubectl delete amf --all
	helm delete --purge armada-operator

install-v2: docker-build-v2
	helm install --name armada-operator chart --set images.tags.operator=${IMG_V2}

install-v3: docker-build-v3
	helm install --name armada-operator chart --set images.tags.operator=${IMG_V3}

# Deploy and purge procedure which do not rely on helm itself
install-kubectl: docker-build
	kubectl apply -f ./chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl apply -f ./chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl apply -f ./chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl apply -f ./chart/templates/role_binding.yaml
	kubectl apply -f ./chart/templates/role.yaml
	kubectl apply -f ./chart/templates/service_account.yaml
	kubectl apply -f ./chart/templates/argo_armada_role.yaml
	kubectl create -f deploy/operator.yaml

purge-kubectl: setup
	kubectl delete -f deploy/operator.yaml
	kubectl delete -f ./chart/templates/armada_v1alpha1_armadachartgroup.yaml
	kubectl delete -f ./chart/templates/armada_v1alpha1_armadachart.yaml
	kubectl delete -f ./chart/templates/armada_v1alpha1_armadamanifest.yaml
	kubectl delete -f ./chart/templates/role_binding.yaml
	kubectl delete -f ./chart/templates/role.yaml
	kubectl delete -f ./chart/templates/service_account.yaml
	kubectl delete -f ./chart/templates/argo_armada_role.yaml

getcrds:
	kubectl get armadacharts.armada.airshipit.org
	kubectl get armadachartgroups.armada.airshipit.org
	kubectl get armadamanifests.armada.airshipit.org

	kubectl get workflows.argoproj.io
