package users

import (
	"golang.clean.architecture/api/configs"
	appUsers "golang.clean.architecture/application/users"
	domainUsers "golang.clean.architecture/domain/users"
	"golang.clean.architecture/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Database) domainUsers.IUserRepository {
	return newUserRepository(persistence.NewConnectionDb(config))
}

func NewUserServiceResolve(config configs.Database) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}
