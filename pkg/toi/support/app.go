// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package support

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/open-component-model/ocm/pkg/toi/install"

	"github.com/open-component-model/ocm/pkg/cobrautils"

	"github.com/open-component-model/ocm/pkg/contexts/ocm"

	common2 "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common"
	"github.com/open-component-model/ocm/pkg/common"
	"github.com/open-component-model/ocm/pkg/contexts/credentials"
	"github.com/open-component-model/ocm/pkg/contexts/datacontext"
	datactg "github.com/open-component-model/ocm/pkg/contexts/datacontext/config"

	"github.com/spf13/pflag"

	"github.com/open-component-model/ocm/pkg/errors"
	"github.com/open-component-model/ocm/pkg/version"

	_ "github.com/open-component-model/ocm/pkg/contexts/clictx/config"
)

type BootstrapperCLIOptions struct {
	ExecutorOptions
	CredentialSettings []string
	Settings           []string
}

func NewCLICommand(ctx ocm.Context, name string, exec func(options *ExecutorOptions) error) *cobra.Command {
	if ctx == nil {
		ctx = ocm.DefaultContext()
	}
	opts := &BootstrapperCLIOptions{
		ExecutorOptions: ExecutorOptions{
			Context: ctx,
		},
	}
	cmd := &cobra.Command{
		Use:                   name + " {<options>} <action> <component version>",
		Short:                 "Bootstrapper using the OCM bootstrap mechanism",
		Version:               version.Get().String(),
		TraverseChildren:      true,
		SilenceUsage:          true,
		DisableFlagsInUseLine: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			action := ""
			if len(args) > 0 {
				action = args[0]
			}
			opts.Action = action
			if len(args) > 1 {
				opts.ComponentVersionName = args[1]
			}
			return opts.Complete()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("This is %s (%s)\n", name, version.Get().String())
			e := &Executor{Completed: true, Options: &opts.ExecutorOptions, Run: exec}
			return e.Execute()
		},
	}
	cobrautils.TweakCommand(cmd, nil)

	cmd.AddCommand(NewVersionCommand())
	opts.AddFlags(cmd.Flags())
	cmd.InitDefaultHelpCmd()

	return cmd
}

func (o *BootstrapperCLIOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.OCMConfig, "ocmconfig", "", "", "ocm configuration file")
	fs.StringArrayVarP(&o.CredentialSettings, "cred", "C", nil, "credential setting")
	fs.StringArrayVarP(&o.Settings, "attribute", "X", nil, "attribute setting")

	fs.StringVarP(&o.Inputs, "inputs", "", "", "input path")
	fs.StringVarP(&o.Outputs, "outputs", "", "", "output path")
	fs.StringVarP(&o.Root, "bootstraproot", "", install.PathTOI, "bootstrapper contract root folder")
	fs.StringVarP(&o.OCMConfig, "config", "", "", "bootstrapper configuration input file")
	fs.StringVarP(&o.Parameters, "parameters", "", "", "bootstrapper parameter input file")
	fs.StringVarP(&o.RepoPath, "ctf", "", "", "bootstrapper transport archive")
}

func (o *BootstrapperCLIOptions) Complete() error {
	if err := o.ExecutorOptions.Complete(); err != nil {
		return fmt.Errorf("unable to complete options: %w", err)
	}

	id := credentials.ConsumerIdentity{}
	attrs := common.Properties{}

	for _, s := range o.CredentialSettings {
		i := strings.Index(s, "=")
		if i < 0 {
			return errors.ErrInvalid("credential setting", s)
		}

		name := s[:i]
		value := s[i+1:]

		if strings.HasPrefix(name, ":") {
			if len(attrs) != 0 {
				o.Context.CredentialsContext().SetCredentialsForConsumer(id, credentials.NewCredentials(attrs))
				id = credentials.ConsumerIdentity{}
				attrs = common.Properties{}
			}
			name = name[1:]
			id[name] = value
		} else {
			attrs[name] = value
		}

		if len(name) == 0 {
			return errors.ErrInvalid("credential setting", s)
		}
	}

	if len(attrs) != 0 {
		o.Context.CredentialsContext().SetCredentialsForConsumer(id, credentials.NewCredentials(attrs))
	} else {
		if len(id) != 0 {
			return errors.Newf("empty credential attribute set for %s", id.String())
		}
	}

	set, err := common2.ParseLabels(o.Settings, "attribute setting")
	if err == nil && len(set) > 0 {
		ctx := o.Context.ConfigContext()
		spec := datactg.New()
		for _, s := range set {
			attr := s.Name
			eff := datacontext.DefaultAttributeScheme.Shortcuts()[attr]
			if eff != "" {
				attr = eff
			}
			err = spec.AddRawAttribute(attr, s.Value)
			if err != nil {
				return errors.Wrapf(err, "attribute %s", s.Name)
			}
		}
		err = ctx.ApplyConfig(spec, "cli")
	}

	if err != nil {
		return fmt.Errorf("unable to parse labels: %w", err)
	}

	return nil
}

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "displays the version",
		Run: func(cmd *cobra.Command, args []string) {
			v := version.Get()
			fmt.Printf("%#v", v)
		},
	}
}
