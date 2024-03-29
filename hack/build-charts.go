//go:generate go run .

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/giantswarm/microerror"
	"github.com/google/go-github/v39/github"
	"golang.org/x/oauth2"

	"github.com/giantswarm/apiextensions/v6/pkg/crd"
)

func main() {
	ctx := context.Background()
	httpClient := http.DefaultClient
	if githubToken := os.Getenv("GITHUB_TOKEN"); githubToken != "" {
		token := oauth2.Token{AccessToken: githubToken}
		ts := oauth2.StaticTokenSource(&token)
		httpClient = oauth2.NewClient(ctx, ts)
	}

	renderer := crd.Renderer{
		GithubClient:       github.NewClient(httpClient),
		LocalCRDDirectory:  "../config/crd",
		OutputDirectory:    "../helm",
		Patches:            patches,
		UpstreamAssets:     upstreamReleaseAssets,
		RemoteRepositories: remoteRepositories,
	}

	for _, provider := range []string{"common", "aws", "azure", "kvm"} {
		err := renderer.Render(ctx, provider)
		if err != nil {
			log.Fatal(microerror.JSON(err))
		}
	}
}
