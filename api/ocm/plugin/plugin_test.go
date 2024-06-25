//go:build unix

package plugin_test

import (
	. "github.com/mandelsoft/goutils/testutils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "ocm.software/ocm-core/api/ocm/plugin/testutils"

	common2 "ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/credentials"
	"ocm.software/ocm-core/api/oci/extensions/actions/oci-repository-prepare"
	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/cpi"
	access "ocm.software/ocm-core/api/ocm/extensions/accessmethods/plugin"
	"ocm.software/ocm-core/api/ocm/extensions/attrs/plugincacheattr"
	"ocm.software/ocm-core/api/ocm/extensions/attrs/plugindirattr"
	"ocm.software/ocm-core/api/ocm/plugin"
	"ocm.software/ocm-core/api/ocm/plugin/cache"
	"ocm.software/ocm-core/api/ocm/plugin/common"
	"ocm.software/ocm-core/api/ocm/plugin/plugins"
	"ocm.software/ocm-core/api/ocm/plugin/registration"
	"ocm.software/ocm-core/api/ocm/valuemergehandler"
	"ocm.software/ocm-core/api/ocm/valuemergehandler/handlers/defaultmerge"
)

var _ = Describe("setup plugin cache", func() {
	var ctx ocm.Context
	var registry plugins.Set
	var plugins TempPluginDir

	BeforeEach(func() {
		cache.DirectoryCache.Reset()
		ctx = ocm.New()
		plugins, registry = Must2(ConfigureTestPlugins2(ctx, "testdata"))
	})

	AfterEach(func() {
		plugins.Cleanup()
	})

	It("finds plugin", func() {
		p := registry.Get("test")
		Expect(p).NotTo(BeNil())
		Expect(p.GetDescriptor().Short).To(Equal("a test plugin"))
	})

	It("executes action", func() {
		p := registry.Get("action")
		Expect(p).NotTo(BeNil())
		Expect(p.GetDescriptor().Short).To(Equal("a test plugin"))

		r := Must(p.Action(oci_repository_prepare.Spec("ghcr.io", "mandelsoft"), nil))
		Expect(r).To(Equal(oci_repository_prepare.Result("all good")))
	})

	It("scans only once", func() {
		ctx = ocm.New()
		plugindirattr.Set(ctx, plugins.Path())
		registry = plugincacheattr.Get(ctx)

		p := registry.Get("test")
		Expect(p).NotTo(BeNil())
		Expect(p.GetDescriptor().Short).To(Equal("a test plugin"))

		Expect(cache.DirectoryCache.Count()).To(Equal(1))
		Expect(cache.DirectoryCache.Requests()).To(Equal(2))
	})

	It("registers access methods", func() {
		p := registry.Get("test")
		Expect(p).NotTo(BeNil())
		Expect(len(p.GetDescriptor().AccessMethods)).To(Equal(2))
		Expect(registration.RegisterExtensions(registry.GetContext())).To(Succeed())
		t := ctx.AccessMethods().GetType("test")
		Expect(t).NotTo(BeNil())
		raw := `
type: test
someattr: value
`
		s, err := ctx.AccessSpecForConfig([]byte(raw), nil)
		Expect(err).To(Succeed())
		spec := s.(*access.AccessSpec)
		h := spec.Handler()
		info, err := h.Info(spec)
		Expect(err).To(Succeed())
		Expect(info).To(Equal(&plugin.AccessSpecInfo{
			Short:     "a test",
			MediaType: "plain/text",
			Hint:      "testfile",
			ConsumerId: credentials.ConsumerIdentity{
				"type":     "test",
				"hostname": "localhost",
			},
		}))
	})

	Context("inexpensive identity method", func() {
		It("inexpensive identity method compatibility test", func() {
			cv := &cpi.DummyComponentVersionAccess{Context: ctx}
			p := registry.Get("test")
			Expect(p).NotTo(BeNil())
			Expect(len(p.GetDescriptor().AccessMethods)).To(Equal(2))
			Expect(registration.RegisterExtensions(registry.GetContext())).To(Succeed())
			t := ctx.AccessMethods().GetType("test")
			Expect(t).NotTo(BeNil())

			raw := `type: test`
			s, err := ctx.AccessSpecForConfig([]byte(raw), nil)
			Expect(err).To(Succeed())
			spec := s.(*access.AccessSpec)
			Expect(spec.GetInexpensiveContentVersionIdentity(cv)).To(Equal(""))
		})

		It("check inexpensive identity method", func() {
			cv := &cpi.DummyComponentVersionAccess{Context: ctx}
			p := registry.Get("identity")
			Expect(p).NotTo(BeNil())
			Expect(len(p.GetDescriptor().AccessMethods)).To(Equal(1))
			Expect(registration.RegisterExtensions(registry.GetContext())).To(Succeed())
			t := ctx.AccessMethods().GetType("identity")
			Expect(t).NotTo(BeNil())

			raw := `type: identity`
			s, err := ctx.AccessSpecForConfig([]byte(raw), nil)
			Expect(err).To(Succeed())
			spec := s.(*access.AccessSpec)
			Expect(spec.GetInexpensiveContentVersionIdentity(cv)).To(Equal("testidentity"))
		})
	})

	Context("valuemergehandler", func() {
		It("finds plugin", func() {
			p := registry.Get("merge")
			Expect(p).NotTo(BeNil())
			Expect(p.Error()).To(Equal(""))
			Expect(p.IsValid()).To(BeTrue())
			Expect(p.GetDescriptor().Short).To(Equal("a test plugin"))
			Expect(len(p.GetDescriptor().ValueMergeHandlers)).To(Equal(1))
		})

		It("merges", func() {
			p := registry.Get("merge")
			Expect(p).NotTo(BeNil())
			Expect(p.Error()).To(Equal(""))
			Expect(p.IsValid()).To(BeTrue())
			spec := Must(valuemergehandler.NewSpecification("acme.org/test", defaultmerge.NewConfig("test")))

			var local, inbound valuemergehandler.Value
			local.SetValue("local")
			inbound.SetValue("local")
			ok, r := Must2(p.MergeValue(spec, local, inbound))
			Expect(ok).To(BeTrue())
			Expect(r.RawMessage).To(YAMLEqual(`{"mode":"resolved"}`))
		})

		It("provider merge specs", func() {
			p := registry.Get("merge")
			Expect(p).NotTo(BeNil())
			Expect(p.Error()).To(Equal(""))
			Expect(p.IsValid()).To(BeTrue())
			s := p.GetLabelMergeSpecification("testlabel", "v1")
			Expect(s).NotTo(BeNil())
			Expect(s.GetDescription()).To(Equal("generic testlabel merge spec"))
			Expect(s.Algorithm).To(Equal("default"))
			s = p.GetLabelMergeSpecification("testlabel", "v2")
			Expect(s).NotTo(BeNil())
			Expect(s.GetDescription()).To(Equal("v2 testlabel merge spec"))
			Expect(s.Algorithm).To(Equal("simpleMapMerge"))
		})

		It("described plugin", func() {
			p := registry.Get("merge")
			Expect(p).NotTo(BeNil())
			pr, buf := common2.NewBufferedPrinter()
			common.DescribePluginDescriptor(nil, p.GetDescriptor(), pr)
			Expect(buf.String()).To(StringEqualTrimmedWithContext(`
Plugin Name:      merge
Plugin Version:   v1
Capabilities:     Value Merge Handlers, Label Merge Specs
Description: 
      a test plugin with value merge algorithm acme.org/test

Value Merge Handlers:
- Name: acme.org/test
    test merger

Label Merge Specifications:
- Name: testlabel
  Algorithm: default
    generic testlabel merge spec
- Name: testlabel@v2
  Algorithm: simpleMapMerge
    v2 testlabel merge spec
`))
		})
	})
})
