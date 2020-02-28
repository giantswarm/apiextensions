package v1alpha2

import "testing"

func Test_NewG8sControlPlaneCRD(t *testing.T) {
	crd := NewG8sControlPlaneCRD()
	if crd == nil {
		t.Error("G8sControlPlane CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("G8sControlPlane CRD name was empty")
	}
}
