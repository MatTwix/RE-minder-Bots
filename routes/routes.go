package routes

import (
	"github.com/MatTwix/RE-minder-Bots/handlers"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	chats := api.Group("/chats")
	chats.Get("/", handlers.GetChats)
}
