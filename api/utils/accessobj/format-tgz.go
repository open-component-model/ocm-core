package accessobj

import (
	"ocm.software/ocm-core/api/utils/accessio"
	"ocm.software/ocm-core/api/utils/compression"
)

var FormatTGZ = NewTarHandlerWithCompression(accessio.FormatTGZ, compression.Gzip)

func init() {
	RegisterFormat(FormatTGZ)
}
