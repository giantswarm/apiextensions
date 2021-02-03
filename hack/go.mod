module github.com/giantswarm/apiextensions/hack

go 1.14

require (
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
	sigs.k8s.io/cluster-api v0.3.13
	sigs.k8s.io/cluster-api-provider-azure v0.4.11
)

replace (
	k8s.io/api => k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.5
	k8s.io/apiserver => k8s.io/apiserver v0.18.5
	k8s.io/client-go => k8s.io/client-go v0.18.5
	sigs.k8s.io/cluster-api v0.3.13 => github.com/giantswarm/cluster-api v0.3.13-gs
	sigs.k8s.io/cluster-api-provider-azure v0.4.11 => github.com/giantswarm/cluster-api-provider-azure v0.4.12-gsalpha1
)
