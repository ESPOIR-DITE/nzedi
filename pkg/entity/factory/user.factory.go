package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

const ISO8601 = "2006-01-02T15:04:05Z0700"
const YYYYMMDD = "2006-01-02"

type UserFactory interface {
	CreateUser(body spec.User) (*entity.User, error)
}

type UserFactoryImpl struct {
}

func (u UserFactoryImpl) CreateUser(body spec.User) (*entity.User, error) {
	dateOfBirth, err := time.Parse(YYYYMMDD, *body.DateOfBirth)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		Id:          body.Id,
		AccountId:   body.AccountId,
		DateOfBirth: dateOfBirth,
		FirstName:   body.FirstName,
		LastName:    body.LastName,
	}, nil
}

func NewUserFactoryImpl() *UserFactoryImpl {
	return &UserFactoryImpl{}
}

var _ UserFactory = &UserFactoryImpl{}
