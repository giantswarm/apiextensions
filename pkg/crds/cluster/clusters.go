package cluster

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const clustersYAML = `apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
  name: clusters.cluster.x-k8s.io
spec:
  conversion:
    strategy: None
  group: cluster.x-k8s.io
  names:
    kind: Cluster
    listKind: ClusterList
    plural: clusters
    singular: cluster
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha2
    served: true
    storage: true
    subresources:
      status: {}
`

func NewClusterCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(clustersYAML), &crd)
	return &crd
}
