package v1alpha1

import "testing"

func Test_NewMemcachedConfigCRD(t *testing.T) {
	crd := NewMemcachedConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
