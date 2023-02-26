package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type Account struct {
	Id        int `gorm:"primarykey"`
	Company   int
	Date      time.Time
	Email     string
	Password  string
	Token     *string
	UserName  *string `sql:"user_name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Account) GetAccount() *entity.Account {
	//TODO implement me
	panic("implement me")
}

var _ models.Account = &Account{}

func (Account) TableName() string {
	return "account"
}
