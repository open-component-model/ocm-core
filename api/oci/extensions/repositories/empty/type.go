package empty

import (
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/datacontext"
	"ocm.software/ocm-core/api/oci/cpi"
	"ocm.software/ocm-core/api/utils/runtime"
)

const (
	Type   = "Empty"
	TypeV1 = Type + runtime.VersionSeparator + "v1"
)

const ATTR_REPOS = "ocm.software/ocm-core/api/oci/extensions/repositories/empty"

func init() {
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](Type))
	cpi.RegisterRepositoryType(cpi.NewRepositoryType[*RepositorySpec](TypeV1))
}

// RepositorySpec describes an OCI registry interface backed by an oci registry.
type RepositorySpec struct {
	runtime.ObjectVersionedType `json:",inline"`
}

// NewRepositorySpec creates a new RepositorySpec.
func NewRepositorySpec() *RepositorySpec {
	return &RepositorySpec{
		ObjectVersionedType: runtime.NewVersionedTypedObject(Type),
	}
}

func (a *RepositorySpec) GetType() string {
	return Type
}

func (a *RepositorySpec) Name() string {
	return Type
}

func (a *RepositorySpec) UniformRepositorySpec() *cpi.UniformRepositorySpec {
	u := &cpi.UniformRepositorySpec{
		Type: Type,
	}
	return u
}

func (a *RepositorySpec) Repository(ctx cpi.Context, creds credentials.Credentials) (cpi.Repository, error) {
	return ctx.GetAttributes().GetOrCreateAttribute(ATTR_REPOS, func(datacontext.Context) interface{} { return NewRepository(ctx) }).(cpi.Repository), nil
}
