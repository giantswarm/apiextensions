package main

import (
	"github.com/giantswarm/to"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/apiextensions/v6/pkg/crd"
)

const (
	Azure                         = "azure"
	InjectCaFromCertificateLegacy = "giantswarm/cluster-api-core-cert"
)

func patchCAPICoreWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		if provider == Azure {
			crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/capi-serving-cert"
		} else {
			crd.Annotations["cert-manager.io/inject-ca-from"] = InjectCaFromCertificateLegacy
		}
	}

	webhookServiceName := "cluster-api-core"
	if provider == Azure {
		webhookServiceName = "capi-webhook-service"
		port = int32(443)
	}

	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      webhookServiceName,
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

// Keep in sync with https://github.com/giantswarm/cluster-api-bootstrap-provider-kubeadm-app/tree/main/helm/cluster-api-bootstrap-provider-kubeadm/templates
func patchCAPIKubeadmBootstrapWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		if provider == Azure {
			crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/capi-kubeadm-bootstrap-serving-cert"
		} else {
			crd.Annotations["cert-manager.io/inject-ca-from"] = InjectCaFromCertificateLegacy
		}
	}

	webhookServiceName := "cluster-api-bootstrap"
	if provider == Azure {
		webhookServiceName = "capi-kubeadm-bootstrap-webhook-service"
		port = int32(443)
	}

	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      webhookServiceName,
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

// Keep in sync with https://github.com/giantswarm/cluster-api-control-plane-app/tree/main/helm/cluster-api-control-plane/templates
func patchCAPIControlPlaneWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		if provider == Azure {
			crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/capi-kubeadm-control-plane-serving-cert"
		} else {
			crd.Annotations["cert-manager.io/inject-ca-from"] = InjectCaFromCertificateLegacy
		}
	}

	webhookServiceName := "cluster-api-controlplane"
	if provider == Azure {
		webhookServiceName = "capi-kubeadm-control-plane-webhook-service"
		port = int32(443)
	}

	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      webhookServiceName,
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
func patchCAPAWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-webhook",
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

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-vsphere-app/tree/master/helm/cluster-api-provider-vsphere/templates
func patchCAPVWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-vsphere-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-vsphere-webhook",
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

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-azure-app/tree/master/helm/cluster-api-provider-azure/templates
func patchCAPZWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok || crd.Name == "azureclusteridentities.infrastructure.cluster.x-k8s.io" {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/capz-serving-cert"
	}

	if crd.Spec.Conversion != nil || crd.Name == "azureclusteridentities.infrastructure.cluster.x-k8s.io" {
		crd.Spec.Conversion = &v1.CustomResourceConversion{
			Strategy: v1.WebhookConverter,
			Webhook: &v1.WebhookConversion{
				ClientConfig: &v1.WebhookClientConfig{
					Service: &v1.ServiceReference{
						Namespace: "giantswarm",
						Name:      "capz-webhook-service",
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
}

// Keep in sync with https://github.com/giantswarm/cluster-api-provider-aws-app/tree/master/helm/cluster-api-provider-aws/templates/eks/control-plane
func patchEKSControlPlaneWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-eks-control-plane-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-eks-control-plane-webhook",
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
func patchEKSConfigWebhook(provider string, crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-provider-aws-eks-bootstrap-webhook"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-provider-aws-eks-bootstrap-webhook",
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
func patchReleaseValidation(provider string, crd *v1.CustomResourceDefinition) {
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

// Upstream CRD contains a string which gitleaks detects as a leak. Applies the normal patchCAPAWebhook patch and then
// edits the description to avoid the false positive.
func patchAWSClusterStaticIdentities(provider string, crd *v1.CustomResourceDefinition) {
	patchCAPAWebhook(provider, crd)
	version := crd.Spec.Versions[0]
	schema := version.Schema.OpenAPIV3Schema
	secretRef := schema.Properties["spec"].Properties["secretRef"]
	secretRef.Description = "Reference to a secret containing the credentials. The secret should contain the following data keys:  AccessKeyID: <access key id>  SecretAccessKey: <secret access key>  SessionToken: Optional"
	schema.Properties["spec"].Properties["secretRef"] = secretRef
	crd.Spec.Versions[0] = version
}

var patches = map[string]crd.Patch{
	// capi
	"clusterclasses.cluster.x-k8s.io":                    patchCAPICoreWebhook,
	"clusters.cluster.x-k8s.io":                          patchCAPICoreWebhook,
	"kubeadmcontrolplanes.controlplane.cluster.x-k8s.io": patchCAPIControlPlaneWebhook,
	"kubeadmconfigs.bootstrap.cluster.x-k8s.io":          patchCAPIKubeadmBootstrapWebhook,
	"kubeadmconfigtemplates.bootstrap.cluster.x-k8s.io":  patchCAPIKubeadmBootstrapWebhook,
	"machinedeployments.cluster.x-k8s.io":                patchCAPICoreWebhook,
	"machinepools.cluster.x-k8s.io":                      patchCAPICoreWebhook,
	"machinehealthchecks.cluster.x-k8s.io":               patchCAPICoreWebhook,
	"machines.cluster.x-k8s.io":                          patchCAPICoreWebhook,
	"machinesets.cluster.x-k8s.io":                       patchCAPICoreWebhook,
	// capa
	"awsclustercontrolleridentities.infrastructure.cluster.x-k8s.io": patchCAPAWebhook,
	"awsclusterroleidentities.infrastructure.cluster.x-k8s.io":       patchCAPAWebhook,
	"awsclusters.infrastructure.cluster.x-k8s.io":                    patchCAPAWebhook,
	"awsclusterstaticidentities.infrastructure.cluster.x-k8s.io":     patchAWSClusterStaticIdentities,
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
	"azuremachinepoolmachines.infrastructure.cluster.x-k8s.io":  patchCAPZWebhook,
	"azuremachinepools.exp.infrastructure.cluster.x-k8s.io":     patchCAPZWebhook,
	"azuremachinepools.infrastructure.cluster.x-k8s.io":         patchCAPZWebhook,
	"azuremanagedclusters.infrastructure.cluster.x-k8s.io":      patchCAPZWebhook,
	"azuremanagedcontrolplanes.infrastructure.cluster.x-k8s.io": patchCAPZWebhook,
	"azuremanagedmachinepools.infrastructure.cluster.x-k8s.io":  patchCAPZWebhook,
	// giantswarm
	"releases.release.giantswarm.io": patchReleaseValidation,
	// capv
	"vsphereclusters.infrastructure.cluster.x-k8s.io":         patchCAPVWebhook,
	"vsphereclustertemplates.infrastructure.cluster.x-k8s.io": patchCAPVWebhook,
	"vspheredeploymentzones.infrastructure.cluster.x-k8s.io":  patchCAPVWebhook,
	"vspherefailuredomains.infrastructure.cluster.x-k8s.io":   patchCAPVWebhook,
	"vspheremachines.infrastructure.cluster.x-k8s.io":         patchCAPVWebhook,
	"vspheremachinetemplates.infrastructure.cluster.x-k8s.io": patchCAPVWebhook,
	"vspherevms.infrastructure.cluster.x-k8s.io":              patchCAPVWebhook,
}
