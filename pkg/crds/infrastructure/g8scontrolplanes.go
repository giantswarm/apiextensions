package infrastructure

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const g8scontrolplanesYAML = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: (devel)
  creationTimestamp: null
  name: g8scontrolplanes.infrastructure.giantswarm.io
spec:
  group: infrastructure.giantswarm.io
  names:
    kind: G8sControlPlane
    listKind: G8sControlPlaneList
    plural: g8scontrolplanes
    singular: g8scontrolplane
  scope: Namespaced
  validation:
    openAPIV3Schema:
      description: "G8sControlPlane defines the Control Plane Nodes (Kubernetes Master
        Nodes) of a Giant Swarm Tenant Cluster \n     apiVersion: infrastructure.giantswarm.io/v1alpha2
        \    kind: G8sControlPlane     metadata:       annotations:         giantswarm.io/docs:
        https://docs.giantswarm.io/reference/g8scontrolplanes.infrastructure.giantswarm.io/v1alpha2/
        \      labels:         aws-operator.giantswarm.io/version: \"6.2.0\"         cluster-operator.giantswarm.io/version:
        \"0.17.0\"         giantswarm.io/cluster: 8y5kc         giantswarm.io/organization:
        giantswarm         release.giantswarm.io/version: \"7.3.1\"       name: 8y5kc
        \    spec:       infrastructureRef:         apiVersion: infrastructure.giantswarm.io/v1alpha2
        \        kind: AWSControlPlane         name: 5f3kb         namespace: default
        \      replicas: 3     status:       readyReplicas: 3       replicas: 3"
      properties:
        apiVersion:
          description: 'APIVersion defines the versioned schema of this representation
            of an object. Servers should convert recognized schemas to the latest
            internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
          type: string
        kind:
          description: 'Kind is a string value representing the REST resource this
            object represents. Servers may infer this from the endpoint the client
            submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
          type: string
        metadata:
          type: object
        spec:
          properties:
            infrastructureRef:
              description: InfrastructureRef is a required reference to provider-specific
                Infrastructure.
              properties:
                apiVersion:
                  description: API version of the referent.
                  type: string
                fieldPath:
                  description: 'If referring to a piece of an object instead of an
                    entire object, this string should contain a valid JSON/Go field
                    access statement, such as desiredState.manifest.containers[2].
                    For example, if the object reference is to a container within
                    a pod, this would take on a value like: "spec.containers{name}"
                    (where "name" refers to the name of the container that triggered
                    the event) or if no container name is specified "spec.containers[2]"
                    (container with index 2 in this pod). This syntax is chosen only
                    to have some well-defined way of referencing a part of an object.
                    TODO: this design is not final and this field is subject to change
                    in the future.'
                  type: string
                kind:
                  description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
                resourceVersion:
                  description: 'Specific resourceVersion to which this reference is
                    made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                  type: string
                uid:
                  description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                  type: string
              type: object
            replicas:
              description: Replicas is the number replicas of the master node.
              type: integer
          required:
          - infrastructureRef
          - replicas
          type: object
        status:
          description: G8sControlPlaneStatus defines the observed state of G8sControlPlane.
          properties:
            readyReplicas:
              description: Total number of fully running and ready control plane machines.
              format: int32
              type: integer
            replicas:
              description: Total number of non-terminated machines targeted by this
                control plane (their labels match the selector).
              format: int32
              type: integer
          type: object
      required:
      - spec
      - status
      type: object
  version: v1alpha2
  versions:
  - name: v1alpha2
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

func NewG8sControlPlaneCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(g8scontrolplanesYAML), &crd)
	return &crd
}
