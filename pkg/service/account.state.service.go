package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateService interface {
	CreateAccountState(accountState entity.AccountState) (*entity.AccountState, error)
	ReadAccountState(id string) (*entity.AccountState, error)
	UpdateAccountState(accountState entity.AccountState) (*entity.AccountState, error)
	DeleteAccountState(accountState entity.AccountState) (bool, error)
	ReadAccountStateAll() ([]entity.AccountState, error)
}
