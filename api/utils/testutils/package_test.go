package testutils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	me "ocm.software/ocm-core/api/utils/testutils"
)

var _ = Describe("package tests", func() {
	It("go module name", func() {
		mod := me.Must(me.GetModuleName())
		Expect(mod).To(Equal("ocm.software/ocm-core"))
	})
})
