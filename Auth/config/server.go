package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port   string
	Secret string

	JWT_ACCESS_SECRET  string
	JWT_REFRESH_SECRET string
	JWT_ACCESS_TIMER   string
	JWT_REFRESH_TIMER  string

	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_SSL      string
	DB_DEBUG    bool
}

func ServerConfig() string {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8000")

	appServer := fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	log.Print("Server Running at :", appServer)
	return appServer
}
