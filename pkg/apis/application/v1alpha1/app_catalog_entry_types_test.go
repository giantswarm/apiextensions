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
	"time"

	"github.com/google/go-cmp/cmp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func Test_GenerateAppCatalogEntryYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_appcatalogentry.yaml", group, version),
			resource: newAppCatalogEntryExampleCR(),
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

func newAppCatalogEntryExampleCR() *AppCatalogEntry {
	cr := NewAppCatalogEntryCR()

	rawTime, _ := time.Parse(time.RFC3339, "2020-09-02T09:40:39.223638219Z")
	timeVal := metav1.NewTime(rawTime)

	cr.ObjectMeta = metav1.ObjectMeta{
		Name:      "giantswarm-nginx-ingress-controller-app-1.9.2",
		Namespace: metav1.NamespaceDefault,
	}
	cr.Spec = AppCatalogEntrySpec{
		AppName:    "nginx-ingress-controller-app",
		AppVersion: "v0.35.0",
		Catalog: AppCatalogEntrySpecCatalog{
			Name:      "giantswarm",
			Namespace: "",
		},
		DateCreated: &timeVal,
		DateUpdated: &timeVal,
		Chart: AppCatalogEntrySpecChart{
			APIVersion: "v1",
			Home:       "https://github.com/giantswarm/nginx-ingress-controller-app",
			Icon:       "https://upload.wikimedia.org/wikipedia/commons/nginx-logo.svg",
		},
		Restrictions: &AppCatalogEntrySpecRestrictions{
			ClusterSingleton:    true,
			CompatibleProviders: []Provider{"aws"},
			FixedNamespace:      "giantswarm",
		},

		Version: "1.9.2",
	}

	return cr
}
