package key

import (
	"fmt"

	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/giantswarm/apiextensions/pkg/crd"
)

const (
	CRDocsAnnotation = "giantswarm.io/docs"
)

func DocumentationLink(crd v1.CustomResourceDefinition) string {
	return fmt.Sprintf("https://docs.giantswarm.io/reference/cp-k8s-api/%s/", crd.Name)
}

func NewCustomResourceMeta(kind metav1.GroupVersionKind, name string, namespace string) metav1.PartialObjectMetadata {
	definition := crd.LoadV1(kind.Group, kind.Kind)
	return metav1.PartialObjectMetadata{
		TypeMeta: metav1.TypeMeta{
			Kind:       kind.Kind,
			APIVersion: schema.GroupVersionKind(kind).GroupVersion().String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				CRDocsAnnotation: DocumentationLink(*definition),
			},
		},
	}
}
