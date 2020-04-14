package crd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"sigs.k8s.io/yaml"

	_ "github.com/giantswarm/apiextensions"
)

func LoadCRD(group, kind string) (found *v1beta1.CustomResourceDefinition) {
	err := pkger.Walk("/config/crd/bases", func(fullPath string, info os.FileInfo, err error) error {
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

		var crd v1beta1.CustomResourceDefinition
		err = yaml.UnmarshalStrict(yamlString, &crd)
		if err != nil {
			return err
		}

		if group == crd.Spec.Group && kind == crd.Spec.Names.Kind {
			found = &crd
		}
		return nil
	})

	if err != nil {
		panic(err)
	}
	return
}
