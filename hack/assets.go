package main

import (
	"github.com/giantswarm/apiextensions/v5/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{}

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
