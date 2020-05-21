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

func NewTypeMeta(kind metav1.GroupVersionKind) metav1.TypeMeta {
	skind := schema.GroupVersionKind(kind)
	return metav1.TypeMeta{
		APIVersion: skind.GroupVersion().String(),
		Kind:       kind.Kind,
	}
}

func NewObjectMeta(kind metav1.GroupVersionKind) metav1.ObjectMeta {
	definition := crd.LoadV1(kind.Group, kind.Kind)
	return metav1.ObjectMeta{
		Annotations: map[string]string{
			CRDocsAnnotation: DocumentationLink(*definition),
		},
	}
}
