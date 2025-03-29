package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DbHost   string
	DbPort   string
	DbUser   string
	DbPass   string
	DbName   string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	return &Config{
		Port:   os.Getenv("PORT"),
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASSWORD"),
		DbName: os.Getenv("DB_NAME"),
	}
}