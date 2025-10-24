// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package container

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/require"
)

func TestScale(t *testing.T) {
	k := types.ServiceConfig{
		Name:  "storagenode",
		Image: "foobar",
	}

	err := scale(&k, "10")
	require.NoError(t, err)

	require.Equal(t, uint64(10), *k.Deploy.Replicas)
}