package routes

import (
	"github.com/MatTwix/RE-minder-Bots/handlers"
	"github.com/MatTwix/RE-minder-Bots/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	chats := api.Group("/chats")
	chats.Use(middleware.APIKeyMiddleware())

	chats.Get("/", handlers.GetChats)
	chats.Put("/", handlers.SetChat)
}
