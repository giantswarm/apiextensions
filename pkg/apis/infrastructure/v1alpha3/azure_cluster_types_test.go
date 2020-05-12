package v1alpha3

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
	//   go test -v ./pkg/apis/infrastructure/v1alpha3 -update
	//
	// to create/overwrite the YAML files in /docs/crd and /docs/cr.
	update = flag.Bool("update", false, "update generated YAMLs")
)

func Test_GenerateAzureClusterYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_azurecluster.yaml", group, version),
			resource: newAzureClusterExampleCR(),
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

func newAzureClusterExampleCR() *AzureCluster {
	cr := NewAzureClusterCR()

	cr.Name = "g8kw3"
	cr.Spec = AzureClusterSpec{
		Cluster: AzureClusterSpecCluster{
			Description: "Dev cluster",
			DNS: AzureClusterSpecClusterDNS{
				Domain: "g8s.example.com",
			},
			KubeProxy: AzureClusterSpecClusterKubeProxy{
				ConntrackMaxPerCore: 100000,
			},
			OIDC: AzureClusterSpecClusterOIDC{
				Claims: AzureClusterSpecClusterOIDCClaims{
					Username: "username-field",
					Groups:   "groups-field",
				},
				ClientID:  "some-example-client-id",
				IssuerURL: "https://idp.example.com/",
			},
		},
		Provider: AzureClusterSpecProvider{
			CredentialSecret: AzureClusterSpecProviderCredentialSecret{
				Name:      "example-credential",
				Namespace: "example-namespace",
			},
			Pods: AzureClusterSpecProviderPods{
				CIDRBlock: "10.2.0.0/16",
			},
			Master: AzureClusterSpecProviderMaster{
				AvailabilityZone: "eu-central-1b",
				InstanceType:     "m5.2xlarge",
			},
			Region: "eu-central-1",
		},
	}

	return cr
}
