package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type UserTypeServiceImpl struct {
	UserTypeRepository repository.UserTypeRepository
}

func NewUserTypeServiceImpl(userTypeRepository repository.UserTypeRepository) *UserTypeServiceImpl {
	return &UserTypeServiceImpl{
		UserTypeRepository: userTypeRepository,
	}
}

var _ service.UserTypeService = &UserTypeServiceImpl{}

func (u UserTypeServiceImpl) CreateUserType(userType entity.UserType) (*entity.UserType, error) {
	userT, err := u.UserTypeRepository.CreateUserType(userType)
	if err != nil {
		return nil, err
	}
	return userT.GetUserType(), nil
}

func (u UserTypeServiceImpl) ReadUserType(id int) (*entity.UserType, error) {
	userType, err := u.UserTypeRepository.ReadUserType(id)
	if err != nil {
		return nil, err
	}
	return userType.GetUserType(), nil
}

func (u UserTypeServiceImpl) UpdateUserType(User entity.UserType) (*entity.UserType, error) {
	UserType, err := u.UserTypeRepository.UpdateUserType(User)
	if err != nil {
		return nil, err
	}
	return UserType.GetUserType(), nil
}

func (u UserTypeServiceImpl) DeleteUserType(User entity.UserType) (bool, error) {
	result, err := u.UserTypeRepository.DeleteUserType(User)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (u UserTypeServiceImpl) ReadUserTypeAll() ([]entity.UserType, error) {
	userTypeList, err := u.UserTypeRepository.ReadUserTypeAll()
	if err != nil {
		return nil, err
	}
	return u.getUserTypeList(userTypeList), nil
}

func (u UserTypeServiceImpl) getUserTypeList(userTypeList []gorm.UserType) []entity.UserType {
	var userTypes []entity.UserType
	if len(userTypeList) < 0 {
		return userTypes
	}
	for _, userType := range userTypeList {
		userTypes = append(userTypes, entity.UserType{
			Id:          userType.Id,
			Name:        userType.Name,
			Description: userType.Description,
		})
	}
	return userTypes
}
