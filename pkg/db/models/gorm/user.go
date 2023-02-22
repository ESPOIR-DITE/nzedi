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
}

func (u User) GetUser() *entity.User {
	//TODO implement me
	panic("implement me")
}

var _ models.User = User{}
