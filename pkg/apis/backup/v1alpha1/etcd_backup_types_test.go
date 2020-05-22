package v1alpha1

import "testing"

func Test_NewETCDBackupCRD(t *testing.T) {
	crd := NewETCDBackupCRD()
	if crd == nil {
		t.Fatal("expected CRD to not be nil")
	}
}
