package builder

import (
	metav1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	"ocm.software/ocm-core/api/ocm/extensions/repositories/comparch"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
)

const T_COMPARCH = "component archive"

func (b *Builder) ComponentArchive(path string, fmt accessio.FileFormat, name, vers string, f ...func()) {
	r, err := comparch.Open(b.OCMContext(), accessobj.ACC_WRITABLE|accessobj.ACC_CREATE, path, 0o777, accessio.PathFileSystem(b.FileSystem()))
	b.failOn(err)
	r.SetName(name)
	r.SetVersion(vers)
	r.GetDescriptor().Provider.Name = metav1.ProviderName("ACME")

	b.configure(&ocmVersion{ComponentVersionAccess: r, kind: T_COMPARCH}, f)
}
