package v1alpha1

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateSparkYAML(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		filename string
		resource runtime.Object
	}{
		{
			name:     fmt.Sprintf("case 1: %s_%s_spark.yaml is generated successfully", group, version),
			category: "cr",
			filename: fmt.Sprintf("%s_%s_spark.yaml", group, version),
			resource: &Spark{
				TypeMeta: NewSparkTypeMeta(),
				ObjectMeta: v1.ObjectMeta{
					Name: "abc12-master",
					Annotations: map[string]string{
						"giantswarm.io/docs": "https://docs.giantswarm.io/reference/cp-k8s-api/sparks.core.giantswarm.io/",
					},
				},
				Spec: SparkSpec{
					Values: map[string]string{
						"dummy": "test",
					},
				},
			},
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			re := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = re.ReplaceAll(rendered, []byte(""))

			path := filepath.Join(docs, tc.category, tc.filename)
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
