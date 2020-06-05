module github.com/giantswarm/apiextensions/hack/tools/openapi-spec-gen

go 1.14

require (
	github.com/giantswarm/apiextensions v0.4.6
	github.com/go-openapi/spec v0.19.8
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	k8s.io/apimachinery v0.16.6
	k8s.io/apiserver v0.16.6
	k8s.io/kube-openapi v0.0.0-20191217135631-a0384dd483d9
)

replace sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v1.0.1
