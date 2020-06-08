package openapi

import (
	"net"

	"github.com/giantswarm/microerror"
	"github.com/go-openapi/spec"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	openapinamer "k8s.io/apiserver/pkg/endpoints/openapi"
	"k8s.io/apiserver/pkg/registry/rest"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"k8s.io/kube-openapi/pkg/builder"
	"k8s.io/kube-openapi/pkg/common"
)

// TypeInfo represents the declaration of a resource.
type TypeInfo struct {
	GroupVersion    schema.GroupVersion
	Resource        string
	Kind            string
	NamespaceScoped bool
}

// VersionResource is the declaration of a version resource.
type VersionResource struct {
	Version  string
	Resource string
}

// Config is the OpenAPI API generator configuration.
type Config struct {
	Scheme *runtime.Scheme
	Codecs serializer.CodecFactory

	Info                spec.InfoProps
	SecurityDefinitions *spec.SecurityDefinitions
	OpenAPIDefinitions  []common.GetOpenAPIDefinitions
	Resources           []TypeInfo
	GetterResources     []TypeInfo
	ListerResources     []TypeInfo
	CDResources         []TypeInfo
	RDResources         []TypeInfo
}

// GetDefinitions extracts the OpenAPI definitions from the configuration.
func (c *Config) GetDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	out := map[string]common.OpenAPIDefinition{}
	for _, def := range c.OpenAPIDefinitions {
		for k, v := range def(ref) {
			out[k] = v
		}
	}
	return out
}

// GenerateSpec creates a Swagger specification out of the configuration.
func GenerateSpec(cfg Config) (*spec.Swagger, error) {
	var err error

	metav1.AddToGroupVersion(cfg.Scheme, schema.GroupVersion{Version: "v1"})

	// Add metav1 types.
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	cfg.Scheme.AddUnversionedTypes(unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)

	var recommendedOptions *genericoptions.RecommendedOptions
	{
		recommendedOptions = genericoptions.NewRecommendedOptions("/registry/foo.com", cfg.Codecs.LegacyCodec(), &genericoptions.ProcessInfo{})
		recommendedOptions.SecureServing.BindPort = 8443
		recommendedOptions.Etcd = nil
		recommendedOptions.Authentication = nil
		recommendedOptions.Authorization = nil
		recommendedOptions.CoreAPI = nil
		recommendedOptions.Admission = nil

		err = recommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")})
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var serverConfig *genericapiserver.RecommendedConfig
	{
		serverConfig = genericapiserver.NewRecommendedConfig(cfg.Codecs)
		err = recommendedOptions.ApplyTo(serverConfig)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		serverConfig.OpenAPIConfig = genericapiserver.DefaultOpenAPIConfig(cfg.GetDefinitions, openapinamer.NewDefinitionNamer(cfg.Scheme))
		serverConfig.OpenAPIConfig.Info.InfoProps = cfg.Info
		serverConfig.OpenAPIConfig.SecurityDefinitions = cfg.SecurityDefinitions
	}

	genericServer, err := serverConfig.Complete().New("stash-server", genericapiserver.NewEmptyDelegate()) // completion is done in Complete, no need for a second time
	if err != nil {
		return nil, microerror.Mask(err)
	}

	table := map[string]map[VersionResource]rest.Storage{}
	{
		for _, ti := range cfg.Resources {
			var resmap map[VersionResource]rest.Storage
			if m, found := table[ti.GroupVersion.Group]; found {
				resmap = m
			} else {
				resmap = map[VersionResource]rest.Storage{}
				table[ti.GroupVersion.Group] = resmap
			}

			gvk := ti.GroupVersion.WithKind(ti.Kind)
			obj, err := cfg.Scheme.New(gvk)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			// Add list type.
			list, err := cfg.Scheme.New(ti.GroupVersion.WithKind(ti.Kind + "List"))
			if err != nil {
				return nil, microerror.Mask(err)
			}

			resmap[VersionResource{Version: ti.GroupVersion.Version, Resource: ti.Resource}] = NewStandardStorage(ResourceInfo{
				gvk:             gvk,
				obj:             obj,
				list:            list,
				namespaceScoped: ti.NamespaceScoped,
			})
		}
	}
	{
		for _, ti := range cfg.GetterResources {
			var resmap map[VersionResource]rest.Storage
			if m, found := table[ti.GroupVersion.Group]; found {
				resmap = m
			} else {
				resmap = map[VersionResource]rest.Storage{}
				table[ti.GroupVersion.Group] = resmap
			}

			gvk := ti.GroupVersion.WithKind(ti.Kind)
			obj, err := cfg.Scheme.New(gvk)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			resmap[VersionResource{Version: ti.GroupVersion.Version, Resource: ti.Resource}] = NewGetterStorage(ResourceInfo{
				gvk:             gvk,
				obj:             obj,
				namespaceScoped: ti.NamespaceScoped,
			})
		}
	}
	{
		for _, ti := range cfg.ListerResources {
			var resmap map[VersionResource]rest.Storage
			if m, found := table[ti.GroupVersion.Group]; found {
				resmap = m
			} else {
				resmap = map[VersionResource]rest.Storage{}
				table[ti.GroupVersion.Group] = resmap
			}

			gvk := ti.GroupVersion.WithKind(ti.Kind)
			obj, err := cfg.Scheme.New(gvk)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			// Add list type.
			list, err := cfg.Scheme.New(ti.GroupVersion.WithKind(ti.Kind + "List"))
			if err != nil {
				return nil, microerror.Mask(err)
			}

			resmap[VersionResource{Version: ti.GroupVersion.Version, Resource: ti.Resource}] = NewListerStorage(ResourceInfo{
				gvk:             gvk,
				obj:             obj,
				list:            list,
				namespaceScoped: ti.NamespaceScoped,
			})
		}
	}
	{
		for _, ti := range cfg.CDResources {
			var resmap map[VersionResource]rest.Storage
			if m, found := table[ti.GroupVersion.Group]; found {
				resmap = m
			} else {
				resmap = map[VersionResource]rest.Storage{}
				table[ti.GroupVersion.Group] = resmap
			}

			gvk := ti.GroupVersion.WithKind(ti.Kind)
			obj, err := cfg.Scheme.New(gvk)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			resmap[VersionResource{Version: ti.GroupVersion.Version, Resource: ti.Resource}] = NewCDStorage(ResourceInfo{
				gvk:             gvk,
				obj:             obj,
				namespaceScoped: ti.NamespaceScoped,
			})
		}
	}
	{
		for _, ti := range cfg.RDResources {
			var resmap map[VersionResource]rest.Storage
			if m, found := table[ti.GroupVersion.Group]; found {
				resmap = m
			} else {
				resmap = map[VersionResource]rest.Storage{}
				table[ti.GroupVersion.Group] = resmap
			}

			gvk := ti.GroupVersion.WithKind(ti.Kind)
			obj, err := cfg.Scheme.New(gvk)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			// Add list type.
			list, err := cfg.Scheme.New(ti.GroupVersion.WithKind(ti.Kind + "List"))
			if err != nil {
				return nil, microerror.Mask(err)
			}

			resmap[VersionResource{Version: ti.GroupVersion.Version, Resource: ti.Resource}] = NewRDStorage(ResourceInfo{
				gvk:             gvk,
				obj:             obj,
				list:            list,
				namespaceScoped: ti.NamespaceScoped,
			})
		}
	}

	for group, resmap := range table {
		apiGroupInfo := genericapiserver.NewDefaultAPIGroupInfo(group, cfg.Scheme, metav1.ParameterCodec, cfg.Codecs)
		for vr, s := range resmap {
			if _, found := apiGroupInfo.VersionedResourcesStorageMap[vr.Version]; !found {
				apiGroupInfo.VersionedResourcesStorageMap[vr.Version] = make(map[string]rest.Storage)
			}
			apiGroupInfo.VersionedResourcesStorageMap[vr.Version][vr.Resource] = s
		}
		if err := genericServer.InstallAPIGroup(&apiGroupInfo); err != nil {
			return nil, microerror.Mask(err)
		}
	}

	// Generate spec.
	spec, err := builder.BuildOpenAPISpec(genericServer.Handler.GoRestfulContainer.RegisteredWebServices(), serverConfig.OpenAPIConfig)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return spec, nil
}
