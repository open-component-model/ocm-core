package bootstrap

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	config "ocm.software/ocm-core/cmds/ocm/commands/toicmds/config/bootstrap"
	_package "ocm.software/ocm-core/cmds/ocm/commands/toicmds/package/bootstrap"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "bootstrap components",
	}, verbs.Bootstrap)
	cmd.AddCommand(_package.NewCommand(ctx))
	cmd.AddCommand(config.NewCommand(ctx))
	return cmd
}
