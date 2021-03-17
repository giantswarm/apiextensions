package v1alpha1

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	goruntime "runtime"
	"sort"
	"testing"
	"time"

	"github.com/go-openapi/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)
	update     = flag.Bool("update", false, "update generated YAMLs")
)

func Test_ReleaseCRValidation(t *testing.T) {
	now := metav1.Now()
	inTenDays := metav1.NewTime(now.Time.Add(240 * time.Hour))

	testCases := []struct {
		name   string
		errors []*errors.Validation
		cr     Release
	}{
		{
			name: "case 0: empty release is invalid",
			cr: Release{
				TypeMeta: NewReleaseTypeMeta(),
			},
			errors: []*errors.Validation{
				{
					Name:  "spec.apps",
					In:    "body",
					Value: "null",
				},
				{
					Name:  "spec.components",
					In:    "body",
					Value: "null",
				},
				{
					Name:  "spec.date",
					In:    "body",
					Value: "null",
				},
				{
					Name:  "spec.state",
					In:    "body",
					Value: "null",
				},
				{
					Name:  "spec.state",
					In:    "body",
					Value: nil,
				},
			},
		},
		{
			name: "case 1: normal release is valid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:         StateActive,
					Date:          &now,
					EndOfLifeDate: &inTenDays,
					Apps: []ReleaseSpecApp{
						{
							Name:             "test-app",
							Version:          "1.0.0",
							ComponentVersion: "2.0.0",
						},
					},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: nil,
		},
		{
			name: "case 2: one component is required",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps: []ReleaseSpecApp{
						{
							Name:             "test-app",
							Version:          "1.0.0",
							ComponentVersion: "2.0.0",
						},
					},
					Components: []ReleaseSpecComponent{},
				},
			},
			errors: []*errors.Validation{
				{
					Name:  "spec.components",
					In:    "body",
					Value: nil,
				},
			},
		},
		{
			name: "case 3: zero apps is valid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: nil,
		},
		{
			name: "case 4: non semver version is invalid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "bad",
						},
					},
				},
			},
			errors: []*errors.Validation{
				{
					Name: "spec.components.version",
					In:   "body",
				},
			},
		},
		{
			name: "case 5: semver with leading v is invalid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "v1.18.0",
						},
					},
				},
			},
			errors: []*errors.Validation{
				{
					Name: "spec.components.version",
					In:   "body",
				},
			},
		},
		{
			name: "case 6: unknown release state is invalid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: "bad",
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: []*errors.Validation{
				{
					Name: "spec.state",
					In:   "body",
				},
			},
		},
		{
			name: "case 7: pre-release component version is valid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0-beta.1",
						},
					},
				},
			},
			errors: nil,
		},
		{
			name: "case 8: non-semver name is invalid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "bad",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: []*errors.Validation{
				{
					Name: "metadata.name",
					In:   "body",
				},
			},
		},
		{
			name: "case 9: semver name without v prefix is invalid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: []*errors.Validation{
				{
					Name: "metadata.name",
					In:   "body",
				},
			},
		},
		{
			name: "case 10: semver name with a suffix is valid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2-beta.10",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: nil,
		},
		{
			name: "case 11: no reference/catalog/releaseOperatorDeploy is valid",
			cr: Release{
				ObjectMeta: metav1.ObjectMeta{
					Name: "v13.1.2",
				},
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State: StateActive,
					Date:  &now,
					Apps:  []ReleaseSpecApp{},
					Components: []ReleaseSpecComponent{
						{
							Name:    "kubernetes",
							Version: "1.18.0",
						},
					},
				},
			},
			errors: nil,
		},
	}
	crd := NewReleaseCRD()

	for _, tc := range testCases {
		for versionCount, crdVersion := range crd.Spec.Versions {
			var v apiextensions.CustomResourceValidation
			err := v1.Convert_v1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crdVersion.Schema, &v, nil)
			if err != nil {
				t.Fatal(err)
			}

			validator, _, err := validation.NewSchemaValidator(&v)
			if err != nil {
				t.Fatal(err)
			}

			opts := []cmp.Option{
				cmpopts.IgnoreUnexported(errors.Validation{}),
			}
			result := validator.Validate(tc.cr)

			if len(result.Errors) != len(tc.errors) {
				t.Fatalf("\n\n%s %s\n", tc.name, cmp.Diff(len(result.Errors), len(tc.errors)))
			}

			var validationErrors []*errors.Validation
			for _, err := range result.Errors {
				validationErrors = append(validationErrors, err.(*errors.Validation))
			}
			if validationErrors == nil {
				continue
			}

			sortErrors(validationErrors)
			sortErrors(tc.errors)

			for i := range result.Errors {
				if !cmp.Equal(validationErrors[i], tc.errors[i], opts...) {
					t.Errorf("\n\n%s - %d %d %s\n", tc.name, versionCount, i, cmp.Diff(validationErrors[i], tc.errors[i], opts...))
				}
			}

		}
	}
}

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

func sortErrors(errors []*errors.Validation) {
	sort.Slice(errors, func(i, j int) bool { return errors[i].Name < errors[j].Name })
}
