package v1alpha1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateAppYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
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
				err := ioutil.WriteFile(path, rendered, 0644) // nolint
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

	cr.ObjectMeta = metav1.ObjectMeta{
		Name:      "prometheus",
		Namespace: "default",
	}
	cr.Spec = AppSpec{
		Name:             "prometheus",
		Namespace:        "monitoring",
		Version:          "1.0.1",
		Catalog:          "my-playground-catalog",
		CatalogNamespace: "giantswarm",
		Config: AppSpecConfig{
			ConfigMap: AppSpecConfigConfigMap{
				Name:      "f2def-cluster-values",
				Namespace: "f2def",
			},
			Secret: AppSpecConfigSecret{
				Name:      "f2def-cluster-values",
				Namespace: "f2def",
			},
		},
		Install: AppSpecInstall{
			SkipCRDs: true,
		},
		KubeConfig: AppSpecKubeConfig{
			InCluster: false,
			Context: AppSpecKubeConfigContext{
				Name: "f2def",
			},
			Secret: AppSpecKubeConfigSecret{
				Name:      "f2def-kubeconfig",
				Namespace: "f2def",
			},
		},
		NamespaceConfig: AppSpecNamespaceConfig{
			Annotations: map[string]string{
				"linkerd.io/inject": "enabled",
			},
		},
		UserConfig: AppSpecUserConfig{
			ConfigMap: AppSpecUserConfigConfigMap{
				Name:      "prometheus-user-values",
				Namespace: "f2def",
			},
			Secret: AppSpecUserConfigSecret{
				Name:      "prometheus-user-values",
				Namespace: "f2def",
			},
		},
	}

	return cr
}
