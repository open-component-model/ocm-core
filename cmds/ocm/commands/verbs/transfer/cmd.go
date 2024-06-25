package transfer

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	artifacts "ocm.software/ocm-core/cmds/ocm/commands/ocicmds/artifacts/transfer"
	comparch "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/componentarchive/transfer"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/transfer"
	ctf "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/ctf/transfer"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Transfer artifacts or components",
	}, verbs.Transfer)
	cmd.AddCommand(comparch.NewCommand(ctx))
	cmd.AddCommand(artifacts.NewCommand(ctx))
	cmd.AddCommand(components.NewCommand(ctx))
	cmd.AddCommand(ctf.NewCommand(ctx))

	return cmd
}
