package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	appConfig := &Config{
		DBUrl: getEnv("DBUrl"),
		Port:  getEnv("PORT"),
	}
	return appConfig
}

func getEnv(key string) string {
	return os.Getenv(key)
}
