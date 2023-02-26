package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type UserType struct {
	Id          int `gorm:"primarykey"`
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u UserType) GetUserType() *entity.UserType {
	//TODO implement me
	panic("implement me")
}

var _ models.UserType = UserType{}

func (UserType) TableName() string {
	return "user_type"
}
