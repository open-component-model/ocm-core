package install

import (
	"github.com/mandelsoft/goutils/errors"

	"ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/ocm"
	metav1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	utils "ocm.software/ocm-core/api/ocm/ocmutils"
	"ocm.software/ocm-core/api/ocm/tools/toi"
	"ocm.software/ocm-core/api/utils/blobaccess/blobaccess"
)

func Execute(p common.Printer, d Driver, name string, rid metav1.Identity, credsrc blobaccess.DataSource, paramsrc blobaccess.DataSource, octx ocm.Context, cv ocm.ComponentVersionAccess, resolver ocm.ComponentVersionResolver) (*OperationResult, error) {
	var creds *Credentials
	var params []byte
	var err error

	if paramsrc != nil {
		params, err = paramsrc.Get()
		if err != nil {
			return nil, errors.Wrapf(err, "parameters")
		}
	}

	if credsrc != nil {
		data, err := credsrc.Get()
		if err == nil {
			creds, err = ParseCredentialSpecification(data, credsrc.Origin())
		}
		if err != nil {
			return nil, errors.Wrapf(err, "credentials")
		}
	}

	ires, _, err := utils.MatchResourceReference(cv, toi.TypeTOIPackage, metav1.NewResourceRef(rid), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "package resource in %s", common.VersionedElementKey(cv).String())
	}

	var spec toi.PackageSpecification

	err = GetResource(ires, &spec)
	if err != nil {
		return nil, errors.ErrInvalidWrap(err, "package spec")
	}
	return ExecuteAction(p, d, name, &spec, creds, params, octx, cv, resolver)
}
