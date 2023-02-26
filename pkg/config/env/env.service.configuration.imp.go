package env

import "github.com/ESPOIR-DITE/nzedi.git/pkg/config"

type EnvServiceConfiguration struct {
	EnvServiceAppConfig AppConfig
	EnvServiceDBConfig  DBConfig
}

func (e EnvServiceConfiguration) AppConfig() config.AppConfig {
	return e.EnvServiceAppConfig
}

func (e EnvServiceConfiguration) DBConfig() config.DBConfig {
	return e.EnvServiceDBConfig
}

var _ config.ServiceConfiguration = &EnvServiceConfiguration{}
