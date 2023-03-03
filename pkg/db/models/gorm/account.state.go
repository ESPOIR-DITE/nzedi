package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type AccountState struct {
	Id          string `gorm:"primarykey"`
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (a AccountState) GetAccountState() *entity.AccountState {
	return &entity.AccountState{
		Id:          a.Id,
		Name:        a.Name,
		Description: a.Description,
	}
}

var _ models.AccountState = AccountState{}

func (AccountState) TableName() string {
	return "account_state"
}
