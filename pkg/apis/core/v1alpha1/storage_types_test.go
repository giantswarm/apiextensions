package v1alpha1

import "testing"

func Test_NewStorageConfigCRD(t *testing.T) {
	crd := NewStorageConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
