package toi

import (
	logging2 "ocm.software/ocm-core/api/utils/logging"
)

var REALM = logging2.DefineSubRealm("TOI logging", "toi")

var Log = logging2.DynamicLogger(REALM)
