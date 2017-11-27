[![CircleCI](https://circleci.com/gh/giantswarm/apiextensions.svg?&style=shield&circle-token=880450a6e0265218c2b1f8540e280599500bb1a6)](https://circleci.com/gh/giantswarm/apiextensions)

# apiextensions

Package apiextensions provides generated Kubernetes clients for the Giant Swarm
infrastructure.

## Contributing

### Adding a New Custom Object

This is example skeleton for adding new object.

- Replace `Object` with your object name.
- If you create a new group or version edit the last argument of
  `generate-groups.sh` call inside `./scripts/gen.sh`. It has format
  `g1:v5 g2:v1 g3:v1`.
- Generate client by calling `./scripts/gen.sh`.
- Commit generated code and all edits to `./scripts/gen.sh`.

```go
// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Object godoc.
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ObjectSpec `json:"spec"`
}

// ObjectSpec godoc.
type ObjectSpec struct {
	FieldName string `json:"fieldName", yaml:"fieldName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ObjectList godoc.
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Object `json:"items"`
}
```

### Names Conventions

Custom object structs are placed in packages corresponding to the endpoints in
Kubernetes API. E.g. structs in package
`github.com/giantswarm/apiextensions/pkg/apis/cluster/v1alpha1` are created
from objects under `/apis/cluster.giantswarm.io/v1alpha1/` endpoint.

As this is common to have name collisions between field type names in different
custom objects sharing the same group and version we prefix all type names
referenced inside custom object with custom object name.

Example:

```go
type AWS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AWSSpec `json:"spec"`
}

type AWSSpec struct {
	Cluster       AWSCluster       `json:"cluster" yaml:"cluster"`
	VersionBundle AWSVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type AWSCluster struct {
	Calico       AWSCalico       `json:"calico" yaml:"calico"`
	DockerDaemon AWSDockerDaemon `json:"dockerDaemon" yaml:"dockerDaemon"`
}
```

