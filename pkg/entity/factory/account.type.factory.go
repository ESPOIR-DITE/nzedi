package factory

import (
	spec "nzedi/api/server/nzedi-api"
	"nzedi/pkg/entity"
)

type AccountType interface {
	CreateAccountType(body spec.AccountType) (*entity.AccountType, error)
}

type AccountTypeImpl struct {
}

func (a AccountTypeImpl) CreateAccountType(body spec.AccountType) (*entity.AccountType, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountTypeImpl() *AccountTypeImpl {
	return &AccountTypeImpl{}
}

var _ AccountType = &AccountTypeImpl{}
