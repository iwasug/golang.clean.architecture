package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.clean.architecture/api/configs"
	controllers_v1 "golang.clean.architecture/api/controllers/v1"
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

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	//Health Check
	e.GET("/api/healthchecks/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "API Server is running..")
	})

	// e.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
	// 	return userService.AuthUser(context2.Background(), username, password)
	// }))

	//User
	v1 := e.Group("/api/v1")
	controllers_v1.CreateUser(v1, userService)
	controllers_v1.GetUserByObjectId(v1, userService)

	//Start Api
	port := strconv.Itoa(config.Host.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
