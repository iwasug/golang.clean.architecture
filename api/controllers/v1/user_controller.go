package controllers_v1

import (
	context2 "context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.clean.architecture/api/middleware"
	"golang.clean.architecture/application/users"
	"golang.clean.architecture/application/users/models"
)

const _prefix = "api/v1/users"

func RegisterUserRoute(group *fiber.App, userService users.UserService) {
	CreateUser(group, userService)
	GetUserByObjectId(group, userService)
	UserAuth(group, userService)
}

func CreateUser(app *fiber.App, userService users.UserService) {
	group := app.Group(_prefix)
	group.Post("/", middleware.JWTProtected(), func(context *fiber.Ctx) error {

		var (
			user *models.NewUserModel
			err  error
		)

		if user, err = userService.AddNewUser(context2.Background(), user); err != nil {
			return err
		}

		return context.Status(http.StatusCreated).JSON(middleware.SuccessData("", user))
	})
}

func GetUserByObjectId(app *fiber.App, userService users.UserService) {
	group := app.Group(_prefix)
	group.Get("/:id", middleware.JWTProtected(), func(context *fiber.Ctx) error {

		var (
			user *models.NewUserModel
			err  error
		)

		id := context.Params("id")
		if user, err = userService.GetUserById(context2.Background(), id); err != nil {
			return context.Status(http.StatusBadRequest).JSON(middleware.Error(err.Error()))
		}

		return context.Status(http.StatusCreated).JSON(middleware.SuccessData("", user))
	})
}

func UserAuth(app *fiber.App, userService users.UserService) {
	group := app.Group(_prefix)
	group.Post("/login", func(context *fiber.Ctx) error {

		var (
			user      *models.NewUserModel
			err       error
			chechAuth bool
		)
		loginRequest := new(models.LoginModel)

		if err = context.BodyParser(loginRequest); err != nil {
			return context.Status(http.StatusBadRequest).JSON(middleware.Error(err.Error()))
		}

		if user, err = userService.GetUserByUsername(context2.Background(), loginRequest.Username); err != nil {
			return context.Status(http.StatusBadRequest).JSON(middleware.Error(err.Error()))
		}

		if chechAuth, err = userService.ComparePasswords(context2.Background(), loginRequest.Username, loginRequest.Password); err != nil {
			return context.Status(http.StatusBadRequest).JSON(middleware.Error(err.Error()))
		}

		if chechAuth {

		}

		return context.Status(http.StatusOK).JSON(middleware.SuccessData("", user))
	})
}
