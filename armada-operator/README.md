# Kubernetes Operator for Armada and Helm

# Introduction

This README is mainly used a log / wiki of results and challenges encounted during the POC

## operator-sdk vs kubebuilder.

Things to clean up:
Did not had time to sort how to completly migrate from operator-sdh to kubebuilder.
Most of the scaffolding is done using the operator-sdk but once the default files are created,
the build process mainly relies on kubebuilder

## Tiller/Helm directories

###  helm v2 vs helm v3

To get the process doing, some of the code for HelmRelease handling is coming from the 
operator-sdk helm-operator. That code relies on tiller component which is gone for Helm3.
Hence the three directory helmif (Interface and Common code), helmv2 (tiller) and helmv3.

### HelmRelease CRDs.

The HelmRelease represent temporalily used until the HelmV3 Release CRD is available.

### HelmRequest CRD

The helm-v3 proposal appendix describes in Appendix the creation of the HelmRequest CRD to
work with a controller and the Helm code provided as a standalone library.
- Note sure the interest of the "Get"....operations since the same result can be achieve
using "kubectl get releases"

### Notes

- JEB: For testing purpose the current Docker file includes a dummy chart deliverd under armada-charts.
This removes the needs to access external chart repository which is also an aspect of helm changing from 2 to 3.

## Armada directories

### Manifest, ChartGroup and Chart CRDs.

Armada CRD are currently inspired from from the structure used by airship-armada.

### ArmadaRequest CRDs.

As for Helm/Tiller, the ArmadaRequest CRD are supposed to represent the "command" oriented operations.
Not sure of the 

### Notes

- JEB: Since we have multiple CRDs (Manifest, CharGroup...Chart), still have to figure out if
we need multiple controller/reconciler each listening only on one CRD or one controller watching multiple CRDs.

## Helm3CRD directories

### Release, Values, Chart... CRDs.

Those CRD and Controller have been build using the helm-3-crd discussion and repository.

### Notes

- JEB: Can't figure out the impact of having thousands of individual LifecycleEvent CR on the performance
and usability of the kubectl.

# Creation of armada-operator

## Initialising the armada-operator

```bash
operator-sdk new armada-operator --skip-git-init
```

## Coding the armada-operator

```bash
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=HelmRelease
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaManifest
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaChart
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaChartGroup
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=HelmRequest
operator-sdk add api --api-version=armada.airshipit.org/v1alpha1 --kind=ArmadaRequest
operator-sdk add api --api-version=helm3crd.airshipit.org/v1beta1 --kind=Release
operator-sdk add api --api-version=helm3crd.airshipit.org/v1beta1 --kind=Values
operator-sdk add api --api-version=helm3crd.airshipit.org/v1beta1 --kind=Manifest
operator-sdk add api --api-version=helm3crd.airshipit.org/v1beta1 --kind=Lifecycle
operator-sdk add api --api-version=helm3crd.airshipit.org/v1beta1 --kind=LifecycleEvent
git add pkg/apis/armada/
git add pkg/apis/helm3crd/
git add pkg/apis/addtoscheme_armada_v1alpha1.go
git add deploy/role.yaml
```

```bash
vi pkg/apis/armada/v1alpha1/*_types.go
operator-sdk generate k8s
```

```bash
operator-sdk add controller --api-version=armada.airshipit.org/v1alpha1 --kind=Tiller
operator-sdk add controller --api-version=armada.airshipit.org/v1alpha1 --kind=Armada
operator-sdk add controller --api-version=helm3crd.airshipit.org/v1beta1 --kind=Helm3CRD
```
## Adjusting crds

Don't understand yet how to build using operator-sdk operator with the same level of detailes than
controller-gen. Big hack that have to be included in Makefile.

```bash
go run vendor/sigs.k8s.io/controller-tools/cmd/controller-gen/main.go crd --output-dir ./chart/templates/
operator-sdk generate k8s
```

or

```bash
make generate
```

## Compiling the armada-operator

```
make docker-build
```

# Deploying

## Deployment of operator using helm

```bash
helm install --name armada-operator chart 
```

# Simple Helm Chart CRD and Controller

## Trigger helm chart deployment

Upon creation of the custom resource, the controller will
- Deploy the Helm Chart described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/tiller/helm-testchart.yaml
kubectl describe hrel/testchart
```

## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe hrel/testchart
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding helm chart should be uninstalled.

```bash
kubectl delete -f examples/tiller/helm-testchart.yaml
```

# Simple ArmadaManifest and Controller

## Trigger ArmadaManifest deployment

Upon creation of the custom resource, the controller will
- Deploy the Armada Manifest described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/armada/simple.yaml
kubectl describe amf/simple-armada
kubectl get amf
kubectl get acg
kubectl get act
```



## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe amf/simple-armada
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding Armada Manifest should be uninstalled.

```bash
kubectl delete act/blog-1
kubectl delete acg/blog-group
kubectl delete amf/simple-armada
kubectl delete -f examples/armada/simple.yaml
```


# Simple Helm3CRD Release and Controller

## Trigger helm chart deployment

Upon creation of the custom resource, the controller will
- Deploy the Helm Chart described in the CRD
- Update status of the custom resources.
- Add events to the custom resources.

```bash
kubectl create -f examples/helm3crd/release.yaml
kubectl describe rel/my-release
```

## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe rel/my-release
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding helm chart should be uninstalled.

```bash
kubectl delete -f examples/helm3crd/release.yaml
```

# Appendix

## Helm hooks for CRDs

Example of annotations hook for a CRD

```yaml
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crontabs.stable.example.com
  annotations:
    "helm.sh/hook": crd-install
spec:
  group: stable.example.com
  version: v1
  scope: Namespaced
  names:
    plural: crontabs
    singular: crontab
    kind: CronTab
    shortNames:
    - ct
```

## Add columns to status

in register.go 

```go
// Package v1alpha1 contains API Schema definitions for the fun v1alpha1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=sigs.k8s.io/controller-tools/pkg/crd/generator/testData/pkg/apis/fun
// +k8s:defaulter-gen=TypeMeta
// +groupName=fun.myk8s.io
```
```go
// ToySpec defines the desired state of Toy
type ToySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:ExclusiveMinimum=true
	Power float32 `json:"power,omitempty"`

	Bricks int32 `json:"bricks,omitempty"`

	// +kubebuilder:validation:MaxLength=15
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty"`

	// This is a comment on an array field.
	// +kubebuilder:validation:MaxItems=500
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=false
	Knights []string `json:"knights,omitempty"`

	// This is a comment on a boolean field.
	Winner bool `json:"winner,omitempty"`

	// +kubebuilder:validation:Enum=Lion,Wolf,Dragon
	Alias string `json:"alias,omitempty"`

	// +kubebuilder:validation:Enum=1,2,3
	Rank int `json:"rank"`

	Comment []byte `json:"comment,omitempty"`

	// This is a comment on an object field.
	Template v1.PodTemplateSpec `json:"template"`

	Claim v1.PersistentVolumeClaim `json:"claim,omitempty"`

	//This is a dummy comment.
	// Just checking if the multi-line comments are working or not.
	Replicas *int32 `json:"replicas"`

	// This is a newly added field.
	// Using this for testing purpose.
	Rook *intstr.IntOrString `json:"rook"`

	// This is a comment on a map field.
	Location map[string]string `json:"location"`
}

// ToyStatus defines the observed state of Toy
type ToyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// It tracks the number of replicas.
	Replicas int32 `json:"replicas"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Toy is the Schema for the toys API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:subresource:scale:specpath=.spec.replicas,statuspath=.status.replicas,selectorpath=
// +kubebuilder:printcolumn:name="toy",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description="descr1",format="date"
// +kubebuilder:printcolumn:name="abc",type="integer",JSONPath="status",description="descr2",format="int32"
// +kubebuilder:printcolumn:name="service",type="string",JSONPath=".status.conditions.ready",description="descr3",format="byte"
// +kubebuilder:resource:path=services,shortName=ty
type Toy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ToySpec   `json:"spec,omitempty"`
	Status ToyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ToyList contains a list of Toy
type ToyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Toy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Toy{}, &ToyList{})
}
```
