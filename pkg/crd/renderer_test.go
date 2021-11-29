package crd

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getLocalCRDs(t *testing.T) {
	renderer := Renderer{
		LocalCRDDirectory: "../../config/crd",
	}
	crds, err := renderer.getLocalCRDs("common")
	require.Nil(t, err, err)
	require.Len(t, crds, 5)
}
