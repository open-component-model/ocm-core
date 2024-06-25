package sign

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/ocm/commands/misccmds/hash/sign"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/sign"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Sign components or hashes",
	}, verbs.Sign)
	cmd.AddCommand(components.NewCommand(ctx))
	cmd.AddCommand(sign.NewCommand(ctx))
	return cmd
}
