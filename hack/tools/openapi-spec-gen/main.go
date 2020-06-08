package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/go-openapi/spec"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kube-openapi/pkg/common"

	applicationv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	backupv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/backup/v1alpha1"
	corev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	examplev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"
	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	providerv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	releasev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	securityv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/security/v1alpha1"
	toolingv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/tooling/v1alpha1"

	"github.com/giantswarm/apiextensions/hack/tools/openapi-spec-gen/pkg/openapi"
)

const (
	clientName    = "giantswarm-cp-client"
	clientVersion = "1.0"
	outputFile    = "swagger.json"
)

var (
	securityDefinition = &spec.SecurityDefinitions{
		"oauth2": spec.OAuth2AccessToken("https://foo.com/authorization", "https://foo.com/token"),
	}
)

func main() {
	s := runtime.NewScheme()
	codec := serializer.NewCodecFactory(s)

	sb := runtime.SchemeBuilder{
		corev1alpha1.AddToScheme,
		applicationv1alpha1.AddToScheme,
		infrastructurev1alpha2.AddToScheme,
		backupv1alpha1.AddToScheme,
		examplev1alpha1.AddToScheme,
		providerv1alpha1.AddToScheme,
		releasev1alpha1.AddToScheme,
		securityv1alpha1.AddToScheme,
		toolingv1alpha1.AddToScheme,
	}

	v1.AddToGroupVersion(s, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(sb.AddToScheme(s))

	var resources []openapi.TypeInfo
	{
		resources = append(resources, getCoreTypes()...)
		resources = append(resources, getApplicationTypes()...)
		resources = append(resources, getBackupTypes()...)
		resources = append(resources, getInfrastructureTypes()...)
		resources = append(resources, getProviderTypes()...)
		resources = append(resources, getReleaseTypes()...)
		resources = append(resources, getSecurityTypes()...)
		resources = append(resources, getToolingTypes()...)
	}

	var getterResources []openapi.TypeInfo
	{
		getterResources = append(getterResources, getExampleTypes()...)
	}

	definitionFactories := []common.GetOpenAPIDefinitions{
		corev1alpha1.GetOpenAPIDefinitions,
		applicationv1alpha1.GetOpenAPIDefinitions,
		backupv1alpha1.GetOpenAPIDefinitions,
		infrastructurev1alpha2.GetOpenAPIDefinitions,
		examplev1alpha1.GetOpenAPIDefinitions,
		providerv1alpha1.GetOpenAPIDefinitions,
		releasev1alpha1.GetOpenAPIDefinitions,
		securityv1alpha1.GetOpenAPIDefinitions,
		toolingv1alpha1.GetOpenAPIDefinitions,
	}

	c := openapi.Config{
		Scheme: s,
		Codecs: codec,
		Info: spec.InfoProps{
			Title:   clientName,
			Version: clientVersion,
		},
		SecurityDefinitions: securityDefinition,
		OpenAPIDefinitions:  definitionFactories,
		Resources:           resources,
		GetterResources:     getterResources,
	}
	apiSpec, err := openapi.GenerateSpec(c)
	if err != nil {
		panic(err)
	}

	err = writeSpec(apiSpec)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("API Spec written to '%s'", outputFile))
}

func writeSpec(apiSpec *spec.Swagger) error {
	data, err := json.MarshalIndent(apiSpec, "", "  ")
	if err != nil {
		return microerror.Mask(err)
	}

	err = ioutil.WriteFile(outputFile, data, 0777)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func getInfrastructureTypes() []openapi.TypeInfo {
	var (
		awsClusterType           = infrastructurev1alpha2.AWSCluster{}
		awsControlPlaneType      = infrastructurev1alpha2.AWSControlPlane{}
		awsMachineDeploymentType = infrastructurev1alpha2.AWSMachineDeployment{}
		g8sControlPlaneType      = infrastructurev1alpha2.G8sControlPlane{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        getResourceName(awsClusterType),
			Kind:            getKindName(awsClusterType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        getResourceName(awsControlPlaneType),
			Kind:            getKindName(awsControlPlaneType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        getResourceName(awsMachineDeploymentType),
			Kind:            getKindName(awsMachineDeploymentType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        getResourceName(g8sControlPlaneType),
			Kind:            getKindName(g8sControlPlaneType),
			NamespaceScoped: true,
		},
	}
}

func getApplicationTypes() []openapi.TypeInfo {
	var (
		appCatalogType = applicationv1alpha1.AppCatalog{}
		appType        = applicationv1alpha1.App{}
		chartType      = applicationv1alpha1.Chart{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(appCatalogType),
			Kind:            getKindName(appCatalogType),
			NamespaceScoped: false,
		}, {
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(appType),
			Kind:            getKindName(appType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(chartType),
			Kind:            getKindName(chartType),
			NamespaceScoped: true,
		},
	}
}

func getBackupTypes() []openapi.TypeInfo {
	var (
		etcdBackupType = backupv1alpha1.ETCDBackup{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    backupv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(etcdBackupType),
			Kind:            getKindName(etcdBackupType),
			NamespaceScoped: false,
		},
	}
}

func getCoreTypes() []openapi.TypeInfo {
	var (
		certConfigType         = corev1alpha1.CertConfig{}
		chartConfigType        = corev1alpha1.ChartConfig{}
		drainerConfigType      = corev1alpha1.DrainerConfig{}
		awsClusterConfigType   = corev1alpha1.AWSClusterConfig{}
		azureClusterConfigType = corev1alpha1.AzureClusterConfig{}
		kvmClusterConfigType   = corev1alpha1.KVMClusterConfig{}
		flannelConfigType      = corev1alpha1.FlannelConfig{}
		ignitionType           = corev1alpha1.Ignition{}
		storageConfigType      = corev1alpha1.StorageConfig{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(certConfigType),
			Kind:            getKindName(certConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(chartConfigType),
			Kind:            getKindName(chartConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(drainerConfigType),
			Kind:            getKindName(drainerConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(awsClusterConfigType),
			Kind:            getKindName(awsClusterConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(azureClusterConfigType),
			Kind:            getKindName(azureClusterConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(kvmClusterConfigType),
			Kind:            getKindName(kvmClusterConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(flannelConfigType),
			Kind:            getKindName(flannelConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(ignitionType),
			Kind:            getKindName(ignitionType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(storageConfigType),
			Kind:            getKindName(storageConfigType),
			NamespaceScoped: true,
		},
	}
}

func getExampleTypes() []openapi.TypeInfo {
	var (
		memcachedConfigType = examplev1alpha1.MemcachedConfig{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    examplev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(memcachedConfigType),
			Kind:            getKindName(memcachedConfigType),
			NamespaceScoped: true,
		},
	}
}

func getProviderTypes() []openapi.TypeInfo {
	var (
		awsConfigType   = providerv1alpha1.AWSConfig{}
		azureConfigType = providerv1alpha1.AzureConfig{}
		kvmConfigType   = providerv1alpha1.KVMConfig{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(awsConfigType),
			Kind:            getKindName(awsConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(azureConfigType),
			Kind:            getKindName(azureConfigType),
			NamespaceScoped: true,
		}, {
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(kvmConfigType),
			Kind:            getKindName(kvmConfigType),
			NamespaceScoped: true,
		},
	}
}

func getReleaseTypes() []openapi.TypeInfo {
	var (
		releaseType = releasev1alpha1.Release{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    releasev1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(releaseType),
			Kind:            getKindName(releaseType),
			NamespaceScoped: false,
		},
	}
}

func getSecurityTypes() []openapi.TypeInfo {
	var (
		organizationType = securityv1alpha1.Organization{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    securityv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(organizationType),
			Kind:            getKindName(organizationType),
			NamespaceScoped: false,
		},
	}
}

func getToolingTypes() []openapi.TypeInfo {
	var (
		azureToolType = toolingv1alpha1.AzureTool{}
	)

	return []openapi.TypeInfo{
		{
			GroupVersion:    toolingv1alpha1.SchemeGroupVersion,
			Resource:        getResourceName(azureToolType),
			Kind:            getKindName(azureToolType),
			NamespaceScoped: true,
		},
	}
}

func getKindName(res interface{}) string {
	return reflect.TypeOf(res).Name()
}

func getResourceName(res interface{}) string {
	name := strings.ToLower(getKindName(res))
	// Transform to plural.
	name += "s"

	return name
}
