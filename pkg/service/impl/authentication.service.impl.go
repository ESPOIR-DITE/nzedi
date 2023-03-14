package impl

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type AuthenticationServiceImpl struct {
	AccountService  service.AccountService
	SecurityService service.JwtSecurityService
}

func NewAuthenticationServiceImpl(accountService service.AccountService,
	securityService service.JwtSecurityService) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{
		AccountService:  accountService,
		SecurityService: securityService,
	}
}

var _ service.AuthenticationService = &AuthenticationServiceImpl{}

func (a AuthenticationServiceImpl) Login(account entity.Account) (*entity.Account, error) {
	accountResult, err := a.AccountService.UserLogin(account)
	if err != nil {
		return nil, err
	}
	token, err := a.SecurityService.GenerateJWTToken(account.Email)
	if err != nil {
		return nil, err
	}
	if token != "" {
		accountResult.Token = &token
	}
	updatedAccount, err := a.AccountService.UpdateToken(accountResult.Id, token)
	if err != nil {
		return nil, err
	}
	if updatedAccount.Token != nil {
		updatedAccount, err = a.AccountService.ReadAccount(accountResult.Id)
		if err != nil {
			return nil, err
		}
	}
	return updatedAccount, nil
}
