package plugin

import (
	"encoding/json"

	"ocm.software/ocm-core/api/ocm/valuemergehandler"
)

type Config struct {
	json.RawMessage
}

var _ valuemergehandler.Config = (*Config)(nil)

func (c Config) Complete(valuemergehandler.Context) error {
	return nil
}
