package v1alpha2

import "testing"

func Test_NewAWSControlPlaneCRD(t *testing.T) {
	crd := NewAWSControlPlaneCRD()
	if crd == nil {
		t.Error("AWSControlPlane CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("AWSControlPlane CRD name was empty")
	}
}
