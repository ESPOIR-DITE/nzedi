package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	gormModel "github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateRepository interface {
	CreateAccountState(accountState entity.AccountState) (models.AccountState, error)
	ReadAccountState(id string) (models.AccountState, error)
	UpdateAccountState(accountState entity.AccountState) (models.AccountState, error)
	DeleteAccountState(accountState entity.AccountState) (bool, error)
	ReadAccountStateAll() ([]gormModel.AccountState, error)
}
