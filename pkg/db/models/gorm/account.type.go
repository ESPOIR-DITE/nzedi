package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type AccountType struct {
	Id           int `gorm:"primarykey"`
	AccountId    int `sql:"account_id"`
	AccountState int `sql:"account_state"`
	UserTypeId   int `sql:"user_type_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (a AccountType) GetAccountType() *entity.Account {
	//TODO implement me
	panic("implement me")
}

var _ models.AccountType = AccountType{}