package v1alpha2

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

func Test_NewAWSControlPlaneCRD(t *testing.T) {
	crd := NewAWSControlPlaneCRD()
	if crd == nil {
		t.Error("AWSControlPlane CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("AWSControlPlane CRD name was empty")
	}
}

func Test_GenerateAWSControlPlaneYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_awscontrolplane.yaml", group, version),
			resource: newAWSControlPlaneExampleCR(),
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

func newAWSControlPlaneExampleCR() *AWSControlPlane {
	cr := NewAWSControlPlaneCR()

	cr.Name = "ier2s"
	cr.Spec = AWSControlPlaneSpec{
		AvailabilityZones: []string{
			"eu-central-1a",
			"eu-central-1b",
			"eu-central-1c",
		},
		InstanceType: "m4.xlarge",
	}

	return cr
}
