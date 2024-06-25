package oci

import (
	"fmt"

	ociidentity "ocm.software/ocm-core/api/credentials/builtin/oci/identity"
	"ocm.software/ocm-core/api/credentials/cpi"
	"ocm.software/ocm-core/api/oci/artdesc"
	"ocm.software/ocm-core/api/oci/grammar"
	"ocm.software/ocm-core/api/utils/runtime"
)

func AsTags(tag string) []string {
	if tag != "" {
		return []string{tag}
	}
	return nil
}

func StandardOCIRef(host, repository, version string) string {
	sep := grammar.TagSeparator
	if ok, _ := artdesc.IsDigest(version); ok {
		sep = grammar.DigestSeparator
	}
	return fmt.Sprintf("%s%s%s%s%s", host, grammar.RepositorySeparator, repository, sep, version)
}

func IsIntermediate(spec RepositorySpec) bool {
	if s, ok := spec.(IntermediateRepositorySpecAspect); ok {
		return s.IsIntermediate()
	}
	return false
}

func IsUnknown(r RepositorySpec) bool {
	return runtime.IsUnknown(r)
}

func GetConsumerIdForRef(ref string) (cpi.ConsumerIdentity, error) {
	r, err := ParseRef(ref)
	if err != nil {
		return nil, err
	}
	return ociidentity.GetConsumerId(r.Host, r.Repository), nil
}
