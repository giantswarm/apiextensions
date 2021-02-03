package v1alpha1

import "testing"

func Test_Config_Spec_Status_App(t *testing.T) {
	t.Log("ensure ConfigSpecApp ConfigStatusApp have identical fields")
	specApp := ConfigSpecApp{}
	_ = ConfigStatusApp(specApp)
}
