package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountType interface {
	CreateAccountTypeFactory(body spec.AccountType) (*entity.AccountType, error)
}

type AccountTypeFactoryImpl struct {
}

func (a AccountTypeFactoryImpl) CreateAccountTypeFactory(body spec.AccountType) (*entity.AccountType, error) {
	return &entity.AccountType{
		Id:           body.Id,
		AccountId:    body.AccountId,
		AccountState: body.AccountState,
		UserTypeId:   body.UserTypeId,
	}, nil
}

func NewAccountTypeFactoryImpl() *AccountTypeFactoryImpl {
	return &AccountTypeFactoryImpl{}
}

var _ AccountType = &AccountTypeFactoryImpl{}
