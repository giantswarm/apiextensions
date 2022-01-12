package crd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/giantswarm/microerror"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	apiyaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

// decodeCRDs reads a slice of CRDs from multi-document YAML-formatted data provided by the given io.ReadCloser and
// closes it when complete or an error occurs.
func decodeCRDs(readCloser io.ReadCloser) ([]v1.CustomResourceDefinition, error) {
	reader := apiyaml.NewYAMLReader(bufio.NewReader(readCloser))
	decoder := scheme.Codecs.UniversalDecoder()

	defer func(contentReader io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			panic(err)
		}
	}(readCloser)

	var crds []v1.CustomResourceDefinition

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

			crds = append(crds, crd)
		}
	}

	return crds, nil
}

func helmChartDirectory(helmDirectory, provider string) string {
	chartName := fmt.Sprintf("crds-%s", provider)
	return filepath.Join(helmDirectory, chartName)
}

// patchCRD applies a patch function to a deep copy of the given CRD if defined in the given patch map. If no patch is
// defined, the CRD will be returned unchanged.
func patchCRD(provider string, patches map[string]Patch, crd v1.CustomResourceDefinition) (v1.CustomResourceDefinition, error) {
	patch, ok := patches[crd.Name]
	if !ok {
		return crd, nil
	}

	crdCopy := crd.DeepCopy()
	patch(provider, crdCopy)

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

func writeObjects(writer io.Writer, objects []runtime.Object) error {
	for _, object := range objects {
		_, err := writer.Write([]byte("\n---\n"))
		if err != nil {
			return microerror.Mask(err)
		}

		crdBytes, err := yaml.Marshal(object)
		if err != nil {
			return microerror.Mask(err)
		}

		_, err = writer.Write(crdBytes)
		if err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}

func writeObjectsToFile(filename string, objects []runtime.Object) error {
	writeBuffer, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return microerror.Mask(err)
	}

	defer func() {
		err = writeBuffer.Close()
		if err != nil {
			panic(microerror.JSON(microerror.Mask(err)))
		}
	}()

	return writeObjects(writeBuffer, objects)
}

func writeCRDsToDirectory(outputDirectory string, crds []v1.CustomResourceDefinition) error {
	if len(crds) == 0 {
		return nil
	}

	for _, crd := range crds {
		filename := filepath.Join(outputDirectory, fmt.Sprintf("%s_%s.yaml", crd.Spec.Group, crd.Spec.Names.Plural))
		if err := writeObjectsToFile(filename, []runtime.Object{crd.DeepCopy()}); err != nil {
			return microerror.Mask(err)
		}
	}

	return nil
}
