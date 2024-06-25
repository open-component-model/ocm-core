package artifactset

import (
	"github.com/mandelsoft/vfs/pkg/vfs"

	"ocm.software/ocm-core/api/oci/cpi"
	"ocm.software/ocm-core/api/utils/accessobj"
)

// NewStateHandler implements the factory interface for the artifact set
// state descriptor handling
// Basically this is an index state.
func NewStateHandler(fs vfs.FileSystem) accessobj.StateHandler {
	return &cpi.IndexStateHandler{}
}
