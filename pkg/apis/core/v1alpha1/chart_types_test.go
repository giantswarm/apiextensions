package v1alpha1

import "testing"

func Test_NewChartConfigCRD(t *testing.T) {
	crd := NewChartConfigCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
