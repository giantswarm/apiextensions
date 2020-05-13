module github.com/giantswarm/apiextensions/hack

go 1.14

require sigs.k8s.io/cluster-api v0.2.10

replace (
	k8s.io/api => k8s.io/api v0.16.6
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.16.6
	k8s.io/apimachinery => k8s.io/apimachinery v0.16.6
	k8s.io/apiserver => k8s.io/apiserver v0.16.6
	k8s.io/client-go => k8s.io/client-go v0.16.6
)
