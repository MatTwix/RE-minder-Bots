package discord

import (
	"log"
	"strconv"

	"github.com/MatTwix/RE-minder-Bots/bot"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/bwmarrin/discordgo"
)

var cfg = config.LoadConfig()

type DiscordBot struct {
	session *discordgo.Session
}

func New() *DiscordBot {
	return &DiscordBot{}
}

func (d *DiscordBot) Platform() string {
	return "discord"
}

func (d *DiscordBot) Start() error {
	log.Println("Discord bot starting...")
	if cfg.DiscordToken == "" {
		log.Println("Discord token not provided, skipping bot start")
		return nil
	}

	var err error

	d.session, err = discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		return err
	}

	err = d.session.Open()
	if err != nil {
		return err
	}

	log.Println("Discord bot started successfully!")

	return nil
}

func (d *DiscordBot) SendMessage(userID int64, message string) error {
	if d.session == nil {
		log.Println("Discord bot is not running, cannot send message")
		return nil
	}

	channel, err := d.session.UserChannelCreate(strconv.Itoa(int(userID)))
	if err != nil {
		return err
	}

	_, err = d.session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	bot.RegisterBot(New())
}
