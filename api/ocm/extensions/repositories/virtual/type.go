package virtual

import (
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/ocm/cpi"
	"ocm.software/ocm-core/api/ocm/internal"
	"ocm.software/ocm-core/api/utils/runtime"
)

const (
	Type   = "Virtual"
	TypeV1 = Type + runtime.VersionSeparator + "v1"
)

type RepositorySpec struct {
	runtime.ObjectVersionedTypedObject
	Access Access `json:"-"`
}

func NewRepositorySpec(acc Access) *RepositorySpec {
	return &RepositorySpec{
		ObjectVersionedTypedObject: runtime.NewVersionedTypedObject(Type),
		Access:                     acc,
	}
}

func (r RepositorySpec) AsUniformSpec(context internal.Context) *cpi.UniformRepositorySpec {
	return nil
}

func (r *RepositorySpec) Repository(ctx cpi.Context, credentials credentials.Credentials) (internal.Repository, error) {
	return NewRepository(ctx, r.Access), nil
}

var _ cpi.RepositorySpec = (*RepositorySpec)(nil)
