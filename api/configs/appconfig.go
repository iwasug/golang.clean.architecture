package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Database      Database
	ServerSetting ServerSetting
	Jwt           Jwt
}

type Database struct {
	DB_HOST     string
	DB_DRIVER   string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     int
}

type ServerSetting struct {
	SERVER_URL          string
	SERVER_READ_TIMEOUT int
}

type Jwt struct {
	JWT_SECRET_KEY                      string
	JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT int
}

func LoadConfig() (config Config) {

	config.Database.DB_HOST = os.Getenv("DB_HOST")
	config.Database.DB_DRIVER = os.Getenv("DB_DRIVER")
	config.Database.DB_USER = os.Getenv("DB_USER")
	config.Database.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	config.Database.DB_NAME = os.Getenv("DB_NAME")
	config.Database.DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))

	config.ServerSetting.SERVER_URL = os.Getenv("SERVER_URL")
	config.ServerSetting.SERVER_READ_TIMEOUT, _ = strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	config.Jwt.JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	config.Jwt.JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT, _ = strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	return
}

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return Fiber configuration.
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
