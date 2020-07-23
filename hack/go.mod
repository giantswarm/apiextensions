module github.com/giantswarm/apiextensions/hack

go 1.14

require (
	sigs.k8s.io/cluster-api v0.3.7
	sigs.k8s.io/cluster-api-provider-azure v0.4.6 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.5
	k8s.io/apiserver => k8s.io/apiserver v0.18.5
	k8s.io/client-go => k8s.io/client-go v0.18.5
	sigs.k8s.io/cluster-api v0.3.7 => github.com/giantswarm/cluster-api v0.3.8-0.20200723145930-f76c9cd8e8d1
)
