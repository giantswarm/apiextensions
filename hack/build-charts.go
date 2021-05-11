//go:generate go run build-charts.go

package main

import (
	"context"
	"log"
	"os"

	"github.com/giantswarm/to"
	"github.com/google/go-github/v35/github"
	"golang.org/x/oauth2"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

var upstreamReleaseAssets = []crd.ReleaseAssetFileDefinition{
	{
		Owner:    "kubernetes-sigs",
		Repo:     "cluster-api",
		Version:  "v0.3.14",
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
		Version:  "v0.4.12",
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
		Version:  "v1.7.4",
		Files:    []string{"deployment.yaml"},
		Provider: "azure",
	},
}

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
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

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-aws-app/tree/master/helm/cluster-api-provider-aws/templates
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

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-aws-app/tree/master/helm/cluster-api-provider-aws/templates/eks/control-plane
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

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-aws-app/tree/master/helm/cluster-api-provider-aws/templates/eks/bootstrap
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

// Kubebuilder comments can't add validation to metadata properties, so we manually specify the validation for release names here.
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

var patches = map[string]crd.Patch{
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

func main() {
	ctx := context.Background()
	token := oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")}
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
