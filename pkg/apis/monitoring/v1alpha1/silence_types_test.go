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
	corev1 "k8s.io/api/core/v1"
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

func Test_GenerateSilenceYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_silence.yaml", group, version),
			resource: newSilenceExampleCR(),
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

func newSilenceExampleCR() *Silence {
	cr := NewSilenceCR()

	cr.ObjectMeta = metav1.ObjectMeta{
		Name:      "custom-silence",
		Namespace: corev1.NamespaceDefault,
	}
	cr.Spec = SilenceSpec{
		TargetTags: []TargetTag{
			{
				Name:  "installation",
				Value: "my_installation_name",
			},
		},
		Matchers: []Matcher{
			{
				IsRegex: false,
				Name:    "cluster_id",
				Value:   "my_cluster_id",
			},
		},
	}

	return cr
}
