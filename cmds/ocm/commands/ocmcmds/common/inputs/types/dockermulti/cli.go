package dockermulti

import (
	"ocm.software/ocm-core/api/utils/cobrautils/flagsets"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/options"
)

func ConfigHandler() flagsets.ConfigOptionTypeSetHandler {
	return flagsets.NewConfigOptionTypeSetHandler(
		TYPE, AddConfig,
		options.VariantsOption,
		options.HintOption,
	)
}

func AddConfig(opts flagsets.ConfigOptions, config flagsets.Config) error {
	flagsets.AddFieldByOptionP(opts, options.VariantsOption, config, "variants")
	flagsets.AddFieldByOptionP(opts, options.HintOption, config, "repository")
	return nil
}
