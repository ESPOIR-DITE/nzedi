package service

import "github.com/ESPOIR-DITE/nzedi.git/pkg/entity"

type AuthenticationService interface {
	Login(account entity.Account) (*entity.Account, error)
}
