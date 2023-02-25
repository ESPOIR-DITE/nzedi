package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type CompanyFactory interface {
	CreateCompany(body spec.Company) (*entity.Company, error)
}

type CompanyFactoryImpl struct {
}

func (c CompanyFactoryImpl) CreateCompany(body spec.Company) (*entity.Company, error) {
	return &entity.Company{
		Id:      body.Id,
		Manager: body.Manager,
		Name:    body.Name,
		Url:     body.Url,
	}, nil
}

func NewCompanyFactoryImpl() *CompanyFactoryImpl {
	return &CompanyFactoryImpl{}
}

var _ CompanyFactory = &CompanyFactoryImpl{}
