//go:generate go run build-charts.go

package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/giantswarm/to"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	capiyaml "sigs.k8s.io/cluster-api/util/yaml"
	"sigs.k8s.io/yaml"
)

type releaseAssetFileDefinition struct {
	owner    string
	repo     string
	version  string
	files    []string
	provider string
}

var notFoundError = errors.New("not found")

var upstreamReleaseAssets = []releaseAssetFileDefinition{
	{
		owner:    "kubernetes-sigs",
		repo:     "cluster-api",
		version:  "v0.3.14",
		files:    []string{"cluster-api-components.yaml"},
		provider: "common",
	},
	{
		owner:   "kubernetes-sigs",
		repo:    "cluster-api-provider-aws",
		version: "v0.6.5",
		files: []string{
			"eks-bootstrap-components.yaml",
			"eks-controlplane-components.yaml",
			"infrastructure-components.yaml",
		},
		provider: "aws",
	},
	{
		owner:    "kubernetes-sigs",
		repo:     "cluster-api-provider-azure",
		version:  "v0.4.12",
		files:    []string{"infrastructure-components.yaml"},
		provider: "azure",
	},
	{
		owner:    "kubernetes-sigs",
		repo:     "cluster-api-provider-vsphere",
		version:  "v0.7.6",
		files:    []string{"infrastructure-components.yaml"},
		provider: "vmware",
	},
	{
		owner:    "Azure",
		repo:     "aad-pod-identity",
		version:  "v1.7.4",
		files:    []string{"deployment.yaml"},
		provider: "azure",
	},
}

func decodeCRDs(readCloser io.ReadCloser) ([]v1.CustomResourceDefinition, error) {
	decoder := capiyaml.NewYAMLDecoder(readCloser)
	defer func(contentReader io.ReadCloser) {
		err := decoder.Close()
		if err != nil {
			panic(err)
		}
	}(readCloser)

	crdGVK := schema.GroupVersionKind{
		Group:   "apiextensions.k8s.io",
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	}
	var crds []v1.CustomResourceDefinition
	for {
		var crd v1.CustomResourceDefinition
		_, decodedGVK, err := decoder.Decode(nil, &crd)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, err
		}
		if *decodedGVK != crdGVK {
			continue
		}
		crds = append(crds, crd)
	}

	return crds, nil
}

func downloadReleaseAssetCRDs(ctx context.Context, client *github.Client, asset releaseAssetFileDefinition) ([]v1.CustomResourceDefinition, error) {
	release, _, err := client.Repositories.GetReleaseByTag(ctx, asset.owner, asset.repo, asset.version)
	if err != nil {
		return nil, err
	}

	var targetAssets []*github.ReleaseAsset
	for _, releaseAsset := range release.Assets {
		for _, file := range asset.files {
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
		contentReader, _, err := client.Repositories.DownloadReleaseAsset(ctx, asset.owner, asset.repo, targetAsset.GetID(), http.DefaultClient)
		if err != nil {
			return nil, err
		}
		crds, err := decodeCRDs(contentReader)
		if err != nil {
			return nil, err
		}
		allCrds = append(allCrds, crds...)
	}

	return allCrds, nil
}

func patchCAPIWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-core-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-core-unique-webhook",
					Path:      to.StringP("/convert"),
					Port:      &port,
				},
				CABundle: []byte("\n"),
			},
			ConversionReviewVersions: []string{
				"v1",
				"v1beta1",
			},
		},
	}
}

func patchCAPAWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-unique-webhook",
					Path:      to.StringP("/convert"),
					Port:      &port,
				},
				CABundle: []byte("\n"),
			},
			ConversionReviewVersions: []string{
				"v1",
				"v1beta1",
			},
		},
	}
}

func patchEKSControlPlaneWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-eks-control-plane-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-eks-control-plane-unique-webhook",
					Path:      to.StringP("/convert"),
					Port:      &port,
				},
				CABundle: []byte("\n"),
			},
			ConversionReviewVersions: []string{
				"v1",
				"v1beta1",
			},
		},
	}
}

func patchEKSConfigWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-eks-bootstrap-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-eks-bootstrap-unique-webhook",
					Path:      to.StringP("/convert"),
					Port:      &port,
				},
				CABundle: []byte("\n"),
			},
			ConversionReviewVersions: []string{
				"v1",
				"v1beta1",
			},
		},
	}
}

func patchReleaseValidation(crd *v1.CustomResourceDefinition) {
	for i := range crd.Spec.Versions {
		crd.Spec.Versions[i].Schema.OpenAPIV3Schema.Properties["metadata"] = v1.JSONSchemaProps{
			Type: "object",
			Properties: map[string]v1.JSONSchemaProps{
				"name": {
					Pattern: "^v(0|[1-9]\\d*)\\.(0|[1-9]\\d*)\\.(0|[1-9]\\d*)(-[\\.0-9a-zA-Z]*)?$",
					Type:    "string",
				},
			},
		}
	}
}

var patches = map[string]func(crd *v1.CustomResourceDefinition){
	"clusters.cluster.x-k8s.io":                                      patchCAPIWebhook,
	"machinedeployments.cluster.x-k8s.io":                            patchCAPIWebhook,
	"machinehealthchecks.cluster.x-k8s.io":                           patchCAPIWebhook,
	"machines.cluster.x-k8s.io":                                      patchCAPIWebhook,
	"machinesets.cluster.x-k8s.io":                                   patchCAPIWebhook,
	"awsclustercontrolleridentities.infrastructure.cluster.x-k8s.io": patchCAPAWebhook,
	"awsclusterroleidentities.infrastructure.cluster.x-k8s.io":       patchCAPAWebhook,
	"awsclusters.infrastructure.cluster.x-k8s.io":                    patchCAPAWebhook,
	"awsclusterstaticidentities.infrastructure.cluster.x-k8s.io":     patchCAPAWebhook,
	"awsfargateprofiles.infrastructure.cluster.x-k8s.io":             patchCAPAWebhook,
	"awsmachinepools.infrastructure.cluster.x-k8s.io":                patchCAPAWebhook,
	"awsmachines.infrastructure.cluster.x-k8s.io":                    patchCAPAWebhook,
	"awsmachinetemplates.infrastructure.cluster.x-k8s.io":            patchCAPAWebhook,
	"awsmanagedclusters.infrastructure.cluster.x-k8s.io":             patchCAPAWebhook,
	"awsmanagedcontrolplanes.controlplane.cluster.x-k8s.io":          patchEKSControlPlaneWebhook,
	"awsmanagedmachinepools.infrastructure.cluster.x-k8s.io":         patchCAPAWebhook,
	"eksconfigs.bootstrap.cluster.x-k8s.io":                          patchEKSConfigWebhook,
	"eksconfigtemplates.bootstrap.cluster.x-k8s.io":                  patchEKSConfigWebhook,
	"releases.release.giantswarm.io":                                 patchReleaseValidation,
}

func patchCRD(crd v1.CustomResourceDefinition) (v1.CustomResourceDefinition, error) {
	patch, ok := patches[crd.Name]
	if !ok {
		return crd, nil
	}

	crdCopy := crd.DeepCopy()
	patch(crdCopy)

	return *crdCopy, nil
}

func getUpstreamCRDs(ctx context.Context, client *github.Client, provider string) ([]v1.CustomResourceDefinition, error) {
	var crds []v1.CustomResourceDefinition
	for _, releaseAsset := range upstreamReleaseAssets {
		if releaseAsset.provider != provider {
			continue
		}

		releaseAssetCRDs, err := downloadReleaseAssetCRDs(ctx, client, releaseAsset)
		if err != nil {
			return nil, err
		}

		crds = append(crds, releaseAssetCRDs...)
	}

	return crds, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getLocalCRDs(category string) ([]v1.CustomResourceDefinition, error) {
	var crds []v1.CustomResourceDefinition
	err := filepath.WalkDir("../config/crd", func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		fileCRDs, err := decodeCRDs(file)
		if err != nil {
			return err
		}

		for _, crd := range fileCRDs {
			if contains(crd.Spec.Names.Categories, category) {
				crds = append(crds, crd)
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return crds, nil
}

func writeCRDsToFile(filename string, crds []v1.CustomResourceDefinition) error {
	if len(crds) == 0 {
		return nil
	}

	writeBuffer, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	defer func() {
		err = writeBuffer.Close()
		if err != nil {
			panic(err)
		}
	}()

	for _, crd := range crds {
		crd, err := patchCRD(crd)
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

func renderChart(ctx context.Context, client *github.Client, provider string) error {
	localCRDs, err := getLocalCRDs(provider)
	if err != nil {
		return err
	}

	err = writeCRDsToFile("../helm/crds-"+provider+"/templates/giantswarm.yaml", localCRDs)
	if err != nil {
		return err
	}

	upstreamCRDs, err := getUpstreamCRDs(ctx, client, provider)
	if err != nil {
		return err
	}

	err = writeCRDsToFile("../helm/crds-"+provider+"/templates/upstream.yaml", upstreamCRDs)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	for _, provider := range []string{"common", "aws", "azure", "kvm", "vmware"} {
		err := renderChart(ctx, client, provider)
		if err != nil {
			log.Fatal(err)
		}
	}
}
