package plugin

import (
	"ocm.software/ocm-core/api/ocm/cpi"
	"ocm.software/ocm-core/api/ocm/extensions/labels/routingslip/spi"
	"ocm.software/ocm-core/api/utils/runtime"
)

type Entry struct {
	runtime.UnstructuredVersionedTypedObject `json:",inline"`
	handler                                  *PluginHandler
}

var _ spi.Entry = &Entry{}

func (s *Entry) Describe(ctx cpi.Context) string {
	return s.handler.Describe(s, ctx)
}

func (s *Entry) Validate(ctx spi.Context) error {
	_, err := s.handler.Validate(s)
	return err
}

func (s *Entry) Handler() *PluginHandler {
	return s.handler
}
