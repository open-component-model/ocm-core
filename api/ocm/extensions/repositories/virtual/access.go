package virtual

import (
	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/api/ocm/cpi"
)

type VersionAccess interface {
	GetDescriptor() *compdesc.ComponentDescriptor
	GetBlob(name string) (cpi.DataAccess, error)
	AddBlob(blob cpi.BlobAccess) (string, error)
	Update() error
	Close() error

	IsReadOnly() bool
	SetReadOnly()
	GetInexpensiveContentVersionIdentity(a cpi.AccessSpec) string
}

type Access interface {
	ComponentLister() cpi.ComponentLister

	ExistsComponentVersion(name string, version string) (bool, error)
	ListVersions(comp string) ([]string, error)

	GetComponentVersion(comp, version string) (VersionAccess, error)

	IsReadOnly() bool
	SetReadOnly()
	Close() error
}

type RepositorySpecProvider interface {
	GetSpecification() cpi.RepositorySpec
}
