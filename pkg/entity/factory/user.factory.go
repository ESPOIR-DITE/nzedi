package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserFactory interface {
	CreateUser(body spec.User) (*entity.User, error)
}

type UserFactoryImpl struct {
}

func (u UserFactoryImpl) CreateUser(body spec.User) (*entity.User, error) {
	return &entity.User{
		Id:          body.Id,
		AccountId:   body.AccountId,
		DateOfBirth: *body.DateOfBirth,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
	}, nil
}

func NewUserFactoryImpl() *UserFactoryImpl {
	return &UserFactoryImpl{}
}

var _ UserFactory = &UserFactoryImpl{}
