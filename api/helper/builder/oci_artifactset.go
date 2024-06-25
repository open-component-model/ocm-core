package builder

import (
	"ocm.software/ocm-core/api/oci/extensions/repositories/artifactset"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
)

const T_OCIARTIFACTSET = "artifact set"

////////////////////////////////////////////////////////////////////////////////

func (b *Builder) ArtifactSet(path string, fmt accessio.FileFormat, f ...func()) {
	r, err := artifactset.Open(accessobj.ACC_WRITABLE|accessobj.ACC_CREATE, path, 0o777, fmt, accessio.PathFileSystem(b.FileSystem()))
	b.failOn(err)

	b.configure(&ociNamespace{NamespaceAccess: r, kind: T_OCIARTIFACTSET, annofunc: func(name, value string) {
		r.Annotate(name, value)
	}}, f)
}
