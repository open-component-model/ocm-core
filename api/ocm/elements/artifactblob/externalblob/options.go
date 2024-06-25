package externalblob

import (
	"github.com/open-component-model/ocm/api/ocm/cpi"
	"github.com/open-component-model/ocm/api/ocm/elements/artifactblob/api"
)

type (
	Option  = api.Option
	Options = api.Options
)

////////////////////////////////////////////////////////////////////////////////
// General Options

func WithHint(h string) Option {
	return api.WrapHint[Options](h)
}

func WithGlobalAccess(a cpi.AccessSpec) Option {
	return api.WrapGlobalAccess[Options](a)
}