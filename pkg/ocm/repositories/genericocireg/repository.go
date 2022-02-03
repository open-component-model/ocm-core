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

package genericocireg

import (
	"github.com/gardener/ocm/pkg/oci"
	"github.com/gardener/ocm/pkg/oci/core"
	"github.com/gardener/ocm/pkg/ocm/compdesc"
	"github.com/gardener/ocm/pkg/ocm/cpi"
)

type Repository struct {
	ctx     cpi.Context
	ocirepo oci.Repository
}

var _ cpi.Repository = &Repository{}

func NewRepository(ctx cpi.Context, ocirepo oci.Repository) (cpi.Repository, error) {
	repo := &Repository{
		ctx:     ctx,
		ocirepo: ocirepo,
	}
	_ = repo
	return repo, nil
}

func (r Repository) ExistsArtefact(name string, version string) (bool, error) {
	panic("implement me")
}

func (r Repository) LookupArtefact(name string, version string) (oci.ArtefactAccess, error) {
	panic("implement me")
}

func (r Repository) LookupNamespace(name string) (core.NamespaceAccess, error) {
	panic("implement me")
}

func (r Repository) Close() error {
	panic("implement me")
}

func (r Repository) GetContext() cpi.Context {
	panic("implement me")
}

func (r Repository) GetSpecification() cpi.RepositorySpec {
	panic("implement me")
}

func (r Repository) ExistsComponent(name string, version string) (bool, error) {
	panic("implement me")
}

func (r Repository) LookupComponent(name string, version string) (cpi.ComponentAccess, error) {
	panic("implement me")
}

func (r Repository) ComposeComponent(name string, version string) (cpi.ComponentComposer, error) {
	panic("implement me")
}

func (r Repository) WriteComponent(access cpi.ComponentAccess) (cpi.ComponentAccess, error) {
	panic("implement me")
}

func (r Repository) LocalSupportForAccessSpec(a compdesc.AccessSpec) bool {
	panic("implement me")
}
