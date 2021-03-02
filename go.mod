module github.com/giantswarm/apiextensions/v3

go 1.14

require (
	github.com/giantswarm/microerror v0.3.0
	github.com/go-openapi/errors v0.19.4
	github.com/google/go-cmp v0.5.4
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/cluster-api v0.3.13
	sigs.k8s.io/yaml v1.2.0
)

replace sigs.k8s.io/cluster-api => sigs.k8s.io/cluster-api v0.3.11-0.20210302171319-f7351b165992
