package internal

import (
	"ocm.software/ocm-core/api/credentials"
)

type UploadTargetSpecInfo struct {
	ConsumerId credentials.ConsumerIdentity `json:"consumerId"`
}
