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

func NewMeta(groupVersion schema.GroupVersion, kind string, name string, namespace string) (typeMeta metav1.TypeMeta, objectMeta metav1.ObjectMeta) {
	definition := crd.LoadV1(groupVersion.Group, kind)
	typeMeta = metav1.TypeMeta{
		Kind:       kind,
		APIVersion: groupVersion.String(),
	}
	objectMeta = metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Annotations: map[string]string{
			CRDocsAnnotation: DocumentationLink(*definition),
		},
	}
	return
}
