package main

import (
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{
	// aws
	{
		Owner:    "giantswarm",
		Repo:     "cluster-api",
		Version:  "v1.1.0-gsalpha.1",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "aws",
	},
	{
		Owner:   "kubernetes-sigs",
		Repo:    "cluster-api-provider-aws",
		Version: "v1.1.0",
		Files: []string{
			"infrastructure-components.yaml",
		},
		Provider: "aws",
	},
	// azure
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.3.22",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "azure",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-azure",
		Version:  "v0.4.15",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "azure",
	},
	{
		Owner:    "Azure",
		Repo:     "aad-pod-identity",
		Version:  "v1.8.0",
		Files:    []string{"deployment.yaml"},
		Provider: "azure",
	},
	// kvm
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.4",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "kvm",
	},
	// openstack
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.4",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "openstack",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-openstack",
		Version:  "v0.4.0",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "openstack",
	},
	// vsphere
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.4",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "vsphere",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-vsphere",
		Version:  "v0.8.1",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "vsphere",
	},
}

var remoteRepositories []crd.RemoteRepositoryDefinition
