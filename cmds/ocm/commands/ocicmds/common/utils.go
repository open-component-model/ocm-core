package common

import (
	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/oci"
	"ocm.software/ocm-core/cmds/ocm/common/options"
)

type OptionCompleter interface {
	CompleteWithSession(ctx clictx.OCI, session oci.Session) error
}

func CompleteOptionsWithContext(ctx clictx.Context, session oci.Session) options.OptionsProcessor {
	return func(opt options.Options) error {
		if c, ok := opt.(OptionCompleter); ok {
			return c.CompleteWithSession(ctx.OCI(), session)
		}
		if c, ok := opt.(options.OptionWithCLIContextCompleter); ok {
			return c.Configure(ctx)
		}
		if c, ok := opt.(options.SimpleOptionCompleter); ok {
			return c.Complete()
		}
		return nil
	}
}
