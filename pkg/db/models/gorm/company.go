package gorm

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"time"
)

type Company struct {
	Id        int `gorm:"primarykey"`
	Manager   int
	Name      string
	Url       *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c Company) GetCompany() *entity.Company {
	return &entity.Company{
		Id:      c.Id,
		Manager: c.Manager,
		Name:    c.Name,
		Url:     c.Url,
	}
}

var _ models.Company = Company{}

func (Company) TableName() string {
	return "company"
}
