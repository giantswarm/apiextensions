package v1alpha2

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
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)
	update     = flag.Bool("update", false, "update generated YAMLs")
)

func Test_NewAWSClusterCRD(t *testing.T) {
	crd := NewAWSClusterCRD()
	if crd == nil {
		t.Error("AWSCluster CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("AWSCluster CRD name was empty.")
	}
}

func Test_GenerateAWSClusterYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_awscluster.yaml", group),
			resource: NewAWSClusterCRD(),
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_awscluster.yaml", group, version),
			resource: newAWSClusterExampleCR(),
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

func newAWSClusterExampleCR() *AWSCluster {
	cr := NewAWSClusterCR()
	cr.Name = "g8kw3"
	cr.Spec = AWSClusterSpec{
		Cluster: AWSClusterSpecCluster{
			Description: "Dev cluster",
			DNS:         AWSClusterSpecClusterDNS{},
			OIDC:        AWSClusterSpecClusterOIDC{},
		},
		Provider: AWSClusterSpecProvider{
			CredentialSecret: AWSClusterSpecProviderCredentialSecret{},
			Master:           AWSClusterSpecProviderMaster{},
			Region:           "eu-central-1",
		},
		//Date:  &DeepCopyTime{time.Date(2020, 3, 3, 11, 12, 13, 0, time.UTC)},
	}
	return cr
}
