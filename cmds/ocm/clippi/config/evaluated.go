package config

import (
	"ocm.software/ocm-core/api/config"
	"ocm.software/ocm-core/api/utils/cobrautils/logopts"
	"ocm.software/ocm-core/cmds/ocm/commands/common/options/keyoption"
)

type EvaluatedOptions struct {
	LogOpts       *logopts.EvaluatedOptions
	Keys          *keyoption.EvaluatedOptions
	ConfigForward config.Config
}
