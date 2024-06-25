package github

import (
	"os"

	"ocm.software/ocm-core/api/common/common"
	"ocm.software/ocm-core/api/credentials/builtin/oci/identity"
	"ocm.software/ocm-core/api/credentials/cpi"
)

const HOST = "ghcr.io"

func init() {
	t := os.Getenv("GITHUB_TOKEN")
	if t != "" {
		host := os.Getenv("GITHUB_HOST")
		if host == "" {
			host = HOST
		}
		id := cpi.NewConsumerIdentity(identity.CONSUMER_TYPE, identity.ID_HOSTNAME, host)
		user := os.Getenv("GITHUB_REPOSITORY_OWNER")
		if user == "" {
			user = "any"
		}
		if src, err := cpi.DefaultContext.GetCredentialsForConsumer(id); err != nil || src == nil {
			creds := cpi.NewCredentials(common.Properties{cpi.ATTR_IDENTITY_TOKEN: t, cpi.ATTR_USERNAME: user})
			cpi.DefaultContext.SetCredentialsForConsumer(id, creds)
		}
	}
}
