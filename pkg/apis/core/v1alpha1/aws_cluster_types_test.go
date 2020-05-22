package v1alpha1

import "testing"

func Test_NewAWSClusterConfigCRD(t *testing.T) {
	crd := NewAWSClusterConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
