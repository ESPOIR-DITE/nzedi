package gorm

import (
	"fmt"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDBConnector struct {
	Config config.DBConfig
}

var _ DBConnector = &PostgresDBConnector{}

func NewPostgresDBConnector(config config.DBConfig) *PostgresDBConnector {
	return &PostgresDBConnector{
		Config: config,
	}
}
func (p PostgresDBConnector) Connect() (*gorm.DB, error) {
	logger, err := logger.CreateGormLogger(p.Config.LogLevel())
	if err != nil {
		return nil, err
	}

	return gorm.Open(postgres.Open(p.connectionString()), &gorm.Config{
		Logger: logger,
	})
}

func (p PostgresDBConnector) connectionString() string {
	config := p.Config
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s Timezone=%s",
		config.Host(), config.UserName(), config.Password(), config.Name(), config.Port(), config.SSLMode(), config.TimeZone())
}
