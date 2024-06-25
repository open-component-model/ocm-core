package install

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	plugins "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/plugins/install"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Install elements.",
	}, verbs.Install)
	cmd.AddCommand(plugins.NewCommand(ctx))
	return cmd
}
