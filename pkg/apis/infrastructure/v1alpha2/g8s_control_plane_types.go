package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

const (
	kindG8sControlPlane              = "G8sControlPlane"
	g8sControlPlaneDocumentationLink = "https://pkg.go.dev/github.com/giantswarm/apiextensions@v0.2.5/pkg/apis/infrastructure/v1alpha2?tab=doc#G8sControlPlane"
)

const g8sControlPlaneCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: g8scontrolplanes.infrastructure.giantswarm.io
spec:
  conversion:
    strategy: None
  group: infrastructure.giantswarm.io
  names:
    kind: G8sControlPlane
    plural: g8scontrolplanes
    singular: g8scontrolplane
  scope: Namespaced
  subresources:
    status: {}
  versions:
    - name: v1alpha1
      served: false
      storage: false
    - name: v1alpha2
      served: true
      storage: true
      schema:
        openAPIV3Schema:
          description: |
            This resource represents an abstract control plane (Kubernetes master node and Etcd server)
            of a tenant cluster in a Giant Swarm installation. It is reconciled by cluster-operator.
            For implementation details it references a provider specific resource.
          type: object
          properties:
            spec:
              type: object
              properties:
                infrastructureRef:
                  description: |
                    Reference to an [AWSControlPlane](https://docs.giantswarm.io/reference/cp-k8s-api/awscontrolplanes.infrastructure.giantswarm.io/)
                    resource defining provider specific details for the c
                  type: object
                  properties:
                    apiVersion:
                      type: string
                    kind:
                      type: string
                    name:
                      type: string
                    namespace:
                      type: string
                replicas:
                  description: |
                    Number of master nodes and Etcd instances to set up.
                  type: integer
                  enum:
                    - 1
                    - 3
`

var g8sControlPlaneCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(g8sControlPlaneCRDYAML), &g8sControlPlaneCRD)
	if err != nil {
		panic(err)
	}
}

func NewG8sControlPlaneCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return g8sControlPlaneCRD.DeepCopy()
}

func NewG8sControlPlaneTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindG8sControlPlane,
	}
}

// NewG8sControlPlaneCR returns a G8sControlPlane Custom Resource.
func NewG8sControlPlaneCR() *G8sControlPlane {
	return &G8sControlPlane{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				crDocsAnnotation: g8sControlPlaneDocumentationLink,
			},
		},
		TypeMeta: NewClusterTypeMeta(),
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// G8sControlPlane defines the Control Plane Nodes (Kubernetes Master Nodes) of
// a Giant Swarm Tenant Cluster.
type G8sControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              G8sControlPlaneSpec   `json:"spec"`
	Status            G8sControlPlaneStatus `json:"status"`
}

type G8sControlPlaneSpec struct {
	// Replicas is the number replicas of the master node.
	Replicas int `json:"replicas" yaml:"replicas"`
	// InfrastructureRef is a required reference to provider-specific
	// Infrastructure.
	InfrastructureRef corev1.ObjectReference `json:"infrastructureRef"`
}

// G8sControlPlaneStatus defines the observed state of G8sControlPlane.
type G8sControlPlaneStatus struct {
	// Total number of non-terminated machines targeted by this control plane
	// (their labels match the selector).
	// +optional
	Replicas int32 `json:"replicas,omitempty"`
	// Total number of fully running and ready control plane machines.
	// +optional
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type G8sControlPlaneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []G8sControlPlane `json:"items"`
}
