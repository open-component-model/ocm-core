package compdesc

import (
	"ocm.software/ocm-core/api/utils/logging"
)

var (
	REALM  = logging.DefineSubRealm("component descriptor handling", "compdesc")
	Logger = logging.DynamicLogger(REALM)
)
