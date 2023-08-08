[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/apiextensions/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/apiextensions/tree/master)

# apiextensions

This library provides generated Kubernetes clients for the Giant Swarm infrastructure.

## Usage

- [`pkg/apis`](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis?tab=doc): Contains data structures for
    custom resources in `*.giantswarm.io` API groups. See full documentation
    [here](https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis?tab=doc).

## Contributing

### Setup

While the generation scripts can run without a GitHub token defined, you may encounter rate limit errors without one. The
token will be loaded from the `GITHUB_TOKEN` environment variable if it exists. Giant Swarm engineers can generally use
`export GITHUB_TOKEN=$OPSCTL_GITHUB_TOKEN` to configure this before running `make`.

### Changing Existing Custom Resources

- Make the desired changes in `pkg/apis/<group>/<version>`.
- Update generated files by calling `make generate`.
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

#### Upstream CRDs

`apiextensions` generates Giant Swarm CRDs based on code in the `pkg/apis` directory and also aggregates CRDs from
various sources. Any repository that publishes CRDs as a YAML-formatted manifest of CRDs attached as an asset to a
GitHub release ("upstream" CRDs) or in a source tree ("remote" CRDs) can be used. The current set of CRDs
including target version is defined in `hack/assets.go` in the `upstreamReleaseAssets` and `remoteRepositories` variables
for upstream and remote CRDs respectively. Update the values and re-run `make` to update the aggregated CRDs.

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
- `generate-deepcopy`: Generates `zz_generated.deepcopy.go` in each package in `pkg/apis` with deep copy functions.
- `generate-manifests`: Generates CRDs in `config/crd` from CRs found in `pkg/apis`.
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

#### [`controller-gen`](https://book.kubebuilder.io/reference/controller-gen.html)

Generates a custom resource definition (CRD) for each custom resource using special comments such as
`// +kubebuilder:validation:Optional`.

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

The `documentation` is later on parsed as Markdown syntax.
Annotations that don't have this YAML format are ignored and not published to the public docs.
Annotations belong to specific API Version of the CRDs but they can be linked to multiple CRDs and multiple API Versions.
