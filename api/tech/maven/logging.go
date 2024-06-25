package maven

import "ocm.software/ocm-core/api/utils/logging"

var REALM = logging.DefineSubRealm("Maven repository", "maven")

var Log = logging.DynamicLogger(REALM)
