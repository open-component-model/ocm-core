package srcs

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/validation/field"

	"ocm.software/ocm-core/api/clictx"
	"ocm.software/ocm-core/api/ocm"
	"ocm.software/ocm-core/api/ocm/compdesc"
	metav1 "ocm.software/ocm-core/api/ocm/compdesc/meta/v1"
	compdescv2 "ocm.software/ocm-core/api/ocm/compdesc/versions/v2"
	"ocm.software/ocm-core/api/utils/runtime"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common"
	"ocm.software/ocm-core/cmds/ocm/commands/ocmcmds/common/addhdlrs"
)

type ResourceSpecHandler struct{}

var _ common.ResourceSpecHandler = (*ResourceSpecHandler)(nil)

func (ResourceSpecHandler) Key() string {
	return "source"
}

func (ResourceSpecHandler) RequireInputs() bool {
	return true
}

func (ResourceSpecHandler) Decode(data []byte) (addhdlrs.ElementSpec, error) {
	var desc ResourceSpec
	err := runtime.DefaultYAMLEncoding.Unmarshal(data, &desc)
	if err != nil {
		return nil, err
	}
	return &desc, nil
}

func (ResourceSpecHandler) Set(v ocm.ComponentVersionAccess, r addhdlrs.Element, acc compdesc.AccessSpec) error {
	spec, ok := r.Spec().(*ResourceSpec)
	if !ok {
		return fmt.Errorf("element spec is not a valid resource spec, failed to assert type %T to ResourceSpec", r.Spec())
	}
	vers := spec.Version
	if vers == "" {
		vers = v.GetVersion()
	}
	meta := &compdesc.SourceMeta{
		ElementMeta: compdesc.ElementMeta{
			Name:          spec.Name,
			Version:       vers,
			ExtraIdentity: spec.ExtraIdentity,
			Labels:        spec.Labels,
		},
		Type: spec.Type,
	}
	return v.SetSource(meta, acc)
}

////////////////////////////////////////////////////////////////////////////////

type ResourceSpec struct {
	compdescv2.SourceMeta `json:",inline"`

	addhdlrs.ResourceInput `json:",inline"`
}

var _ addhdlrs.ElementSpec = (*ResourceSpec)(nil)

func (r *ResourceSpec) GetRawIdentity() metav1.Identity {
	return r.ElementMeta.GetRawIdentity()
}

func (r *ResourceSpec) Info() string {
	return fmt.Sprintf("source %s: %s", r.Type, r.GetRawIdentity())
}

func (r *ResourceSpec) Validate(ctx clictx.Context, input *addhdlrs.ResourceInput) error {
	allErrs := field.ErrorList{}
	var fldPath *field.Path

	src := compdescv2.Source{
		SourceMeta: r.SourceMeta,
	}
	if err := compdescv2.ValidateSource(fldPath, src, false); err != nil {
		allErrs = append(allErrs, err...)
	}
	if r.Access != nil {
		if r.Access.GetType() == "" {
			allErrs = append(allErrs, field.Required(fldPath.Child("access", "type"), "type of access required"))
		} else {
			acc, err := r.Access.Evaluate(ctx.OCMContext())
			if err != nil {
				raw, _ := r.Access.GetRaw()
				allErrs = append(allErrs, field.Invalid(fldPath.Child("access"), string(raw), err.Error()))
			} else if acc.IsLocal(ctx.OCMContext()) {
				kind := runtime.GetKind(r.Access)
				allErrs = append(allErrs, field.Invalid(fldPath.Child("access", "type"), kind, "local access no possible"))
			}
		}
	}
	return allErrs.ToAggregate()
}
