module github.com/giantswarm/apiextensions/hack

go 1.14

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/onsi/ginkgo v1.14.0 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	k8s.io/cluster-bootstrap v0.18.5 // indirect
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29 // indirect
	k8s.io/utils v0.0.0-20200731180307-f00132d28269 // indirect
	sigs.k8s.io/cluster-api v0.3.9
	sigs.k8s.io/cluster-api-provider-azure v0.4.7 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.18.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.5
	k8s.io/apiserver => k8s.io/apiserver v0.18.5
	k8s.io/client-go => k8s.io/client-go v0.18.5
	sigs.k8s.io/cluster-api v0.3.9 => github.com/giantswarm/cluster-api v0.3.9-gs
	sigs.k8s.io/cluster-api-provider-azure v0.4.7 => github.com/giantswarm/cluster-api-provider-azure v0.4.7-gs
)
