package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cert struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              CertSpec `json:"spec"`
}

type CertSpec struct {
	Cert          CertSpecCert          `json:"cert" yaml:"cert"`
	VersionBundle CertSpecVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type CertSpecCert struct {
	AllowBareDomains bool     `json:"allowBareDomains" yaml:"allowBareDomains"`
	AltNames         []string `json:"altNames" yaml:"altNames"`
	ClusterComponent string   `json:"clusterComponent" yaml:"clusterComponent"`
	ClusterID        string   `json:"clusterID" yaml:"clusterID"`
	CommonName       string   `json:"commonName" yaml:"commonName"`
	IPSANs           []string `json:"ipSans" yaml:"ipSans"`
	Organizations    []string `json:"organizations" yaml:"organizations"`
	TTL              string   `json:"ttl" yaml:"ttl"`
}

type CertSpecVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Cert `json:"items"`
}
