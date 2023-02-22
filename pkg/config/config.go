package config

type AppConfig interface {
	Port() int
	AwsRegion() string
	LogLevel() string
}
type DBConfig interface {
	Port() int
	Host() string
	UserName() string
	Password() string
	Name() string
	SSLMode() string
	TimeZone() string
	MigrationFolder() string
	LogLevel() string
}
