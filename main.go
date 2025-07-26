package main

import (
	"github.com/MatTwix/RE-minder-Bots/bot"
	_ "github.com/MatTwix/RE-minder-Bots/bot/discord"
	_ "github.com/MatTwix/RE-minder-Bots/bot/telegram"
	_ "github.com/MatTwix/RE-minder-Bots/bot/vk"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/consumer"
	"github.com/MatTwix/RE-minder-Bots/database"
	"github.com/MatTwix/RE-minder-Bots/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	config.InitValidator()

	database.ConnectDB()
	// defer database.DB.Close()

	bot.StartAllBots()
	consumer.Start()

	routes.SetupRoutes(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
