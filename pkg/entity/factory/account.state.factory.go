package factory

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateFactory interface {
	CreateAccountState(body spec.AccountState) (*entity.AccountState, error)
}

type AccountStateFactoryImpl struct {
}

func (a AccountStateFactoryImpl) CreateAccountState(body spec.AccountState) (*entity.AccountState, error) {
	return &entity.AccountState{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
	}, nil
}

func NewAccountStateFactoryImpl() *AccountStateFactoryImpl {
	return &AccountStateFactoryImpl{}
}

var _ AccountStateFactory = &AccountStateFactoryImpl{}
