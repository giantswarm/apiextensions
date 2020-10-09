package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/apiextensions/v3/pkg/annotation"
	"github.com/giantswarm/apiextensions/v3/pkg/id"
	"github.com/giantswarm/apiextensions/v3/pkg/label"
)

// +k8s:deepcopy-gen=false

type NetworkPoolCRsConfig struct {
	CIDRBlock     string
	NetworkPoolID string
	Namespace     string
	Owner         string
}

// +k8s:deepcopy-gen=false

type NetworkPoolCRs struct {
	NetworkPool *NetworkPool
}

func NewNetworkPoolCRs(config NetworkPoolCRsConfig) (NetworkPoolCRs, error) {
	// Default some essentials in case certain information are not given. E.g.
	// the Tenant NetworkPoolID may be provided by the user.
	{
		if config.NetworkPoolID == "" {
			config.NetworkPoolID = id.Generate()
		}
		if config.Namespace == "" {
			config.Namespace = metav1.NamespaceDefault
		}
	}

	networkPoolCR := newNetworkPoolCR(config)

	crs := NetworkPoolCRs{
		NetworkPool: networkPoolCR,
	}

	return crs, nil
}

func newNetworkPoolCR(c NetworkPoolCRsConfig) *NetworkPool {
	return &NetworkPool{
		TypeMeta: metav1.TypeMeta{
			Kind:       kindNetworkPool,
			APIVersion: SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      c.NetworkPoolID,
			Namespace: c.Namespace,
			Annotations: map[string]string{
				annotation.Docs: "https://docs.giantswarm.io/reference/cp-k8s-api/networkpools.infrastructure.giantswarm.io/",
			},
			Labels: map[string]string{
				label.Organization: c.Owner,
			},
		},
		Spec: NetworkPoolSpec{
			CIDRBlock: c.CIDRBlock,
		},
	}
}
