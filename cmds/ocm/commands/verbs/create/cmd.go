package create

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	rsakeypair "ocm.software/ocm-core/cmds/ocm/commands/misccmds/rsakeypair"
	ctf "ocm.software/ocm-core/cmds/ocm/commands/ocicmds/ctf/create"
	comparch "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/componentarchive/create"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Create transport or component archive",
	}, verbs.Create)
	cmd.AddCommand(comparch.NewCommand(ctx))
	cmd.AddCommand(ctf.NewCommand(ctx))
	cmd.AddCommand(rsakeypair.NewCommand(ctx))
	return cmd
}
