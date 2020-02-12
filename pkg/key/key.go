package key

import (
	"fmt"
	v1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

// KindToCRDNameSingular returns a Kubernetes kind as a CRD singular form.
// For example, AWSCluster would become awscluster.
func KindToCRDNameSingular(kind string) string {
	return strings.ToLower(kind)
}

// KindToCRDNamePlural returns a Kubernetes kind as a CRD plural form.
// For example, AWSCluster would become awsclusters.
func KindToCRDNamePlural(kind string) string {
	return KindToCRDNameSingular(kind) + "s"
}

// KindNames returns an object used for CRD names based on the provided kind.
func KindNames(kind string) v1beta1.CustomResourceDefinitionNames {
	return v1beta1.CustomResourceDefinitionNames{
		Plural:   KindToCRDNamePlural(kind),
		Singular: KindToCRDNameSingular(kind),
		Kind:     kind,
	}
}

// CRDObjectMeta returns a Kubernetes kind to a CRD plural form.
// For example, AWSCluster would become awsclusters.
func CRDObjectMeta(kind, group string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name: fmt.Sprintf("%s.%s", kind, group),
	}
}

func NewCRD(kind, group, version string, scope v1beta1.ResourceScope, schema v1beta1.JSONSchemaProps) *v1beta1.CustomResourceDefinition {
	preserve := false
	return &v1beta1.CustomResourceDefinition{
		TypeMeta: metav1.TypeMeta{
			APIVersion: v1beta1.SchemeGroupVersion.String(),
			Kind:       "CustomResourceDefinition",
		},
		ObjectMeta: CRDObjectMeta(kind, group),
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:                 group,
			Scope:                 scope,
			PreserveUnknownFields: &preserve, // Deprecated, remove when moving to v1.CustomResourceDefinition
			Names:                 KindNames(kind),
			Versions: []v1beta1.CustomResourceDefinitionVersion{
				{
					Name:    version,
					Served:  true,
					Storage: true,
					Schema: &v1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &schema,
					},
				},
			},
		},
	}
}
