package hash

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/hash"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Hash and normalization operations",
	}, verbs.Hash)
	cmd.AddCommand(components.NewCommand(ctx))
	return cmd
}
