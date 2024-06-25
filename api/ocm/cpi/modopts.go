package cpi

import (
	"ocm.software/ocm-core/api/ocm/extensions/attrs/hashattr"
	"ocm.software/ocm-core/api/ocm/extensions/attrs/signingattr"
	"ocm.software/ocm-core/api/ocm/internal"
	"ocm.software/ocm-core/api/tech/signing/hasher/sha256"
)

type (
	ModificationOption  = internal.ModificationOption
	ModificationOptions = internal.ModificationOptions

	BlobModificationOption  = internal.BlobModificationOption
	BlobModificationOptions = internal.BlobModificationOptions

	BlobUploadOption  = internal.BlobUploadOption
	BlobUploadOptions = internal.BlobUploadOptions

	AddVersionOption  = internal.AddVersionOption
	AddVersionOptions = internal.AddVersionOptions
)

////////////////////////////////////////////////////////////////////////////////

func NewAddVersionOptions(list ...AddVersionOption) *AddVersionOptions {
	return internal.NewAddVersionOptions(list...)
}

// Overwrite enabled the overwrite mode for adding a component version.
func Overwrite(flag ...bool) AddVersionOption {
	return internal.Overwrite(flag...)
}

////////////////////////////////////////////////////////////////////////////////

func NewBlobModificationOptions(list ...BlobModificationOption) *BlobModificationOptions {
	return internal.NewBlobModificationOptions(list...)
}

////////////////////////////////////////////////////////////////////////////////

func NewBlobUploadOptions(list ...BlobUploadOption) *BlobUploadOptions {
	return internal.NewBlobUploadOptions(list...)
}

func UseBlobHandlers(h BlobHandlerProvider) internal.BlobOptionImpl {
	return internal.UseBlobHandlers(h)
}

////////////////////////////////////////////////////////////////////////////////

func NewModificationOptions(list ...ModificationOption) *ModificationOptions {
	return internal.NewModificationOptions(list...)
}

func ModifyResource(flag ...bool) internal.ModOptionImpl {
	return internal.ModifyResource(flag...)
}

func AcceptExistentDigests(flag ...bool) internal.ModOptionImpl {
	return internal.AcceptExistentDigests(flag...)
}

func WithDefaultHashAlgorithm(algo ...string) internal.ModOptionImpl {
	return internal.WithDefaultHashAlgorithm(algo...)
}

func WithHasherProvider(prov HasherProvider) internal.ModOptionImpl {
	return internal.WithHasherProvider(prov)
}

func SkipVerify(flag ...bool) internal.ModOptionImpl {
	return internal.SkipVerify(flag...)
}

// SkipDigest disables digest creation if enabled.
//
// Deprecated: for legacy code, only.
func SkipDigest(flag ...bool) internal.ModOptionImpl {
	return internal.SkipDigest(flag...)
}

///////////////////////////////////////////////////////

func CompleteModificationOptions(ctx ContextProvider, m *ModificationOptions) {
	attr := hashattr.Get(ctx.OCMContext())
	if m.DefaultHashAlgorithm == "" {
		m.DefaultHashAlgorithm = attr.DefaultHasher
	}
	if m.DefaultHashAlgorithm == "" {
		m.DefaultHashAlgorithm = sha256.Algorithm
	}
	if m.HasherProvider == nil {
		m.HasherProvider = signingattr.Get(ctx.OCMContext())
	}
}
