package pluginhdlr

import (
	"strings"

	"github.com/mandelsoft/goutils/errors"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/ocm/extensions/attrs/plugincacheattr"
	"ocm.software/ocm-core/api/ocm/plugin"
	"ocm.software/ocm-core/api/ocm/plugin/descriptor"
	"ocm.software/ocm-core/api/ocm/plugin/plugins"
	"ocm.software/ocm-core/cmds/ocm/common/output"
	"ocm.software/ocm-core/cmds/ocm/common/utils"
)

func Elem(e interface{}) plugin.Plugin {
	return e.(*Object).Plugin
}

////////////////////////////////////////////////////////////////////////////////

type Object struct {
	Plugin plugin.Plugin
}

type Manifest struct {
	Element *plugin.Descriptor `json:"element"`
}

func (o *Object) AsManifest() interface{} {
	return &Manifest{
		o.Plugin.GetDescriptor(),
	}
}

////////////////////////////////////////////////////////////////////////////////

type TypeHandler struct {
	octx clictx.OCM
}

func NewTypeHandler(octx clictx.OCM) utils.TypeHandler {
	return &TypeHandler{
		octx: octx,
	}
}

func (h *TypeHandler) Close() error {
	return nil
}

func (h *TypeHandler) All() ([]output.Object, error) {
	cache := plugincacheattr.Get(h.octx.Context())
	result := []output.Object{}

	for _, n := range cache.PluginNames() {
		result = append(result, &Object{cache.Get(n)})
	}
	return result, nil
}

func (h *TypeHandler) Get(elemspec utils.ElemSpec) ([]output.Object, error) {
	cache := plugincacheattr.Get(h.octx.Context())

	p := cache.Get(elemspec.String())
	if p == nil {
		objs := Lookup(elemspec.String(), cache)
		if len(objs) == 0 {
			return nil, errors.ErrNotFound(descriptor.KIND_PLUGIN, elemspec.String())
		}
		return objs, nil
	}
	return []output.Object{&Object{p}}, nil
}

func Lookup(prefix string, cache plugins.Set) []output.Object {
	var objs []output.Object
	prefix += "."
	for _, n := range cache.PluginNames() {
		if strings.HasPrefix(n, prefix) {
			objs = append(objs, &Object{cache.Get(n)})
		}
	}
	return objs
}
