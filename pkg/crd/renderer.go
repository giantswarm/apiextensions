package crd

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/google/go-github/v39/github"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// Kubernetes API group, version and kind for v1 CRDs
	crdV1GVK = schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	}
)

func (r Renderer) patchCRDs(provider string, crds []v1.CustomResourceDefinition) ([]v1.CustomResourceDefinition, error) {
	patchedCRDs := make([]v1.CustomResourceDefinition, 0, len(crds))
	for _, crd := range crds {
		patched, err := patchCRD(provider, r.Patches, crd)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		patchedCRDs = append(patchedCRDs, patched)
	}
	return patchedCRDs, nil
}

// Render creates helm chart templates for the given provider by downloading upstream CRDs, merging them with local
// CRDs, patching them, and writing them to the corresponding provider helm template directory.
func (r Renderer) Render(ctx context.Context, provider string) error {
	outputDirectory := filepath.Join(helmChartDirectory(r.OutputDirectory, provider), "templates")

	{
		localCRDs, err := r.getLocalCRDs(provider)
		if err != nil {
			return microerror.Mask(err)
		}

		remoteCRDs, err := r.getRemoteCRDs(ctx, provider)
		if err != nil {
			return microerror.Mask(err)
		}

		internalCRDs, err := r.patchCRDs(provider, append(localCRDs, remoteCRDs...))
		if err != nil {
			return microerror.Mask(err)
		}

		err = writeCRDsToDirectory(outputDirectory, internalCRDs)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	{
		upstreamCRDs, err := r.getUpstreamCRDs(ctx, provider)
		if err != nil {
			return microerror.Mask(err)
		}

		upstreamCRDs, err = r.patchCRDs(provider, upstreamCRDs)
		if err != nil {
			return microerror.Mask(err)
		}

		err = writeCRDsToDirectory(outputDirectory, upstreamCRDs)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

// downloadReleaseAssetCRDs returns a slice of CRDs by downloading the given GitHub release asset, parsing it as YAML,
// and filtering for only CRD objects.
func (r Renderer) downloadReleaseAssetCRDs(ctx context.Context, asset ReleaseAssetFileDefinition) ([]v1.CustomResourceDefinition, error) {
	release, _, err := r.GithubClient.Repositories.GetReleaseByTag(ctx, asset.Owner, asset.Repo, asset.Version)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var targetAssets []*github.ReleaseAsset
	for _, releaseAsset := range release.Assets {
		for _, file := range asset.Files {
			if releaseAsset.GetName() == file {
				targetAssets = append(targetAssets, releaseAsset)
			}
		}
	}
	if targetAssets == nil {
		return nil, notFoundError
	}

	var allCrds []v1.CustomResourceDefinition
	for _, targetAsset := range targetAssets {
		contentReader, _, err := r.GithubClient.Repositories.DownloadReleaseAsset(ctx, asset.Owner, asset.Repo, targetAsset.GetID(), http.DefaultClient)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		crds, err := decodeCRDs(contentReader)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		allCrds = append(allCrds, crds...)
	}

	return allCrds, nil
}

// getUpstreamCRDs returns all upstream CRDs for a provider based on the Renderer's upstream asset configuration.
func (r Renderer) getUpstreamCRDs(ctx context.Context, provider string) ([]v1.CustomResourceDefinition, error) {
	var crds []v1.CustomResourceDefinition
	for _, releaseAsset := range r.UpstreamAssets {
		if releaseAsset.Provider != provider {
			continue
		}

		releaseAssetCRDs, err := r.downloadReleaseAssetCRDs(ctx, releaseAsset)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		crds = append(crds, releaseAssetCRDs...)
	}

	return crds, nil
}

// getLocalCRDs reads the configured local directory and returns a slice of CRDs that have the given category.
func (r Renderer) getLocalCRDs(category string) ([]v1.CustomResourceDefinition, error) {
	var crds []v1.CustomResourceDefinition
	err := filepath.WalkDir(r.LocalCRDDirectory, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return microerror.Mask(walkErr)
		}
		if entry.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return microerror.Mask(err)
		}

		fileCRDs, err := decodeCRDs(file)
		if err != nil {
			return microerror.Mask(err)
		}

		for _, crd := range fileCRDs {
			if contains(crd.Spec.Names.Categories, category) {
				crds = append(crds, crd)
			}
		}

		return nil
	})
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return crds, nil
}

// getRemoteCRDs returns all remote CRDs for a provider based on the Renderer's remote repository configuration.
func (r Renderer) getRemoteCRDs(ctx context.Context, provider string) ([]v1.CustomResourceDefinition, error) {
	var crds []v1.CustomResourceDefinition
	for _, releaseAsset := range r.RemoteRepositories {
		if releaseAsset.Provider != provider {
			continue
		}

		remoteCRDs, err := r.downloadRepositoryCRDs(ctx, releaseAsset)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		crds = append(crds, remoteCRDs...)
	}

	return crds, nil
}

// downloadRepositoryCRDs returns a slice of CRDs by downloading the given GitHub repository tree, listing files in the
// given path, parsing them as YAML, and filtering for only CRD objects.
func (r Renderer) downloadRepositoryCRDs(ctx context.Context, repo RemoteRepositoryDefinition) ([]v1.CustomResourceDefinition, error) {
	refString := fmt.Sprintf("tags/%s", repo.Reference)
	ref, response, err := r.GithubClient.Git.GetRef(ctx, repo.Owner, repo.Name, refString)
	if err != nil && response.StatusCode == 404 {
		refString = fmt.Sprintf("heads/%s", repo.Reference)
		ref, _, err = r.GithubClient.Git.GetRef(ctx, repo.Owner, repo.Name, refString)
	}
	if err != nil {
		return nil, microerror.Mask(err)
	}

	commit, _, err := r.GithubClient.Git.GetCommit(ctx, repo.Owner, repo.Name, ref.Object.GetSHA())
	if err != nil {
		return nil, microerror.Mask(err)
	}

	tree, _, err := r.GithubClient.Git.GetTree(ctx, repo.Owner, repo.Name, commit.Tree.GetSHA(), true)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var targetEntries []*github.TreeEntry
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" && strings.HasPrefix(entry.GetPath(), repo.Path) {
			targetEntries = append(targetEntries, entry)
		}
	}
	if targetEntries == nil {
		return nil, notFoundError
	}

	var allCrds []v1.CustomResourceDefinition
	for _, entry := range targetEntries {
		blob, _, err := r.GithubClient.Git.GetBlob(ctx, repo.Owner, repo.Name, entry.GetSHA())
		if err != nil {
			return nil, microerror.Mask(err)
		}

		content, err := base64.StdEncoding.DecodeString(blob.GetContent())
		if err != nil {
			return nil, microerror.Mask(err)
		}

		contentReader := io.NopCloser(bytes.NewReader(content))

		crds, err := decodeCRDs(contentReader)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		allCrds = append(allCrds, crds...)
	}

	return allCrds, nil
}
