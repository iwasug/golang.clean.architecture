package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ConnectionStrings ConnectionStrings
	Host              Host
}

type ConnectionStrings struct {
	DefaultConnection string
}

type Host struct {
	Port int
}

func LoadConfig(path, env string) (config Config, err error) {

	filePath := fmt.Sprintf("%s/appsettings.%s.json", path, env)
	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
