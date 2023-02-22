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

type AccountStateRepositoryImpl struct {
	GormDB              *gorm.DB
	AccountStateFactory factory.AccountStateFactory
}

func NewAccountStateRepositoryImpl(gormDB *gorm.DB, accountStateFactory factory.AccountStateFactory) *AccountStateRepositoryImpl {
	return &AccountStateRepositoryImpl{
		GormDB:              gormDB,
		AccountStateFactory: accountStateFactory,
	}
}

var _ repository.AccountStateRepository = &AccountStateRepositoryImpl{}

func (a AccountStateRepositoryImpl) CreateAccountState(accountState entity.AccountState) (models.AccountState, error) {
	var gormAccountState *gormModel.AccountState = a.AccountStateFactory.CreateAccountState(accountState).(*gormModel.AccountState)
	if err := a.GormDB.Create(&gormAccountState).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get accountState with name: %s, err: %s", accountState.Name, err))
		return nil, err
	}
	return gormAccountState, nil
}

func (a AccountStateRepositoryImpl) ReadAccountState(id int) (models.AccountState, error) {
	gormAccountState := &gormModel.AccountState{}
	if err := a.GormDB.First(&gormAccountState, id).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get accountState with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormAccountState, nil
}

func (a AccountStateRepositoryImpl) UpdateAccountState(accountState entity.AccountState) (models.AccountState, error) {
	var gormAccountState *gormModel.AccountState = a.AccountStateFactory.CreateAccountState(accountState).(*gormModel.AccountState)
	if err := a.GormDB.Where("id = ?", accountState.Id).Updates(gormAccountState).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update AccountState id: %d", accountState.Id))
		return nil, err
	}
	return gormAccountState, nil
}

func (a AccountStateRepositoryImpl) DeleteAccountState(accountState entity.AccountState) (bool, error) {
	gormMediaState := &gormModel.AccountState{}
	if err := a.GormDB.Where("id = ?", accountState.Id).Delete(&gormMediaState).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update AccountState id: %d", accountState.Id))
		return false, err
	}
	return true, nil
}

func (a AccountStateRepositoryImpl) ReadAccountStateAll() ([]gormModel.AccountState, error) {
	gormAccountState := []gormModel.AccountState{}
	if err := a.GormDB.Find(&gormAccountState).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads account state"))
		return nil, err
	}
	return gormAccountState, nil
}
