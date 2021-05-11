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
	update     = flag.Bool("update", false, "update generated YAMLs")
)

func newReleaseExampleCR() *Release {
	cr := NewReleaseCR()
	cr.Name = "v11.2.0"
	cr.Spec = ReleaseSpec{
		Apps: []ReleaseSpecApp{
			{
				Name:    "cert-exporter",
				Version: "1.2.1",
			},
			{
				Name:    "chart-operator",
				Version: "0.11.4",
			},
			{
				Name:             "coredns",
				ComponentVersion: "1.6.5",
				Version:          "1.1.3",
			},
			{
				Name:             "kube-state-metrics",
				ComponentVersion: "1.9.2",
				Version:          "1.0.2",
			},
			{
				Name:             "metrics-server",
				ComponentVersion: "0.3.3",
				Version:          "1.0.0",
			},
			{
				Name:    "net-exporter",
				Version: "1.6.0",
			},
			{
				Name:             "nginx-ingress-controller",
				ComponentVersion: "0.29.0",
				Version:          "1.5.0",
			},
			{
				Name:             "node-exporter",
				ComponentVersion: "0.18.1",
				Version:          "1.2.0",
			},
		},
		Components: []ReleaseSpecComponent{
			{
				Name:    "app-operator",
				Version: "1.0.0",
			},
			{
				Name:    "cert-operator",
				Version: "0.1.0",
			},
			{
				Name:    "cluster-operator",
				Version: "0.23.1",
			},
			{
				Name:    "flannel-operator",
				Version: "0.2.0",
			},
			{
				Name:    "kvm-operator",
				Version: "3.10.0",
			},
			{
				Name:    "kubernetes",
				Version: "1.16.3",
			},
			{
				Name:    "containerlinux",
				Version: "2247.6.0",
			},
			{
				Name:    "coredns",
				Version: "1.6.5",
			},
			{
				Name:    "calico",
				Version: "3.10.1",
			},
			{
				Name:    "etcd",
				Version: "3.3.17",
			},
		},
		Date:          &metav1.Time{Time: time.Date(2020, 3, 3, 11, 12, 13, 0, time.UTC)},
		EndOfLifeDate: &metav1.Time{Time: time.Date(2020, 10, 3, 0, 0, 0, 0, time.UTC)},
		State:         StateActive,
	}
	return cr
}

func Test_GenerateReleaseYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_release.yaml", group, version),
			resource: newReleaseExampleCR(),
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
