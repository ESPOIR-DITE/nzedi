package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserTypeFactory interface {
	CreateUserType(body spec.UserType) (*entity.UserType, error)
}
type UserTypeFactoryImpl struct {
}

func (u UserTypeFactoryImpl) CreateUserType(body spec.UserType) (*entity.UserType, error) {
	return &entity.UserType{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
	}, nil
}

func NewUserTypeFactoryImpl() *UserTypeFactoryImpl {
	return &UserTypeFactoryImpl{}
}

var _ UserTypeFactory = UserTypeFactoryImpl{}
