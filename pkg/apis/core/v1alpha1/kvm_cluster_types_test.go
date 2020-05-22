package v1alpha1

import "testing"

func Test_NewKVMClusterConfigCRD(t *testing.T) {
	crd := NewKVMClusterConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
