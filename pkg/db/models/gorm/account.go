package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type Account struct {
	Id        string `gorm:"primarykey"`
	Date      time.Time
	Email     string
	Password  string
	Token     *string
	Username  string `sql:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a Account) GetAccount() *entity.Account {
	return &entity.Account{
		Id:       a.Id,
		Email:    a.Email,
		Date:     a.Date,
		Password: a.Password,
		Token:    a.Token,
		Username: a.Username,
	}
}

var _ models.Account = &Account{}

func (Account) TableName() string {
	return "account"
}
