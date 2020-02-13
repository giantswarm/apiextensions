package v1alpha1

import (
	"sort"
	"testing"

	"github.com/go-openapi/errors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
)

func Test_NewReleaseCRD(t *testing.T) {
	crd := NewReleaseCRD()
	if crd == nil {
		t.Error("Release CRD was nil.")
		return
	}
	if crd.Name == "" {
		t.Error("Release CRD name was empty.")
	}
}

func sortErrors(errors []*errors.Validation) {
	sort.Slice(errors, func(i, j int) bool { return errors[i].Name < errors[j].Name })
}

func Test_ReleaseCRValidation(t *testing.T) {
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
					Name:  "spec.state",
					In:    "body",
					Value: nil,
				},
				{
					Name:  "spec.version",
					In:    "body",
					Value: nil,
				},
			},
		},
		{
			name: "case 1: normal release is valid",
			cr: Release{
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   StateActive,
					Version: "13.1.2",
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
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   StateActive,
					Version: "13.1.2",
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
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   StateActive,
					Version: "13.1.2",
					Apps:    []ReleaseSpecApp{},
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
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   StateActive,
					Version: "13.1.2",
					Apps:    []ReleaseSpecApp{},
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
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   StateActive,
					Version: "v13.1.2",
					Apps:    []ReleaseSpecApp{},
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
					Name: "spec.version",
					In:   "body",
				},
				{
					Name: "spec.components.version",
					In:   "body",
				},
			},
		},
		{
			name: "case 6: unexpected release state is invalid",
			cr: Release{
				TypeMeta: NewReleaseTypeMeta(),
				Spec: ReleaseSpec{
					State:   "bad",
					Version: "13.1.2",
					Apps:    []ReleaseSpecApp{},
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
	}
	crd := NewReleaseCRD()

	var v apiextensions.CustomResourceValidation
	err := v1beta1.Convert_v1beta1_CustomResourceValidation_To_apiextensions_CustomResourceValidation(crd.Spec.Versions[0].Schema, &v, nil)
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

	for _, tc := range testCases {
		result := validator.Validate(tc.cr)

		if !cmp.Equal(len(result.Errors), len(tc.errors)) {
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
				t.Errorf("\n\n%s %d %s\n", tc.name, i, cmp.Diff(validationErrors[i], tc.errors[i], opts...))
			}
		}
	}
}
