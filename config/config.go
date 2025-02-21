package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	SSLMode          string
	ServerPort       string
	JWTSecret        string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️ No .env file found, using environment variables")
	}

	AppConfig = Config{
		DatabaseHost:     viper.GetString("DATABASE_HOST"),
		DatabasePort:     viper.GetString("DATABASE_PORT"),
		DatabaseUser:     viper.GetString("DATABASE_USER"),
		DatabasePassword: viper.GetString("DATABASE_PASSWORD"),
		DatabaseName:     viper.GetString("DATABASE_NAME"),
		SSLMode:          viper.GetString("SSL_MODE"),
		ServerPort:       viper.GetString("PORT"),
		JWTSecret:        viper.GetString("JWT_SECRET"),
	}
}
