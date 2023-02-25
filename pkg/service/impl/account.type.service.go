package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type AccountTypeServiceImpl struct {
	AccountTypeRepository repository.AccountTypeRepository
}

func NewAccountTypeServiceImpl(accountTypeRepository repository.AccountTypeRepository) *AccountTypeServiceImpl {
	return &AccountTypeServiceImpl{AccountTypeRepository: accountTypeRepository}
}

var _ service.AccountTypeService = &AccountTypeServiceImpl{}

func (a AccountTypeServiceImpl) CreateAccountType(accountType entity.AccountType) (*entity.AccountType, error) {
	AccountType, err := a.AccountTypeRepository.CreateAccountType(accountType)
	if err != nil {
		return nil, err
	}
	return AccountType.GetAccountType(), nil
}

func (a AccountTypeServiceImpl) ReadAccountType(id int) (*entity.AccountType, error) {
	AccountType, err := a.AccountTypeRepository.ReadAccountType(id)
	if err != nil {
		return nil, err
	}
	return AccountType.GetAccountType(), nil
}

func (a AccountTypeServiceImpl) ReadAccountTypeWithAccountId(id int) (*entity.AccountType, error) {
	AccountType, err := a.AccountTypeRepository.ReadAccountTypeWithAccountId(id)
	if err != nil {
		return nil, err
	}
	return AccountType.GetAccountType(), nil
}

func (a AccountTypeServiceImpl) UpdateAccountType(accountState entity.AccountType) (*entity.AccountType, error) {
	AccountType, err := a.AccountTypeRepository.UpdateAccountType(accountState)
	if err != nil {
		return nil, err
	}
	return AccountType.GetAccountType(), nil
}

func (a AccountTypeServiceImpl) DeleteAccountType(accountState entity.AccountType) (bool, error) {
	AccountType, err := a.AccountTypeRepository.DeleteAccountType(accountState)
	if err != nil {
		return false, err
	}
	return AccountType, nil
}

func (a AccountTypeServiceImpl) ReadAccountTypeAll() ([]entity.AccountType, error) {
	AccountTypes, err := a.AccountTypeRepository.ReadAccountTypeAll()
	if err != nil {
		return nil, err
	}
	return a.getAccountTypeList(AccountTypes), nil
}

func (a AccountTypeServiceImpl) getAccountTypeList(accountList []gorm.AccountType) []entity.AccountType {
	var accountTypeAll []entity.AccountType
	if accountList == nil {
		return accountTypeAll
	}
	for _, company := range accountList {
		accountTypeAll = append(accountTypeAll, entity.AccountType{
			Id:           company.Id,
			AccountId:    company.AccountId,
			AccountState: company.AccountState,
			UserTypeId:   company.UserTypeId,
		})
	}
	return accountTypeAll
}
