package v1alpha2

import (
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v2/pkg/crd"
)

const (
	kindNetworkPool              = "NetworkPool"
)

func NewNetworkPoolCRD() *v1.CustomResourceDefinition {
	return crd.LoadV1(group, kindNetworkPool)
}

func NewNetworkPoolTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindNetworkPool,
	}
}
