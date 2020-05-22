package v1alpha1

import "testing"

func Test_NewKVMConfigCRD(t *testing.T) {
	crd := NewKVMConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}

