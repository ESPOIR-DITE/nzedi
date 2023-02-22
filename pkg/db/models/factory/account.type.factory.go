package factory

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountTypeFactory interface {
	CreateAccountType(accountType entity.AccountType) models.AccountType
}
