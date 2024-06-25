package dockerconfig

import (
	ocmlog "ocm.software/ocm-core/api/utils/logging"
)

var REALM = ocmlog.DefineSubRealm("docker config handling as credential repository", "credentials/dockerconfig")
