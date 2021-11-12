package crd

import (
	"github.com/google/go-github/v39/github"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Patch func(crd *v1.CustomResourceDefinition)

type Renderer struct {
	GithubClient *github.Client

	LocalCRDDirectory string
	OutputDirectory   string

	Patches map[string]Patch

	UpstreamAssets     []ReleaseAssetFileDefinition
	RemoteRepositories []RemoteRepositoryDefinition
}

type ReleaseAssetFileDefinition struct {
	Files    []string
	Owner    string
	Provider string
	Repo     string
	Version  string
}

type RemoteRepositoryDefinition struct {
	Path      string // path to CRDs in repo, usually "config/crd"
	Owner     string // repo owner
	Provider  string // provider the CRD should be added to in helm charts
	Name      string // name of repo
	Reference string // tag or branch of repo to list files from
}
