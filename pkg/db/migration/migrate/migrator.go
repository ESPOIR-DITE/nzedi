package migrate

import "github.com/golang-migrate/migrate/v4"

type Migrator interface {
	NewMigrator() (*migrate.Migrate, error)
}
