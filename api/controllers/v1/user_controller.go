package controllers_v1

import (
	context2 "context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	apiResult "golang.clean.architecture/api/middleware"
	"golang.clean.architecture/application/users"
	"golang.clean.architecture/application/users/models"
)

const _prefix = "/users"

func RegisterRoute(group *echo.Group, userService users.UserService) {
	CreateUser(group, userService)
	GetUserByObjectId(group, userService)
}

func CreateUser(group *echo.Group, userService users.UserService) {
	path := fmt.Sprintf("%s/User", _prefix)
	group.POST(path, func(context echo.Context) error {

		var (
			user *models.NewUserModel
			err  error
		)

		if user, err = userService.AddNewUser(context2.Background(), user); err != nil {
			return err
		}

		return context.JSON(http.StatusCreated, user)
	})
}

func GetUserByObjectId(group *echo.Group, userService users.UserService) {
	path := fmt.Sprintf("%s/id/:id", _prefix)
	group.GET(path, func(context echo.Context) error {

		var (
			user *models.NewUserModel
			err  error
		)

		id := context.Param("id")
		if user, err = userService.GetUserById(context2.Background(), id); err != nil {
			return context.String(http.StatusBadRequest, err.Error())
		}

		return context.JSON(http.StatusCreated, user)
	})
}

func UserAuth(group *echo.Group, userService users.UserService) {
	path := fmt.Sprintf("%s/login", _prefix)
	group.POST(path, func(context echo.Context) error {

		var (
			user *models.NewUserModel
			err  error
		)
		loginRequest := new(models.LoginModel)
		if err = context.Bind(loginRequest); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if user, err = userService.GetUserById(context2.Background(), loginRequest.Username); err != nil {
			return context.JSON(http.StatusNotFound, apiResult.Error(err.Error()))
		}
		// if chechAuth, err = userService.AuthUser(context2.Background(), loginRequest.Username, loginRequest.Password); err != nil {
		// 	return context.JSON(http.StatusInternalServerError, apiResult.Error(err.Error()))
		// }

		// if chechAuth {
		// 	tok, time, err := apiResult.generateAccessToken(user)
		// }

		return context.JSON(http.StatusCreated, user)
	})
}
