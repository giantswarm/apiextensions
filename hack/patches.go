package main

import (
	"github.com/giantswarm/to"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/crd"
)

func patchCAPICoreWebhook(crd *v1.CustomResourceDefinition) {
	if len(crd.Spec.Versions) < 2 {
		// If we don't have at least 2 versions, there is no need to have a conversion webhook.
		return
	}

	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-core-cert"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-core",
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

func patchCAPIKubeadmBootstrapWebhook(crd *v1.CustomResourceDefinition) {
	if len(crd.Spec.Versions) < 2 {
		// If we don't have at least 2 versions, there is no need to have a conversion webhook.
		return
	}

	var hasV1alpha4 bool
	for _, v := range crd.Spec.Versions {
		if v.Name == "v1alpha4" {
			hasV1alpha4 = true
			break
		}
	}

	// The name of the certificate and conversion webhook service changed between v1alpha3 and v1alpha4 (the app also
	// changed from cluster-api-bootstrap-provider-kubeadm-app to cluster-api-app) so we check which versions are present and apply the correct
	// patch.
	if hasV1alpha4 {
		patchCAPIKubeadmBootstrapWebhookV1Alpha4(crd)
	} else {
		patchCAPIKubeadmBootstrapWebhookV1Alpha3(crd)
	}
}

// Keep in sync with https://github.com/giantswarm/cluster-api-bootstrap-provider-kubeadm-app/tree/main/helm/cluster-api-bootstrap-provider-kubeadm/templates
func patchCAPIKubeadmBootstrapWebhookV1Alpha3(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-bootstrap-cert"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-bootstrap",
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

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
func patchCAPIKubeadmBootstrapWebhookV1Alpha4(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-core-cert"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-bootstrap",
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

func patchCAPIControlPlaneWebhook(crd *v1.CustomResourceDefinition) {
	if len(crd.Spec.Versions) < 2 {
		// If we don't have at least 2 versions, there is no need to have a conversion webhook.
		return
	}

	var isV1alpha4 bool
	for _, v := range crd.Spec.Versions {
		if v.Name == "v1alpha4" {
			isV1alpha4 = true
			break
		}
	}

	// The name of the certificate and conversion webhook service changed between v1alpha3 and v1alpha4 (the app also
	// changed from cluster-api-control-plane-app to cluster-api-app) so we check which versions are present and apply the correct
	// patch.
	if isV1alpha4 {
		patchCAPIControlPlaneWebhookV1Alpha4(crd)
	} else {
		patchCAPIControlPlaneWebhookV1Alpha3(crd)
	}
}

// Keep in sync with https://github.com/giantswarm/cluster-api-control-plane-app/tree/main/helm/cluster-api-control-plane/templates
func patchCAPIControlPlaneWebhookV1Alpha3(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-controlplane-cert"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-controlplane",
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

// Keep in sync with https://github.com/giantswarm/cluster-api-core-app/tree/main/helm/cluster-api-core/templates
func patchCAPIControlPlaneWebhookV1Alpha4(crd *v1.CustomResourceDefinition) {
	port := int32(9443)
	if _, ok := crd.Annotations["cert-manager.io/inject-ca-from"]; ok {
		crd.Annotations["cert-manager.io/inject-ca-from"] = "giantswarm/cluster-api-core-cert"
	}
	crd.Spec.Conversion = &v1.CustomResourceConversion{
		Strategy: v1.WebhookConverter,
		Webhook: &v1.WebhookConversion{
			ClientConfig: &v1.WebhookClientConfig{
				Service: &v1.ServiceReference{
					Namespace: "giantswarm",
					Name:      "cluster-api-controlplane",
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
	if len(crd.Spec.Versions) < 2 {
		// If we don't have at least 2 versions, there is no need to have a conversion webhook.
		return
	}

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
func patchCAPVWebhook(crd *v1.CustomResourceDefinition) {
	if len(crd.Spec.Versions) < 2 {
		// If we don't have at least 2 versions, there is no need to have a conversion webhook.
		return
	}

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
func patchCAPZWebhook(crd *v1.CustomResourceDefinition) {
	delete(crd.Annotations, "cert-manager.io/inject-ca-from")
	crd.Spec.Conversion = nil
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
	"clusterclasses.cluster.x-k8s.io":                            patchCAPICoreWebhook,
	"clusters.cluster.x-k8s.io":                                  patchCAPICoreWebhook,
	"clusterresourcesetbindings.addons.cluster.x-k8s.io":         patchCAPICoreWebhook,
	"clusterresourcesets.addons.cluster.x-k8s.io":                patchCAPICoreWebhook,
	"kubeadmcontrolplanes.controlplane.cluster.x-k8s.io":         patchCAPIControlPlaneWebhook,
	"kubeadmcontrolplanetemplates.controlplane.cluster.x-k8s.io": patchCAPIControlPlaneWebhook,
	"kubeadmconfigs.bootstrap.cluster.x-k8s.io":                  patchCAPIKubeadmBootstrapWebhook,
	"kubeadmconfigtemplates.bootstrap.cluster.x-k8s.io":          patchCAPIKubeadmBootstrapWebhook,
	"machinedeployments.cluster.x-k8s.io":                        patchCAPICoreWebhook,
	"machinehealthchecks.cluster.x-k8s.io":                       patchCAPICoreWebhook,
	"machinepools.cluster.x-k8s.io":                              patchCAPICoreWebhook,
	"machines.cluster.x-k8s.io":                                  patchCAPICoreWebhook,
	"machinesets.cluster.x-k8s.io":                               patchCAPICoreWebhook,
	// capa
	"awsclustercontrolleridentities.infrastructure.cluster.x-k8s.io": patchCAPAWebhook,
	"awsclusterroleidentities.infrastructure.cluster.x-k8s.io":       patchCAPAWebhook,
	"awsclusters.infrastructure.cluster.x-k8s.io":                    patchCAPAWebhook,
	"awsclusterstaticidentities.infrastructure.cluster.x-k8s.io":     patchCAPAWebhook,
	"awsclustertemplates.infrastructure.cluster.x-k8s.io":            patchCAPAWebhook,
	"awsfargateprofiles.infrastructure.cluster.x-k8s.io":             patchCAPAWebhook,
	"awsmachinepools.infrastructure.cluster.x-k8s.io":                patchCAPAWebhook,
	"awsmachines.infrastructure.cluster.x-k8s.io":                    patchCAPAWebhook,
	"awsmachinetemplates.infrastructure.cluster.x-k8s.io":            patchCAPAWebhook,
	"awsmanagedclusters.infrastructure.cluster.x-k8s.io":             patchCAPAWebhook,
	"awsmanagedcontrolplanes.controlplane.cluster.x-k8s.io":          patchCAPAWebhook,
	"awsmanagedmachinepools.infrastructure.cluster.x-k8s.io":         patchCAPAWebhook,
	"eksconfigs.bootstrap.cluster.x-k8s.io":                          patchCAPAWebhook,
	"eksconfigtemplates.bootstrap.cluster.x-k8s.io":                  patchCAPAWebhook,
	// capz
	"azureclusteridentities.infrastructure.cluster.x-k8s.io":   patchCAPZWebhook,
	"azureclusters.infrastructure.cluster.x-k8s.io":            patchCAPZWebhook,
	"azuremachines.infrastructure.cluster.x-k8s.io":            patchCAPZWebhook,
	"azuremachinetemplates.infrastructure.cluster.x-k8s.io":    patchCAPZWebhook,
	"azuremachinepoolmachines.infrastructure.cluster.x-k8s.io": patchCAPZWebhook,
	"azuremachinepools.exp.infrastructure.cluster.x-k8s.io":    patchCAPZWebhook,
	"azuremachinepools.infrastructure.cluster.x-k8s.io":        patchCAPZWebhook,
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
