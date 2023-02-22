package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserType struct {
	Id          int `gorm:"primarykey"`
	Name        string
	Description *string
}

func (u UserType) GetUserType() *entity.UserType {
	//TODO implement me
	panic("implement me")
}

var _ models.UserType = UserType{}
