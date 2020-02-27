package v1alpha1

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	goruntime "runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)
	update     = flag.Bool("update", false, "update generated YAMLs")
)

func Test_GenerateYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_ignition.yaml", group),
			resource: NewIgnitionCRD(),
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_ignition.yaml", group, version),
			resource: &Ignition{
				ObjectMeta: v1.ObjectMeta{
					Name: "example",
				},
				TypeMeta: NewIgnitionTypeMeta(),
			},
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	if *update {
		if _, err := os.Stat(docs); os.IsNotExist(err) {
			err = os.Mkdir(docs, 0755)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: generates %s successfully", i, tc.name), func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			directory := filepath.Join(docs, tc.category)
			path := filepath.Join(directory, tc.name)

			if *update {
				if _, err := os.Stat(directory); os.IsNotExist(err) {
					err = os.Mkdir(directory, 0755)
					if err != nil {
						t.Fatal(err)
					}
				}
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
