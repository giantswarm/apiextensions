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
//   name: githubsearchindices.example.giantswarm.io
//   spec:
//     group: example.giantswarm.io
//     scope: Namespaced
//     version: v1alpha1
//     names:
//       kind: GitHubSearchIndex
//       plural: githubsearchindices
//       singular: githubsearchindex
//
// An example CR:
//
//  apiVersion: example.giantswarm.io/v1alpha1
//  kind: GitHubSearchIndex
//  metadata:
//      name: giantswarm-giantswarm-indexer
//      labels: github-search-index-operator.giantswarm.io/version: "1.0.0"
//  spec:
//      repository: https://github.com/giantswarm/giantswarm.git
//      esIndexName: documents
//  status:
//      lastCommitSHA: 1f6baaad653e433e2b6e78bd3fb6c062d8c52679
//      lastCommitTime: 2019-12-16T11:02:14Z
//      indexerJobStarted: 2019-12-18T12:00:00Z
//      indexerJobName: indexer-giantswarm-giantswarm
//
func NewGithubSearchIndexCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: apiextensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "githubsearchindices.example.giantswarm.io",
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   "example.giantswarm.io",
			Scope:   "Namespaced",
			Version: "v1alpha1",
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Kind:     "GitHubSearchIndex",
				Plural:   "githubsearchindices",
				Singular: "githubsearchindex",
			},
		},
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GitHubSearchIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              GitHubSearchIndexSpec   `json:"spec"`
	Status            GitHubSearchIndexStatus `json:"status"`
}

type GitHubSearchIndexSpec struct {
	// Repository is the full repository URL, including .git suffix.
	// Example: "https://github.com/giantswarm/giantswarm.git"
	Repository string `json:"repository"`

	// EsIndexName is the name of the index to use in Elasticsearch.
	EsIndexName string `json:"esIndexName"`
}

type GitHubSearchIndexStatus struct {
	LastCommitSHA     string       `json:"lastCommitSHA"`
	LastCommitTime    DeepCopyTime `json:"lastCommitTime"`
	IndexerJobStarted DeepCopyTime `json:"indexerJobStarted"`
	IndexerJobName    string       `json:"indexerJobName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GitHubSearchIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []GitHubSearchIndex `json:"items"`
}
