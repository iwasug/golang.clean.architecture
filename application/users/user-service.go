package users

import (
	"context"

	"golang.clean.architecture/application/users/mappers"
	"golang.clean.architecture/application/users/models"
	"golang.clean.architecture/domain/users"
)

type (
	UserService interface {
		AddNewUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error)
		AddNewAdminUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error)
		AddNewGuestUser(ctx context.Context) (*models.NewUserModel, error)
		GetUserById(ctx context.Context, id string) (*models.NewUserModel, error)
		AuthUser(ctx context.Context, username, password string) (bool, error)
	}
	userService struct {
		Repository users.IUserRepository
	}
)

func NewUserService(repository users.IUserRepository) UserService {
	return &userService{Repository: repository}
}

func (service userService) GetUserById(ctx context.Context, id string) (*models.NewUserModel, error) {

	var (
		user *users.User
		err  error
	)

	if user, err = service.Repository.FindOneById(ctx, id); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AddNewUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error) {

	user := users.NewUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AddNewAdminUser(ctx context.Context, newUserModel *models.NewUserModel) (*models.NewUserModel, error) {

	user := users.NewAdminUser(newUserModel.FirstName, newUserModel.LastName, newUserModel.UserName, newUserModel.Password)

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil

}

func (service userService) AddNewGuestUser(ctx context.Context) (*models.NewUserModel, error) {
	user := users.NewGuestUser()

	if err := service.Repository.Add(ctx, user); err != nil {
		return nil, err
	}

	return mappers.MapNewUserModel(user), nil
}

func (service userService) AuthUser(ctx context.Context, username, password string) (bool, error) {

	var (
		user *users.User
		err  error
	)

	if user, err = service.Repository.FindOneByUsername(ctx, username); err != nil {
		return false, err
	}

	return users.ComparePasswords(user.Password, []byte(password)), nil
}
