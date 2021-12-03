package mappers

import (
	"golang.clean.architecture/application/users/models"
	"golang.clean.architecture/domain/users"
)

func MapNewUserModel(user *users.User) *models.NewUserModel {
	return &models.NewUserModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
	}
}
