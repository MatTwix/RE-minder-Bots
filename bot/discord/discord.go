package discord

import (
	"log"

	"github.com/MatTwix/RE-minder-Bots/bot"
)

type DiscordBot struct {
}

func New() *DiscordBot {
	return &DiscordBot{}
}

func (d *DiscordBot) Platform() string {
	return "discord"
}

func (d *DiscordBot) Start() error {
	log.Println("Discord bot starting...")
	return nil
}

func (d *DiscordBot) SendMessage(chatID int64, message string) error {
	log.Printf("DISCORD: Sending message to %d, %s", chatID, message)
	return nil
}

func init() {
	bot.RegisterBot(New())
}
