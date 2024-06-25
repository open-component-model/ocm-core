package vault

import (
	ocmlog "ocm.software/ocm-core/api/utils/logging"
)

var (
	REALM = ocmlog.DefineSubRealm("HashiCorp Vault Access", "credentials", "vault")
	log   = ocmlog.DynamicLogger(REALM)
)
