package nzedi

type ConfigManager interface {
	Load() (ServiceConfiguration, error)
}
