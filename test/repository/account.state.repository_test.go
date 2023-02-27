package repository

import (
	"fmt"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config/env"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/connector/gorm"
	factory2 "github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	gorm2 "github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/brianvoe/gofakeit/v6"
	"os"
	"testing"
)

func IniAccountStateRepositoryTest() (repository.AccountStateRepository, error) {
	initDb := env.DBConfig{
		DBPort: 5432, DBHost: "localhost", DBUserName: "root", DBPassword: "root", DBName: "nzedi_eshop", DBLogLevel: "INFO", DBSSLMode: "disable", DBTimeZone: "Europe/London",
	}
	gormDB, err := gorm.NewPostgresDBConnector(&initDb).Connect()
	if err != nil {
		return nil, err
	}
	factory := factory2.NewAccountStateFactoryImpl()
	repository := gorm2.NewAccountStateRepositoryImpl(gormDB, factory)
	return repository, nil
}
func TestAccountStateCreate(t *testing.T) {
	repository, err := IniAccountStateRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var account entity.AccountState
	gofakeit.Struct(&account)
	account.Id = 002
	result, err := repository.CreateAccountState(account)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestAccountStateUpdate(t *testing.T) {
	repository, err := IniAccountStateRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var accountState entity.AccountState
	gofakeit.Struct(&accountState)
	accountState.Id = 1
	result, err := repository.UpdateAccountState(accountState)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestGetAccountState(t *testing.T) {
	repository, err := IniAccountStateRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := repository.ReadAccountState(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestGetAccountStates(t *testing.T) {
	repository, err := IniAccountStateRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := repository.ReadAccountStateAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
