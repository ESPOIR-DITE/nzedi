package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountFactoryImpl struct{}

func NewAccountFactory() *AccountFactoryImpl {
	return &AccountFactoryImpl{}
}

func (a AccountFactoryImpl) CreateAccount(account entity.Account) models.Account {
	return &gorm.Account{
		Id:       account.Id,
		Date:     account.Date,
		Email:    account.Email,
		Password: account.Password,
		Token:    account.Token,
		UserName: account.UserName,
	}
}

var _ factory.AccountFactory = &AccountFactoryImpl{}
