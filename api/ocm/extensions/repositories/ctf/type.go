package ctf

import (
	"ocm.software/ocm-core/api/oci/extensions/repositories/ctf"
	"ocm.software/ocm-core/api/ocm/extensions/repositories/genericocireg"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
)

const Type = ctf.Type

func NewRepositorySpec(acc accessobj.AccessMode, path string, opts ...accessio.Option) (*genericocireg.RepositorySpec, error) {
	spec, err := ctf.NewRepositorySpec(acc, path, opts...)
	if err != nil {
		return nil, err
	}
	return genericocireg.NewRepositorySpec(spec, nil), nil
}
