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

type UserRepositoryImpl struct {
	GormDB      *gorm.DB
	UserFactory factory.UserFactory
}

var _ repository.UserRepository = &UserRepositoryImpl{}

func NewUserRepositoryImpl(gormDB *gorm.DB,
	userFactory factory.UserFactory) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		GormDB:      gormDB,
		UserFactory: userFactory,
	}
}

func (u UserRepositoryImpl) CreateUser(User entity.User) (models.User, error) {
	var gormUser *gormModel.User = u.UserFactory.CreateUser(User).(*gormModel.User)
	if err := u.GormDB.Create(&gormUser).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create User with account id: %d, err: %s", User.AccountId, err))
		return nil, err
	}
	return gormUser, nil
}

func (u UserRepositoryImpl) ReadUser(id int) (models.User, error) {
	gormUser := &gormModel.User{}
	if err := u.GormDB.First(&gormUser, id).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get User with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormUser, nil
}

func (u UserRepositoryImpl) ReadUserWithAccountId(id int) (models.User, error) {
	gormUser := &gormModel.User{}
	if err := u.GormDB.Where("account_id = ?", id).First(&gormUser).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to get user Type with accountId: %d : %s", id, err))
		return nil, err
	}
	return gormUser, nil
}

func (u UserRepositoryImpl) UpdateUser(User entity.User) (models.User, error) {
	var gormUser *gormModel.User = u.UserFactory.CreateUser(User).(*gormModel.User)
	if err := u.GormDB.Where("id = ?", User.Id).Updates(gormUser).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update User with id: %d", User.Id))
		return nil, err
	}
	return gormUser, nil
}

func (u UserRepositoryImpl) DeleteUser(User entity.User) (bool, error) {
	gormUser := &gormModel.User{}
	if err := u.GormDB.Where("id = ?", User.Id).Delete(&gormUser).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update User with id : %d", User.Id))
		return false, err
	}
	return true, nil
}

func (u UserRepositoryImpl) ReadUserAll() ([]gormModel.User, error) {
	gormUser := []gormModel.User{}
	if err := u.GormDB.Find(&gormUser).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to read Users"))
		return nil, err
	}
	return gormUser, nil
}
