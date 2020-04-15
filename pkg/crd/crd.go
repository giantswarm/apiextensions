package crd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/yaml"

	_ "github.com/giantswarm/apiextensions"
)

var (
	v1Kind = schema.GroupVersionKind{
		Group:   apiextensions.GroupName,
		Version: "v1",
		Kind:    "CustomResourceDefinition",
	}
	v1beta1Kind = schema.GroupVersionKind{
		Group:   apiextensions.GroupName,
		Version: "v1beta1",
		Kind:    "CustomResourceDefinition",
	}
)

func FindCRD(group, kind string) (interface{}, error) {
	var found interface{}
	walkFunc := func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if found != nil || info.IsDir() {
			return nil
		}

		split := strings.Split(fullPath, ":")
		path := split[1]
		extension := filepath.Ext(path)
		if extension != ".yaml" {
			return nil
		}

		yamlFile, err := pkger.Open(path)
		if err != nil {
			return err
		}
		yamlString, err := ioutil.ReadAll(yamlFile)
		if err != nil {
			return err
		}

		var object unstructured.Unstructured
		err = yaml.UnmarshalStrict(yamlString, &object)
		if err != nil {
			return err
		}

		switch object.GetObjectKind().GroupVersionKind() {
		case v1beta1Kind:
			var crd v1beta1.CustomResourceDefinition
			err = yaml.UnmarshalStrict(yamlString, &crd)
			if err != nil {
				return err
			}
			if group == crd.Spec.Group && kind == crd.Spec.Names.Kind {
				found = &crd
			}
			return nil
		case v1Kind:
			var crd v1.CustomResourceDefinition
			err = yaml.UnmarshalStrict(yamlString, &crd)
			if err != nil {
				return err
			}
			if group == crd.Spec.Group && kind == crd.Spec.Names.Kind {
				found = &crd
			}
			return nil
		}
		return nil
	}

	err := pkger.Walk("/config/crd/bases", walkFunc)
	if err != nil {
		return nil, err
	}

	return found, nil
}

func LoadV1Beta1(group, kind string) *v1beta1.CustomResourceDefinition {
	found, err := FindCRD(group, kind)
	if err != nil {
		panic(err)
	}
	if found == nil {
		panic("not found")
	}
	crd, ok := found.(*v1beta1.CustomResourceDefinition)
	if !ok {
		panic("conversion failed")
	}
	return crd
}

func LoadV1(group, kind string) (out *v1.CustomResourceDefinition) {
	found, err := FindCRD(group, kind)
	if err != nil {
		panic(err)
	}
	if found == nil {
		panic("not found")
	}
	crd, ok := found.(*v1.CustomResourceDefinition)
	if !ok {
		panic("conversion failed")
	}
	return crd
}
