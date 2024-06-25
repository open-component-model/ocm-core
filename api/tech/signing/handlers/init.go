package handlers

import (
	_ "github.com/open-component-model/ocm/api/tech/signing/handlers/rsa"
	_ "github.com/open-component-model/ocm/api/tech/signing/handlers/rsa-pss"
	_ "github.com/open-component-model/ocm/api/tech/signing/handlers/rsa-pss-signingservice"
	_ "github.com/open-component-model/ocm/api/tech/signing/handlers/rsa-signingservice"
	_ "github.com/open-component-model/ocm/api/tech/signing/handlers/sigstore"
	_ "github.com/sigstore/cosign/v2/pkg/providers/all"
)
