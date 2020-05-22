package v1alpha1

import "testing"

func Test_NewAzureClusterConfigCRD(t *testing.T) {
	crd := NewAzureClusterConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}

