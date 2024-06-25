package binary

import (
	"ocm.software/ocm-core/api/utils/cobrautils/flagsets"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/cpi"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/options"
)

func ConfigHandler() flagsets.ConfigOptionTypeSetHandler {
	set := flagsets.NewConfigOptionTypeSetHandler(TYPE, AddConfig, options.DataOption)
	cpi.AddProcessSpecOptionTypes(set)
	return set
}

func AddConfig(opts flagsets.ConfigOptions, config flagsets.Config) error {
	flagsets.AddFieldByOptionP(opts, options.DataOption, config, "data")
	return cpi.AddProcessSpecConfig(opts, config)
}
