package users

import (
	"golang.clean.architecture/api/configs"
	appUsers "golang.clean.architecture/application/users"
	domainUsers "golang.clean.architecture/domain/users"
	common_di "golang.clean.architecture/infrastructure/common"
	"golang.clean.architecture/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Config) domainUsers.IUserRepository {
	rbt := common_di.NewRabbitMQResolve(config)
	eventHandler := common_di.NewEventHandlerResolve(rbt)
	return newUserRepository(persistence.NewMongoDb(config.User.MongoDb, config.User.Database), eventHandler)
}

func NewUserServiceResolve(config configs.Config) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}
