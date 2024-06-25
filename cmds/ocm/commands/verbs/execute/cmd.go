package execute

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	action "ocm.software/ocm-core/cmds/ocm/commands/misccmds/action/execute"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Execute an element.",
	}, verbs.Execute)
	cmd.AddCommand(action.NewCommand(ctx))
	return cmd
}
