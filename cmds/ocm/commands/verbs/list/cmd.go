package list

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/list"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "List information about components",
	}, verbs.List)
	cmd.AddCommand(components.NewCommand(ctx))
	return cmd
}
