package controller

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/ocm/commands/controllercmds/install"
	"ocm.software/ocm-core/cmds/ocm/commands/controllercmds/names"
	"ocm.software/ocm-core/cmds/ocm/commands/controllercmds/uninstall"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Commands acting on the ocm-controller",
	}, names.Controller...)
	cmd.AddCommand(install.NewCommand(ctx, install.Verb))
	cmd.AddCommand(uninstall.NewCommand(ctx, uninstall.Verb))
	return cmd
}
