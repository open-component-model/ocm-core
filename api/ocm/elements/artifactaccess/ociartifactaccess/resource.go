package ociartifactaccess

import (
	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/api/ocm/cpi"
	"ocm.software/ocm-core/api/ocm/elements/artifactaccess/genericaccess"
	access "ocm.software/ocm-core/api/ocm/extensions/accessmethods/ociartifact"
	"ocm.software/ocm-core/api/ocm/extensions/resourcetypes"
)

const TYPE = resourcetypes.OCI_IMAGE

func Access[M any, P compdesc.ArtifactMetaPointer[M]](ctx ocm.Context, meta P, refname string) cpi.ArtifactAccess[M] {
	if meta.GetType() == "" {
		meta.SetType(TYPE)
	}

	spec := access.New(refname)
	// is global access, must work, otherwise there is an error in the lib.
	return genericaccess.MustAccess(ctx, meta, spec)
}

func ResourceAccess(ctx ocm.Context, meta *cpi.ResourceMeta, path string) cpi.ResourceAccess {
	return Access(ctx, meta, path)
}

func SourceAccess(ctx ocm.Context, meta *cpi.SourceMeta, path string) cpi.SourceAccess {
	return Access(ctx, meta, path)
}
