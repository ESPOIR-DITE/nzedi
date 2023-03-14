package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type CompanyServiceImpl struct {
	CompanyServiceRepository repository.CompanyRepository
}

func NewCompanyServiceImpl(companyServiceRepository repository.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{CompanyServiceRepository: companyServiceRepository}
}

var _ service.CompanyService = &CompanyServiceImpl{}

func (c CompanyServiceImpl) CreateCompany(Company entity.Company) (*entity.Company, error) {
	company, err := c.CompanyServiceRepository.CreateCompany(Company)
	if err != nil {
		return nil, err
	}
	return company.GetCompany(), err
}

func (c CompanyServiceImpl) ReadCompany(id string) (*entity.Company, error) {
	company, err := c.CompanyServiceRepository.ReadCompany(id)
	if err != nil {
		return nil, err
	}
	return company.GetCompany(), err
}

func (c CompanyServiceImpl) ReadCompanyWithUserId(id string) (*entity.Company, error) {
	company, err := c.CompanyServiceRepository.ReadCompanyWithUserId(id)
	if err != nil {
		return nil, err
	}
	return company.GetCompany(), err
}

func (c CompanyServiceImpl) UpdateCompany(Company entity.Company) (*entity.Company, error) {
	company, err := c.CompanyServiceRepository.CreateCompany(Company)
	if err != nil {
		return nil, err
	}
	return company.GetCompany(), err
}

func (c CompanyServiceImpl) DeleteCompany(Company entity.Company) (bool, error) {
	company, err := c.CompanyServiceRepository.DeleteCompany(Company)
	if err != nil {
		return false, err
	}
	return company, err
}

func (c CompanyServiceImpl) ReadCompanyAll() ([]entity.Company, error) {
	companies, err := c.CompanyServiceRepository.ReadCompanyAll()
	if err != nil {
		return nil, err
	}
	return c.getCompanyList(companies), err
}

func (c CompanyServiceImpl) getCompanyList(accountList []gorm.Company) []entity.Company {
	var companyAll []entity.Company
	if accountList == nil {
		return companyAll
	}
	for _, company := range accountList {
		companyAll = append(companyAll, entity.Company{
			Id:      company.Id,
			Name:    company.Name,
			Manager: company.Manager,
			Url:     company.Url,
		})
	}
	return companyAll
}
