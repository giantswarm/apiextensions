package crd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/giantswarm/microerror"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/apiextensions/pkg/crd/internal"
)

const (
	crdKind = "CustomResourceDefinition"
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

type objectHandler func(data []byte) error

func iterateResources(groupVersionKind schema.GroupVersionKind, handle objectHandler) error {
	crdDirectory := fmt.Sprintf("/config/crd/%s", groupVersionKind.Version)
	fs := internal.FS(false)
	directory, err := fs.Open(crdDirectory)
	if err != nil {
		return microerror.Mask(err)
	}
	files, err := directory.Readdir(0)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, info := range files {
		if info.IsDir() {
			continue
		}
		if filepath.Ext(info.Name()) != ".yaml" {
			continue
		}

		// Read the file to a string.
		file, err := fs.Open(filepath.Join(crdDirectory, info.Name()))
		if err != nil {
			return microerror.Mask(err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			return microerror.Mask(err)
		}

		// Unmarshal into Unstructured since we don't know if this is a v1 or v1beta1 CRD yet.
		var object unstructured.Unstructured
		err = yaml.UnmarshalStrict(contents, &object)
		if err != nil {
			return microerror.Mask(err)
		}
		if object.GetObjectKind().GroupVersionKind() != groupVersionKind {
			continue
		}

		if err := handle(contents); err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

var cache []v1.CustomResourceDefinition
var cacheV1Beta1 []v1beta1.CustomResourceDefinition

// ListV1Beta1 loads all v1beta1 CRDs from the virtual filesystem.
func ListV1Beta1() ([]v1beta1.CustomResourceDefinition, error) {
	if cacheV1Beta1 != nil {
		return cacheV1Beta1, nil
	}

	handler := func(data []byte) error {
		var crd v1beta1.CustomResourceDefinition
		err := yaml.UnmarshalStrict(data, &crd)
		if err != nil {
			return microerror.Mask(err)
		}
		cacheV1Beta1 = append(cacheV1Beta1, crd)
		return nil
	}

	err := iterateResources(v1beta1GroupVersionKind, handler)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return cacheV1Beta1, nil
}

// ListV1 loads all v1 CRDs from the virtual filesystem.
func ListV1() ([]v1.CustomResourceDefinition, error) {
	if cache != nil {
		return cache, nil
	}

	handler := func(data []byte) error {
		var crd v1.CustomResourceDefinition
		err := yaml.UnmarshalStrict(data, &crd)
		if err != nil {
			return microerror.Mask(err)
		}
		cache = append(cache, crd)
		return nil
	}

	err := iterateResources(v1GroupVersionKind, handler)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return cache, nil
}

// LoadV1Beta1 loads a v1beta1 CRD from the virtual filesystem.
func LoadV1Beta1(group, kind string) *v1beta1.CustomResourceDefinition {
	crds, err := ListV1Beta1()
	if err != nil {
		panic(microerror.Mask(err))
	}

	for _, crd := range crds {
		if crd.Spec.Names.Kind == kind && crd.Spec.Group == group {
			return &crd
		}
	}
	panic(microerror.Mask(notFoundError))
}

// LoadV1 loads a v1 CRD from the virtual filesystem
func LoadV1(group, kind string) *v1.CustomResourceDefinition {
	crds, err := ListV1()
	if err != nil {
		panic(microerror.Mask(err))
	}

	for _, crd := range crds {
		if crd.Spec.Names.Kind == kind && crd.Spec.Group == group {
			return &crd
		}
	}
	panic(microerror.Mask(notFoundError))
}

// ExtractV1CRDVersions takes instance of CRD and variable number of versions
// and returns CRD with those versions. This function does not guarantee that
// all defined versions are present in returned CRDs.
//
// Example:
//	crd contains versions for v1alpha1 and v1alpha3.
//	versions is v1alpha3.
// Returned:
//	crd with only v1alpha3 version.
//
func ExtractV1CRDVersions(crd *v1.CustomResourceDefinition, versions ...string) *v1.CustomResourceDefinition {
	crd = crd.DeepCopy()

VERSION_LOOP:
	for i := 0; i < len(crd.Spec.Versions); i++ {
		v := crd.Spec.Versions[i]
		for _, name := range versions {
			if v.Name == name {
				// Keep current version and proceed to next.
				continue VERSION_LOOP
			}
		}

		// Remove current element since its version was not specified in versions.
		crd.Spec.Versions = append(crd.Spec.Versions[:i], crd.Spec.Versions[i+1:]...)
		i--
	}

	return crd
}

// WithStorageVersion takes CRD and version and sets storage flag for this
// version as true while setting others to false.
//
// Example:
//	crd contains versions v1alpha2 and v1alpha3 and v1alpha3 has storage == true.
//	version is v1alpha2.
// Returned:
//	crd with v1alpha2.Storage as true and v1alpha3.Storage as false.
//
func WithStorageVersion(crd *v1.CustomResourceDefinition, version string) *v1.CustomResourceDefinition {
	crd = crd.DeepCopy()

	versionSet := false
	for i := 0; i < len(crd.Spec.Versions); i++ {
		if crd.Spec.Versions[i].Name == version {
			crd.Spec.Versions[i].Storage = true
			versionSet = true
		} else {
			crd.Spec.Versions[i].Storage = false
		}
	}

	if !versionSet {
		panic(microerror.Maskf(notFoundError, "version: %q", version))
	}

	return crd
}
