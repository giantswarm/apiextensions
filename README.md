[![CircleCI](https://circleci.com/gh/giantswarm/apiextensions.svg?&style=shield)](https://circleci.com/gh/giantswarm/apiextensions)

# apiextensions

This library provides generated Kubernetes clients for the Giant Swarm infrastructure.

## Usage

- [`pkg/apis`](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis?tab=doc): Contains data structures for
    custom resources in `*.giantswarm.io` API groups. See full documentation
    [here](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis?tab=doc).
- [`pkg/clientset/versioned`](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/clientset/versioned?tab=doc):
    Contains a clientset, a client for each custom resource, and a fake client for unit testing. See full documentation
    [here](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/clientset/versioned?tab=doc).
- [`pkg/crd`](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/crd?tab=doc): Contains an interface for
    accessing individual CRDs or listing all CRDs. See full documentation
    [here](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/crd?tab=doc).

## Contributing

### Changing Existing Custom Resources

- Make the desired changes in `pkg/apis/<group>/<version>`
- Update generated files by calling `make`.
    - `make generate && goimports -local github.com/giantswarm/apiextensions -w ./pkg`
- Review and commit all changes including generated code.

#### Naming Convention

Custom resource structs are placed in packages corresponding to the endpoints in
Kubernetes API. For example, structs in package
`github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2` are created
from objects under `/apis/infrastructure.giantswarm.io/v1alpha2/` endpoint.

As this is common to have name collisions between field type names in different
custom objects sharing the same group and version (e.g. `Spec` or `Status`) we prefix all type names
referenced inside custom object with the name of the parent object.

Example:

```go
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NewObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NewObjSpec `json:"spec"`
}

type NewObjSpec struct {
	Field string `json:"field"`
}
```

### Adding a New Custom Resource

This is example skeleton for adding new object.

- Make sure group and version of the object to add exists (described in
  [previous paragraph](#adding-a-new-group-andor-version)).
- Replace `NewObj` with your object name.
- Put struct definitions inside a proper package denoted by group and version
  in a file named `new_obj_types.go`. Replace `new_obj` with snake_case-formatted object name.
- Add `NewObj` and `NewObjList` to `knownTypes` slice in `register.go`
- Generate code for the resource by calling `make`.
- Commit changes and create a release.

```go
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NewObj godoc.
type NewObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NewObjSpec `json:"spec"`
}

// NewObjSpec godoc.
type NewObjSpec struct {
	FieldName string `json:"fieldName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NewObjList godoc.
type NewObjList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NewObj `json:"items"`
}
```

### Adding a New Group and/or Version

This is example skeleton for adding new group and/or version.

- Replace `GROUP` with new group name and `VERSION` with new version name.
- Create a new package `/pkg/apis/GROUP/VERSION/`.
- Inside the package create a file `doc.go` (content below).
- Inside the package create a file `register.go` (content below).
- Add a new object (described in [next paragraph](#adding-a-new-custom-object)).

Example `doc.go` content.

```go
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register

// +groupName=GROUP.giantswarm.io
package VERSION
```

Example `register.go` content.

```go
package VERSION

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	group   = "GROUP.giantswarm.io"
	version = "VERSION"
)

// knownTypes is the full list of objects to register with the scheme. It
// should contain pointers of zero values of all custom objects and custom
// object lists in the group version.
var knownTypes = []runtime.Object{
		//&Object{},
		//&ObjectList{},
}

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   group,
	Version: version,
}

var (
	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme is used by the generated client.
	AddToScheme = schemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, knownTypes...)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
```

### Updating dependencies

#### Cluster API

Cluster API CRDs are also exported by this library using `controller-gen`. The version used is determined by the value
of `sigs.k8s.io/cluster-api` in `hack/go.mod`.

#### Code Generation Tools

To change the version of a tool, edit the version manually in `hack/tools/<tool>/go.mod` and run `go mod tidy` in
that directory so that `go.sum` is updated.

### Versioning

This library uses standard semantic versioning. Versioning of CRDs is a separate issue covered in the [Kubernetes
    deprecation policy](https://kubernetes.io/docs/reference/using-api/deprecation-policy/).
    In short, if an API field needs to be removed, a new version must be created, and any versions served concurrently
    must be convertible in both directions without data loss.

## Code generation

This library uses code generation to generate several components so that they do not need to be maintained manually
and to reduce the chance of mistakes such as when, for example, defining an OpenAPIV3 schema for a CRD.

### Makefile

The `Makefile` at the root of the repository ensures that required tools (defined below) are installed in
`hack/tools/bin` and then runs each step of the code generation pipeline sequentially.

The main code generation steps are as follows:
- `generate-clientset`: Generates the clientset for accessing custom resources in a Kubernetes cluster.
- `generate-deepcopy`: Generates `zz_generated.deepcopy.go` in each package in `pkg/apis` with deep copy functions.
- `generate-manifests`: Generates CRDs in `config/crd/v1` and `config/crd/v1beta1` from CRs found in `pkg/apis`.
- `generate-fs`: Generates `pkg/crd/internal` package containing a filesystem holding all files in `config/crd`.
- `imports`: Sorts imports in all source files under `./pkg`.
- `patch`: Applies the git patch `hack/generated.patch` to work around limitations in code generators.

These can all be run with `make generate` or simply `make` as `generate` is the default rule.

Extra commands are provided including:
- `clean-tools`: Deletes all tools from the tools binary directory.
- `clean-generated`: Deletes all generated files.
- `verify`: Regenerates files and exits with a non-zero exit code if generated files don't match `HEAD` in source control.

### Tools

Tools are third-party executables which perform a particular action as part of the code generation pipeline. They are
defined in `hack/tools` in separate directories. Versions for the tools are defined in the `go.mod` file in their
respective directories. A common `go.mod` isn't used so that their dependencies don't interfere.

#### [`deepcopy-gen`](https://godoc.org/k8s.io/code-generator/cmd/deepcopy-gen)

Generates `DeepCopy` and `DeepCopyInto` functions for all custom resources to satisfy the `runtime.Object` interface.

#### [`client-gen`](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-api-machinery/generating-clientset.md)

Generates a "client set" which provides CRUD interfaces for each custom resource.

#### [`esc`](https://github.com/mjibson/esc)

Encodes local filesystem trees into a Go source file containing an `http.FileSystem` which provides access to the
files at runtime. This allows these files to be accessed from a binary outside of the source tree containing those files.

#### [`controller-gen`](https://book.kubebuilder.io/reference/controller-gen.html)

Generates a custom resource definition (CRD) for each custom resource using special comments such as
`// +kubebuilder:validation:Optional`.

#### [`kustomize`](https://github.com/kubernetes-sigs/kustomize)

Provides an extra patch step for generated CRD YAML files because certain CRD fields can't be modified with
`controller-gen` directly.

#### [`goimports`](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

Updates Go import lines by adding missing ones and removing unreferenced ones. This is required because CI checks for
imports ordered in three sections (standard library, third-party, local) but certain code generators only generate
source files with two sections.

### Annotation Documentation

Docs generator has support for documenting annotations. This is achieved by creating code annotations in YAML on the definition of the variables like:

```
// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha2
//     release: Since 14.0.0
// documentation:
//   Here is the documentation related to this annotation
```

Check [AWS Annotations](pkg/annotation/aws.go) for examples.

Annotations that don't have this YAML format are ignored and not published to the public docs.
Annotations belong to specific API Version of the CRDs but they can be linked to multiple CRDs and multiple API Versions.
