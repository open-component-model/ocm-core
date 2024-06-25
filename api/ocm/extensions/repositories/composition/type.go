package composition

import (
	"github.com/open-component-model/ocm/api/credentials"
	"github.com/open-component-model/ocm/api/ocm/cpi"
	"github.com/open-component-model/ocm/api/ocm/internal"
	"github.com/open-component-model/ocm/api/utils/runtime"
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