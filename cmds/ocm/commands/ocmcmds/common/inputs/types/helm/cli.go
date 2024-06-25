package helm

import (
	"ocm.software/ocm-core/api/utils/cobrautils/flagsets"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/cpi"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/options"
)

func ConfigHandler() flagsets.ConfigOptionTypeSetHandler {
	return cpi.NewMediaFileSpecOptionType(TYPE, AddConfig,
		options.PathOption, options.VersionOption, options.HintOption, options.HelmRepositoryOption)
}

func AddConfig(opts flagsets.ConfigOptions, config flagsets.Config) error {
	if err := cpi.AddPathSpecConfig(opts, config); err != nil {
		return err
	}
	flagsets.AddFieldByOptionP(opts, options.VersionOption, config, "version")
	flagsets.AddFieldByOptionP(opts, options.HintOption, config, "repository")
	flagsets.AddFieldByOptionP(opts, options.HelmRepositoryOption, config, "helmRepository")
	return nil
}
