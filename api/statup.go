package api

import (
	context2 "context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.clean.architecture/api/configs"
	controllers_v1 "golang.clean.architecture/api/controllers/v1"
	"golang.clean.architecture/application/users/consumers"
	common_di "golang.clean.architecture/infrastructure/common"
	infUsers "golang.clean.architecture/infrastructure/users"
)

func Init() {

	var (
		config configs.Config
		err    error
	)

	if config, err = configs.LoadConfig("./api", os.Getenv("Environment")); err != nil {
		panic(err)
	}

	var userService = infUsers.NewUserServiceResolve(config)

	BindConsumers(config)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		return userService.AuthUser(context2.Background(), username, password)
	}))

	v1 := e.Group("/api/v1")
	controllers_v1.CreateGuestUser(v1, userService)
	controllers_v1.GetUserByObjectId(v1, userService)

	e.Start(":8080")
}

func BindConsumers(config configs.Config) {
	rbt := common_di.NewRabbitMQResolve(config)
	rbt.BindConsumer(consumers.NewUserCreatedConsumer())
	rbt.Start()
}
