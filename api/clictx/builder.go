package clictx

import (
	"context"
	"io"

	"github.com/mandelsoft/vfs/pkg/vfs"

	"ocm.software/ocm-core/api/clictx/internal"
	"ocm.software/ocm-core/api/datacontext"
	"ocm.software/ocm-core/api/ocm"
)

func WithContext(ctx context.Context) internal.Builder {
	return internal.Builder{}.WithContext(ctx)
}

func WithSharedAttributes(ctx datacontext.AttributesContext) internal.Builder {
	return internal.Builder{}.WithSharedAttributes(ctx)
}

func WithOCM(ctx ocm.Context) internal.Builder {
	return internal.Builder{}.WithOCM(ctx)
}

func WithFileSystem(fs vfs.FileSystem) internal.Builder {
	return internal.Builder{}.WithFileSystem(fs)
}

func WithOutput(w io.Writer) internal.Builder {
	return internal.Builder{}.WithOutput(w)
}

func WithErrorOutput(w io.Writer) internal.Builder {
	return internal.Builder{}.WithErrorOutput(w)
}

func WithInput(r io.Reader) internal.Builder {
	return internal.Builder{}.WithInput(r)
}

func New(mode ...datacontext.BuilderMode) internal.Context {
	return internal.Builder{}.New(mode...)
}
