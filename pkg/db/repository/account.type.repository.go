package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	gormModel "github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountTypeRepository interface {
	CreateAccountType(accountType entity.AccountType) (models.AccountType, error)
	ReadAccountType(id int) (models.AccountType, error)
	ReadAccountTypeWithAccountId(id int) (models.AccountType, error)
	UpdateAccountType(accountState entity.AccountType) (models.AccountType, error)
	DeleteAccountType(accountState entity.AccountType) (bool, error)
	ReadAccountTypeAll() ([]gormModel.AccountType, error)
}
