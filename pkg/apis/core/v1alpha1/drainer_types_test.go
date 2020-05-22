package v1alpha1

import "testing"

func Test_NewDrainerConfigCRD(t *testing.T) {
	crd := NewDrainerConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
