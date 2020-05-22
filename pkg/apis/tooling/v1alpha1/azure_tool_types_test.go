package v1alpha1

import "testing"

func Test_NewAzureToolCRD(t *testing.T) {
	crd := NewAzureToolCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
