package testhelper

import (
	"io"

	"github.com/mandelsoft/vfs/pkg/vfs"
	"github.com/spf13/cobra"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/config"
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/datacontext"
	"ocm.software/ocm-core/api/helper/builder"
	"ocm.software/ocm-core/api/helper/env"
	"ocm.software/ocm-core/api/oci"
	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/cmds/ocm/app"
)

type CLI struct {
	clictx.Context
}

func NewCLI(ctx clictx.Context) *CLI {
	if ctx == nil {
		ctx = clictx.DefaultContext()
	}
	return &CLI{ctx}
}

func (c *CLI) Execute(args ...string) error {
	cmd, err := app.NewCliCommandForArgs(c, args)
	if err != nil {
		return err
	}
	return cmd.Execute()
}

func (c *CLI) ExecuteModified(mod func(ctx clictx.Context, cmd *cobra.Command), args ...string) error {
	cmd, err := app.NewCliCommandForArgs(c, args, mod)
	if err != nil {
		return err
	}
	return cmd.Execute()
}

type TestEnv struct {
	*builder.Builder
	CLI
}

func NewTestEnv(opts ...env.Option) *TestEnv {
	b := builder.NewBuilder(opts...)
	ctx := clictx.WithOCM(b.OCMContext()).WithSharedAttributes(datacontext.New(nil)).New()
	return &TestEnv{
		Builder: b,
		CLI:     *NewCLI(ctx),
	}
}

func (e *TestEnv) ApplyOption(opts accessio.Options) error {
	return e.Builder.ApplyOption(opts)
}

func (e *TestEnv) ConfigContext() config.Context {
	return e.Builder.ConfigContext()
}

func (e *TestEnv) CredentialsContext() credentials.Context {
	return e.Builder.CredentialsContext()
}

func (e *TestEnv) OCMContext() ocm.Context {
	return e.Builder.OCMContext()
}

func (e *TestEnv) OCIContext() oci.Context {
	return e.Builder.OCIContext()
}

func (e *TestEnv) FileSystem() vfs.FileSystem {
	return e.Builder.FileSystem()
}

func (e *TestEnv) ReadTextFile(path string) (string, error) {
	data, err := e.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e TestEnv) CatchOutput(w io.Writer) *TestEnv {
	e.Context = e.Context.WithStdIO(nil, w, nil)
	return &e
}

func (e TestEnv) CatchErrorOutput(w io.Writer) *TestEnv {
	e.Context = e.Context.WithStdIO(nil, nil, w)
	return &e
}

func (e TestEnv) WithInput(r io.Reader) *TestEnv {
	e.Context = e.Context.WithStdIO(r, nil, nil)
	return &e
}
