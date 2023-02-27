package gorm

import (
	"fmt"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	gormModel "github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/logger"
	"gorm.io/gorm"
)

type CompanyRepositoryImpl struct {
	GormDB         *gorm.DB
	CompanyFactory factory.CompanyFactory
}

func NewCompanyRepositoryImpl(gormDB *gorm.DB,
	companyFactory factory.CompanyFactory) *CompanyRepositoryImpl {
	return &CompanyRepositoryImpl{
		GormDB:         gormDB,
		CompanyFactory: companyFactory,
	}
}

var _ repository.CompanyRepository = &CompanyRepositoryImpl{}

func (c CompanyRepositoryImpl) CreateCompany(company entity.Company) (models.Company, error) {
	var gormCompany gormModel.Company = c.CompanyFactory.CreateCompany(company).(gormModel.Company)
	if err := c.GormDB.Create(&gormCompany).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create Company with Name: %s, err: %s", company.Name, err))
		return nil, err
	}
	return gormCompany, nil
}

func (c CompanyRepositoryImpl) ReadCompany(id int) (models.Company, error) {
	gormCompany := &gormModel.Company{}
	if err := c.GormDB.First(&gormCompany, id).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get Company with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormCompany, nil
}

func (c CompanyRepositoryImpl) ReadCompanyWithUserId(id int) (models.Company, error) {
	gormCompany := &gormModel.Company{}
	if err := c.GormDB.Where("manager = ?", id).First(&gormCompany).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to get Company Type with manager id: %d : %s", id, err))
		return nil, err
	}
	return gormCompany, nil
}

func (c CompanyRepositoryImpl) UpdateCompany(company entity.Company) (models.Company, error) {
	var gormCompany gormModel.Company = c.CompanyFactory.CreateCompany(company).(gormModel.Company)
	if err := c.GormDB.Where("id = ?", company.Id).Updates(gormCompany).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update company id: %d", company.Id))
		return nil, err
	}
	return gormCompany, nil
}

func (c CompanyRepositoryImpl) DeleteCompany(Company entity.Company) (bool, error) {
	gormCompany := &gormModel.Company{}
	if err := c.GormDB.Where("id = ?", Company.Id).Delete(&gormCompany).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update Company id: %d", Company.Id))
		return false, err
	}
	return true, nil
}

func (c CompanyRepositoryImpl) ReadCompanyAll() ([]gormModel.Company, error) {
	gormCompany := []gormModel.Company{}
	if err := c.GormDB.Find(&gormCompany).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads Companies"))
		return nil, err
	}
	return gormCompany, nil
}
