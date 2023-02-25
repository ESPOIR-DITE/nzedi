package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateRepository interface {
	CreateAccountState(accountState entity.AccountState) (*entity.AccountState, error)
	ReadAccountState(id int) (*entity.AccountState, error)
	UpdateAccountState(accountState entity.AccountState) (*entity.AccountState, error)
	DeleteAccountState(accountState entity.AccountState) (bool, error)
	ReadAccountStateAll() ([]entity.AccountState, error)
}
