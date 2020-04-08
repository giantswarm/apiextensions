package v1alpha1

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	goruntime "runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)

	// This flag allows to call the tests like
	//
	//   go test -v ./pkg/apis/infrastructure/v1alpha2 -update
	//
	// to create/overwrite the YAML files in /docs/crd and /docs/cr.
	update = flag.Bool("update", false, "update generated YAMLs")
)

func Test_NewAppCatalogCRD(t *testing.T) {
	crd := NewAppCatalogCRD()
	if crd == nil {
		t.Error("AppCatalog CRD was nil.")
	}
}

func Test_GenerateAppCatalogYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_appcatalog.yaml", group),
			resource: NewAppCatalogCRD(),
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_appcatalog.yaml", group, version),
			resource: newAppCatalogExampleCR(),
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: generates %s successfully", i, tc.name), func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			directory := filepath.Join(docs, tc.category)
			path := filepath.Join(directory, tc.name)

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			statusRegex := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = statusRegex.ReplaceAll(rendered, []byte(""))

			if *update {
				err := ioutil.WriteFile(path, rendered, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(rendered, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(rendered)))
			}
		})
	}
}

func newAppCatalogExampleCR() *AppCatalog {
	cr := NewAppCatalogCR()

	cr.Name = "my-playground-catalog"
	cr.Spec = AppCatalogSpec{
		Title:       "My Playground Catalog",
		Description: "A catalog to store all new application packages.",
		Config: AppCatalogSpecConfig{
			ConfigMap: AppCatalogSpecConfigConfigMap{
				Name:      "my-playground-catalog",
				Namespace: "my-namespace",
			},
			Secret: AppCatalogSpecConfigSecret{
				Name:      "my-playground-catalog",
				Namespace: "my-namespace",
			},
		},
		LogoURL: "https://my-org.github.com/logo.png",
		Storage: AppCatalogSpecStorage{
			Type: "helm",
			URL:  "https://my-org.github.com/my-playground-catalog/",
		},
	}

	return cr
}
