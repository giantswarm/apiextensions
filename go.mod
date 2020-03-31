module github.com/giantswarm/apiextensions

go 1.13

require (
	github.com/go-openapi/errors v0.19.4
	github.com/google/go-cmp v0.3.0
	k8s.io/api v0.16.6
	k8s.io/apiextensions-apiserver v0.16.6
	k8s.io/apimachinery v0.18.0
	k8s.io/client-go v0.16.6
	k8s.io/code-generator v0.16.6
	sigs.k8s.io/yaml v1.1.0
)

replace golang.org/x/tools => golang.org/x/tools v0.0.0-20190821162956-65e3620a7ae7
