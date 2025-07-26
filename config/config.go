package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV  string
	Port string

	DatabaseURL    string
	InternalAPIKey string

	MainAppUrl string

	RabbitMQUrl string
}

func LoadConfig() Config {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load(".env"); err != nil {
			panic("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return Config{
		ENV:  os.Getenv("ENV"),
		Port: port,

		DatabaseURL:    os.Getenv("DATABASE_URL"),
		InternalAPIKey: os.Getenv("INTERNAL_API_KEY"),

		MainAppUrl: os.Getenv("MAIN_APP_URL"),

		RabbitMQUrl: os.Getenv("RABBITMQ_URL"),
	}
}
