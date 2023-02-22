package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type Company struct {
	Id      int `gorm:"primarykey"`
	Manager int
	Name    string
	Url     *string
}

func (c Company) GetCompany() *entity.Company {
	//TODO implement me
	panic("implement me")
}

var _ models.Company = Company{}
