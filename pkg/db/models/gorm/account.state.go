package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type AccountState struct {
	Id          int `gorm:"primarykey"`
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (a AccountState) GetAccountState() *entity.AccountState {
	//TODO implement me
	panic("implement me")
}

var _ models.AccountState = AccountState{}