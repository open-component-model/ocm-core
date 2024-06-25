package internal

import (
	"fmt"

	"github.com/mandelsoft/goutils/errors"

	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/api/utils/errkind"
)

const (
	KIND_COMPONENT        = errkind.KIND_COMPONENT
	KIND_COMPONENTVERSION = "component version"
	KIND_RESOURCE         = "component resource"
	KIND_SOURCE           = "component source"
	KIND_REFERENCE        = compdesc.KIND_REFERENCE
	KIND_REPOSITORYSPEC   = "repository specification"
)

func ErrComponentVersionNotFound(name, version string) error {
	return errors.ErrNotFound(KIND_COMPONENTVERSION, fmt.Sprintf("%s:%s", name, version))
}

func ErrComponentVersionNotFoundWrap(err error, name, version string) error {
	return errors.ErrNotFoundWrap(err, KIND_COMPONENTVERSION, fmt.Sprintf("%s:%s", name, version))
}
