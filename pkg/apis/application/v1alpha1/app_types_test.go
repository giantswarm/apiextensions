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
	//   go test -v ./pkg/apis/application/v1alpha1 -update
	//
	// to create/overwrite the YAML files in /docs/crd and /docs/cr.
	update = flag.Bool("update", false, "update generated YAMLs")
)

func Test_NewAppCRD(t *testing.T) {
	crd := NewAppCRD()
	if crd == nil {
		t.Error("App CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("App CRD name was empty.")
	}
}

func Test_GenerateAppYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_app.yaml", group),
			resource: NewAppCRD(),
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_app.yaml", group, version),
			resource: newAppExampleCR(),
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

func newAppExampleCR() *App {
	cr := NewAppCR()

	cr.Name = "prometheus"
	cr.Spec = AppSpec{
		Name:      "prometheus",
		Namespace: "monitoring",
		Version:   "1.0.0",
		Catalog:   "my-catalog",
		Config: AppSpecConfig{
			ConfigMap: AppSpecConfigConfigMap{
				Name:      "my-configmap",
				Namespace: "monitoring",
			},
			Secret: AppSpecConfigSecret{
				Name:      "my-secret",
				Namespace: "monitoring",
			},
		},
		KubeConfig: AppSpecKubeConfig{
			InCluster: false,
			Context: AppSpecKubeConfigContext{
				Name: "my-context-name",
			},
			Secret: AppSpecKubeConfigSecret{
				Name:      "my-kubeconfig-secret",
				Namespace: "monitoring",
			},
		},
		UserConfig: AppSpecUserConfig{
			ConfigMap: AppSpecUserConfigConfigMap{
				Name:      "my-userconfig-configmap",
				Namespace: "monitoring",
			},
			Secret: AppSpecUserConfigSecret{
				Name:      "my-userconfig-secret",
				Namespace: "monitoring",
			},
		},
	}

	return cr
}
