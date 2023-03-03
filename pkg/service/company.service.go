package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type CompanyService interface {
	CreateCompany(Company entity.Company) (*entity.Company, error)
	ReadCompany(id string) (*entity.Company, error)
	ReadCompanyWithUserId(id string) (*entity.Company, error)
	UpdateCompany(Company entity.Company) (*entity.Company, error)
	DeleteCompany(Company entity.Company) (bool, error)
	ReadCompanyAll() ([]entity.Company, error)
}
