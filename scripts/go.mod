module github.com/giantswarm/apiextensions/scripts

go 1.14

require (
	github.com/Azure/go-autorest v11.1.2+incompatible // indirect
	github.com/gophercloud/gophercloud v0.3.0 // indirect
	github.com/markbates/pkger v0.15.1
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	gopkg.in/yaml.v1 v1.0.0-20140924161607-9f9df34309c0 // indirect
	k8s.io/code-generator v0.17.2
	sigs.k8s.io/cluster-api v0.3.5
	sigs.k8s.io/controller-tools v0.2.4
	sigs.k8s.io/kustomize/kustomize/v3 v3.5.4
	sigs.k8s.io/testing_frameworks v0.1.2-0.20190130140139-57f07443c2d4 // indirect
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
