package maventest

import (
	"ocm.software/ocm-core/api/helper/env"
)

func TestData(dest ...string) env.Option {
	return env.ProjectTestDataForCaller(dest...)
}

func ModifiableTestData(dest ...string) env.Option {
	return env.ModifiableProjectTestDataForCaller(dest...)
}
