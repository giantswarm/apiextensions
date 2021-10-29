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
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

var (
	// Kubernetes API group, version and kind for v1 CRDs
	crdV1GVK = schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	}
)

// Render creates helm chart templates for the given provider by downloading upstream CRDs, merging them with local
// CRDs, patching them, and writing them to the corresponding provider helm template directory.
func (r Renderer) Render(ctx context.Context, provider string) error {
	internalCRDs, err := r.getLocalCRDs(provider)
	if err != nil {
		return microerror.Mask(err)
	}

	remoteCRDs, err := r.getRemoteCRDs(ctx, provider)
	if err != nil {
		return microerror.Mask(err)
	}

	internalCRDs = append(internalCRDs, remoteCRDs...)

	giantswarmFilename := helmChartTemplateFile(r.OutputDirectory, provider, "giantswarm.yaml")
	err = r.writeCRDsToFile(giantswarmFilename, internalCRDs)
	if err != nil {
		return microerror.Mask(err)
	}

	upstreamCRDs, err := r.getUpstreamCRDs(ctx, provider)
	if err != nil {
		return microerror.Mask(err)
	}

	upstreamFilename := helmChartTemplateFile(r.OutputDirectory, provider, "upstream.yaml")
	err = r.writeCRDsToFile(upstreamFilename, upstreamCRDs)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// downloadReleaseAssetCRDs returns a slice of CRDs by downloading the given GitHub release asset, parsing it as YAML,
// and filtering for only CRD objects.
func (r Renderer) downloadReleaseAssetCRDs(ctx context.Context, asset ReleaseAssetFileDefinition) ([]runtime.Object, error) {
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

	var allCrds []runtime.Object
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
func (r Renderer) getUpstreamCRDs(ctx context.Context, provider string) ([]runtime.Object, error) {
	var crds []runtime.Object
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

func (r Renderer) writeCRDsToFile(filename string, crds []runtime.Object) error {
	if len(crds) == 0 {
		return nil
	}

	writeBuffer, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return microerror.Mask(err)
	}

	defer func() {
		err = writeBuffer.Close()
		if err != nil {
			panic(microerror.JSON(microerror.Mask(err)))
		}
	}()

	for _, crd := range crds {
		if crdV1, ok := crd.(*v1.CustomResourceDefinition); ok {
			crd, err = patchCRD(r.Patches, crdV1)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		crdBytes, err := yaml.Marshal(crd)
		if err != nil {
			return microerror.Mask(err)
		}

		_, err = writeBuffer.Write(crdBytes)
		if err != nil {
			return microerror.Mask(err)
		}

		_, err = writeBuffer.Write([]byte("\n---\n"))
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

// getLocalCRDs reads the configured local directory and returns a slice of CRDs that have the given category.
func (r Renderer) getLocalCRDs(category string) ([]runtime.Object, error) {
	var crds []runtime.Object
	err := filepath.WalkDir("../config/crd", func(path string, entry os.DirEntry, walkErr error) error {
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
			var categories []string
			if crdV1, ok := crd.(*v1.CustomResourceDefinition); ok {
				categories = crdV1.Spec.Names.Categories
			} else if crdV1Beta1, ok := crd.(*v1beta1.CustomResourceDefinition); ok {
				categories = crdV1Beta1.Spec.Names.Categories
			}
			if contains(categories, category) {
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

// getRemoteCRDs returns all upstream CRDs for a provider based on the Renderer's upstream asset configuration.
func (r Renderer) getRemoteCRDs(ctx context.Context, provider string) ([]runtime.Object, error) {
	var crds []runtime.Object
	for _, releaseAsset := range r.RemoteRepositories {
		if releaseAsset.Provider != provider {
			continue
		}

		releaseAssetCRDs, err := r.downloadRepositoryCRDs(ctx, releaseAsset)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		crds = append(crds, releaseAssetCRDs...)
	}

	return crds, nil
}

// downloadRepositoryCRDs returns a slice of CRDs by downloading the given GitHub release asset, parsing it as YAML,
// and filtering for only CRD objects.
func (r Renderer) downloadRepositoryCRDs(ctx context.Context, asset RemoteRepositoryDefinition) ([]runtime.Object, error) {
	refString := fmt.Sprintf("tags/%s", asset.Version)
	ref, response, err := r.GithubClient.Git.GetRef(ctx, asset.Owner, asset.Repo, refString)
	if err != nil && response.StatusCode == 404 {
		refString = fmt.Sprintf("heads/%s", asset.Version)
		ref, _, err = r.GithubClient.Git.GetRef(ctx, asset.Owner, asset.Repo, refString)
	}
	if err != nil {
		return nil, microerror.Mask(err)
	}

	commit, _, err := r.GithubClient.Git.GetCommit(ctx, asset.Owner, asset.Repo, ref.Object.GetSHA())
	if err != nil {
		return nil, microerror.Mask(err)
	}

	tree, _, err := r.GithubClient.Git.GetTree(ctx, asset.Owner, asset.Repo, commit.Tree.GetSHA(), true)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var targetEntries []*github.TreeEntry
	for _, entry := range tree.Entries {
		if entry.GetType() == "blob" && strings.HasPrefix(entry.GetPath(), asset.Path) {
			targetEntries = append(targetEntries, entry)
		}
	}
	if targetEntries == nil {
		return nil, notFoundError
	}

	var allCrds []runtime.Object
	for _, entry := range targetEntries {
		blob, _, err := r.GithubClient.Git.GetBlob(ctx, asset.Owner, asset.Repo, entry.GetSHA())
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
