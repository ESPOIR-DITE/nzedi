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

type AccountRepositoryImpl struct {
	GormDB         *gorm.DB
	AccountFactory factory.AccountFactory
}

func NewAccountRepositoryImpl(gormDB *gorm.DB, accountFactory factory.AccountFactory) *AccountRepositoryImpl {
	return &AccountRepositoryImpl{
		GormDB: gormDB, AccountFactory: accountFactory,
	}
}

var _ repository.AccountRepository = &AccountRepositoryImpl{}

func (a AccountRepositoryImpl) CreateAccount(account entity.Account) (models.Account, error) {
	var gormAccount *gormModel.Account = a.AccountFactory.CreateAccount(account).(*gormModel.Account)
	if err := a.GormDB.Create(&gormAccount).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create account with Email: %s, err: %s", account.Email, err))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) ReadAccount(id string) (models.Account, error) {
	gormAccount := &gormModel.Account{}
	if err := a.GormDB.Where("id = ?", id).First(&gormAccount).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get account with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) UpdateAccount(account entity.Account) (models.Account, error) {
	var gormAccount *gormModel.Account = a.AccountFactory.CreateAccount(account).(*gormModel.Account)
	if err := a.GormDB.Where("id = ?", account.Id).Updates(gormAccount).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update Account id: %d", account.Id))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) DeleteAccount(account entity.Account) (bool, error) {
	gormMediaType := &gormModel.Account{}
	if err := a.GormDB.Where("id = ?", account.Id).Delete(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update Account id: %d", account.Id))
		return false, err
	}
	return true, nil
}

func (a AccountRepositoryImpl) ReadAccountAll() ([]gormModel.Account, error) {
	gormAccount := []gormModel.Account{}
	if err := a.GormDB.Find(&gormAccount).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to reads Accounts"))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) UpdateToken(id, token string) (models.Account, error) {
	gormAccount := &gormModel.Account{}
	if err := a.GormDB.Model(gormAccount).Where("id = ?", id).Update("token", token).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update token with Account id: %s : %s", id, err))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) LoginWithEmail(account entity.Account) (models.Account, error) {
	gormAccount := &gormModel.Account{}
	if err := a.GormDB.Where("email = ? AND password = ? ", account.Email, account.Password).First(&gormAccount).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to login with Account email: %s : %s", account.Email, err))
		return nil, err
	}
	return gormAccount, nil
}

func (a AccountRepositoryImpl) LoginWithUserName(account entity.Account) (models.Account, error) {
	gormMediaType := &gormModel.Account{}
	if err := a.GormDB.Where("username = ? AND password = ? ", account.Username, account.Password).First(&gormMediaType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to login with username Account username: %d", account.Username))
		return nil, err
	}
	return gormMediaType, nil
}
