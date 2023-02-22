package env

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/config/env"
)

type EshopEnvServiceConfiguration struct {
	env.EnvServiceConfiguration
}

var _ config.ServiceConfiguration = &EshopEnvServiceConfiguration{}
