package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserTypeFactoryImpl struct{}

func NewUserTypeFactoryImpl() *UserTypeFactoryImpl {
	return &UserTypeFactoryImpl{}
}
func (u UserTypeFactoryImpl) CreateUserType(userType entity.UserType) models.UserType {
	return &gorm.UserType{
		Id:          userType.Id,
		Name:        userType.Name,
		Description: userType.Description,
	}
}

var _ factory.UserTypeFactory = &UserTypeFactoryImpl{}
