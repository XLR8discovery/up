// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package modify

import (
	"testing"

	"github.com/compose-spec/compose-go/types"
	"github.com/stretchr/testify/require"

	recipe2 "xlr8d.io/oss-up/pkg/recipe"
	"xlr8d.io/oss-up/pkg/runtime/compose"
	"xlr8d.io/oss-up/pkg/runtime/runtime"
)

func TestAddFlag(t *testing.T) {
	dir := t.TempDir()
	st := recipe2.Stack([]recipe2.Recipe{
		{
			Name: "base",
			Add: []*recipe2.Service{
				{
					Command: []string{
						"one", "--flag=sg",
					},
				},
			},
		},
	})

	rt := compose.NewEmptyCompose(dir)
	err := runtime.ApplyRecipes(st, rt, []string{"base"}, 0)
	require.NoError(t, err)

	err = addFlag(st, rt, []string{"base", "nf=2"})
	require.NoError(t, err)

	s := rt.GetServices()[0]
	rawService := s.(*compose.Service)
	err = rawService.TransformRaw(func(config *types.ServiceConfig) error {
		require.Equal(t, types.ShellCommand{"one", "--flag=sg", "--nf=2"}, config.Command)
		return nil
	})
	require.NoError(t, err)

}

func TestRemoveFlag(t *testing.T) {
	dir := t.TempDir()
	st := recipe2.Stack([]recipe2.Recipe{
		{
			Name: "base",
			Add: []*recipe2.Service{
				{
					Command: []string{
						"one", "--flag=sg",
					},
				},
			},
		},
	})

	rt := compose.NewEmptyCompose(dir)
	err := runtime.ApplyRecipes(st, rt, []string{"base"}, 0)
	require.NoError(t, err)

	err = removeFlag(st, rt, []string{"base", "flag"})
	require.NoError(t, err)

	s := rt.GetServices()[0]
	rawService := s.(*compose.Service)
	err = rawService.TransformRaw(func(config *types.ServiceConfig) error {
		require.Equal(t, types.ShellCommand{"one"}, config.Command)
		return nil
	})
	require.NoError(t, err)

}