package users

import (
	"golang.clean.architecture/api/configs"
	appUsers "golang.clean.architecture/application/users"
	domainUsers "golang.clean.architecture/domain/users"
	"golang.clean.architecture/infrastructure/common/persistence"
)

func NewUserRepositoryResolve(config configs.Config) domainUsers.IUserRepository {
	return newUserRepository(persistence.NewPosgreSqlDb(config.ConnectionStrings.DefaultConnection))
}

func NewUserServiceResolve(config configs.Config) appUsers.UserService {
	return appUsers.NewUserService(NewUserRepositoryResolve(config))
}
