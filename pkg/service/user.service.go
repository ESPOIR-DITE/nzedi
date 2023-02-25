package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserService interface {
	CreateUser(User entity.User) (*entity.User, error)
	ReadUser(id int) (*entity.User, error)
	ReadUserWithAccountId(id int) (*entity.User, error)
	UpdateUser(User entity.User) (*entity.User, error)
	DeleteUser(User entity.User) (bool, error)
	ReadUserAll() ([]entity.User, error)
}
