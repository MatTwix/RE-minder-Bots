package telegram

import (
	"log"

	"github.com/MatTwix/RE-minder-Bots/bot"
)

type TelegramBot struct {
}

func New() *TelegramBot {
	return &TelegramBot{}
}

func (d *TelegramBot) Platform() string {
	return "telegram"
}

func (d *TelegramBot) Start() error {
	log.Println("Telegram bot starting...")
	return nil
}

func (d *TelegramBot) SendMessage(chatID int64, message string) error {
	log.Printf("Telegram: Sending message to %d, %s", chatID, message)
	return nil
}

func init() {
	bot.RegisterBot(New())
}
