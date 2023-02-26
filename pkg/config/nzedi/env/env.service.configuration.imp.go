package env

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config/env"
)

type EshopEnvServiceConfiguration struct {
	env.EnvServiceConfiguration
}

var _ config.ServiceConfiguration = &EshopEnvServiceConfiguration{}
