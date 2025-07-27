package vk

import (
	"log"
	"math/rand"

	"github.com/MatTwix/RE-minder-Bots/bot"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
)

var cfg = config.LoadConfig()

type VKBot struct {
	vk *api.VK
}

func New() *VKBot {
	return &VKBot{}
}

func (v *VKBot) Platform() string {
	return "vk"
}

func (v *VKBot) Start() error {
	log.Println("VK bot starting...")
	if cfg.VKToken == "" {
		log.Println("VK token not provided, skipping bot start")
		return nil
	}

	v.vk = api.NewVK(cfg.VKToken)
	log.Println("VK bot started successfully!")

	return nil
}

func (v *VKBot) SendMessage(chatID int64, message string) error {
	if v.vk == nil {
		log.Println("VK bot is not running, cannot send message")
		return nil
	}

	b := params.NewMessagesSendBuilder()
	b.PeerID(int(chatID))
	b.Message(message)
	b.RandomID(rand.Int())

	_, err := v.vk.MessagesSend(b.Params)

	return err
}

func init() {
	bot.RegisterBot(New())
}
