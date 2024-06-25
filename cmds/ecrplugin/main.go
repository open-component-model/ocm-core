package main

import (
	"os"

	"ocm.software/ocm-core/api/ocm/plugin/ppi"
	"ocm.software/ocm-core/api/ocm/plugin/ppi/cmds"
	"ocm.software/ocm-core/api/version"
	"ocm.software/ocm-core/cmds/ecrplugin/actions"
	"ocm.software/ocm-core/cmds/ecrplugin/config"
)

func main() {
	p := ppi.NewPlugin("ecrplugin", version.Get().String())

	p.SetShort("AWS ecr repository creation")
	p.SetLong("plugin assuring the existence of required AWS ECR repositories")
	p.SetConfigParser(config.GetConfig)
	p.SetDescriptorTweaker(func(d ppi.Descriptor) ppi.Descriptor {
		cfg, _ := p.GetConfig()
		if cfg == nil {
			return d
		}
		return config.TweakDescriptor(d, cfg.(*config.Config))
	})

	p.RegisterAction(actions.New())
	err := cmds.NewPluginCommand(p).Execute(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
