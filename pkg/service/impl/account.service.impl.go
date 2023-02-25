package impl

import (
	"errors"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
}

func NewAccountServiceImpl(accountRepository repository.AccountRepository) *AccountServiceImpl {
	return &AccountServiceImpl{AccountRepository: accountRepository}
}

var _ service.AccountService = &AccountServiceImpl{}

func (a AccountServiceImpl) CreateAccount(account entity.Account) (*entity.Account, error) {
	acc, err := a.AccountRepository.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	return acc.GetAccount(), nil
}

func (a AccountServiceImpl) ReadAccount(id int) (*entity.Account, error) {
	acc, err := a.AccountRepository.ReadAccount(id)
	if err != nil {
		return nil, err
	}
	return acc.GetAccount(), nil
}

func (a AccountServiceImpl) UpdateAccount(account entity.Account) (*entity.Account, error) {
	acc, err := a.AccountRepository.UpdateAccount(account)
	if err != nil {
		return nil, err
	}
	return acc.GetAccount(), nil
}

func (a AccountServiceImpl) DeleteAccount(account entity.Account) (bool, error) {
	acc, err := a.AccountRepository.DeleteAccount(account)
	if err != nil {
		return false, err
	}
	return acc, nil
}

func (a AccountServiceImpl) ReadAccountAll() ([]entity.Account, error) {
	acc, err := a.AccountRepository.ReadAccountAll()
	if err != nil {
		return nil, err
	}
	return a.getAccountList(acc), nil
}

func (a AccountServiceImpl) UserLogin(account entity.Account) (*entity.Account, error) {
	if account.Email != "" {
		Account, err := a.AccountRepository.LoginWithEmail(account)
		if err != nil {
			return Account.GetAccount(), err
		}
		return Account.GetAccount(), nil
	}
	if account.UserName != nil {
		Account, err := a.AccountRepository.LoginWithUserName(account)
		if err != nil {
			return Account.GetAccount(), err
		}
		return Account.GetAccount(), nil
	}
	return nil, errors.New("Missing necessary details")
}

func (a AccountServiceImpl) getAccountList(accountList []gorm.Account) []entity.Account {
	var accountAll []entity.Account
	if accountList == nil {
		return accountAll
	}
	for _, account := range accountList {
		accountAll = append(accountAll, entity.Account{
			Id:       account.Id,
			Company:  account.Company,
			Date:     account.Date,
			Email:    account.Email,
			Password: account.Password,
			Token:    account.Token,
			UserName: account.UserName,
		})
	}
	return accountAll
}
