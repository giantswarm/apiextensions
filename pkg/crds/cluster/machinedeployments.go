package cluster

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"
)

const machinedeploymentsYAML = `apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  creationTimestamp: null
  labels:
    controller-tools.k8s.io: "1.0"
  name: machinedeployments.cluster.x-k8s.io
spec:
  conversion:
    strategy: None
  group: cluster.x-k8s.io
  names:
    kind: MachineDeployment
    listKind: MachineDeploymentList
    plural: machinedeployments
    singular: machinedeployment
  preserveUnknownFields: true
  scope: Namespaced
  versions:
  - name: v1alpha2
    served: true
    storage: true
    subresources:
      status: {}
`

func NewMachineDeploymentCRD() *v1beta1.CustomResourceDefinition {
	var crd v1beta1.CustomResourceDefinition
	_ = yaml.Unmarshal([]byte(machinedeploymentsYAML), &crd)
	return &crd
}
