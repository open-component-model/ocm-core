package plugin

import (
	"encoding/json"

	"github.com/mandelsoft/goutils/errors"

	"github.com/open-component-model/ocm/api/credentials"
	"github.com/open-component-model/ocm/api/credentials/identity/hostpath"
	"github.com/open-component-model/ocm/api/ocm/cpi"
	"github.com/open-component-model/ocm/api/ocm/plugin"
	"github.com/open-component-model/ocm/api/ocm/plugin/descriptor"
	"github.com/open-component-model/ocm/api/utils/accessio"
)

// pluginHandler delegates storage of blobs to a plugin based handler.
type pluginHandler struct {
	plugin     plugin.Plugin
	name       string
	target     json.RawMessage
	targetinfo *plugin.UploadTargetSpecInfo
}

func New(p plugin.Plugin, name string, target json.RawMessage) (cpi.BlobHandler, error) {
	var err error

	ud := p.GetUploaderDescriptor(name)
	if ud == nil {
		return nil, errors.ErrUnknown(descriptor.KIND_UPLOADER, name, p.Name())
	}

	var info *plugin.UploadTargetSpecInfo
	if target != nil {
		info, err = p.ValidateUploadTarget(name, target)
		if err != nil {
			return nil, err
		}
	}
	return &pluginHandler{
		plugin:     p,
		name:       name,
		target:     target,
		targetinfo: info,
	}, nil
}

func (b *pluginHandler) StoreBlob(blob cpi.BlobAccess, artType, hint string, global cpi.AccessSpec, ctx cpi.StorageContext) (acc cpi.AccessSpec, err error) {
	var creds credentials.Credentials

	if b.targetinfo != nil {
		if len(b.targetinfo.ConsumerId) > 0 {
			creds, err = credentials.CredentialsForConsumer(ctx.GetContext(), b.targetinfo.ConsumerId, hostpath.IdentityMatcher(b.targetinfo.ConsumerId.Type()))
			if err != nil {
				return nil, err
			}
		}
	}

	target := b.target

	if b.target == nil {
		target, err = json.Marshal(ctx.TargetComponentRepository().GetSpecification())
		if err != nil {
			return nil, errors.Wrapf(err, "cannot marshal target repo spec")
		}
	}

	cpi.BlobHandlerLogger(ctx.GetContext()).Debug("plugin blob handler",
		"plugin", b.plugin.Name(),
		"uploader", b.name,
		"arttype", artType,
		"mediatype", blob.MimeType(),
		"hint", hint,
		"target", string(target),
	)

	var creddata json.RawMessage
	if creds != nil {
		creddata, err = json.Marshal(creds.Properties())
		if err != nil {
			return nil, errors.Wrapf(err, "cannot marshal credentials")
		}
	}

	r := accessio.NewOndemandReader(blob)
	defer errors.PropagateError(&err, r.Close)

	return b.plugin.Put(b.name, r, artType, blob.MimeType(), hint, creddata, target)
}