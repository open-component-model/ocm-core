package signing_test

import (
	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"

	"ocm.software/ocm-core/api/ocm"
	v1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	"ocm.software/ocm-core/api/ocm/extensions/repositories/composition"
	"ocm.software/ocm-core/api/ocm/extensions/resourcetypes"
	"ocm.software/ocm-core/api/ocm/tools/signing"
	"ocm.software/ocm-core/api/tech/signing/handlers/rsa"
	rsa_pss "ocm.software/ocm-core/api/tech/signing/handlers/rsa-pss"
	"ocm.software/ocm-core/api/tech/signing/signutils"
	"ocm.software/ocm-core/api/utils/blobaccess"
	"ocm.software/ocm-core/api/utils/mime"
)

var _ = Describe("Simple signing handlers", func() {
	Context("", func() {
		ctx := ocm.DefaultContext()

		var cv ocm.ComponentVersionAccess
		var pub signutils.GenericPublicKey
		var priv signutils.GenericPrivateKey

		BeforeEach(func() {
			priv, pub = Must2(rsa.CreateKeyPair())
			cv = composition.NewComponentVersion(ctx, COMPONENTA, VERSION)
			MustBeSuccessful(cv.SetResourceBlob(ocm.NewResourceMeta("blob", resourcetypes.PLAIN_TEXT, v1.LocalRelation), blobaccess.ForString(mime.MIME_TEXT, "test data"), "", nil))
		})

		DescribeTable("rsa handlers", func(kind string) {
			Must(signing.SignComponentVersion(cv, "signature", signing.PrivateKey("signature", priv)))
			Must(signing.VerifyComponentVersion(cv, "signature", signing.PublicKey("signature", pub)))
		},
			Entry("rsa", rsa.Algorithm),
			Entry("rsapss", rsa_pss.Algorithm),
		)
	})
})
