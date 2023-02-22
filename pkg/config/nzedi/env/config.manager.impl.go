package env

import (
	"github.com/kelseyhightower/envconfig"
	"nzedi/pkg/config/nzedi"
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
