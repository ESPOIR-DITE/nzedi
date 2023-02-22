package factory

import (
	"github.com/ESPOIR-DITE/ditkay-eshop/pkg/entity"
	spec "nzedi/api/server/nzedi-api"
)

type UserTypeFactory interface {
	CreateUserType(body spec.UserType) (*entity.Media, error)
}
type UserTypeFactoryImpl struct {
}

func (u UserTypeFactoryImpl) CreateUserType(body spec.UserType) (*entity.Media, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserTypeFactoryImpl() *UserTypeFactoryImpl {
	return &UserTypeFactoryImpl{}
}

var _ UserTypeFactory = UserTypeFactoryImpl{}
