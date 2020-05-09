package key

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	CRDocsAnnotation         = "giantswarm.io/docs"
	KindAppCatalog           = "AppCatalog"
	KindApp                  = "App"
	KindChart                = "Chart"
	KindETCDBackup           = "ETCDBackup"
	KindAWSClusterConfig     = "AWSClusterConfig"
	KindAzureClusterConfig   = "AzureClusterConfig"
	KindCertConfig           = "CertConfig"
	KindChartConfig          = "ChartConfig"
	KindDrainerConfig        = "DrainerConfig"
	KindFlannelConfig        = "FlannelConfig"
	KindIgnition             = "Ignition"
	KindKVMClusterConfig     = "KVMClusterConfig"
	KindStorageConfig        = "StorageConfig"
	KindMemcachedConfig      = "MemcachedConfig"
	KindAWSCluster           = "AWSCluster"
	KindAWSControlPlane      = "AWSControlPlane"
	KindAWSMachineDeployment = "AWSMachineDeployment"
	KindG8sControlPlane      = "G8sControlPlane"
	KindAWSConfig            = "AWSConfig"
	KindAzureConfig          = "AzureConfig"
	KindKVMConfig            = "KVMConfig"
	KindRelease              = "Release"
	KindAzureTool            = "AzureTool"
	GroupApplication         = "application.giantswarm.io"
)

func DocumentationLink(groupVersionKind metav1.GroupVersionKind) string {
	shortGroup := strings.TrimSuffix(groupVersionKind.Group, ".giantswarm.io")
	return fmt.Sprintf("https://pkg.go.dev/github.com/giantswarm/apiextensions/pkg/apis/%s/%s?tab=doc#%s", shortGroup, groupVersionKind.Version, groupVersionKind.Kind)
}

func NewTypeMeta(kind schema.GroupVersionKind) metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: kind.GroupVersion().String(),
		Kind:       kind.Kind,
	}
}

func NewObjectMeta(groupVersionKind metav1.GroupVersionKind, name, namespace string) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Annotations: map[string]string{
			CRDocsAnnotation: DocumentationLink(groupVersionKind),
		},
		Name:      name,
		Namespace: namespace,
	}
}
