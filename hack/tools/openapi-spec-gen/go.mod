module github.com/giantswarm/apiextensions/hack/tools/openapi-spec-gen

go 1.14

require (
	github.com/giantswarm/apiextensions v0.4.7-0.20200605183009-6034049fbc37
	github.com/go-openapi/spec v0.19.8
	k8s.io/apimachinery v0.16.6
	k8s.io/apiserver v0.16.6
	k8s.io/kube-openapi v0.0.0-20191217135631-a0384dd483d9
)

replace sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v1.0.1
