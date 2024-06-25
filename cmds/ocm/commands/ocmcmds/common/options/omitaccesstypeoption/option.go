package omitaccesstypeoption

import (
	"github.com/spf13/pflag"

	"github.com/open-component-model/ocm/api/ocm/tools/transfer/transferhandler"
	"github.com/open-component-model/ocm/api/ocm/tools/transfer/transferhandler/standard"
	"github.com/open-component-model/ocm/cmds/ocm/common/options"
)

func From(o options.OptionSetProvider) *Option {
	var opt *Option
	o.AsOptionSet().Get(&opt)
	return opt
}

func New() *Option {
	return &Option{}
}

type Option struct {
	standard.TransferOptionsCreator
	Types []string
}

var _ transferhandler.TransferOption = (*Option)(nil)

func (o *Option) AddFlags(fs *pflag.FlagSet) {
	fs.StringSliceVarP(&o.Types, "omit-access-types", "N", nil, "omit by-value transfer for resource types")
}

func (o *Option) Usage() string {
	s := `
It the option <code>--omit-access-types</code> is given, by-value transfer
is omitted completely for the given resource types.
`
	return s
}

func (o *Option) ApplyTransferOption(opts transferhandler.TransferOptions) error {
	if len(o.Types) > 0 {
		return standard.OmitAccessTypes(o.Types...).ApplyTransferOption(opts)
	}
	return nil
}
