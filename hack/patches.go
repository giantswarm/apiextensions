package main

import (
	"github.com/giantswarm/to"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

const (
	v1alpha3 = "v1alpha3"
)

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
func patchCAPICoreWebhook(crd *v1.CustomResourceDefinition) {
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
	}
}

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
func patchCAPIKubeadmBootstrapWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-kubeadm-bootstrap-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-kubeadm-bootstrap-unique-webhook",
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
	}
}

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
func patchCAPIControlPlaneWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-controlplane-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-controlplane-unique-webhook",
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
	}
}

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-azure-app/tree/master/helm/cluster-api-provider-azure/templates
func patchCAPZWebhook(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-azure-unique-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-azure-unique-webhook",
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
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
	// We only want to set v1alpha4 as not stored when there is also v1alpha3
	if len(crd.Spec.Versions) > 1 {
		for i, apiversion := range crd.Spec.Versions {
			if apiversion.Name == v1alpha3 {
				crd.Spec.Versions[i].Storage = true
			} else {
				crd.Spec.Versions[i].Storage = false
			}
		}
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
	// capi
	"clusters.cluster.x-k8s.io":                          patchCAPICoreWebhook,
	"kubeadmcontrolplanes.controlplane.cluster.x-k8s.io": patchCAPIControlPlaneWebhook,
	"kubeadmconfigs.bootstrap.cluster.x-k8s.io":          patchCAPIKubeadmBootstrapWebhook,
	"kubeadmconfigtemplates.bootstrap.cluster.x-k8s.io":  patchCAPIKubeadmBootstrapWebhook,
	"machinedeployments.cluster.x-k8s.io":                patchCAPICoreWebhook,
	"machinehealthchecks.cluster.x-k8s.io":               patchCAPICoreWebhook,
	"machines.cluster.x-k8s.io":                          patchCAPICoreWebhook,
	"machinesets.cluster.x-k8s.io":                       patchCAPICoreWebhook,
	// capa
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
	// capz
	"azureclusteridentities.infrastructure.cluster.x-k8s.io":    patchCAPZWebhook,
	"azureidentities.aadpodidentity.k8s.io":                     patchCAPZWebhook,
	"azureidentitybindings.aadpodidentity.k8s.io":               patchCAPZWebhook,
	"azurepodidentityexceptions.aadpodidentity.k8s.io":          patchCAPZWebhook,
	"azureassignedidentities.aadpodidentity.k8s.io":             patchCAPZWebhook,
	"azureclusters.infrastructure.cluster.x-k8s.io":             patchCAPZWebhook,
	"azuremachines.infrastructure.cluster.x-k8s.io":             patchCAPZWebhook,
	"azuremachinetemplates.infrastructure.cluster.x-k8s.io":     patchCAPZWebhook,
	"azuremachinepools.exp.infrastructure.cluster.x-k8s.io":     patchCAPZWebhook,
	"azuremachinepools.infrastructure.cluster.x-k8s.io":         patchCAPZWebhook,
	"azuremanagedclusters.infrastructure.cluster.x-k8s.io":      patchCAPZWebhook,
	"azuremanagedcontrolplanes.infrastructure.cluster.x-k8s.io": patchCAPZWebhook,
	"azuremanagedmachinepools.infrastructure.cluster.x-k8s.io":  patchCAPZWebhook,
	"azuremachinepoolmachines.infrastructure.cluster.x-k8s.io":  patchCAPZWebhook,
	// giantswarm
	"releases.release.giantswarm.io": patchReleaseValidation,
}
