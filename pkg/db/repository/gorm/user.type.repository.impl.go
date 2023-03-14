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

type UserTypeRepositoryImpl struct {
	GormDB          *gorm.DB
	UserTypeFactory factory.UserTypeFactory
}

func NewUserTypeRepositoryImpl(gormDB *gorm.DB,
	userTypeFactory factory.UserTypeFactory) *UserTypeRepositoryImpl {
	return &UserTypeRepositoryImpl{
		GormDB:          gormDB,
		UserTypeFactory: userTypeFactory,
	}
}

func (u UserTypeRepositoryImpl) CreateUserType(userType entity.UserType) (models.UserType, error) {
	var gormUserType *gormModel.UserType = u.UserTypeFactory.CreateUserType(userType).(*gormModel.UserType)
	if err := u.GormDB.Create(&gormUserType).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to create UserType with name %d, err: %s", userType.Name, err))
		return nil, err
	}
	return gormUserType, nil
}

func (u UserTypeRepositoryImpl) ReadUserType(id string) (models.UserType, error) {
	gormUserType := &gormModel.UserType{}
	if err := u.GormDB.Where("id = ?", id).First(&gormUserType).Error; err != nil {
		logger.Log.Error(fmt.Printf("faile to get UserType with id: %d, err: %s", id, err))
		return nil, err
	}
	return gormUserType, nil
}

func (u UserTypeRepositoryImpl) UpdateUserType(UserType entity.UserType) (models.UserType, error) {
	var gormUserType *gormModel.UserType = u.UserTypeFactory.CreateUserType(UserType).(*gormModel.UserType)
	if err := u.GormDB.Where("id = ?", UserType.Id).Updates(gormUserType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update UserType with id: %d", UserType.Id))
		return nil, err
	}
	return gormUserType, nil
}

func (u UserTypeRepositoryImpl) DeleteUserType(User entity.UserType) (bool, error) {
	gormUser := &gormModel.UserType{}
	if err := u.GormDB.Where("id = ?", User.Id).Delete(&gormUser).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to update UserType with id : %d", User.Id))
		return false, err
	}
	return true, nil
}

func (u UserTypeRepositoryImpl) ReadUserTypeAll() ([]gormModel.UserType, error) {
	gormUserType := []gormModel.UserType{}
	if err := u.GormDB.Find(&gormUserType).Error; err != nil {
		logger.Log.Error(fmt.Errorf("failed to read UserTypes"))
		return nil, err
	}
	return gormUserType, nil
}

var _ repository.UserTypeRepository = &UserTypeRepositoryImpl{}
