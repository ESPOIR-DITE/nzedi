package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateRepository interface {
	CreateAccountState(accountState entity.AccountState) (models.AccountState, error)
	ReadAccountState(id int) (models.AccountState, error)
	UpdateAccountState(accountState entity.AccountState) (models.AccountState, error)
	DeleteAccountState(accountState entity.AccountState) (bool, error)
	ReadAccountStateAll() ([]models.AccountState, error)
}
