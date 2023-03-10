package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountTypeService interface {
	CreateAccountType(accountType entity.AccountType) (*entity.AccountType, error)
	ReadAccountType(id string) (*entity.AccountType, error)
	ReadAccountTypeWithAccountId(id string) (*entity.AccountType, error)
	UpdateAccountType(accountState entity.AccountType) (*entity.AccountType, error)
	DeleteAccountType(accountState entity.AccountType) (bool, error)
	ReadAccountTypeAll() ([]entity.AccountType, error)
}
