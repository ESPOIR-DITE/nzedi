package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type CompanyRepository interface {
	CreateCompany(Company entity.Company) (models.Company, error)
	ReadCompany(id string) (models.Company, error)
	ReadCompanyWithUserId(id string) (models.Company, error)
	UpdateCompany(Company entity.Company) (models.Company, error)
	DeleteCompany(Company entity.Company) (bool, error)
	ReadCompanyAll() ([]gorm.Company, error)
}
