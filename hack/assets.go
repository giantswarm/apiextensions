package main

import (
	"github.com/giantswarm/apiextensions/v6/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{}

var remoteRepositories = []crd.RemoteRepositoryDefinition{
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "apiextensions-application",
		Reference: "v0.6.0",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "silence-operator",
		Reference: "v0.6.1",
	},
	{
		Path:      "config/crd",
		Owner:     "giantswarm",
		Provider:  "common",
		Name:      "release-operator",
		Reference: "v4.0.0",
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
		Name:      "organization-operator",
		Reference: "v1.0.0",
	},
}
