package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountRepository interface {
	CreateAccount(account entity.Account) (models.Account, error)
	ReadAccount(id int) (models.Account, error)
	UpdateAccount(account entity.Account) (models.Account, error)
	DeleteAccount(account entity.Account) (bool, error)
	ReadAccountAll() ([]gorm.Account, error)
	LoginWithEmail(account entity.Account) (models.Account, error)
	LoginWithUserName(account entity.Account) (models.Account, error)
}
