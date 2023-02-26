package wire

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config/nzedi"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/connector/gorm"
	migrate2 "github.com/ESPOIR-DITE/nzedi.git/pkg/db/migration/migrate"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"github.com/golang-migrate/migrate/v4"
)

func Run(ConfigManager nzedi.ConfigManager) int {
	config, err := ConfigManager.Load()
	if err != nil {
		logger.Log.Fatalf("Failed to load application configurations: %s", err.Error())
		return 1
	}
	if err := logger.Configure(config.AppConfig().LogLevel()); err != nil {
		logger.Log.Fatalf("Failed to configure logger: %s", err.Error())
	}
	gormDB, err := gorm.NewPostgresDBConnector(config.DBConfig()).Connect()
	if err != nil {
		logger.Log.Fatalf(err.Error())
		return 2
	}
	migrator, err := migrate2.NewPostgresMigrator(config.DBConfig()).NewMigrator()
	if err != nil {
		logger.Log.Fatalf("Failed to initiate db migration: %s", err.Error())
		return 3
	}
	err = migrator.Up()
	if err != migrate.ErrNoChange && err != nil {
		logger.Log.Fatalf("fail to run migration: %s", err)
		return 4
	}
	httpServer := wire(config, gormDB)

	if err = httpServer.Start(); err != nil {
		logger.Log.Fatal(err)
		return 1
	}
	return 0
}
