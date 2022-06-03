package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.clean.architecture/api/configs"
	"golang.clean.architecture/api/middleware"
)

func Init() {

	var (
		configSetting configs.Config
	)

	configSetting = configs.LoadConfig()

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	//var userService = infUsers.NewUserServiceResolve(configSetting.Database)
	// e := echo.New()
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// }))

	// //Health Check
	// e.GET("/api/healthchecks/status", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "API Server is running..")
	// })

	// //User
	// v1 := e.Group("/api/v1")

	// controllers_v1.UserAuth(v1, userService)
	// v1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	Claims:                  &auth.Claims{},
	// 	SigningKey:              []byte(auth.GetJWTSecret()),
	// 	TokenLookup:             "cookie:access-token",
	// 	ErrorHandlerWithContext: auth.JWTErrorChecker,
	// }))

	// //Register Route User
	// controllers_v1.RegisterRoute(v1, userService)

	// //Start Api
	// port := strconv.Itoa(config.Host.Port)
	// e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))

	if err := app.Listen(configSetting.ServerSetting.SERVER_URL); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
