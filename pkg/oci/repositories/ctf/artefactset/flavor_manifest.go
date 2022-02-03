// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package artefactset

import (
	"compress/gzip"
	"sync"

	"github.com/gardener/ocm/pkg/errors"
	"github.com/gardener/ocm/pkg/oci/artdesc"
	"github.com/gardener/ocm/pkg/oci/cpi"
	"github.com/opencontainers/go-digest"
)

type Manifest struct {
	access   ArtefactSetContainer
	manifest *artdesc.Manifest
	lock     sync.RWMutex
	handler  *BlobHandler // not inherited because only blob access should be offered
}

var _ cpi.ManifestAccess = (*Manifest)(nil)

func NewManifest(access ArtefactSetContainer, def ...*artdesc.Manifest) *Manifest {
	var manifest *artdesc.Manifest
	if len(def) == 0 || def[0] == nil {
		manifest = artdesc.NewManifest()
	} else {
		manifest = def[0]
	}
	m := &Manifest{
		access:   access,
		manifest: manifest,
	}
	m.handler = NewBlobHandler(access, m)
	return m
}

func (m *Manifest) IsManifest() bool {
	return true
}

func (m *Manifest) IsIndex() bool {
	return false
}

func (m *Manifest) Manifest() (*artdesc.Manifest, error) {
	return m.manifest, nil
}

func (m *Manifest) Index() (*artdesc.Index, error) {
	return nil, errors.ErrInvalid()
}

func (m *Manifest) Artefact() *artdesc.Artefact {
	a := artdesc.New()
	_ = a.SetManifest(m.manifest)
	return a
}

func (m *Manifest) GetDescriptor() *artdesc.Manifest {
	return m.manifest
}

func (m *Manifest) GetBlobDescriptor(digest digest.Digest) *cpi.Descriptor {
	d := m.manifest.GetBlobDescriptor(digest)
	if d != nil {
		return d
	}
	return m.access.GetBlobDescriptor(digest)
}

func (m *Manifest) GetBlob(digest digest.Digest) (cpi.BlobAccess, error) {
	return m.handler.GetBlob(digest)
}

func (m *Manifest) AddBlob(blob cpi.BlobAccess) error {
	return m.access.AddBlob(blob)
}

func (m *Manifest) AddLayer(blob cpi.BlobAccess, d *artdesc.Descriptor) (int, error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if d == nil {
		d = &artdesc.Descriptor{}
	}
	d.Digest = blob.Digest()
	d.Size = blob.Size()
	if d.MediaType == "" {
		d.MediaType = blob.MimeType()
		if d.MediaType == "" {
			d.MediaType = artdesc.MediaTypeImageLayer
			r, err := blob.Reader()
			if err != nil {
				return -1, err
			}
			defer r.Close()
			zr, err := gzip.NewReader(r)
			if err == nil {
				err = zr.Close()
				if err == nil {
					d.MediaType = artdesc.MediaTypeImageLayerGzip
				}
			}
		}
	}

	err := m.access.AddBlob(blob)
	if err != nil {
		return -1, err
	}

	m.manifest.Layers = append(m.manifest.Layers, *d)
	return len(m.manifest.Layers) - 1, nil
}
