package wire

import (
	"github.com/ESPOIR-DITE/nzedi.git/pkg/config"
	factory2 "github.com/ESPOIR-DITE/nzedi.git/pkg/db/models/gorm/factory"
	gorm2 "github.com/ESPOIR-DITE/nzedi.git/pkg/db/repository/gorm"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/entity/factory"
	controller2 "github.com/ESPOIR-DITE/nzedi.git/pkg/http/controller"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/http/server"
	"github.com/ESPOIR-DITE/nzedi.git/pkg/service/impl"

	"gorm.io/gorm"
)

func wire(config config.ServiceConfiguration, db *gorm.DB) server.HttpServer {
	accountFactory := factory2.NewAccountFactory()
	accountTypeFactory := factory2.NewAccountTypeFactory()
	accountStateFactory := factory2.NewAccountStateFactoryImpl()
	companyFactory := factory2.NewCompanyFactoryImpl()
	userFactory := factory2.NewUserFactoryImpl()
	UserTypeFactory := factory2.NewUserTypeFactoryImpl()

	accountRepository := gorm2.NewAccountRepositoryImpl(db, accountFactory)
	accountStateRepository := gorm2.NewAccountStateRepositoryImpl(db, accountStateFactory)
	accountTypeRepository := gorm2.NewAccountTypeRepositoryImpl(db, accountTypeFactory)
	companyRepository := gorm2.NewCompanyRepositoryImpl(db, companyFactory)
	userTypeRepository := gorm2.NewUserTypeRepositoryImpl(db, UserTypeFactory)
	userRepository := gorm2.NewUserRepositoryImpl(db, userFactory)

	accountService := impl.NewAccountServiceImpl(accountRepository)
	accountTypeService := impl.NewAccountTypeServiceImpl(accountTypeRepository)
	accountStateService := impl.NewAccountStateServiceImpl(accountStateRepository)
	companyService := impl.NewCompanyServiceImpl(companyRepository)
	userTypeService := impl.NewUserTypeServiceImpl(userTypeRepository)
	userService := impl.NewUserServiceImpl(userRepository)

	accountFactoryImpl := factory.NewAccountFactoryImpl()
	accountTypeFactoryImpl := factory.NewAccountTypeFactoryImpl()
	accountStateFactoryImpl := factory.NewAccountStateFactoryImpl()
	companyFactoryImpl := factory.NewCompanyFactoryImpl()
	userFactoryImpl := factory.NewUserFactoryImpl()
	userTypeFactoryImpl := factory.NewUserTypeFactoryImpl()

	controller := controller2.NzediApiController{
		AccountService:      accountService,
		AccountStateService: accountStateService,
		AccountTypeService:  accountTypeService,
		CompanyService:      companyService,
		UserTypeService:     userTypeService,
		UserService:         userService,

		AccountFactory:      accountFactoryImpl,
		AccountStateFactory: accountStateFactoryImpl,
		AccountTypeFactory:  accountTypeFactoryImpl,
		CompanyFactory:      companyFactoryImpl,
		UserFactory:         userFactoryImpl,
		UserTypeFactory:     userTypeFactoryImpl,
	}

	return server.NewHttpServerImpl(config.AppConfig(), controller)
}
