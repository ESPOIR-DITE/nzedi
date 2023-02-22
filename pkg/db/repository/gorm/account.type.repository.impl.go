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

type AccountTypeRepositoryImpl struct {
	GormDB             *gorm.DB
	AccountTypeFactory factory.AccountTypeFactory
}

func NewAccountTypeRepositoryImpl(gormDB *gorm.DB,
	accountStateFactory factory.AccountTypeFactory) *AccountTypeRepositoryImpl {
	return &AccountTypeRepositoryImpl{
		GormDB:             gormDB,
		AccountTypeFactory: accountStateFactory,
	}
}

var _ repository.AccountTypeRepository = &AccountTypeRepositoryImpl{}

func (a AccountTypeRepositoryImpl) CreateAccountType(accountType entity.AccountType) (models.AccountType, error) {
	var gormAccountType *gormModel.AccountType = a.AccountTypeFactory.CreateAccountType(accountType).(*gormModel.AccountType)
	if err := a.GormDB.Create(&gormAccountType).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create accountType with account id: %d, err: %s", accountType.AccountId, err))
		return nil, err
	}
	return gormAccountType, nil
}

func (a AccountTypeRepositoryImpl) ReadAccountType(id int) (models.AccountType, error) {
	gormAccountType := &gormModel.AccountType{}
	if err := a.GormDB.First(&gormAccountType, id).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get AccountType with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormAccountType, nil
}

func (a AccountTypeRepositoryImpl) ReadAccountTypeWithAccountId(id string) (models.AccountType, error) {
	//TODO implement me
	panic("implement me")
}

func (a AccountTypeRepositoryImpl) UpdateAccountType(accountState entity.AccountType) (models.AccountType, error) {
	var gormAccountType *gormModel.AccountType = a.AccountTypeFactory.CreateAccountType(accountState).(*gormModel.AccountType)
	if err := a.GormDB.Where("id = ?", accountState.Id).Updates(gormAccountType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update AccountType id: %d", accountState.Id))
		return nil, err
	}
	return gormAccountType, nil
}

func (a AccountTypeRepositoryImpl) DeleteAccountType(accountState entity.AccountType) (bool, error) {
	gormAccountType := &gormModel.AccountType{}
	if err := a.GormDB.Where("id = ?", accountState.Id).Delete(&gormAccountType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update AccountType id: %d", accountState.Id))
		return false, err
	}
	return true, nil
}

func (a AccountTypeRepositoryImpl) ReadAccountTypeAll() ([]models.AccountType, error) {
	//TODO implement me
	panic("implement me")
}
