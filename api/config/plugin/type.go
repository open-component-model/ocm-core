package plugin

import (
	"ocm.software/ocm-core/api/config/cpi"
	"ocm.software/ocm-core/api/config/internal"
	"ocm.software/ocm-core/api/utils/runtime"
)

var _ cpi.Config = (*Config)(nil)

type Config struct {
	runtime.UnstructuredVersionedTypedObject `json:",inline"`
}

func (c *Config) ApplyTo(context internal.Context, i interface{}) error {
	return nil
}

func New(name string, desc string) cpi.ConfigType {
	return cpi.NewConfigType[*Config](name, desc)
}
