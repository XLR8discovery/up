// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package container

import (
	"github.com/compose-spec/compose-go/types"
	"github.com/spf13/cobra"

	"xlr8d.io/oss-up/cmd"
	"xlr8d.io/oss-up/pkg/common"
	"xlr8d.io/oss-up/pkg/recipe"
	"xlr8d.io/oss-up/pkg/runtime/runtime"
)

func entryPointCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "local-entrypoint <selector>... local**remote",
		Short: "bind mount entrypoint.sh to use local modifications",
		Args:  cobra.MinimumNArgs(2),
		RunE: cmd.ExecuteOSSUP(func(st recipe.Stack, rt runtime.Runtime, args []string) error {
			// TODO: doesn't look right args is unused
			return cmd.ChangeCompose(st, rt, []string{"satellite-api"}, updateEntryPoint)
		}),
	}
}

func init() {
	cmd.RootCmd.AddCommand(entryPointCmd())
}

// updateEntrypoint sets the entrypoint of the docker image.
func updateEntryPoint(composeService *types.ServiceConfig) error {
	const scriptName = "entrypoint.sh"
	const source = "./" + scriptName
	const target = "/var/lib/oss/entrypoint.sh"

	// Ensure the script exists
	if err := common.IsRegularFile(scriptName); err != nil {
		return err
	}

	// Check if the bind mount already exists before adding.
	for _, volume := range composeService.Volumes {
		if volume.Type == "bind" && volume.Target == target {
			return nil
		}
	}

	composeService.Volumes = append(composeService.Volumes, common.CreateBind(source, target))
	return nil
}