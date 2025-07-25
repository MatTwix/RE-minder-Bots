package main

import (
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/consumer"
	"github.com/MatTwix/RE-minder-Bots/database"
	"github.com/MatTwix/RE-minder-Bots/routes"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	cfg := config.LoadConfig()

	database.ConnectDB()
	// defer database.DB.Close()

	consumer.Start()

	routes.SetupRoutes(app)

	if err := app.Listen(":" + cfg.Port); err != nil {
		panic("Error starting server: " + err.Error())
	}
}
