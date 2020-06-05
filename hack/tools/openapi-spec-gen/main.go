package main

import (
	"fmt"
	"io/ioutil"

	applicationv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/application/v1alpha1"
	backupv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/backup/v1alpha1"
	corev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	examplev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/example/v1alpha1"
	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	providerv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	releasev1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	securityv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/security/v1alpha1"
	toolingv1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/tooling/v1alpha1"
	"github.com/go-openapi/spec"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kube-openapi/pkg/common"

	"github.com/giantswarm/apiextensions/hack/tools/openapi-spec-gen/pkg/openapi"
)

const (
	clientName    = "giantswarm-cp-client"
	clientVersion = "1.0"
	outputFile    = "swagger.json"
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

	c := openapi.Config{
		Scheme: s,
		Codecs: codec,
		Info: spec.InfoProps{
			Title:   clientName,
			Version: clientVersion,
		},
		SecurityDefinitions: &spec.SecurityDefinitions{
			"oauth2": spec.OAuth2AccessToken("https://foo.com/authorization", "https://foo.com/token"),
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			corev1alpha1.GetOpenAPIDefinitions,
			applicationv1alpha1.GetOpenAPIDefinitions,
			backupv1alpha1.GetOpenAPIDefinitions,
			infrastructurev1alpha2.GetOpenAPIDefinitions,
			examplev1alpha1.GetOpenAPIDefinitions,
			providerv1alpha1.GetOpenAPIDefinitions,
			releasev1alpha1.GetOpenAPIDefinitions,
			securityv1alpha1.GetOpenAPIDefinitions,
			toolingv1alpha1.GetOpenAPIDefinitions,
		},
		Resources:       resources,
		GetterResources: getterResources,
	}
	apiSpec, err := openapi.RenderOpenAPISpec(c)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(outputFile, []byte(apiSpec), 0777)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("API Spec written to '%s'", outputFile))
}

func getInfrastructureTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        "awsclusters",
			Kind:            "AWSCluster",
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        "awscontrolplanes",
			Kind:            "AWSControlPlane",
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        "awsmachinedeployments",
			Kind:            "AWSMachineDeployment",
			NamespaceScoped: true,
		}, {
			GroupVersion:    infrastructurev1alpha2.SchemeGroupVersion,
			Resource:        "g8scontrolplanes",
			Kind:            "G8sControlPlane",
			NamespaceScoped: true,
		},
	}
}

func getApplicationTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        "appcatalogs",
			Kind:            "AppCatalog",
			NamespaceScoped: false,
		}, {
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        "apps",
			Kind:            "App",
			NamespaceScoped: true,
		}, {
			GroupVersion:    applicationv1alpha1.SchemeGroupVersion,
			Resource:        "charts",
			Kind:            "Chart",
			NamespaceScoped: true,
		},
	}
}

func getBackupTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    backupv1alpha1.SchemeGroupVersion,
			Resource:        "etcdbackups",
			Kind:            "ETCDBackup",
			NamespaceScoped: false,
		},
	}
}

func getCoreTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "certconfigs",
			Kind:            "CertConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "chartconfigs",
			Kind:            "ChartConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "drainerconfigs",
			Kind:            "DrainerConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "awsclusterconfigs",
			Kind:            "AWSClusterConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "azureclusterconfigs",
			Kind:            "AzureClusterConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "kvmclusterconfigs",
			Kind:            "KVMClusterConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "flannelconfigs",
			Kind:            "FlannelConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "ignitions",
			Kind:            "Ignition",
			NamespaceScoped: true,
		}, {
			GroupVersion:    corev1alpha1.SchemeGroupVersion,
			Resource:        "storageconfigs",
			Kind:            "StorageConfig",
			NamespaceScoped: true,
		},
	}
}

func getExampleTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    examplev1alpha1.SchemeGroupVersion,
			Resource:        "memcachedconfigs",
			Kind:            "MemcachedConfig",
			NamespaceScoped: true,
		},
	}
}

func getProviderTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        "awsconfigs",
			Kind:            "AWSConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        "azureconfigs",
			Kind:            "AzureConfig",
			NamespaceScoped: true,
		}, {
			GroupVersion:    providerv1alpha1.SchemeGroupVersion,
			Resource:        "kvmconfigs",
			Kind:            "KVMConfig",
			NamespaceScoped: true,
		},
	}
}

func getReleaseTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    releasev1alpha1.SchemeGroupVersion,
			Resource:        "releases",
			Kind:            "Release",
			NamespaceScoped: false,
		},
	}
}

func getSecurityTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    securityv1alpha1.SchemeGroupVersion,
			Resource:        "organizations",
			Kind:            "Organization",
			NamespaceScoped: false,
		},
	}
}

func getToolingTypes() []openapi.TypeInfo {
	return []openapi.TypeInfo{
		{
			GroupVersion:    toolingv1alpha1.SchemeGroupVersion,
			Resource:        "azuretools",
			Kind:            "AzureTool",
			NamespaceScoped: true,
		},
	}
}
