package env

import "github.com/ESPOIR-DITE/ditkay-eshop/pkg/config"

type AppConfig struct {
	AppPort      int    `envconfig:"HTTP_SERVER_PORT" required:"true"`
	AppAwsRegion string `envconfig:"AWS_REGION" required:"true"`
	AppLogLevel  string `envconfig:"APP_LOG_LEVEL" default:"INFO"`
}

func (a AppConfig) Port() int {
	return a.AppPort
}

func (a AppConfig) AwsRegion() string {
	return a.AppAwsRegion
}

func (a AppConfig) LogLevel() string {
	return a.AppLogLevel
}

var _ config.AppConfig = &AppConfig{}

type DBConfig struct {
	DBPort            int    `envconfig:"DB_PORT" required:"true"`
	DBHost            string `envconfig:"DB_HOST" required:"true"`
	DBUserName        string `envconfig:"DB_USERNAME" required:"true"`
	DBPassword        string `envconfig:"DB_PASSWORD" required:"true"`
	DBName            string `envconfig:"DB_NAME" required:"true"`
	DBSSLMode         string `envconfig:"DB_SSL_MODE" default:"disable"`
	DBTimeZone        string `envconfig:"DB_TIMEZONE" default:"Europe/London"`
	DBMigrationFolder string `envconfig:"DB_MIGRATION_FOLDER" default:"db/migration"`
	DBLogLevel        string `envconfig:"DB_LOGLEVEL" default:"silent"`
}

func (D DBConfig) Port() int {
	return D.DBPort
}

func (D DBConfig) Host() string {
	return D.DBHost
}

func (D DBConfig) UserName() string {
	return D.DBUserName
}

func (D DBConfig) Password() string {
	return D.DBPassword
}

func (D DBConfig) Name() string {
	return D.DBName
}

func (D DBConfig) SSLMode() string {
	return D.DBSSLMode
}

func (D DBConfig) TimeZone() string {
	return D.DBTimeZone
}

func (D DBConfig) MigrationFolder() string {
	return D.DBMigrationFolder
}

func (D DBConfig) LogLevel() string {
	return D.DBLogLevel
}

var _ config.DBConfig = &DBConfig{}
