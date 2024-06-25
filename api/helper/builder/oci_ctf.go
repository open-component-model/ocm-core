package builder

import (
	"ocm.software/ocm-core/api/oci/extensions/repositories/ctf"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
)

const T_OCI_CTF = "oci common transport format"

func (b *Builder) OCICommonTransport(path string, fmt accessio.FileFormat, f ...func()) {
	r, err := ctf.Open(b.OCMContext().OCIContext(), accessobj.ACC_WRITABLE|accessobj.ACC_CREATE, path, 0o777, accessio.PathFileSystem(b.FileSystem()))
	b.failOn(err)
	b.configure(&ociRepository{Repository: r, kind: T_OCI_CTF}, f)
}
