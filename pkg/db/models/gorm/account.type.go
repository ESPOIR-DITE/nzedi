package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type AccountType struct {
	Id           string `gorm:"primarykey"`
	AccountId    string `sql:"account_id"`
	AccountState string `sql:"account_state"`
	UserTypeId   string `sql:"user_type_id"`
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
