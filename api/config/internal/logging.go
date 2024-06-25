package internal

import (
	ocmlog "ocm.software/ocm-core/api/utils/logging"
)

var Realm = ocmlog.DefineSubRealm("configuration management", "config")

var Logger = ocmlog.DynamicLogger(Realm)
