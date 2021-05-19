//go:generate go run .

package main

import (
	"context"
	"log"
	"os"

	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

func main() {
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		githubToken = os.Getenv("GIANTSWARM_GITHUB_TOKEN")
	}

	ctx := context.Background()
	token := oauth2.Token{AccessToken: githubToken}
	ts := oauth2.StaticTokenSource(&token)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	renderer := crd.Renderer{
		GithubClient:      client,
		LocalCRDDirectory: "../config/crd",
		OutputDirectory:   "../helm",
		Patches:           patches,
		UpstreamAssets:    upstreamReleaseAssets,
	}

	for _, provider := range []string{"common", "aws", "azure", "kvm", "vmware"} {
		err := renderer.Render(ctx, provider)
		if err != nil {
			log.Fatal(err)
		}
	}
}
