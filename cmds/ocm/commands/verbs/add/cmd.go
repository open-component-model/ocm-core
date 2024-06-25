package add

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/add"
	references "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/references/add"
	resourceconfig "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/resourceconfig/add"
	resources "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/resources/add"
	routingslips "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/routingslips/add"
	sourceconfig "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/sourceconfig/add"
	sources "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/sources/add"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Add elements to a component repository or component version",
	}, verbs.Add)
	cmd.AddCommand(resourceconfig.NewCommand(ctx))
	cmd.AddCommand(sourceconfig.NewCommand(ctx))

	cmd.AddCommand(resources.NewCommand(ctx))
	cmd.AddCommand(sources.NewCommand(ctx))
	cmd.AddCommand(references.NewCommand(ctx))
	cmd.AddCommand(components.NewCommand(ctx))
	cmd.AddCommand(routingslips.NewCommand(ctx))
	return cmd
}
