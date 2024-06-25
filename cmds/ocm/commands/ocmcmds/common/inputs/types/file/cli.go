package file

import (
	"ocm.software/ocm-core/api/utils/cobrautils/flagsets"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs/cpi"
)

func ConfigHandler() flagsets.ConfigOptionTypeSetHandler {
	return cpi.NewMediaFileSpecOptionType(TYPE, AddConfig)
}

func AddConfig(opts flagsets.ConfigOptions, config flagsets.Config) error {
	return cpi.AddMediaFileSpecConfig(opts, config)
}
