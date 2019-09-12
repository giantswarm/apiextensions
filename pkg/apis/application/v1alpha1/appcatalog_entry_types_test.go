package v1alpha1

import "testing"

func Test_NewAppCatalogEntryCRD(t *testing.T) {
	crd := NewAppCatalogEntryCRD()
	if crd == nil {
		t.Error("AppCatalogEntry CRD was nil.")
	}
}
