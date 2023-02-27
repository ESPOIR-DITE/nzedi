package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type User struct {
	Id          int       `gorm:"primarykey"`
	AccountId   int       `sql:"account_id"`
	DateOfBirth time.Time `sql:"date_of_birth"`
	Firstname   string
	Lastname    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u User) GetUser() *entity.User {
	return &entity.User{
		Id:          u.Id,
		AccountId:   u.AccountId,
		DateOfBirth: u.DateOfBirth,
		FirstName:   u.Firstname,
		LastName:    u.Lastname,
	}
}

var _ models.User = User{}

func (User) TableName() string {
	return "user"
}
