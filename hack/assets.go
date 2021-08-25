package main

import (
	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.4.2",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "common",
	},
	{
		Owner:   "kubernetes-sigs",
		Repo:    "cluster-api-provider-aws",
		Version: "v0.6.5",
		Files: []string{
			"eks-bootstrap-components.yaml",
			"eks-controlplane-components.yaml",
			"infrastructure-components.yaml",
		},
		Provider: "aws",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-azure",
		Version:  "v0.5.2",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "azure",
	},
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api-provider-vsphere",
		Version:  "v0.7.6",
		Files:    []string{"infrastructure-components.yaml"},
		Provider: "vmware",
	},
	{
		Owner:    "Azure",
		Repo:     "aad-pod-identity",
		Version:  "v1.8.0",
		Files:    []string{"deployment.yaml"},
		Provider: "azure",
	},
}
