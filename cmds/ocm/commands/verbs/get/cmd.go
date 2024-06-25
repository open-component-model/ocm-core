package get

import (
	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/api/clictx"
	config "github.com/open-component-model/ocm/cmds/ocm/commands/misccmds/config/get"
	credentials "github.com/open-component-model/ocm/cmds/ocm/commands/misccmds/credentials/get"
	artifacts "github.com/open-component-model/ocm/cmds/ocm/commands/ocicmds/artifacts/get"
	components "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/components/get"
	plugins "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/plugins/get"
	references "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/references/get"
	resources "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/resources/get"
	routingslips "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/routingslips/get"
	sources "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/sources/get"
	"github.com/open-component-model/ocm/cmds/ocm/commands/verbs"
	"github.com/open-component-model/ocm/cmds/ocm/common/utils"
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
