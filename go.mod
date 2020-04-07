module github.com/giantswarm/apiextensions

go 1.13

require (
	github.com/go-openapi/errors v0.19.4
	github.com/google/go-cmp v0.4.0
	k8s.io/api v0.17.0
	k8s.io/apiextensions-apiserver v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
	k8s.io/code-generator v0.17.0
	k8s.io/gengo v0.0.0-20190822140433-26a664648505
	k8s.io/klog v1.0.0
	sigs.k8s.io/cluster-api v0.2.10
	sigs.k8s.io/controller-tools v0.2.8
	sigs.k8s.io/yaml v1.2.0
)

replace golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7
