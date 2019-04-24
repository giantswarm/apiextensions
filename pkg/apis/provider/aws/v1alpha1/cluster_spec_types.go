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
	Cluster           AWSProviderSpecCluster  `json:"cluster" yaml:"cluster"`
	Provider          AWSProviderSpecProvider `json:"provider" yaml:"provider"`
}

type AWSProviderSpecCluster struct {
	Description   string                              `json:"description" yaml:"description"`
	DNS           AWSProviderSpecClusterDNS           `json:"dns" yaml:"dns"`
	OIDC          AWSProviderSpecClusterOIDC          `json:"oidc" yaml:"oidc"`
	VersionBundle AWSProviderSpecClusterVersionBundle `json:"versionBundle" yaml:"versionBundle"`
}

type AWSProviderSpecClusterDNS struct {
	Domain string `json:"domain" yaml:"domain"`
}

type AWSProviderSpecClusterOIDC struct {
	Claims    AWSProviderSpecClusterOIDCClaims `json:"claims" yaml:"claims"`
	ClientID  string                           `json:"clientID" yaml:"clientID"`
	IssuerURL string                           `json:"issuerURL" yaml:"issuerURL"`
}

type AWSProviderSpecClusterOIDCClaims struct {
	Username string `json:"username" yaml:"username"`
	Groups   string `json:"groups" yaml:"groups"`
}

type AWSProviderSpecClusterVersionBundle struct {
	Version string `json:"version" yaml:"version"`
}

type AWSProviderSpecProvider struct {
	CredentialSecret AWSProviderSpecProviderCredentialSecret `json:"credentialSecret" yaml:"credentialSecret"`
	Master           AWSProviderSpecProviderMaster           `json:"master" yaml:"master"`
	Region           string                                  `json:"region" yaml:"region"`
}

type AWSProviderSpecProviderCredentialSecret struct {
	Name      string `json:"name" yaml:"name"`
	Namespace string `json:"namespace" yaml:"namespace"`
}

type AWSProviderSpecProviderMaster struct {
	InstanceType string `json:"instanceType" yaml:"instanceType"`
}
