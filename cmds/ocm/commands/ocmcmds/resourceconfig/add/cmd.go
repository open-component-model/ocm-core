package add

import (
	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/api/clictx"
	"github.com/open-component-model/ocm/api/ocm"
	"github.com/open-component-model/ocm/api/utils/template"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/addhdlrs/rscs"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/names"
	rscadd "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/resources/add"
	"github.com/open-component-model/ocm/cmds/ocm/commands/verbs"
	"github.com/open-component-model/ocm/cmds/ocm/common/utils"
)

var (
	Names = names.ResourceConfig
	Verb  = verbs.Add
)

type Command struct {
	handler common.ResourceSpecHandler
	common.ResourceConfigAdderCommand
}

// NewCommand creates a new ctf command.
func NewCommand(ctx clictx.Context, names ...string) *cobra.Command {
	h := rscs.New()
	return utils.SetupCommand(
		&Command{
			handler:                    h,
			ResourceConfigAdderCommand: common.NewResourceConfigAdderCommand(ctx, rscadd.NewResourceSpecificationsProvider(ctx, ""), h),
		},
		utils.Names(Names, names...)...,
	)
}

func (o *Command) ForName(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "[<options>] <target> {<configfile> | <var>=<value>}",
		Args:  cobra.MinimumNArgs(1),
		Short: "add a resource specification to a resource config file",
		Example: `
$ ocm add resource-configuration resources.yaml --name myresource --type PlainText --input '{ "type": "file", "path": "testdata/testcontent", "mediaType": "text/plain" }'
`,
	}
}

func (o *Command) Long() string {
	return `
Add a resource specification to a resource config file used by <CMD>ocm add resources</CMD>.
` + o.Adder.Description() + ` Elements must follow the resource meta data
description scheme of the component descriptor.

If expressions/templates are used in the specification file an appropriate
templater and the required settings might be required to provide
a correct input validation.

This command accepts additional resource specification files describing the sources
to add to a component version.

` + (&template.Options{}).Usage() +
		inputs.Usage(inputs.DefaultInputTypeScheme) +
		ocm.AccessUsage(o.OCMContext().AccessMethods(), true)
}

func (o *Command) Run() error {
	return o.ProcessResourceDescriptions(o.handler)
}
