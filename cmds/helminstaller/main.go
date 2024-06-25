package main

import (
	"os"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/cmds/helminstaller/app"
)

func main() {
	c := app.NewCliCommand(clictx.New(), nil)
	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
