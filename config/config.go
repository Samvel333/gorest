// Configuration module

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host 	 string
	Port     string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
}

// This function reads configs from .env file, which is in base path
func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error when Loading env file")
	}

	return &Config{
		Host:   os.Getenv("HOST"),
		Port:   os.Getenv("PORT"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}
}