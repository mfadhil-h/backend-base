package config

import (
	"log"

	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️  No .env file found, relying on environment variables")
	}
}
