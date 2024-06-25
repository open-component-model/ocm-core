package composition_test

import (
	"runtime"
	"time"

	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ocm.software/ocm-core/api/datacontext"
	"ocm.software/ocm-core/api/ocm"
	me "ocm.software/ocm-core/api/ocm/extensions/repositories/composition"
	"ocm.software/ocm-core/api/utils/refmgmt"
)

var _ = Describe("repository", func() {
	It("finalizes with context", func() {
		ctx := ocm.New(datacontext.MODE_EXTENDED)

		repo := me.NewRepository(ctx, "test")
		MustBeSuccessful(repo.Close())

		ctx = nil
		runtime.GC()
		time.Sleep(time.Second)
		runtime.GC()
		time.Sleep(time.Second)
		Expect(refmgmt.ReferenceCount(repo)).To(Equal(0))
	})
})
