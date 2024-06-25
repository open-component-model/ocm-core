package components

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	ocmcomp "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/names"
	toicomp "ocm.software/ocm-core/cmds/ocm/commands/toicmds/package"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

var Names = names.Components

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Commands acting on components",
	}, Names...)
	ocmcomp.AddCommands(ctx, cmd)
	toicomp.AddCommands(ctx, cmd)
	return cmd
}
