package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{UserRepository: userRepository}
}

var _ service.UserService = &UserServiceImpl{}

func (u UserServiceImpl) CreateUser(User entity.User) (*entity.User, error) {
	user, err := u.UserRepository.CreateUser(User)
	if err != nil {
		return nil, err
	}
	return user.GetUser(), nil
}

func (u UserServiceImpl) ReadUser(id string) (*entity.User, error) {
	user, err := u.UserRepository.ReadUser(id)
	if err != nil {
		return nil, err
	}
	return user.GetUser(), nil
}

func (u UserServiceImpl) ReadUserWithAccountId(id string) (*entity.User, error) {
	user, err := u.UserRepository.ReadUserWithAccountId(id)
	if err != nil {
		return nil, err
	}
	return user.GetUser(), nil
}

func (u UserServiceImpl) UpdateUser(User entity.User) (*entity.User, error) {
	user, err := u.UserRepository.UpdateUser(User)
	if err != nil {
		return nil, err
	}
	return user.GetUser(), nil
}

func (u UserServiceImpl) DeleteUser(User entity.User) (bool, error) {
	result, err := u.UserRepository.DeleteUser(User)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (u UserServiceImpl) ReadUserAll() ([]entity.User, error) {
	users, err := u.UserRepository.ReadUserAll()
	if err != nil {
		return nil, err
	}
	return u.getUserList(users), nil
}

func (u UserServiceImpl) getUserList(userList []gorm.User) []entity.User {
	var userAll []entity.User
	if userList == nil {
		return userAll
	}
	for _, users := range userList {
		userAll = append(userAll, entity.User{
			Id:          users.Id,
			AccountId:   users.AccountId,
			DateOfBirth: users.DateOfBirth,
			FirstName:   users.Lastname,
			LastName:    users.Lastname,
		})
	}
	return userAll
}
