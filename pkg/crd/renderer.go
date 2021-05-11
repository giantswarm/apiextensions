package crd

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/giantswarm/microerror"
	"github.com/google/go-github/v35/github"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

// Kubernetes API group, version and kind for v1 CRDs
var crdGroupVersionKind = schema.GroupVersionKind{
	Group:   "apiextensions.k8s.io",
	Version: "v1",
	Kind:    "CustomResourceDefinition",
}

// Render creates helm chart templates for the given provider by downloading upstream CRDs, merging them with local
// CRDs, patching them, and writing them to the corresponding provider helm template directory.
func (r Renderer) Render(ctx context.Context, provider string) error {
	localCRDs, err := r.getLocalCRDs(provider)
	if err != nil {
		return microerror.Mask(err)
	}

	giantswarmFilename := helmChartTemplateFile(r.OutputDirectory, provider, "giantswarm.yaml")
	err = r.writeCRDsToFile(giantswarmFilename, localCRDs)
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
		return nil, microerror.Mask(notFoundError)
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

func (r Renderer) writeCRDsToFile(filename string, crds []v1.CustomResourceDefinition) error {
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
		crd, err := patchCRD(r.Patches, crd)
		if err != nil {
			return err
		}

		crdBytes, err := yaml.Marshal(crd)
		if err != nil {
			return err
		}

		_, err = writeBuffer.Write(crdBytes)
		if err != nil {
			return err
		}

		_, err = writeBuffer.Write([]byte("\n---\n"))
		if err != nil {
			return err
		}
	}

	return nil
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
