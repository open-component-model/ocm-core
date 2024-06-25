package oci_repository_prepare

import (
	"github.com/mandelsoft/goutils/generics"

	"ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/datacontext/action/handlers"
)

func Execute(hdlrs handlers.Registry, host, repo string, creds common.Properties) (*ActionResult, error) {
	return generics.CastR[*ActionResult](hdlrs.Execute(Spec(host, repo), creds))
}
