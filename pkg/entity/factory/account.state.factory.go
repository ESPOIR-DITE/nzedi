package factory

import (
	spec "nzedi/api/server/nzedi-api"
	"nzedi/pkg/entity"
)

type AccountStateFactory interface {
	CreateAccountState(body spec.AccountState) (*entity.AccountState, error)
}

type AccountStateFactoryImpl struct {
}

func (a AccountStateFactoryImpl) CreateAccountState(body spec.AccountState) (*entity.AccountState, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountStateFactoryImpl() *AccountStateFactoryImpl {
	return &AccountStateFactoryImpl{}
}

var _ AccountStateFactory = &AccountStateFactoryImpl{}
