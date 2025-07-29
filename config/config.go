package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV                string
	Port               string
	RateLimiterEnabled bool

	DatabaseURL    string
	InternalAPIKey string

	DiscordToken  string
	VKToken       string
	TelegramToken string

	MainAppUrl string

	RabbitMQUrl string

	SMTPHost      string
	SMTPPort      string
	SMTPUsername  string
	SMTPPassword  string
	SMTPFromEmail string
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
		ENV:                os.Getenv("ENV"),
		Port:               port,
		RateLimiterEnabled: os.Getenv("RATE_LIMITER_ENABLED") == "true",

		DatabaseURL:    os.Getenv("DATABASE_URL"),
		InternalAPIKey: os.Getenv("INTERNAL_API_KEY"),

		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
		VKToken:       os.Getenv("VK_TOKEN"),
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),

		MainAppUrl: os.Getenv("MAIN_APP_URL"),

		RabbitMQUrl: os.Getenv("RABBITMQ_URL"),

		SMTPHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		SMTPUsername:  os.Getenv("SMTP_USERNAME"),
		SMTPPassword:  os.Getenv("SMTP_PASSWORD"),
		SMTPFromEmail: os.Getenv("SMTP_FROM_EMAIL"),
	}
}
