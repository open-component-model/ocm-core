package testhelper

import (
	"github.com/mandelsoft/logging"

	ocmlog "ocm.software/ocm-core/api/utils/logging"
	"ocm.software/ocm-core/api/utils/refmgmt"
)

func EnableRefMgmtLog() {
	ocmlog.Context().AddRule(logging.NewConditionRule(logging.TraceLevel, refmgmt.ALLOC_REALM))
}
