package ocm

import (
	_ "ocm.software/ocm-core/api/datacontext/attrs"
	_ "ocm.software/ocm-core/api/oci"
	_ "ocm.software/ocm-core/api/ocm/compdesc/normalizations"
	_ "ocm.software/ocm-core/api/ocm/compdesc/versions"
	_ "ocm.software/ocm-core/api/ocm/config"
	_ "ocm.software/ocm-core/api/ocm/extensions/accessmethods"
	_ "ocm.software/ocm-core/api/ocm/extensions/blobhandler/config"
	_ "ocm.software/ocm-core/api/ocm/extensions/blobhandler/handlers"
	_ "ocm.software/ocm-core/api/ocm/extensions/digester/digesters"
	_ "ocm.software/ocm-core/api/ocm/extensions/download/config"
	_ "ocm.software/ocm-core/api/ocm/extensions/download/handlers"
	_ "ocm.software/ocm-core/api/ocm/extensions/repositories"
	_ "ocm.software/ocm-core/api/ocm/valuemergehandler/handlers"
)
