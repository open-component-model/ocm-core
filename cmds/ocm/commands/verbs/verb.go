package verbs

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context, name string, short string) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: short,
	}, name)
	return cmd
}
