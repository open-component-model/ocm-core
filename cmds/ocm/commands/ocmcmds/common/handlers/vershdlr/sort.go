package vershdlr

import (
	"strings"

	"ocm.software/ocm-core/api/utils/semverutils"
	"ocm.software/ocm-core/cmds/ocm/common/processing"
)

func Compare(a, b interface{}) int {
	aa := a.(*Object)
	ab := b.(*Object)

	c := strings.Compare(aa.Component, ab.Component)
	if c != 0 {
		return c
	}
	return semverutils.Compare(aa.Version, ab.Version)
}

// Sort is a processing chain sorting original objects provided by type handler.
var Sort = processing.Sort(Compare)
