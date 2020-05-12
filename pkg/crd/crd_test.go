package crd

import (
	"testing"
)

func Test_List(t *testing.T) {
	crdV1 := List()
	t.Run("case 0: crd slice should not be nil", func(t *testing.T) {
		if crdV1 == nil {
			t.Fatal("expected crd slice to not be nil")
		}
	})
	t.Run("case 1: crd slice should contain at least one item", func(t *testing.T) {
		if len(crdV1) == 0 {
			t.Fatal("expected crd slice to contain at least one item")
		}
	})
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
