package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserRepository interface {
	CreateUser(User entity.User) (models.User, error)
	ReadUser(id string) (models.User, error)
	ReadUserWithAccountId(id string) (models.User, error)
	UpdateUser(User entity.User) (models.User, error)
	DeleteUser(User entity.User) (bool, error)
	ReadUserAll() ([]models.User, error)
}
