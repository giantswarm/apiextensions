package v1alpha1

import (
	"github.com/go-openapi/errors"
	_ "github.com/go-openapi/validate"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	"sort"
	"testing"
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
					Name:  "spec.status",
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
					Status:  StatusActive,
					Version: "13.1.2",
				},
			},
			errors: nil,
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
			t.Fatalf("\n\n%s\n", cmp.Diff(len(result.Errors), len(tc.errors)))
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
				t.Errorf("\n\n%s\n", cmp.Diff(validationErrors[i], tc.errors[i], opts...))
			}
		}
	}
}
