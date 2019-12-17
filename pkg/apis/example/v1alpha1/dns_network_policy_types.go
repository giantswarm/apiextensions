package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewDNSNetworkPolicyCRD returns a new custom resource definition for
// DNSNetworkPolicy. This might look something like the following.
//
//	apiVersion: apiextensions.k8s.io/v1beta1
//	kind: CustomResourceDefinition
//	metadata:
//	  name: dnsnetworkpolicies.example.giantswarm.io
//	spec:
//	  group: example.giantswarm.io
//	  scope: Namespaced
//	  version: v1alpha1
//	  names:
//	    kind: DNSNetworkPolicy
//	    plural: dnsnetworkpolicies
//	    singular: dnsnetworkpolicy
//    subresources:
//      status: {}
//
func NewDNSNetworkPolicyCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "dnsnetworkpolicies.example.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "example.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "DNSNetworkPolicy",
				Plural:   "dnsnetworkpolicies",
				Singular: "dnsnetworkpolicy",
			},
			Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
				Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DNSNetworkPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DNSNetworkPolicySpec   `json:"spec" yaml:"status"`
	Status            DNSNetworkPolicyStatus `json:"status" yaml:"status"`
}

type DNSNetworkPolicySpec struct {
	// Domains is the list of domain names, which should be resolved
	// before updating reference network policy.
	// e.g. ["kubernetes.local", "http://google.com", "https://twitter.com"]
	Domains []string `json:"domains" yaml:"domains"`
	// targetNetworkPolicy is an existing network policy in the object namespace,
	// which is updated with resolved domains IP addresses.
	// e.g. memcached-network-policy
	TargetNetworkPolicy string `json:"targetNetworkPolicy" yaml:"targetNetworkPolicy"`
}

// DNSNetworkPolicyStatus holds information about last successful update of the target
// network policy.
// Fields here are automatically filled and can only ever be read.
type DNSNetworkPolicyStatus struct {
	// LastTargetNetworkPolicyUpdateTime is the last time targetNetworkPolicy was updated.
	LastTargetNetworkPolicyUpdateTime DeepCopyTime `json:"lastTargetNetworkPolicyUpdateTime" yaml:"lastTargetNetworkPolicyUpdateTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DNSNetworkPolicyList godoc.
type DNSNetworkPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []DNSNetworkPolicy `json:"items"`
}
