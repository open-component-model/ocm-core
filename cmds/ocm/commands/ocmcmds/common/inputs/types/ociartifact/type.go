package ociartifact

import (
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/inputs"
)

const (
	TYPE        = "ociArtifact"
	LEGACY_TYPE = "ociImage"
)

func init() {
	inputs.DefaultInputTypeScheme.Register(inputs.NewInputType(TYPE, &Spec{}, usage, ConfigHandler()))
	inputs.DefaultInputTypeScheme.Register(inputs.NewInputType(LEGACY_TYPE, &Spec{}, legacy_usage, ConfigHandler()))
}

const legacy_usage = `
DEPRECATED: This type is deprecated, please use ` + TYPE + ` instead.
`

const usage = `
The path must denote an OCI image reference.

This blob type specification supports the following fields: 
- **<code>path</code>** *string*

  This REQUIRED property describes the OVI image reference of the image to
  import.

- **<code>repository</code>** *string*

  This OPTIONAL property can be used to specify the repository hint for the
  generated local artifact access. It is prefixed by the component name if
  it does not start with slash "/".

- **<code>platforms</code>** *[]string*

  This OPTIONAL property can be used to filter index artifacts to include
  only images for dedicated operating systems/architectures.
  Elements must meet the syntax [&lt;os>]/[&lt;architecture>].
`
