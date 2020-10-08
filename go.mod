module github.com/giantswarm/apiextensions/v2

go 1.14

require (
	github.com/giantswarm/microerror v0.2.0
	github.com/go-openapi/errors v0.19.4
	github.com/google/go-cmp v0.4.1
	k8s.io/api v0.18.9
	k8s.io/apiextensions-apiserver v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v0.18.9
	sigs.k8s.io/cluster-api v0.3.10
	sigs.k8s.io/yaml v1.2.0
)

replace (
	sigs.k8s.io/cluster-api v0.3.10 => github.com/giantswarm/cluster-api v0.3.10-gs
	sigs.k8s.io/cluster-api v0.3.7 => github.com/giantswarm/cluster-api v0.3.7-gs
)
