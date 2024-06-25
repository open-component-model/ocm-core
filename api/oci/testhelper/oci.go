package testhelper

import (
	. "github.com/onsi/gomega"

	"ocm.software/ocm-core/api/helper/builder"
	"ocm.software/ocm-core/api/oci/extensions/repositories/ctf"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
)

func FakeOCIRepo(env *builder.Builder, path string, host string) string {
	spec, err := ctf.NewRepositorySpec(accessobj.ACC_READONLY, path, accessio.PathFileSystem(env.FileSystem()))
	ExpectWithOffset(1, err).To(Succeed())
	env.OCIContext().SetAlias(host, spec)
	return host + ".alias"
}
