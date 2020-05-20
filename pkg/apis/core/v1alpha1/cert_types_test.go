package v1alpha1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateCertConfigYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_certconfig.yaml", group, version),
			resource: newCertConfigExampleCR(),
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

func newCertConfigExampleCR() *CertConfig {
	cr := NewCertConfigCR()

	cr.Name = "c68pn-prometheus"
	cr.Spec = CertConfigSpec{
		Cert: CertConfigSpecCert{
			AllowBareDomains: false,
			AltNames: []string{
				"api.c68pn.gollum.westeurope.azure.gigantic.io",
			},
			ClusterComponent:    "prometheus",
			ClusterID:           "c68pn",
			CommonName:          "api.c68pn.k8s.gollum.westeurope.azure.gigantic.io",
			DisableRegeneration: false,
			Organizations: []string{
				"giantswarm",
			},
			TTL: "4320h",
		},
		VersionBundle: CertConfigSpecVersionBundle{
			Version: "0.1.0",
		},
	}

	return cr
}
