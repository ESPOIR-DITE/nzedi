package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserFactoryImpl struct{}

func NewUserFactoryImpl() *UserFactoryImpl {
	return &UserFactoryImpl{}
}

func (u UserFactoryImpl) CreateUser(user entity.User) models.User {
	return &gorm.User{
		Id:          user.Id,
		AccountId:   user.AccountId,
		DateOfBirth: user.DateOfBirth,
		Firstname:   user.FirstName,
		Lastname:    user.LastName,
	}
}

var _ factory.UserFactory = &UserFactoryImpl{}
