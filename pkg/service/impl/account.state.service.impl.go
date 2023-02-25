package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
)

type AccountStateServiceImpl struct {
	AccountStateRepository repository.AccountStateRepository
}

func NewAccountStateServiceImpl(accountStateRepository repository.AccountStateRepository) *AccountStateServiceImpl {
	return &AccountStateServiceImpl{
		AccountStateRepository: accountStateRepository,
	}
}

var _ repository.AccountStateRepository = &AccountStateServiceImpl{}

func (a AccountStateServiceImpl) CreateAccountState(accountState entity.AccountState) (*entity.AccountState, error) {
	AccountState, err := a.AccountStateRepository.CreateAccountState(accountState)
	if err != nil {
		return nil, err
	}
	return AccountState, nil
}

func (a AccountStateServiceImpl) ReadAccountState(id int) (*entity.AccountState, error) {
	AccountState, err := a.AccountStateRepository.ReadAccountState(id)
	if err != nil {
		return nil, err
	}
	return AccountState, nil
}

func (a AccountStateServiceImpl) UpdateAccountState(accountState entity.AccountState) (*entity.AccountState, error) {
	AccountState, err := a.AccountStateRepository.UpdateAccountState(accountState)
	if err != nil {
		return nil, err
	}
	return AccountState, nil
}

func (a AccountStateServiceImpl) DeleteAccountState(accountState entity.AccountState) (bool, error) {
	AccountState, err := a.AccountStateRepository.DeleteAccountState(accountState)
	if err != nil {
		return false, err
	}
	return AccountState, nil
}

func (a AccountStateServiceImpl) ReadAccountStateAll() ([]entity.AccountState, error) {
	AccountState, err := a.AccountStateRepository.ReadAccountStateAll()
	if err != nil {
		return nil, err
	}
	return a.getAccountStateList(AccountState), nil
}

func (a AccountStateServiceImpl) getAccountStateList(accountList []entity.AccountState) []entity.AccountState {
	var accountStateAll []entity.AccountState
	if accountList == nil {
		return accountStateAll
	}
	for _, accountState := range accountList {
		accountStateAll = append(accountStateAll, entity.AccountState{
			Id:          accountState.Id,
			Name:        accountState.Name,
			Description: accountState.Description,
		})
	}
	return accountStateAll
}
