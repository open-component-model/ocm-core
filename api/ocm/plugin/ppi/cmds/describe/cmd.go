package describe

import (
	"os"

	"github.com/spf13/cobra"

	common2 "ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/datacontext/action"
	"ocm.software/ocm-core/api/ocm/plugin/common"
	"ocm.software/ocm-core/api/ocm/plugin/ppi"
)

const NAME = "describe"

func New(p ppi.Plugin) *cobra.Command {
	return &cobra.Command{
		Use:   NAME,
		Short: "describe plugin",
		Long:  "Display a detailed description of the capabilities of this OCM plugin.",
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			d := p.Descriptor()
			common.DescribePluginDescriptor(action.DefaultRegistry(), &d, common2.NewPrinter(os.Stdout))
			return nil
		},
	}
}
