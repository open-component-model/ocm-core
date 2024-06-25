package composition

import (
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/ocm/cpi"
	"ocm.software/ocm-core/api/ocm/internal"
	"ocm.software/ocm-core/api/utils/runtime"
)

const (
	Type   = "Composition"
	TypeV1 = Type + runtime.VersionSeparator + "v1"
)

func init() {
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](Type, nil))
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](TypeV1, nil))
}

type RepositorySpec struct {
	runtime.ObjectVersionedTypedObject
	Name string `json:"name"`
}

func NewRepositorySpec(name string) *RepositorySpec {
	return &RepositorySpec{
		ObjectVersionedTypedObject: runtime.NewVersionedTypedObject(Type),
		Name:                       name,
	}
}

func (r RepositorySpec) AsUniformSpec(context internal.Context) *cpi.UniformRepositorySpec {
	return nil
}

func (r *RepositorySpec) Repository(ctx cpi.Context, credentials credentials.Credentials) (internal.Repository, error) {
	return NewRepository(ctx, r.Name), nil
}

var _ cpi.RepositorySpec = (*RepositorySpec)(nil)
