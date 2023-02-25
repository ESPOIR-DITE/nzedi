package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
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
	return &entity.Account{
		Id:       body.Id,
		Company:  body.Company,
		Date:     body.Date,
		Email:    body.Email,
		Password: body.Password,
		Token:    body.Token,
		UserName: body.UserName,
	}, nil
}
