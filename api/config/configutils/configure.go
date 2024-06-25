package configutils

import (
	_ "ocm.software/ocm-core/api/datacontext/config"

	"github.com/mandelsoft/vfs/pkg/vfs"

	"ocm.software/ocm-core/api/config"
	utils "ocm.software/ocm-core/api/ocm/ocmutils"
)

func Configure(path string, fss ...vfs.FileSystem) error {
	_, err := utils.Configure(config.DefaultContext(), path, fss...)
	return err
}

func ConfigureContext(ctxp config.ContextProvider, path string, fss ...vfs.FileSystem) error {
	_, err := utils.Configure(ctxp, path, fss...)
	return err
}
