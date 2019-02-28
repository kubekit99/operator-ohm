# Kubernetes Operator for Openstack HELM

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
kubectl create -f examples/armada/manifest.yaml
kubectl describe amf/simple-armada
```

## Test controller reconcilation logic (for depending resources)

Upon deletion of its depending resources, the controller will recreate it,

```bash
kubectl describe amf/simple-armada
```

## Test controller reconcilation logic (for CRD)

When deleting the CRD, the corresponding Armada Manifest should be uninstalled.

```bash
kubectl delete -f examples/armada/manifest.yaml
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
// +kubebuilder:printcolumn:name="toy",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description="descr1",format="date",priority=3
// +kubebuilder:printcolumn:name="abc",type="integer",JSONPath="status",description="descr2",format="int32",priority=1
// +kubebuilder:printcolumn:name="service",type="string",JSONPath=".status.conditions.ready",description="descr3",format="byte",priority=2
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
