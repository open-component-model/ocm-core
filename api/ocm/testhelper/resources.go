package testhelper

import (
	"github.com/mandelsoft/goutils/testutils"

	"ocm.software/ocm-core/api/helper/builder"
	metav1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	"ocm.software/ocm-core/api/ocm/extensions/digester/digesters/blob"
	"ocm.software/ocm-core/api/tech/signing/hasher/sha256"
	"ocm.software/ocm-core/api/utils/mime"
)

func TextResourceDigestSpec(d string) *metav1.DigestSpec {
	return &metav1.DigestSpec{
		HashAlgorithm:          sha256.Algorithm,
		NormalisationAlgorithm: blob.GenericBlobDigestV1,
		Value:                  d,
	}
}

var Digests = testutils.Substitutions{
	"D_TESTDATA":  D_TESTDATA,
	"D_OTHERDATA": D_OTHERDATA,
}

const S_TESTDATA = "testdata"

const D_TESTDATA = "810ff2fb242a5dee4220f2cb0e6a519891fb67f2f828a6cab4ef8894633b1f50"

var DS_TESTDATA = TextResourceDigestSpec(D_TESTDATA)

func TestDataResource(env *builder.Builder, funcs ...func()) {
	env.Resource("testdata", "", "PlainText", metav1.LocalRelation, func() {
		env.BlobStringData(mime.MIME_TEXT, S_TESTDATA)
		env.Configure(funcs...)
	})
}

const S_OTHERDATA = "otherdata"

const D_OTHERDATA = "54b8007913ec5a907ca69001d59518acfd106f7b02f892eabf9cae3f8b2414b4"

var DS_OTHERDATA = TextResourceDigestSpec(D_OTHERDATA)

func OtherDataResource(env *builder.Builder, funcs ...func()) {
	env.Resource("otherdata", "", "PlainText", metav1.LocalRelation, func() {
		env.BlobStringData(mime.MIME_TEXT, S_OTHERDATA)
		env.Configure(funcs...)
	})
}
