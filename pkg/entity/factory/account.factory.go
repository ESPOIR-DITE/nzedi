package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type AccountFactory interface {
	CreateAccount(body spec.Account) (*entity.Account, error)
}
type AccountFactoryImpl struct {
}

func NewAccountFactoryImpl() *AccountFactoryImpl {
	return &AccountFactoryImpl{}
}

var _ AccountFactory = &AccountFactoryImpl{}

func (a AccountFactoryImpl) CreateAccount(body spec.Account) (*entity.Account, error) {
	if body.UserName == nil {
		body.UserName = &body.Email
	}
	date, err := time.Parse(ISO8601, body.Date)
	if err != nil {
		return nil, err
	}
	return &entity.Account{
		Id:       body.Id,
		Date:     date,
		Email:    body.Email,
		Password: body.Password,
		Token:    body.Token,
		Username: *body.UserName,
	}, nil
}
