package internal

import (
	"encoding/json"
)

// +kubebuilder:validation:Type=number
// +kubebuilder:validation:Format=double
type Float struct {
	Value float64 `json:"-"`
}

// UnmarshalJSON implements the json.Unmarshaller interface.
func (f *Float) UnmarshalJSON(value []byte) error {
	return json.Unmarshal(value, &f.Value)
}

// MarshalJSON implements the json.Marshaller interface.
func (f Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Value)
}

// OpenAPISchemaType is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
//
// See: https://github.com/kubernetes/kube-openapi/tree/master/pkg/generators
func (Float) OpenAPISchemaType() []string { return []string{"number"} }

// OpenAPISchemaFormat is used by the kube-openapi generator when constructing
// the OpenAPI spec of this type.
func (Float) OpenAPISchemaFormat() string { return "double" }
