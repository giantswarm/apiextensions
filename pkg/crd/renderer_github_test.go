//go:build github

package crd

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-github/v39/github"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func buildTestRenderer() Renderer {
	ctx := context.Background()
	httpClient := http.DefaultClient
	if githubToken := os.Getenv("GITHUB_TOKEN"); githubToken != "" {
		token := oauth2.Token{AccessToken: githubToken}
		ts := oauth2.StaticTokenSource(&token)
		httpClient = oauth2.NewClient(ctx, ts)
	}

	return Renderer{
		GithubClient: github.NewClient(httpClient),
	}
}

func Test_downloadReleaseAssetCRDs(t *testing.T) {
	renderer := buildTestRenderer()
	asset := ReleaseAssetFileDefinition{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.3.22",
		Files:    []string{"cluster-api-components.yaml"},
		Provider: "aws",
	}
	crds, err := renderer.downloadReleaseAssetCRDs(context.Background(), asset)
	require.Nil(t, err, err)
	require.Len(t, crds, 11)
}

func Test_downloadRepositoryCRDs(t *testing.T) {
	renderer := buildTestRenderer()
	crds, err := renderer.downloadRepositoryCRDs(context.Background(), RemoteRepositoryDefinition{
		Path:     "config/crd",
		Owner:    "giantswarm",
		Provider: "common",
		Repo:     "apiextensions",
		Version:  "v3.35.0",
	})
	require.Nil(t, err, err)
	require.Len(t, crds, 30)
}
