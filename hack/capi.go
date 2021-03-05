// +build capi

package hack

import (
	_ "github.com/Azure/aad-pod-identity"
	_ "sigs.k8s.io/cluster-api"
	_ "sigs.k8s.io/cluster-api-provider-aws"
	_ "sigs.k8s.io/cluster-api-provider-azure"
	_ "sigs.k8s.io/cluster-api-provider-vsphere"
)
