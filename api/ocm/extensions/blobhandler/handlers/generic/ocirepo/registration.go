package ocirepo

import (
	"fmt"

	"github.com/mandelsoft/goutils/errors"

	"github.com/open-component-model/ocm/api/oci/artdesc"
	"github.com/open-component-model/ocm/api/ocm/cpi"
	"github.com/open-component-model/ocm/api/ocm/extensions/attrs/ociuploadattr"
	"github.com/open-component-model/ocm/api/utils/listformat"
	"github.com/open-component-model/ocm/api/utils/registrations"
)

type Config = ociuploadattr.Attribute

func init() {
	cpi.RegisterBlobHandlerRegistrationHandler("ocm/ociArtifacts", &RegistrationHandler{})
}

type RegistrationHandler struct{}

var _ cpi.BlobHandlerRegistrationHandler = (*RegistrationHandler)(nil)

func (r *RegistrationHandler) RegisterByName(handler string, ctx cpi.Context, config cpi.BlobHandlerConfig, olist ...cpi.BlobHandlerOption) (bool, error) {
	if handler != "" {
		return true, fmt.Errorf("invalid ociArtifact handler %q", handler)
	}
	if config == nil {
		return true, fmt.Errorf("oci target specification required")
	}
	attr, err := registrations.DecodeConfig[Config](config, ociuploadattr.AttributeType{}.Decode)
	if err != nil {
		return true, errors.Wrapf(err, "blob handler configuration")
	}

	var mimes []string
	opts := cpi.NewBlobHandlerOptions(olist...)
	if opts.MimeType != "" {
		found := false
		for _, a := range artdesc.ArchiveBlobTypes() {
			if a == opts.MimeType {
				found = true
				break
			}
		}
		if !found {
			return true, fmt.Errorf("unexpected type mime type %q for oci blob handler target", opts.MimeType)
		}
		mimes = append(mimes, opts.MimeType)
	} else {
		mimes = artdesc.ArchiveBlobTypes()
	}

	h := NewArtifactHandler(attr)
	for _, m := range mimes {
		opts.MimeType = m
		ctx.BlobHandlers().Register(h, opts)
	}

	return true, nil
}

func (r *RegistrationHandler) GetHandlers(ctx cpi.Context) registrations.HandlerInfos {
	return registrations.NewLeafHandlerInfo("downloading OCI artifacts", `
The <code>ociArtifacts</code> downloader is able to download OCI artifacts
as artifact archive according to the OCI distribution spec.
The following artifact media types are supported:
`+listformat.FormatList("", artdesc.ArchiveBlobTypes()...)+`
By default, it is registered for these mimetypes.

It accepts a config with the following fields:
`+listformat.FormatMapElements("", ociuploadattr.AttributeDescription())+`
Alternatively, a single string value can be given representing an OCI repository
reference.`,
	)
}