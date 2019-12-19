package v1alpha1

import (
	"github.com/ghodss/yaml"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	kindAppPRDeployer = "AppPRDeployer"
)

const appPRDeployerCRDYAML = `
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: appprdeployers.application.giantswarm.io
spec:
  group: application.giantswarm.io
  scope: Namespaced
  version: v1alpha1
  names:
    kind: AppPrDeployer
    plural: appprdeployers
    singular: appprdeployer
  subresources:
    status: {}
  validation:
    openAPIV3Schema:
      properties:
`

var appPRDeployerCRD *apiextensionsv1beta1.CustomResourceDefinition

func init() {
	err := yaml.Unmarshal([]byte(appPRDeployerCRDYAML), &appPRDeployerCRD)
	if err != nil {
		panic(err)
	}
}

func NewAppPRDeployerCRD() *apiextensionsv1beta1.CustomResourceDefinition {
	return appPRDeployerCRD.DeepCopy()
}

func NewAppPRDeployerTypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		APIVersion: SchemeGroupVersion.String(),
		Kind:       kindAppPRDeployer,
	}
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AppPRDeployer CRs might look something like the following.
//
//    apiVersion: application.giantswarm.io/v1alpha1
//    kind: AppPRDeployer
//    metadata:
//      name: "happa"
//
//    spec:
//      appName: "happa"
//      appCatalog: "control-plane-test"
//      repo: "https://github.com/giantswarm/happa"
//			repoType: "github"
//
//    status:
// 			deployments:
//				- id: 469b5d
//
//					github:
// 						checksPassed: true
// 						commentID: 354177635
// 						commitSHA: 469b5dc57b23b59cfcf25954d8c57ce35e75015a
// 						pullRequest: https://github.com/giantswarm/happa/pull/908
//
// 					appCR:
// 						name: happa-469b5d
//
// 				- id: 32b1331
//
// 					github:
// 						checksPassed: false
// 						commentID: 454579241
// 						commitSHA: 32b13317e803e94fc5a3f3aff649e9ae12c0cf0d
// 						pullRequest: https://github.com/giantswarm/happa/pull/910
// 					appCR:
// 						name: ""  // No App CR created yet since checksPassed is false
//
type AppPRDeployer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              AppPRDeployerSpec   `json:"spec"`
	Status            AppPRDeployerStatus `json:"status" yaml:"status"`
}

type AppPRDeployerSpec struct {
	AppName    string `json:"appName" yaml:"appName"`
	AppCatalog string `json:"appCatalog" yaml:"appCatalog"`
	Repo       string `json:"repo" yaml:"repo"`
	RepoType   string `json:"repoType" yaml:"repoType"`
}

type AppPRDeployerStatus struct {
	Deployments []AppPRDeployerStatusDeployment `json:"deployments"`
}

type AppPRDeployerStatusDeployment struct {
	ID     string
	Github AppPRDeployerStatusDeploymentGithub
	AppCR  AppPRDeployerStatusDeploymentAppCR
}

type AppPRDeployerStatusDeploymentGithub struct {
	ChecksPassed bool   `json:"checksPassed" yaml:"checksPassed"`
	CommentID    string `json:"commentID" yaml:"commentID"`
	CommitSHA    string `json:"commitSHA" yaml:"commitSHA"`
	PullRequest  string `json:"pullRequest" yaml:"pullRequest"`
}

type AppPRDeployerStatusDeploymentAppCR struct {
	Name string `json:"name" yaml:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AppPRDeployerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []AppPRDeployer `json:"items"`
}
