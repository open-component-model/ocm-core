package hpi

import (
	ocmlog "ocm.software/ocm-core/api/utils/logging"
)

var REALM = ocmlog.DefineSubRealm("value marge handling", "valuemerge")

var Log = ocmlog.DynamicLogger(REALM)
