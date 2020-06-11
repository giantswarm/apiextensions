package v1alpha1

// +k8s:openapi-gen=true
type CredentialSecret struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
