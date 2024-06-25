package downloaderoption

import (
	_ "ocm.software/ocm-core/api/ocm/extensions/download/handlers"

	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/extensions/download"
	"ocm.software/ocm-core/api/utils/listformat"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/options/optutils"
	"ocm.software/ocm-core/cmds/ocm/common/options"
)

type Registration = optutils.Registration

func From(o options.OptionSetProvider) *Option {
	var opt *Option
	o.AsOptionSet().Get(&opt)
	return opt
}

func New(ctx ocm.Context) *Option {
	return &Option{optutils.NewRegistrationOption("downloader", "", "artifact downloader", Usage(ctx))}
}

type Option struct {
	optutils.RegistrationOption
}

func (o *Option) Register(ctx ocm.ContextProvider) error {
	for _, s := range o.Registrations {
		err := download.RegisterHandlerByName(ctx.OCMContext(), s.Name, s.Config,
			download.ForArtifactType(s.ArtifactType), download.ForMimeType(s.MediaType))
		if err != nil {
			return err
		}
	}
	return nil
}

func Usage(ctx ocm.Context) string {
	list := download.For(ctx).GetHandlers(ctx)
	return listformat.FormatListElements("", list) + `

See <CMD>ocm ocm-downloadhandlers</CMD> for further details on using
download handlers.
`
}
