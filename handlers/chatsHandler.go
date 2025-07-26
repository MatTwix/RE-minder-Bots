package handlers

import (
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/services"
	"github.com/gofiber/fiber/v3"
)

type ChatInput struct {
	ReminderUserID int    `json:"reminder_user_id" validate:"required,number"`
	Platform       string `json:"platform" validate:"required,oneof=discord vk telegram"`
	ChatID         int64  `json:"chat_id" validate:"required"`
}

func GetChats(c fiber.Ctx) error {
	chats, err := services.GetChats(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error while getting chats: " + err.Error()})
	}

	return c.JSON(chats)
}

func SetChat(c fiber.Ctx) error {
	var input ChatInput

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Incorrect data format: " + err.Error()})
	}

	if err := config.Validator.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation error: " + err.Error()})
	}

	createdChat, err := services.SetChat(c, input.ReminderUserID, input.Platform, input.ChatID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating chat: " + err.Error()})
	}

	return c.JSON(createdChat)
}
