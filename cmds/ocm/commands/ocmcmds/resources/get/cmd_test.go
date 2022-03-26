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

package get_test

import (
	"bytes"

	. "github.com/gardener/ocm/cmds/ocm/testhelper"
	"github.com/gardener/ocm/cmds/ocm/testhelper/builder"
	"github.com/gardener/ocm/pkg/common/accessio"
	"github.com/gardener/ocm/pkg/mime"
	metav1 "github.com/gardener/ocm/pkg/ocm/compdesc/meta/v1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const ARCH = "/tmp/ca"
const VERSION = "v1"
const COMP = "test.de/x"
const PROVIDER = "mandelsoft"

var _ = Describe("Test Environment", func() {
	var env *builder.Builder

	BeforeEach(func() {
		env = builder.NewBuilder(NewTestEnv(TestData()))

	})

	AfterEach(func() {
		env.Cleanup()
	})

	It("lists single resource in component archive", func() {
		env.ComponentArchive(ARCH, accessio.FormatDirectory, COMP, VERSION, func() {
			env.Provider(PROVIDER)
			env.Resource("testdata", "", "PlainText", metav1.LocalRelation, func() {
				env.BlobStringData(mime.MIME_TEXT, "testdata")
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "resources", ARCH)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
NAME     VERSION IDENTITY          TYPE      RELATION
testdata v1      "name"="testdata" PlainText local
`))
	})

	It("lists single resource in ctf file", func() {
		env.CommonTransport(ARCH, accessio.FormatDirectory, func() {
			env.Component(COMP, func() {
				env.Version(VERSION, func() {
					env.Provider(PROVIDER)
					env.Resource("testdata", "", "PlainText", metav1.LocalRelation, func() {
						env.BlobStringData(mime.MIME_TEXT, "testdata")
					})
				})
			})
		})

		buf := bytes.NewBuffer(nil)
		Expect(env.CatchOutput(buf).Execute("get", "resources", ARCH)).To(Succeed())
		Expect("\n" + buf.String()).To(Equal(
			`
NAME     VERSION IDENTITY          TYPE      RELATION
testdata v1      "name"="testdata" PlainText local
`))
	})
})
