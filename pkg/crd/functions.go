package crd

import (
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	capiyaml "sigs.k8s.io/cluster-api/util/yaml"
)

// decodeCRDs reads a slice of CRDs from multi-document YAML-formatted data provided by the given io.ReadCloser and
// closes it when complete or an error occurs.
func decodeCRDs(readCloser io.ReadCloser) ([]v1.CustomResourceDefinition, error) {
	decoder := capiyaml.NewYAMLDecoder(readCloser)
	defer func(contentReader io.ReadCloser) {
		err := decoder.Close()
		if err != nil {
			panic(microerror.JSON(microerror.Mask(err)))
		}
	}(readCloser)

	var crds []v1.CustomResourceDefinition
	for {
		var crd v1.CustomResourceDefinition
		_, decodedGroupVersionKind, err := decoder.Decode(nil, &crd)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, microerror.Mask(err)
		}
		if *decodedGroupVersionKind != crdGroupVersionKind {
			continue
		}
		crds = append(crds, crd)
	}

	return crds, nil
}

func helmChartTemplateFile(helmDirectory, provider, templateFilename string) string {
	return filepath.Join(helmChartDirectory(helmDirectory, provider), "templates", templateFilename)
}

func helmChartDirectory(helmDirectory, provider string) string {
	chartName := fmt.Sprintf("crds-%s", provider)
	return filepath.Join(helmDirectory, chartName)
}

// patchCRD applies a patch function to a deep copy of the given CRD if defined in the given patch map. If no patch is
// defined, the CRD will be returned unchanged.
func patchCRD(patches map[string]Patch, crd v1.CustomResourceDefinition) (v1.CustomResourceDefinition, error) {
	patch, ok := patches[crd.Name]
	if !ok {
		return crd, nil
	}

	crdCopy := crd.DeepCopy()
	patch(crdCopy)

	return *crdCopy, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
