package vk

import (
	"log"

	"github.com/MatTwix/RE-minder-Bots/bot"
)

type VKBot struct {
}

func New() *VKBot {
	return &VKBot{}
}

func (d *VKBot) Platform() string {
	return "vk"
}

func (d *VKBot) Start() error {
	log.Println("VK bot starting...")
	return nil
}

func (d *VKBot) SendMessage(chatID int64, message string) error {
	log.Printf("VK: Sending message to %d, %s", chatID, message)
	return nil
}

func init() {
	bot.RegisterBot(New())
}
