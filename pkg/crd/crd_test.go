package crd

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Test_List(t *testing.T) {
	crdV1, err := ListV1()
	if err != nil {
		t.Fatalf("expected err to be nil: %s", err)
	}

	if crdV1 == nil {
		t.Fatal("expected crd slice to not be nil")
	}

	if len(crdV1) == 0 {
		t.Fatal("expected crd slice to contain at least one item")
	}
}

func Test_Load(t *testing.T) {
	nilPanicMatcher := func(ret interface{}) bool {
		return ret == nil
	}
	notFoundPanicMatcher := func(ret interface{}) bool {
		err, ok := ret.(error)
		return ok && IsNotFound(err)
	}
	testCases := []struct {
		name            string
		inputGroup      string
		inputKind       string
		inputCRDVersion string
		panicMatcher    func(ret interface{}) bool
	}{
		{
			name:            "case 0: v1beta1 CRD loads normally",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "App",
			inputCRDVersion: "v1beta1",
			panicMatcher:    nilPanicMatcher,
		},
		{
			name:            "case 1: non-existent v1beta1 CRD panics with notFoundError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "Bad",
			inputCRDVersion: "v1beta1",
			panicMatcher:    notFoundPanicMatcher,
		},
		{
			name:            "case 2: v1 CRD loads normally",
			inputGroup:      "infrastructure.giantswarm.io",
			inputKind:       "AWSCluster",
			inputCRDVersion: "v1",
			panicMatcher:    nilPanicMatcher,
		},
		{
			name:            "case 3: non-existent v1 CRD panics with notFoundError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "Bad",
			inputCRDVersion: "v1",
			panicMatcher:    notFoundPanicMatcher,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if !tc.panicMatcher(err) {
					t.Errorf("unexpected panic: %#v", err)
				}
			}()
			var crd interface{}
			switch tc.inputCRDVersion {
			case "v1beta1":
				crd = LoadV1Beta1(tc.inputGroup, tc.inputKind)
			case "v1":
				crd = LoadV1(tc.inputGroup, tc.inputKind)
			}
			if crd == nil {
				t.Errorf("nil crd")
			}
		})
	}
}

func Test_ExtractV1CRDVersions(t *testing.T) {
	testCases := []struct {
		name        string
		crd         *v1.CustomResourceDefinition
		versions    []string
		expectedCRD *v1.CustomResourceDefinition
	}{
		{
			name: "case 0: simple case, drop one of two",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha2",
						},
					},
				},
			},
			versions: []string{"v1alpha1"},
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
					},
				},
			},
		},
		{
			name: "case 1: drop one of three",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha2",
						},
						{
							Name: "v1alpha3",
						},
					},
				},
			},
			versions: []string{"v1alpha1", "v1alpha3"},
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha3",
						},
					},
				},
			},
		},
		{
			name: "case 2: return all three",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha2",
						},
						{
							Name: "v1alpha3",
						},
					},
				},
			},
			versions: []string{"v1alpha1", "v1alpha2", "v1alpha3"},
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha2",
						},
						{
							Name: "v1alpha3",
						},
					},
				},
			},
		},
		{
			name: "case 3: drop all three",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name: "v1alpha1",
						},
						{
							Name: "v1alpha2",
						},
						{
							Name: "v1alpha3",
						},
					},
				},
			},
			versions: []string{},
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			crd := ExtractV1CRDVersions(tc.crd, tc.versions...)

			if !cmp.Equal(crd, tc.expectedCRD) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedCRD, crd))
			}
		})
	}
}

func Test_WithStorageVersion(t *testing.T) {
	nilPanicMatcher := func(ret interface{}) bool {
		return ret == nil
	}
	notFoundPanicMatcher := func(ret interface{}) bool {
		err, ok := ret.(error)
		return ok && IsNotFound(err)
	}

	testCases := []struct {
		name         string
		crd          *v1.CustomResourceDefinition
		version      string
		expectedCRD  *v1.CustomResourceDefinition
		panicMatcher func(ret interface{}) bool
	}{
		{
			name: "case 0: switch storage version",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: false,
						},
						{
							Name:    "v1alpha2",
							Storage: true,
						},
					},
				},
			},
			version: "v1alpha1",
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: true,
						},
						{
							Name:    "v1alpha2",
							Storage: false,
						},
					},
				},
			},
			panicMatcher: nilPanicMatcher,
		},
		{
			name: "case 1: keep the set one",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: false,
						},
						{
							Name:    "v1alpha2",
							Storage: true,
						},
					},
				},
			},
			version: "v1alpha2",
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: false,
						},
						{
							Name:    "v1alpha2",
							Storage: true,
						},
					},
				},
			},
			panicMatcher: nilPanicMatcher,
		},
		{
			name: "case 2: panic not found",
			crd: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: false,
						},
						{
							Name:    "v1alpha2",
							Storage: true,
						},
					},
				},
			},
			version: "v1alpha3",
			expectedCRD: &v1.CustomResourceDefinition{
				Spec: v1.CustomResourceDefinitionSpec{
					Versions: []v1.CustomResourceDefinitionVersion{
						{
							Name:    "v1alpha1",
							Storage: false,
						},
						{
							Name:    "v1alpha2",
							Storage: false,
						},
					},
				},
			},
			panicMatcher: notFoundPanicMatcher,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if !tc.panicMatcher(err) {
					t.Errorf("unexpected panic: %#v", err)
				}
			}()

			crd := WithStorageVersion(tc.crd, tc.version)

			if !cmp.Equal(crd, tc.expectedCRD) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.expectedCRD, crd))
			}
		})
	}
}
