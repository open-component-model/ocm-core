package compatattr_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"ocm.software/ocm-core/api/config"
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/datacontext"
	"ocm.software/ocm-core/api/oci"
	"ocm.software/ocm-core/api/ocm"
	me "ocm.software/ocm-core/api/ocm/extensions/attrs/compatattr"
	"ocm.software/ocm-core/api/utils/runtime"
)

var _ = Describe("attribute", func() {
	var ctx ocm.Context
	var cfgctx config.Context

	BeforeEach(func() {
		cfgctx = config.WithSharedAttributes(datacontext.New(nil)).New()
		credctx := credentials.WithConfigs(cfgctx).New()
		ocictx := oci.WithCredentials(credctx).New()
		ctx = ocm.WithOCIRepositories(ocictx).New()
	})
	It("local setting", func() {
		Expect(me.Get(ctx)).To(BeFalse())
		Expect(me.Set(ctx, true)).To(Succeed())
		Expect(me.Get(ctx)).To(BeTrue())
	})

	It("global setting", func() {
		Expect(me.Get(cfgctx)).To(BeFalse())
		Expect(me.Set(ctx, true)).To(Succeed())
		Expect(me.Get(ctx)).To(BeTrue())
	})

	It("parses string", func() {
		Expect(me.AttributeType{}.Decode([]byte("true"), runtime.DefaultJSONEncoding)).To(BeTrue())
	})
})
