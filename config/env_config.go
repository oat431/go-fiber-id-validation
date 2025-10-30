package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Error while reading .env file")
	} else {
		log.Println(".env loaded")
	}
}

func GetEnv(key string) string {
	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatal(key + " failed to load")
	}

	return value
}
