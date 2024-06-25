package relativeociref_test

import (
	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "ocm.software/ocm-core/api/helper/builder"
	. "ocm.software/ocm-core/api/oci/testhelper"

	"github.com/mandelsoft/goutils/finalizer"

	v1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	"ocm.software/ocm-core/api/ocm/cpi/accspeccpi"
	"ocm.software/ocm-core/api/ocm/extensions/accessmethods/relativeociref"
	"ocm.software/ocm-core/api/ocm/extensions/repositories/ctf"
	utils "ocm.software/ocm-core/api/ocm/ocmutils"
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/accessobj"
	"ocm.software/ocm-core/api/utils/blobaccess/blobaccess"
)

const (
	OCIPATH = "/tmp/oci"
	OCIHOST = "alias"
)

const (
	COMP     = "acme.org/compo"
	COMPVERS = "v1.0.0"
	RES      = "ref"
)

var _ = Describe("Method", func() {
	var env *Builder

	BeforeEach(func() {
		env = NewBuilder()
	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("accesses artifact", func() {
		var finalize finalizer.Finalizer
		defer Defer(finalize.Finalize)

		env.OCICommonTransport(OCIPATH, accessio.FormatDirectory, func() {
			OCIManifest1(env)
		})
		FakeOCIRepo(env, OCIPATH, OCIHOST)

		env.OCMCommonTransport(OCIPATH, accessio.FormatDirectory, func() {
			env.ComponentVersion(COMP, COMPVERS, func() {
				env.Resource(RES, COMPVERS, "testtyp", v1.LocalRelation, func() {
					env.Access(relativeociref.New(OCINAMESPACE + ":" + OCIVERSION))
				})
			})
		})

		repo := Must(ctf.Open(env, accessobj.ACC_READONLY, OCIPATH, 0, env))
		finalize.Close(repo)
		vers := Must(repo.LookupComponentVersion(COMP, COMPVERS))
		finalize.Close(vers)
		res := Must(vers.GetResourceByIndex(0))
		m := Must(res.AccessMethod())
		finalize.With(func() error {
			return m.Close()
		})
		data := Must(m.Get())
		Expect(len(data)).To(Equal(628))
		Expect(accspeccpi.GetAccessMethodImplementation(m).(blobaccess.DigestSource).Digest().String()).To(Equal("sha256:0c4abdb72cf59cb4b77f4aacb4775f9f546ebc3face189b2224a966c8826ca9f"))
		Expect(utils.GetOCIArtifactRef(env, res)).To(Equal("ocm/value:v2.0"))
	})
})
