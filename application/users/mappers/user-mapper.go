package mappers

import (
	"golang.clean.architecture/application/users/models"
	"golang.clean.architecture/domain/users"
)

func MapNewUserModel(user *users.User) *models.NewUserModel {
	return &models.NewUserModel{
		Fullname: user.Fullname,
		Username: user.UserName,
		Password: user.Password,
	}
}
