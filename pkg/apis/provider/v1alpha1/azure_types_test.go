package v1alpha1

import "testing"

func Test_NewAzureConfigCRD(t *testing.T) {
	crd := NewAzureConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
