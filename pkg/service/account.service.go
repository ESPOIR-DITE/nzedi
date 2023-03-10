package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountService interface {
	CreateAccount(account entity.Account) (*entity.Account, error)
	ReadAccount(id string) (*entity.Account, error)
	UpdateAccount(account entity.Account) (*entity.Account, error)
	DeleteAccount(account entity.Account) (bool, error)
	ReadAccountAll() ([]entity.Account, error)
	UpdateToken(id, token string) (*entity.Account, error)
	UserLogin(account entity.Account) (*entity.Account, error)
}
