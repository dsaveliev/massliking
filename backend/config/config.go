package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const RELEASE = "production"
const DEVELOP = "development"

var APP_ENV string

func Init() {
	APP_ENV = os.Getenv("APP_ENV")

	viper.SetConfigName(APP_ENV)
	viper.AddConfigPath("./config/")
	viper.AddConfigPath("/srv/release/config/")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}

func GetInt(key string) int {
	return viper.GetInt(key)
}

func IsRelease() bool {
	return APP_ENV == RELEASE
}

func IsDevelop() bool {
	return APP_ENV == DEVELOP
}
