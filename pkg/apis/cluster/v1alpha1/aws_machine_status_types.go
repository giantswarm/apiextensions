package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSMachineStatus is the structure put into the provider status of the Cluster
// API's Machine type. There it is tracked as serialized raw extension.
//
//     kind: AWSMachineStatus
//     apiVersion: cluster.giantswarm.io/v1alpha1
//     metadata:
//       labels:
//         "giantswarm.io/cluster": "8y5kc"
//         "giantswarm.io/node-pool": "al9qy"
//       name: p36xn
//     machine:
//       versions:
//       - lastTransitionTime: "2019-03-25T17:10:09.995948706Z"
//         version: 4.9.0
//
type AWSMachineStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Machine           AWSMachineStatusMachine `json:"machine" yaml:"machine"`
}

type AWSMachineStatusMachine struct {
	Versions []AWSMachineStatusMachineVersion `json:"versions" yaml:"versions"`
}

type AWSMachineStatusMachineVersion struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Version            string       `json:"version" yaml:"version"`
}
