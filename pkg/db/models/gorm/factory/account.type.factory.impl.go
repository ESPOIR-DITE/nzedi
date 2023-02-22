package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountTypeFactoryImpl struct{}

func NewAccountTypeFactory() *AccountTypeFactoryImpl {
	return &AccountTypeFactoryImpl{}
}

func (a AccountTypeFactoryImpl) CreateAccountType(accountType entity.AccountType) models.AccountType {
	return &gorm.AccountType{
		Id:           accountType.Id,
		AccountId:    accountType.AccountId,
		AccountState: accountType.AccountState,
		UserTypeId:   accountType.UserTypeId,
	}
}

var _ factory.AccountTypeFactory = &AccountTypeFactoryImpl{}
