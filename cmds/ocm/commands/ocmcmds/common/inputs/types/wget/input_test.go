// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Open Component Model contributors.
//
// SPDX-License-Identifier: Apache-2.0

package wget_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs/testutils"

	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs/options"
	"github.com/open-component-model/ocm/cmds/ocm/commands/ocmcmds/common/inputs/types/wget"
	"github.com/open-component-model/ocm/pkg/mime"
)

var _ = Describe("Input Type", func() {
	var env *InputTest

	BeforeEach(func() {
		env = NewInputTest(wget.TYPE)
	})

	It("simple decode", func() {
		env.Set(options.URLOption, "https://example.com/test")
		env.Set(options.MediaTypeOption, mime.MIME_TEXT)
		env.Set(options.HTTPHeaderOption, "Host: developer.mozilla.org")
		env.Set(options.HTTPHeaderOption, "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0")
		env.Set(options.HTTPHeaderOption, "Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8")
		env.Set(options.HTTPVerbOption, http.MethodPost)
		env.Set(options.HTTPBodyOption, "hello world")
		env.Set(options.HTTPRedirectOption, "true")
		env.Check(&wget.Spec{
			InputSpecBase: inputs.InputSpecBase{},
			URL:           "https://example.com/test",
			MimeType:      mime.MIME_TEXT,
			Header: map[string][]string{
				"Host":       {"developer.mozilla.org"},
				"User-Agent": {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:50.0) Gecko/20100101 Firefox/50.0"},
				"Accept":     {"text/html", "application/xhtml+xml", "application/xml;q=0.9", "*/*;q=0.8"},
			},
			Verb:       http.MethodPost,
			Body:       "hello world",
			NoRedirect: true,
		})
	})
})
