package directcreds

import (
	"ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/credentials/cpi"
)

func NewCredentials(props common.Properties) cpi.CredentialsSpec {
	return cpi.NewCredentialsSpec(Type, NewRepositorySpec(props))
}
