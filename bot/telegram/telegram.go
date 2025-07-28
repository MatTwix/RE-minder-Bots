package telegram

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MatTwix/RE-minder-Bots/bot"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/services"
	"gopkg.in/telebot.v3"
)

var cfg = config.LoadConfig()

type TelegramBot struct {
	b *telebot.Bot
}

func New() *TelegramBot {
	return &TelegramBot{}
}

func (t *TelegramBot) Platform() string {
	return "telegram"
}

func (t *TelegramBot) Start() error {
	if cfg.TelegramToken == "" {
		log.Printf("Telegram token not provided, skkipping bot start")
		return nil
	}

	var err error

	pref := telebot.Settings{
		Token:  cfg.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	t.b, err = telebot.NewBot(pref)
	if err != nil {
		return err
	}

	t.b.Handle("/start", t.startCommandHandler)

	log.Println("Telegram bot starting listener...")
	go t.b.Start()

	return nil
}

func (t *TelegramBot) startCommandHandler(c telebot.Context) error {
	args := c.Args()
	if len(args) != 1 {
		return c.Send("Wrong format. Please, use command like: /start <your reminder_user_id>")
	}

	userIDStr := args[0]
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return c.Send("ID must be number")
	}

	platform := t.Platform()

	_, err = services.SetChat(context.Background(), userID, platform, strconv.Itoa(int(c.Chat().ID)))
	if err != nil {
		log.Printf("Failed to create chat link for user %d: %v", userID, err)
		return c.Send("Internat server error. Please, try later")
	}

	successMessage := fmt.Sprintf("Your account successfully linked to reminder user %d!", userID)
	return c.Send(successMessage)
}

func (t *TelegramBot) SendMessage(chatID string, message string) error {
	if t.b == nil {
		return errors.New("telegram bot is not initialized")
	}

	recipientID, err := strconv.Atoi(chatID)
	if err != nil {
		return errors.New("invalid chat id")
	}

	recipient := &telebot.User{ID: int64(recipientID)}
	_, err = t.b.Send(recipient, message)
	return err
}

func init() {
	bot.RegisterBot(New())
}
