package common

import (
	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/handlers/elemhdlr"
	"ocm.software/ocm-core/cmds/ocm/common/output"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

func Elem(e interface{}) *compdesc.ComponentReference {
	return e.(*elemhdlr.Object).Element.(*compdesc.ComponentReference)
}

var OptionsFor = elemhdlr.OptionsFor

////////////////////////////////////////////////////////////////////////////////

type TypeHandler struct {
	*elemhdlr.TypeHandler
}

func NewTypeHandler(octx clictx.OCM, opts *output.Options, repo ocm.Repository, session ocm.Session, compspecs []string, hopts ...elemhdlr.Option) (utils.TypeHandler, error) {
	return elemhdlr.NewTypeHandler(octx, opts, repo, session, ocm.KIND_REFERENCE, compspecs, func(access ocm.ComponentVersionAccess) compdesc.ElementAccessor {
		return access.GetDescriptor().References
	}, hopts...)
}
