package factory

import (
	spec "nzedi/api/server/nzedi-api"
	"nzedi/pkg/entity"
)

type CompanyFactory interface {
	CreateCompany(body spec.Company) (*entity.Company, error)
}

type CompanyFactoryImpl struct {
}

func (c CompanyFactoryImpl) CreateCompany(body spec.Company) (*entity.Company, error) {
	//TODO implement me
	panic("implement me")
}

func NewCompanyFactoryImpl() *CompanyFactoryImpl {
	return &CompanyFactoryImpl{}
}

var _ CompanyFactory = &CompanyFactoryImpl{}
