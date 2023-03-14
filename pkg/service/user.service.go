package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserService interface {
	CreateUser(User entity.User) (*entity.User, error)
	ReadUser(id string) (*entity.User, error)
	ReadUserWithAccountId(id string) (*entity.User, error)
	UpdateUser(User entity.User) (*entity.User, error)
	DeleteUser(User entity.User) (bool, error)
	ReadUserAll() ([]entity.User, error)
}
