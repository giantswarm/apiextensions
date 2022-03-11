package main

import (
	"github.com/giantswarm/apiextensions/v5/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{
	// azure
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-azure",
		Version:  "v1.1.0",
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

var remoteRepositories = []crd.RemoteRepositoryDefinition{
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "apiextensions-application",
		Reference: "v0.3.1",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "silence-operator",
		Reference: "v0.4.0",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "release-operator",
		Reference: "v3.2.0",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "config-controller",
		Reference: "v0.5.1",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "apiextensions-backup",
		Reference: "v0.2.0",
	},
}
