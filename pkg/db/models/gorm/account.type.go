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

func (a AccountType) GetAccountType() *entity.AccountType {
	return &entity.AccountType{
		Id:           a.Id,
		AccountId:    a.AccountId,
		UserTypeId:   a.UserTypeId,
		AccountState: a.AccountState,
	}
}

var _ models.AccountType = AccountType{}

func (AccountType) TableName() string {
	return "account_type"
}
