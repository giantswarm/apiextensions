module github.com/giantswarm/apiextensions/hack

go 1.14

require (
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.8 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	sigs.k8s.io/cluster-api v0.3.6
	sigs.k8s.io/cluster-api-provider-azure v0.4.5 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.18.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.0
	k8s.io/apiserver => k8s.io/apiserver v0.18.0
	k8s.io/client-go => k8s.io/client-go v0.18.0
)
