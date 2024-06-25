// Package jsonv1 provides a normalization which uses schema specific
// normalizations.
// It creates the requested schema for the component descriptor
// and just forwards the normalization to this version.
package jsonv1

import (
	"github.com/mandelsoft/goutils/errors"

	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/api/utils/errkind"
)

const Algorithm = compdesc.JsonNormalisationV1

func init() {
	compdesc.Normalizations.Register(Algorithm, normalization{})
}

type normalization struct{}

func (m normalization) Normalize(cd *compdesc.ComponentDescriptor) ([]byte, error) {
	cv := compdesc.DefaultSchemes[cd.SchemaVersion()]
	if cv == nil {
		if cv == nil {
			return nil, errors.ErrNotSupported(errkind.KIND_SCHEMAVERSION, cd.SchemaVersion())
		}
	}
	v, err := cv.ConvertFrom(cd)
	if err != nil {
		return nil, err
	}
	return v.Normalize(Algorithm)
}
