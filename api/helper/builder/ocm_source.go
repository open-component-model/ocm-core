package builder

import (
	"github.com/mandelsoft/goutils/errors"

	"ocm.software/ocm-core/api/ocm/compdesc"
	"ocm.software/ocm-core/api/utils/blobaccess/blobaccess"
)

type ocmSource struct {
	base

	meta   compdesc.SourceMeta
	access compdesc.AccessSpec
	blob   blobaccess.BlobAccess
}

const T_OCMSOURCE = "source"

func (r *ocmSource) Type() string {
	return T_OCMSOURCE
}

func (r *ocmSource) Set() {
	r.Builder.ocm_src = &r.meta
	r.Builder.ocm_acc = &r.access
	r.Builder.ocm_meta = &r.meta.ElementMeta
	r.Builder.ocm_labels = &r.meta.ElementMeta.Labels
	r.Builder.blob = &r.blob
}

func (r *ocmSource) Close() error {
	switch {
	case r.access != nil:
		return r.Builder.ocm_vers.SetSource(&r.meta, r.access)
	case r.blob != nil:
		return r.Builder.ocm_vers.SetSourceBlob(&r.meta, r.blob, "", nil)
	}
	return errors.New("access or blob required")
}

////////////////////////////////////////////////////////////////////////////////

func (b *Builder) Source(name, vers, typ string, f ...func()) {
	b.expect(b.ocm_vers, T_OCMVERSION)
	r := &ocmSource{}
	r.meta.Name = name
	r.meta.Type = typ
	r.meta.Version = vers
	b.configure(r, f)
}
