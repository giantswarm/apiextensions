package v1alpha3

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateG8sControlPlaneYAML(t *testing.T) {
	crdGroup := SchemeGroupVersion.Group
	crdKindLower := strings.ToLower(kindG8sControlPlane)

	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_%s.yaml", crdGroup, version, crdKindLower),
			resource: newG8sControlPlaneExampleCR(),
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

func newG8sControlPlaneExampleCR() *G8sControlPlane {
	cr := NewG8sControlPlaneCR()

	cr.Name = "0p8h5"
	cr.Spec = G8sControlPlaneSpec{
		// ClusterNetwork does not occur in our practice, so leaving it empty.
		//ClusterNetwork:    &apiv1alpha3.ClusterNetwork{},
		InfrastructureRef: corev1.ObjectReference{
			APIVersion: "infrastructure.giantswarm.io/v1alpha3",
			Kind:       "AWSControlPlane",
			Name:       "0p8h5",
			Namespace:  "default",
		},
		Replicas: 1,
	}

	return cr
}
