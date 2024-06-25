package maventest

import (
	"github.com/open-component-model/ocm/api/helper/env"
)

func TestData(dest ...string) env.Option {
	return env.ProjectTestDataForCaller(dest...)
}

func ModifiableTestData(dest ...string) env.Option {
	return env.ModifiableProjectTestDataForCaller(dest...)
}