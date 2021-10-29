package crd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"path/filepath"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	apiyaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

// decodeCRDs reads a slice of CRDs from multi-document YAML-formatted data provided by the given io.ReadCloser and
// closes it when complete or an error occurs.
func decodeCRDs(readCloser io.ReadCloser) ([]runtime.Object, error) {
	reader := apiyaml.NewYAMLReader(bufio.NewReader(readCloser))
	decoder := scheme.Codecs.UniversalDecoder()

	defer func(contentReader io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			panic(err)
		}
	}(readCloser)

	var crds []runtime.Object

	for {
		doc, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, microerror.Mask(err)
		}

		//  Skip over empty documents, i.e. a leading `---`
		if len(bytes.TrimSpace(doc)) == 0 {
			continue
		}

		var object unstructured.Unstructured
		_, decodedGVK, err := decoder.Decode(doc, nil, &object)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		switch *decodedGVK {
		case crdV1GVK:
			var crd v1.CustomResourceDefinition
			_, _, err = decoder.Decode(doc, nil, &crd)
			if err != nil {
				return nil, microerror.Mask(err)
			}

			crds = append(crds, &crd)
		}
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
func patchCRD(patches map[string]Patch, crd *v1.CustomResourceDefinition) (*v1.CustomResourceDefinition, error) {
	patch, ok := patches[crd.Name]
	if !ok {
		return crd, nil
	}

	crdCopy := crd.DeepCopy()
	patch(crdCopy)

	return crdCopy, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
