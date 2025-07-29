package main

import (
	"github.com/MatTwix/RE-minder-Bots/bot"
	_ "github.com/MatTwix/RE-minder-Bots/bot/discord"
	_ "github.com/MatTwix/RE-minder-Bots/bot/google"
	_ "github.com/MatTwix/RE-minder-Bots/bot/telegram"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/consumer"
	"github.com/MatTwix/RE-minder-Bots/database"
	"github.com/MatTwix/RE-minder-Bots/routes"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	config.InitValidator()

	database.ConnectDB()
	// defer database.DB.Close()

	bot.StartAllBots()
	consumer.Start()

	allowedOrigins := []string{
		cfg.MainAppUrl,
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           86400,
	}))

	if cfg.RateLimiterEnabled {
		app.Use(limiter.New(limiter.Config{
			Max:        1000,
			Expiration: 60 * 1000,
			LimitReached: func(c fiber.Ctx) error {
				return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
					"error": "Too many requests, please try later.",
				})
			},
		}))
	}

	routes.SetupRoutes(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
