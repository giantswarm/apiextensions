package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CertConfigSpec `json:"spec"`
}

type ChartConfigSpec struct {
	Chart         ChartConfigSpecChart         `json:"chart" yaml:"chart"`
	VersionBundle ChartConfigSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type ChartConfigSpecChart struct {
	Channel string `json:"channel" yaml:"channel"`
	Name    string `json:"name" yaml:"name"`
}

type ChartConfigSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ChartConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []ChartConfig `json:"items"`
}
