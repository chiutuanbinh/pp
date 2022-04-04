package utils

import (
	"log"

	"github.com/spf13/viper"
)

type config struct {
	ConnStr string
}

var Config config

func init() {
	viper.AddConfigPath("../")
	viper.AddConfigPath("./")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("ERROR READING CONFIG FILE")
	}
	Config = config{}
	Config.ConnStr = viper.GetString("POSTGRES_CONNECTION_STRING") //"postgresql://postgres:123456@0.0.0.0:5432/postgres?sslmode=disable"
}
