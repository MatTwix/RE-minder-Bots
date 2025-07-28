package google

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/MatTwix/RE-minder-Bots/bot"
	"github.com/MatTwix/RE-minder-Bots/config"
)

var cfg = config.LoadConfig()

type GoogleEmailBot struct {
	auth smtp.Auth
}

func New() *GoogleEmailBot {
	return &GoogleEmailBot{}
}

func (b *GoogleEmailBot) Platform() string {
	return "google"
}

func (b *GoogleEmailBot) Start() error {
	log.Println("Google Email service starting..")
	if cfg.SMTPUsername == "" || cfg.SMTPPassword == "" {
		log.Println("SMTP credintials not provided,  skipping email service start...")
		return nil
	}

	b.auth = smtp.PlainAuth("", cfg.SMTPUsername, cfg.SMTPPassword, cfg.SMTPHost)

	log.Println("Google Email service started successfully")
	return nil
}

func (b *GoogleEmailBot) SendMessage(userEmail, message string) error {
	if b.auth == nil {
		log.Println("Email service is not running, cannot send message")
		return nil
	}

	fromHeader := fmt.Sprintf("From: reminder<%s>\r\n", cfg.SMTPFromEmail)
	toHeader := fmt.Sprintf("To: %s\r\n", userEmail)
	subjectHeader := "Subject: New Reminding!\r\n"

	msg := []byte(fromHeader + toHeader + subjectHeader + "\r\n" + message + "\r\n")

	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)

	err := smtp.SendMail(addr, b.auth, cfg.SMTPFromEmail, []string{userEmail}, msg)
	if err != nil {
		log.Printf("Error sending email to %s: %v", userEmail, err)
		return err
	}

	return nil
}

func init() {
	bot.RegisterBot(New())
}
