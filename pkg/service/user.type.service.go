package service

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type UserTypeService interface {
	CreateUserType(userType entity.UserType) (*entity.UserType, error)
	ReadUserType(id string) (*entity.UserType, error)
	UpdateUserType(User entity.UserType) (*entity.UserType, error)
	DeleteUserType(User entity.UserType) (bool, error)
	ReadUserTypeAll() ([]entity.UserType, error)
}
