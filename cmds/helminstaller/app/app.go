package app

import (
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/ocm/tools/toi/support"
	"ocm.software/ocm-core/cmds/helminstaller/app/driver"
	"ocm.software/ocm-core/cmds/helminstaller/app/driver/helm"
)

func NewCliCommand(ctx clictx.Context, d driver.Driver) *cobra.Command {
	if d == nil {
		d = helm.New()
	}
	return support.NewCLICommand(ctx.OCMContext(), "helmbootstrapper", New(d))
}
