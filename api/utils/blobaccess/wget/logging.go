package wget

import (
	ocmlog "ocm.software/ocm-core/api/utils/logging"
)

var REALM = ocmlog.DefineSubRealm("blob access for wget", "blobaccess/wget")
