package routingslips

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/names"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/routingslips/add"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/routingslips/get"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

var Names = names.RoutingSlips

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Commands working on routing slips",
	}, Names...)
	AddCommands(ctx, cmd)
	return cmd
}

func AddCommands(ctx clictx.Context, cmd *cobra.Command) {
	cmd.AddCommand(add.NewCommand(ctx, add.Verb))
	cmd.AddCommand(get.NewCommand(ctx, get.Verb))
}
