module github.com/giantswarm/apiextensions/v3

go 1.16

require (
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/to v0.3.0
	github.com/google/go-cmp v0.5.4
	github.com/google/go-github/v35 v35.2.0
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	k8s.io/api v0.18.19
	k8s.io/apiextensions-apiserver v0.18.19
	k8s.io/apimachinery v0.18.19
	k8s.io/client-go v0.18.19
	sigs.k8s.io/cluster-api v0.3.13
	sigs.k8s.io/yaml v1.2.0
)

replace (
	// v3.3.10 is required by spf13/viper. Can remove this replace when updated.
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible

	github.com/dgrijalva/jwt-go => github.com/form3tech-oss/jwt-go v3.2.1+incompatible

	// Use v1.3.2 of gogo/protobuf to fix nancy alert.
	github.com/gogo/protobuf v1.3.1 => github.com/gogo/protobuf v1.3.2

	github.com/gorilla/websocket v1.4.0 => github.com/gorilla/websocket v1.4.2

	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.13-gs
)
