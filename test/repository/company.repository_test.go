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

func IniCompanyRepositoryTest() (repository.CompanyRepository, error) {
	initDb := env.DBConfig{
		DBPort: 5432, DBHost: "localhost", DBUserName: "root", DBPassword: "root", DBName: "nzedi_eshop", DBLogLevel: "INFO", DBSSLMode: "disable", DBTimeZone: "Europe/London",
	}
	gormDB, err := gorm.NewPostgresDBConnector(&initDb).Connect()
	if err != nil {
		return nil, err
	}
	factory := factory2.NewCompanyFactoryImpl()
	repository := gorm2.NewCompanyRepositoryImpl(gormDB, factory)
	return repository, nil
}
func TestCompanyCreate(t *testing.T) {
	repository, err := IniCompanyRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var account entity.Company
	gofakeit.Struct(&account)
	account.Id = 001
	account.Manager = 1
	result, err := repository.CreateCompany(account)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestCompanyUpdate(t *testing.T) {
	repository, err := IniCompanyRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var account entity.Company
	gofakeit.Struct(&account)
	account.Id = 1
	account.Manager = 1
	result, err := repository.UpdateCompany(account)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestGetCompany(t *testing.T) {
	repository, err := IniCompanyRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := repository.ReadCompany(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestGetCompanies(t *testing.T) {
	repository, err := IniCompanyRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := repository.ReadCompanyAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
func TestGetWithManager(t *testing.T) {
	repository, err := IniCompanyRepositoryTest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result, err := repository.ReadCompanyWithUserId(1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(result)
	}
}
