[![CircleCI](https://circleci.com/gh/giantswarm/apiextensions.svg?&style=shield&circle-token=880450a6e0265218c2b1f8540e280599500bb1a6)](https://circleci.com/gh/giantswarm/apiextensions)

# apiextensions

Package apiextensions provides generated Kubernetes clients for the Giant Swarm
infrastructure.

## Contributing

### Adding a New Custom Object

This is example skeleton for adding new object.

- Replace `NewObj` with your object name.
- Put struct definitions inside a proper package denoted by group and version
  in file named `newobj_types.go`. Replace `newobj` with lowercased object
  name.
- If you create a new group or version edit the last argument of
  `generate-groups.sh` call inside `./scripts/gen.sh`. It has format
  `g1:v5 g2:v1 g3:v1`.
- Generate client by calling `./scripts/gen.sh`.
- Commit generated code and all edits to `./scripts/gen.sh`.

```go
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NewObj godoc.
type NewObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.NewObjMeta `json:"metadata"`
	Spec              NewObjSpec `json:"spec"`
}

// NewObjSpec godoc.
type NewObjSpec struct {
	FieldName string `json:"fieldName", yaml:"fieldName"`
}

// ...
// ...
// ...

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NewObjList godoc.
type NewObjList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []NewObj `json:"items"`
}
```

### Naming Convention

Custom object structs are placed in packages corresponding to the endpoints in
Kubernetes API. E.g. structs in package
`github.com/giantswarm/apiextensions/pkg/apis/cluster/v1alpha1` are created
from objects under `/apis/cluster.giantswarm.io/v1alpha1/` endpoint.

As this is common to have name collisions between field type names in different
custom objects sharing the same group and version we prefix all type names
referenced inside custom object with custom object name.

Example:

```go
type NewObj struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              NewObjSpec `json:"spec"`
}

type NewObjSpec struct {
	Cluster       NewObjCluster       `json:"cluster" yaml:"cluster"`
	VersionBundle NewObjVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type NewObjCluster struct {
	Calico       NewObjCalico       `json:"calico" yaml:"calico"`
	DockerDaemon NewObjDockerDaemon `json:"dockerDaemon" yaml:"dockerDaemon"`
}
```

