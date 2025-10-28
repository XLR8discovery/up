// Copyright (C) 2025 XLR8discovery PBC
// See LICENSE for copying information.

package history

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"xlr8d.io/oss-up/cmd"
	"xlr8d.io/oss-up/pkg/common"
)

func undoCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "undo",
		Short: "revert to a previous version of the generated docker compose file",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			newTemplateBytes, err := common.Store.RestoreLatestVersion()
			if err != nil {
				return err
			}
			if newTemplateBytes == nil {
				return fmt.Errorf("no previous version of the compose file found")
			}
			newTemplate, err := common.LoadComposeFromBytes(newTemplateBytes)
			if err != nil {
				return err
			}
			pwd, _ := os.Getwd()
			err = common.WriteComposeFileNoHistory(pwd, newTemplate)
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func init() {
	cmd.RootCmd.AddCommand(undoCmd())
}
