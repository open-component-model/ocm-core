package oci

import (
	"context"

	"ocm.software/ocm-core/api/oci/internal"
	"ocm.software/ocm-core/api/utils/blobaccess/blobaccess"
)

const (
	KIND_OCIARTIFACT = internal.KIND_OCIARTIFACT
	KIND_MEDIATYPE   = blobaccess.KIND_MEDIATYPE
	KIND_BLOB        = blobaccess.KIND_BLOB
)

const CONTEXT_TYPE = internal.CONTEXT_TYPE

const CommonTransportFormat = internal.CommonTransportFormat

type (
	Context                          = internal.Context
	ContextProvider                  = internal.ContextProvider
	Repository                       = internal.Repository
	RepositorySpecHandlers           = internal.RepositorySpecHandlers
	RepositorySpecHandler            = internal.RepositorySpecHandler
	UniformRepositorySpec            = internal.UniformRepositorySpec
	RepositoryTypeScheme             = internal.RepositoryTypeScheme
	RepositorySpec                   = internal.RepositorySpec
	RepositoryType                   = internal.RepositoryType
	IntermediateRepositorySpecAspect = internal.IntermediateRepositorySpecAspect
	GenericRepositorySpec            = internal.GenericRepositorySpec
	ArtifactAccess                   = internal.ArtifactAccess
	NamespaceLister                  = internal.NamespaceLister
	NamespaceAccess                  = internal.NamespaceAccess
	ManifestAccess                   = internal.ManifestAccess
	IndexAccess                      = internal.IndexAccess
	BlobAccess                       = internal.BlobAccess
	DataAccess                       = internal.DataAccess
	ConsumerIdentityProvider         = internal.ConsumerIdentityProvider
)

func DefaultContext() internal.Context {
	return internal.DefaultContext
}

func FromContext(ctx context.Context) Context {
	return internal.ForContext(ctx)
}

func FromProvider(p ContextProvider) Context {
	return internal.FromProvider(p)
}

func DefinedForContext(ctx context.Context) (Context, bool) {
	return internal.DefinedForContext(ctx)
}

func IsErrBlobNotFound(err error) bool {
	return blobaccess.IsErrBlobNotFound(err)
}

func ToGenericRepositorySpec(spec RepositorySpec) (*GenericRepositorySpec, error) {
	return internal.ToGenericRepositorySpec(spec)
}
