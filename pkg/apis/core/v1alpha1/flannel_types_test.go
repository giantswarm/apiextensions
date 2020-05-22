package v1alpha1

import "testing"

func Test_NewFlannelConfigCRD(t *testing.T) {
	crd := NewFlannelConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
