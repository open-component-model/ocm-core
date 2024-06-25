package main

import (
	"os"

	"ocm.software/ocm-core/api/ocm/plugin/ppi"
	"ocm.software/ocm-core/api/ocm/plugin/ppi/cmds"
	"ocm.software/ocm-core/api/version"
	"ocm.software/ocm-core/cmds/demoplugin/accessmethods"
	"ocm.software/ocm-core/cmds/demoplugin/config"
	"ocm.software/ocm-core/cmds/demoplugin/uploaders"
	"ocm.software/ocm-core/cmds/demoplugin/valuesets"
)

func main() {
	p := ppi.NewPlugin("demo", version.Get().String())

	p.SetShort("demo plugin")
	p.SetLong("plugin providing access to temp files and a check routing slip entry.")
	p.SetConfigParser(config.GetConfig)

	p.RegisterAccessMethod(accessmethods.New())
	u := uploaders.New()
	p.RegisterUploader("testArtifact", "", u)
	p.RegisterValueSet(valuesets.New())
	err := cmds.NewPluginCommand(p).Execute(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
}
