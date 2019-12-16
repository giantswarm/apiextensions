package v1alpha1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewGithubSearchIndexCRD returns a new custom resource definition for
// GithubSearchIndex. This might look something like the following:
//
//   apiVersion: apiextensions.k8s.io/v1beta1
//   kind: CustomResourceDefinition
//   metadata:
//   name: searchindex.githubsearch.giantswarm.io
//   spec:
//     group: githubsearch.giantswarm.io
//     scope: Namespaced
//     version: v1alpha1
//     names:
//       kind: SearchIndex
//       plural: searchindices
//       singular: searchindex
//
// An example CR:
//
//  apiVersion: githubsearch.giantswarm.io/v1alpha1
//  kind: SearchIndex
//  metadata:
//      name: giantswarm-giantswarm-indexer
//      labels: github-search-index-operator.giantswarm.io/version: "1.0.0"
//  spec:
//      repository: giantswarm/giantswarm
//      esIndexName: documents
//  status:
//      lastCommitSHA: 1f6baaad653e433e2b6e78bd3fb6c062d8c52679
//      lastCommitTime: 2019-12-16T11:02:14Z
//
func NewGithubSearchIndexCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "searchindex.githubsearch.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "githubsearch.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "SearchIndex",
				Plural:   "searchindices",
				Singular: "searchindex",
			},
		},
	}
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SearchIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              SearchIndexSpec   `json:"spec"`
	Status            SearchIndexStatus `json:"status"`
}

type SearchIndexSpec struct {
	// Repository is the org/repo name combination.
	// For a git URL of https://github.com/giantswarm/giantswarm.git, use "giantswarm/giantswarm".
	Repository string `json:"repository"`

	// EsIndexName is the name of the index to use in Elasticsearch.
	EsIndexName string `json:"esIndexName"`
}

type SearchIndexStatus struct {
	LastCommitSHA  string       `json:"lastCommitSHA"`
	LastCommitTime DeepCopyTime `json:"lastCommitTime"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SearchIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []SearchIndex `json:"items"`
}
