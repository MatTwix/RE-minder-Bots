package handlers

import (
	"github.com/MatTwix/RE-minder-Bots/services"
	"github.com/gofiber/fiber/v3"
)

func GetChats(c fiber.Ctx) error {
	chats, err := services.GetChats(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error while getting chats: " + err.Error()})
	}

	return c.JSON(chats)
}
