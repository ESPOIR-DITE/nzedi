package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateFactoryImpl struct{}

func NewAccountStateFactoryImpl() *AccountStateFactoryImpl {
	return &AccountStateFactoryImpl{}
}

func (a AccountStateFactoryImpl) CreateAccountState(accountState entity.AccountState) models.AccountState {
	return &gorm.AccountState{
		Id:          accountState.Id,
		Name:        accountState.Name,
		Description: accountState.Description,
	}
}

var _ factory.AccountStateFactory = &AccountStateFactoryImpl{}
