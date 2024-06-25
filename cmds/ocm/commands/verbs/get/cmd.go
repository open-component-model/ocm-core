package get

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	config "ocm.software/ocm-core/cmds/ocm/commands/misccmds/config/get"
	credentials "ocm.software/ocm-core/cmds/ocm/commands/misccmds/credentials/get"
	artifacts "ocm.software/ocm-core/cmds/ocm/commands/ocicmds/artifacts/get"
	components "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/components/get"
	plugins "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/plugins/get"
	references "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/references/get"
	resources "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/resources/get"
	routingslips "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/routingslips/get"
	sources "ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/sources/get"
	"ocm.software/ocm-core/cmds/ocm/commands/verbs"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

// NewCommand creates a new command.
func NewCommand(ctx clictx.Context) *cobra.Command {
	cmd := utils.MassageCommand(&cobra.Command{
		Short: "Get information about artifacts and components",
	}, verbs.Get)
	cmd.AddCommand(artifacts.NewCommand(ctx))
	cmd.AddCommand(components.NewCommand(ctx))
	cmd.AddCommand(resources.NewCommand(ctx))
	cmd.AddCommand(references.NewCommand(ctx))
	cmd.AddCommand(sources.NewCommand(ctx))
	cmd.AddCommand(credentials.NewCommand(ctx))
	cmd.AddCommand(plugins.NewCommand(ctx))
	cmd.AddCommand(routingslips.NewCommand(ctx))
	cmd.AddCommand(config.NewCommand(ctx))
	return cmd
}
