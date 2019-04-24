package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AWSProviderSpec is the structure put into the provider spec of the Cluster
// API's Cluster type. There it is tracked as serialized raw extension.
//
//     kind: CustomSpec
//     apiVersion: aws.provider.giantswarm.io/v1beta1
//     metadata:
//       name: 8y5kc
//     cluster:
//       description: my fancy cluster
//       dns:
//         domain: gauss.eu-central-1.aws.gigantic.io
//       oidc:
//         claims:
//           username: email
//           groups: groups
//         clientId: foobar-dex-client
//         issuerUrl: https://dex.8y5kc.fr-east-1.foobar.example.com
//       versionBundle:
//         version: 4.9.0
//     provider:
//       credentialSecret:
//         name: credential-default
//         namespace: giantswarm
//       master:
//         instanceType: m4.large
//       region: eu-central-1
//
type AWSProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Cluster           ClusterSpecCluster  `json:"cluster" yaml:"cluster"`
	Provider          ClusterSpecProvider `json:"provider" yaml:"provider"`
}

type ClusterSpecCluster struct {
	Description   string                          `json:"description" yaml:"description"`
	DNS           ClusterSpecClusterDNS           `json:"dns" yaml:"dns"`
	OIDC          ClusterSpecClusterOIDC          `json:"oidc" yaml:"oidc"`
	VersionBundle ClusterSpecClusterVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type ClusterSpecClusterDNS struct {
	Domain string `json:"domain" yaml:"domain"`
}

type ClusterSpecClusterOIDC struct {
	Claims    ClusterSpecClusterOIDCClaims `json:"claims" yaml:"claims"`
	ClientID  string                       `json:"clientID" yaml:"clientID"`
	IssuerURL string                       `json:"issuerURL" yaml:"issuerURL"`
}

type ClusterSpecClusterOIDCClaims struct {
	Username string `json:"username" yaml:"username"`
	Groups   string `json:"groups" yaml:"groups"`
}

type ClusterSpecClusterVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

type ClusterSpecProvider struct {
	CredentialSecret ClusterSpecProviderCredentialSecret `json:"credentialSecret" yaml:"credentialSecret"`
	Master           ClusterSpecProviderMaster           `json:"master" yaml:"master"`
	Region           string                              `json:"region" yaml:"region"`
}

type ClusterSpecProviderCredentialSecret struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

type ClusterSpecProviderMaster struct {
	InstanceType string `json:"instanceType" yaml:"instanceType"`
}
