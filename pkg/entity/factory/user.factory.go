package factory

import (
	spec "nzedi/api/server/nzedi-api"
	"nzedi/pkg/entity"
)

type UserFactory interface {
	CreateUser(body spec.User) (*entity.User, error)
}

type UserFactoryImpl struct {
}

func (u UserFactoryImpl) CreateUser(body spec.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserFactoryImpl() *UserFactoryImpl {
	return &UserFactoryImpl{}
}

var _ UserFactory = &UserFactoryImpl{}
