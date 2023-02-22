package repository

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserTypeRepository interface {
	CreateUserType(userType entity.UserType) (models.UserType, error)
	ReadUserType(id string) (models.UserType, error)
	UpdateUserType(User entity.UserType) (models.UserType, error)
	DeleteUserType(User entity.UserType) (bool, error)
	ReadUserTypeAll() ([]models.UserType, error)
}
