package crd

import (
	"io/ioutil"
	"path/filepath"

	"github.com/giantswarm/microerror"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"
)

const (
	crdDirectoryV1      = "/config/crd/v1"
	crdDirectoryV1Beta1 = "/config/crd/v1beta1"
	crdKind             = "CustomResourceDefinition"
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

func find(crdKind schema.GroupVersionKind, crGroup, crKind string) (interface{}, error) {
	var path string
	switch crdKind.Version {
	case v1GroupVersionKind.Version:
		path = crdDirectoryV1
	case v1beta1GroupVersionKind.Version:
		path = crdDirectoryV1Beta1
	}

	fs := _escFS(false)
	directory, err := fs.Open(path)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	files, err := directory.Readdir(0)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	for _, info := range files {
		if info.IsDir() {
			continue
		}
		if filepath.Ext(info.Name()) != ".yaml" {
			continue
		}

		// Read the file to a string.
		file, err := fs.Open(filepath.Join(path, info.Name()))
		if err != nil {
			return nil, microerror.Mask(err)
		}
		contents, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		// Unmarshal into Unstructured since we don't know if this is a v1 or v1beta1 CRD yet.
		var object unstructured.Unstructured
		err = yaml.UnmarshalStrict(contents, &object)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if object.GetObjectKind().GroupVersionKind() != crdKind {
			continue
		}

		switch crdKind {
		case v1beta1GroupVersionKind:
			var crd v1beta1.CustomResourceDefinition
			err = yaml.UnmarshalStrict(contents, &crd)
			if err != nil {
				return nil, microerror.Mask(err)
			}
			if crGroup == crd.Spec.Group && crKind == crd.Spec.Names.Kind {
				return &crd, nil
			}
		case v1GroupVersionKind:
			var crd v1.CustomResourceDefinition
			err = yaml.UnmarshalStrict(contents, &crd)
			if err != nil {
				return nil, microerror.Mask(err)
			}
			if crGroup == crd.Spec.Group && crKind == crd.Spec.Names.Kind {
				return &crd, nil
			}
		}
	}

	return nil, microerror.Mask(notFoundError)
}

// LoadV1Beta1 loads a v1beta1 CRD from the virtual filesystem.
func LoadV1Beta1(group, kind string) *v1beta1.CustomResourceDefinition {
	found, err := find(v1beta1GroupVersionKind, group, kind)
	if err != nil {
		panic(microerror.Mask(err))
	}
	crd, ok := found.(*v1beta1.CustomResourceDefinition)
	if !ok {
		panic(microerror.Mask(conversionFailedError))
	}
	return crd
}

// LoadV1Beta1 loads a v1 CRD from the virtual filesystem
func LoadV1(group, kind string) (out *v1.CustomResourceDefinition) {
	found, err := find(v1GroupVersionKind, group, kind)
	if err != nil {
		panic(microerror.Mask(err))
	}
	crd, ok := found.(*v1.CustomResourceDefinition)
	if !ok {
		panic(microerror.Mask(conversionFailedError))
	}
	return crd
}
