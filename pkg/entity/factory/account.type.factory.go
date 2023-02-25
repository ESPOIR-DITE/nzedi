package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountType interface {
	CreateAccountType(body spec.AccountType) (*entity.AccountType, error)
}

type AccountTypeImpl struct {
}

func (a AccountTypeImpl) CreateAccountType(body spec.AccountType) (*entity.AccountType, error) {
	return &entity.AccountType{
		Id:           body.Id,
		AccountId:    body.AccountId,
		AccountState: body.AccountState,
		UserTypeId:   body.UserTypeId,
	}, nil
}

func NewAccountTypeImpl() *AccountTypeImpl {
	return &AccountTypeImpl{}
}

var _ AccountType = &AccountTypeImpl{}
