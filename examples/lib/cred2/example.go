package main

import (
	"fmt"

	"github.com/mandelsoft/goutils/errors"

	"github.com/open-component-model/ocm/examples/lib/helper"
	"github.com/open-component-model/ocm/api/utils/blobaccess"
	"github.com/open-component-model/ocm/api/credentials"
	"github.com/open-component-model/ocm/api/credentials/builtin/oci/identity"
	"github.com/open-component-model/ocm/api/ocm"
	"github.com/open-component-model/ocm/api/ocm/extensions/accessmethods/ociartifact"
	"github.com/open-component-model/ocm/api/ocm/compdesc"
	metav1 "github.com/open-component-model/ocm/api/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/api/ocm/extensions/repositories/ocireg"
	"github.com/open-component-model/ocm/api/ocm/extensions/resourcetypes"
	"github.com/open-component-model/ocm/api/utils/mime"
)

func SimpleWriteWithCredentials() (ferr error) {
	cfg, err := helper.ReadConfig(CFG)
	if err != nil {
		return err
	}

	octx := ocm.DefaultContext()

	octx.CredentialsContext().SetCredentialsForConsumer(
		credentials.NewConsumerIdentity(identity.CONSUMER_TYPE,
			identity.ID_HOSTNAME, "ghcr.io",
			identity.ID_PATHPREFIX, "mandelsoft",
		),
		cfg.GetCredentials(),
	)
	repoSpec := ocireg.NewRepositorySpec(cfg.Repository, nil)

	repo, err := octx.RepositoryForSpec(repoSpec)
	if err != nil {
		return err
	}
	defer repo.Close()

	comp, err := repo.LookupComponent(cfg.Component)
	if err != nil {
		return errors.Wrapf(err, "cannot lookup component %s", cfg.Component)
	}
	defer comp.Close()

	compvers, err := comp.NewVersion(cfg.Version, true)
	if err != nil {
		return errors.Wrapf(err, "cannot create new version %s", cfg.Version)
	}
	defer errors.PropagateError(&ferr, compvers.Close)

	// add provider information
	compvers.GetDescriptor().Provider = metav1.Provider{Name: "mandelsoft"}

	// add a new resource artifact with the local identity `name="test"`.
	err = compvers.SetResourceBlob(
		&compdesc.ResourceMeta{
			ElementMeta: compdesc.ElementMeta{
				Name: "test",
			},
			Type:     resourcetypes.BLOB,
			Relation: metav1.LocalRelation,
		},
		blobaccess.ForString(mime.MIME_TEXT, "testdata"),
		"", nil,
	)
	if err != nil {
		return errors.Wrapf(err, "cannot add resource")
	}

	imageAccess := ociartifact.New("ghcr.io/open-component-model/ocm/ocm.software/toi/installers/helminstaller/helminstaller:0.4.0")
	err = compvers.SetResource(&compdesc.ResourceMeta{
		ElementMeta: compdesc.ElementMeta{
			Name:    "helminstaller",
			Version: "0.4.0",
		},
		Type:     resourcetypes.OCI_IMAGE,
		Relation: metav1.ExternalRelation,
	}, imageAccess)
	if err != nil {
		return errors.Wrapf(err, "cannot add image resource")
	}

	// finally push the new component version
	if err = comp.AddVersion(compvers); err != nil {
		return errors.Wrapf(err, "cannot add new version")
	}
	fmt.Printf("added component %s version %s to %s\n", cfg.Component, cfg.Version, cfg.Repository)
	return nil
}
