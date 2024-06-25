package config

import (
	"context"

	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/plugin/ppi/cmds/command"
	"ocm.software/ocm-core/api/utils/runtime"
)

func init() {
	command.RegisterCommandConfigHandler(&commandHandler{})
}

type commandHandler struct{}

func (c commandHandler) HandleConfig(ctx context.Context, data []byte) (context.Context, error) {
	var err error

	octx := ocm.DefaultContext()
	ctx = octx.BindTo(ctx)
	if len(data) != 0 {
		_, err = octx.ConfigContext().ApplyData(data, runtime.DefaultYAMLEncoding, " cli config")
		// Ugly, enforce configuration update
		octx.GetResolver()
	}
	return ctx, err
}
