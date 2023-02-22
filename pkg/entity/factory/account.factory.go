package factory

import (
	spec "nzedi/api/server/nzedi-api"
	"nzedi/pkg/entity"
)

type AccountFactory interface {
	CreateAccount(body spec.Account) (entity.Account, error)
}
type AccountFactoryImpl struct {
}

func (a AccountFactoryImpl) CreateAccount(body spec.Account) (entity.Account, error) {
	//TODO implement me
	panic("implement me")
}

func NewAccountFactoryImpl() *AccountFactoryImpl {
	return &AccountFactoryImpl{}
}

var _ AccountFactory = &AccountFactoryImpl{}
