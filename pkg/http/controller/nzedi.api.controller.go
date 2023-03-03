package controller

import (
	spec "github.com/ESPOIR-DITE/nzedi.git/api/server/nzedi-api"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity/factory"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service"
)

type NzediApiController struct {
	AccountService      service.AccountService
	AccountStateService service.AccountStateService
	AccountTypeService  service.AccountTypeService
	CompanyService      service.CompanyService
	UserService         service.UserService
	UserTypeService     service.UserTypeService

	AccountFactory        factory.AccountFactory
	AccountStateFactory   factory.AccountStateFactory
	AccountTypeFactory    factory.AccountType
	CompanyFactory        factory.CompanyFactory
	UserFactory           factory.UserFactory
	UserTypeFactory       factory.UserTypeFactory
	AuthenticationService service.AuthenticationService
}

func NewNzediApiController(
	accountService service.AccountService,
	accountStateService service.AccountStateService,
	accountTypeService service.AccountTypeService,
	companyService service.CompanyService,
	userService service.UserService,
	userTypeService service.UserTypeService,
	accountFactory factory.AccountFactory,
	authenticationService service.AuthenticationService,
) *NzediApiController {
	return &NzediApiController{
		AccountService:        accountService,
		AccountStateService:   accountStateService,
		AccountTypeService:    accountTypeService,
		CompanyService:        companyService,
		UserService:           userService,
		UserTypeService:       userTypeService,
		AccountFactory:        accountFactory,
		AuthenticationService: authenticationService,
	}
}

var _ spec.ServerInterface = &NzediApiController{}
