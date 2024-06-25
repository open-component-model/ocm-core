package credentials

import (
	"context"

	"ocm.software/ocm-core/api/config"
	"ocm.software/ocm-core/api/credentials/internal"
	"ocm.software/ocm-core/api/datacontext"
)

func WithContext(ctx context.Context) internal.Builder {
	return internal.Builder{}.WithContext(ctx)
}

func WithConfigs(ctx config.Context) internal.Builder {
	return internal.Builder{}.WithConfig(ctx)
}

func WithRepositoyTypeScheme(scheme RepositoryTypeScheme) internal.Builder {
	return internal.Builder{}.WithRepositoyTypeScheme(scheme)
}

func WithStandardConumerMatchers(matchers internal.IdentityMatcherRegistry) internal.Builder {
	return internal.Builder{}.WithStandardConumerMatchers(matchers)
}

func New(mode ...datacontext.BuilderMode) Context {
	return internal.Builder{}.New(mode...)
}
