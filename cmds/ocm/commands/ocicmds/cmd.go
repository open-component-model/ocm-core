package ocicmds

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/ocm/commands/ocicmds/artifacts"
	"ocm.software/ocm-core/cmds/ocm/commands/ocicmds/ctf"
	"ocm.software/ocm-core/cmds/ocm/commands/ocicmds/tags"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
	topicocirefs "ocm.software/ocm-core/cmds/ocm/topics/oci/refs"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Dedicated command flavors for the OCI layer",
	}, "oci")
	cmd.AddCommand(artifacts.NewCommand(ctx))
	cmd.AddCommand(ctf.NewCommand(ctx))
	cmd.AddCommand(tags.NewCommand(ctx))

	cmd.AddCommand(topicocirefs.New(ctx))
	return cmd
}
