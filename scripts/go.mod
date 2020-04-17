module github.com/giantswarm/apiextensions/scripts

go 1.14

require (
	github.com/markbates/pkger v0.15.1
	k8s.io/code-generator v0.16.6
	sigs.k8s.io/cluster-api v0.2.10
	sigs.k8s.io/controller-tools v0.2.4
	sigs.k8s.io/kustomize/kustomize/v3 v3.5.4
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
