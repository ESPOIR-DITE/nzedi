package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type CompanyFactoryImpl struct{}

func (c CompanyFactoryImpl) CreateCompany(company entity.Company) models.Company {
	return gorm.Company{
		Id:      company.Id,
		Manager: company.Manager,
		Name:    company.Name,
		Url:     company.Url,
	}
}

var _ factory.CompanyFactory = CompanyFactoryImpl{}
