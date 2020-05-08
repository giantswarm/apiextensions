package crd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/giantswarm/microerror"
	"github.com/markbates/pkger"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

const (
	crdDirectory = "/config/crd"
	crdKind      = "CustomResourceDefinition"
)

var (
	// GroupVersionKind of CustomResourceDefinition in apiextensions.k8s.io/v1beta1.
	v1beta1GroupVersionKind = schema.GroupVersionKind{
		Group:   apiextensions.GroupName,
		Version: "v1beta1",
		Kind:    crdKind,
	}
	// GroupVersionKind of CustomResourceDefinition in apiextensions.k8s.io/v1.
	v1GroupVersionKind = schema.GroupVersionKind{
		Group:   apiextensions.GroupName,
		Version: "v1",
		Kind:    crdKind,
	}
)

type objectHandler func(unstructured.Unstructured)

func iterateResources(groupVersionKind schema.GroupVersionKind, handle objectHandler) error {
	// Function called for every file in the CRD directory.
	walkFunc := func(fullPath string, info os.FileInfo, err error) error {
		// An unknown error, stop walking.
		if err != nil {
			return microerror.Mask(err)
		}
		if info.IsDir() {
			return nil
		}

		// pkger files have a path like github.com/giantswarm/apiextensions:/config/crd/bases/release.giantswarm.io_releases.yaml.
		split := strings.Split(fullPath, ":")
		path := split[1]
		extension := filepath.Ext(path)
		// Skip non-yaml files.
		if extension != ".yaml" {
			return nil
		}

		// Read the file to a string.
		yamlFile, err := pkger.Open(path)
		if err != nil {
			return microerror.Mask(err)
		}
		yamlString, err := ioutil.ReadAll(yamlFile)
		if err != nil {
			return microerror.Mask(err)
		}

		// Unmarshal into Unstructured since we don't know if this is a v1 or v1beta1 CRD yet.
		var object unstructured.Unstructured
		err = yaml.UnmarshalStrict(yamlString, &object)
		if err != nil {
			return microerror.Mask(err)
		}

		if object.GroupVersionKind() == groupVersionKind {
			handle(object)
		}

		return nil
	}

	// Entry point for walking the CRD YAML directory.
	err := pkger.Walk(crdDirectory, walkFunc)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

var cache []v1.CustomResourceDefinition
var cacheV1Beta1 []v1beta1.CustomResourceDefinition

// ListV1Beta1 loads all v1beta1 CRDs from the virtual filesystem.
func ListV1Beta1() []v1beta1.CustomResourceDefinition {
	if cacheV1Beta1 != nil {
		return cacheV1Beta1
	}
	handler := func(unstructured unstructured.Unstructured) {
		var crd v1beta1.CustomResourceDefinition
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructured.UnstructuredContent(), &crd)
		if err != nil {
			return
		}
		cacheV1Beta1 = append(cacheV1Beta1, crd)
	}
	err := iterateResources(v1beta1GroupVersionKind, handler)
	if err != nil {
		panic(microerror.Mask(err))
	}
	return cacheV1Beta1
}

// ListV1 loads all v1 CRDs from the virtual filesystem.
func List() []v1.CustomResourceDefinition {
	if cache != nil {
		return cache
	}
	handler := func(unstructured unstructured.Unstructured) {
		var crd v1.CustomResourceDefinition
		err := runtime.DefaultUnstructuredConverter.
			FromUnstructured(unstructured.UnstructuredContent(), &crd)
		if err != nil {
			return
		}
		cache = append(cache, crd)
	}
	err := iterateResources(v1GroupVersionKind, handler)
	if err != nil {
		panic(microerror.Mask(err))
	}
	return cache
}

// LoadV1Beta1 loads a v1beta1 CRD from the virtual filesystem.
func LoadV1Beta1(group, kind string) *v1beta1.CustomResourceDefinition {
	for _, crd := range ListV1Beta1() {
		if crd.Spec.Names.Kind == kind && crd.Spec.Group == group {
			return &crd
		}
	}
	panic(microerror.Mask(notFoundError))
}

// LoadV1 loads a v1 CRD from the virtual filesystem
func Load(group, kind string) *v1.CustomResourceDefinition {
	for _, crd := range List() {
		if crd.Spec.Names.Kind == kind && crd.Spec.Group == group {
			return &crd
		}
	}
	panic(microerror.Mask(notFoundError))
}
