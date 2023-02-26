package env

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config/nzedi"
	"github.com/kelseyhightower/envconfig"
)

type NzediApiConfigManagerImpl struct{}

func NewNzediApiConfigManagerImpl() *NzediApiConfigManagerImpl {
	return &NzediApiConfigManagerImpl{}
}

func (e NzediApiConfigManagerImpl) Load() (config nzedi.ServiceConfiguration, err error) {
	config = new(EshopEnvServiceConfiguration)
	if err = envconfig.Process("", config); err != nil {
		return
	}
	return
}

var _ nzedi.ConfigManager = &NzediApiConfigManagerImpl{}
