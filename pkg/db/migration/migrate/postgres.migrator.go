package migrate

import (
	"fmt"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config"
	"github.com/golang-migrate/migrate/v4"
)

type PostgresMigrator struct {
	Config config.DBConfig
}

func NewPostgresMigrator(config config.DBConfig) *PostgresMigrator {
	return &PostgresMigrator{
		Config: config,
	}
}

func (p PostgresMigrator) NewMigrator() (*migrate.Migrate, error) {
	return migrate.New("file://"+p.Config.MigrationFolder(), p.connectionString())
}
func (p PostgresMigrator) connectionString() string {
	config := p.Config
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", config.UserName(), config.Password(), config.Host(), config.Port(), config.Name(), config.SSLMode())
}

var _ Migrator = &PostgresMigrator{}
