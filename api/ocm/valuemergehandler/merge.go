package valuemergehandler

import (
	"ocm.software/ocm-core/api/ocm/cpi"
	"ocm.software/ocm-core/api/ocm/valuemergehandler/hpi"
)

func Merge(ctx cpi.Context, m *Specification, hint hpi.Hint, local Value, inbound *Value) (bool, error) {
	return hpi.Merge(ctx, m, hint, local, inbound)
}

func LabelHint(name string, optversion ...string) hpi.Hint {
	return hpi.LabelHint(name, optversion...)
}
