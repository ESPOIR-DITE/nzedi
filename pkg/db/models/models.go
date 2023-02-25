package models

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type Account interface {
	GetAccount() *entity.Account
}

type AccountState interface {
	GetAccountState() *entity.AccountState
}

type AccountType interface {
	GetAccountType() *entity.AccountType
}

type Company interface {
	GetCompany() *entity.Company
}

type User interface {
	GetUser() *entity.User
}

type UserType interface {
	GetUserType() *entity.UserType
}
