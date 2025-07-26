package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/MatTwix/RE-minder-Bots/bot"
	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/MatTwix/RE-minder-Bots/services"
	"github.com/rabbitmq/amqp091-go"
)

type NotificationTask struct {
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

type NotificationSettings struct {
	ID                   int  `json:"id"`
	UserID               int  `json:"user_id"`
	TelegramNotification bool `json:"telegram_notification"`
	DiscordNotification  bool `json:"discord_notification"`
	VKNotification       bool `json:"vk_notification"`
}

func Start() {
	cfg := config.LoadConfig()

	conn, err := amqp091.Dial(cfg.RabbitMQUrl)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error openning a channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		"notification_tasks",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Error declaring a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Error registering new consumer")
	}

	go func() {
		log.Println("RabbitMQ Consumer started. Waiting gor messages...")
		for d := range msgs {
			var task NotificationTask

			err := json.Unmarshal(d.Body, &task)
			if err != nil {
				log.Printf("Error unmarshalling task: %v. Discarding message.", err)
				d.Nack(false, false)
				continue
			}

			log.Printf("Successfully parsed task for user_id: %d", task.UserID)

			var settings NotificationSettings
			url := fmt.Sprintf("%s/internal/notifications_settings/user/%d", cfg.MainAppUrl, task.UserID)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("X-API-Key", cfg.InternalAPIKey)

			resp, err := http.DefaultClient.Do(req)
			if err != nil || resp.StatusCode != http.StatusOK {
				bodyBites, _ := io.ReadAll(resp.Body)
				log.Printf("Error getting notifications settings for user %d: %v, StatusCode: %d, Body: %s", task.UserID, err, resp.StatusCode, string(bodyBites))
				d.Nack(false, true)
				continue
			}
			json.NewDecoder(resp.Body).Decode(&settings)
			resp.Body.Close()

			chats, err := services.GetChats(context.Background(), services.Condition{
				Field:    "reminder_user_id",
				Operator: services.Equal,
				Value:    task.UserID,
			})

			if err != nil {
				log.Printf("Error getting user %d chats: %v", task.UserID, err)
				d.Nack(false, true)
				continue
			}

			for _, chat := range chats {
				platform := chat.Platform

				botInstance, ok := bot.GetBot(platform)
				if !ok {
					log.Printf("No bot registered for platform: %s", platform)
					continue
				}

				shouldSend := false

				switch platform {
				case "discord":
					if settings.DiscordNotification {
						shouldSend = true
					}
				case "vk":
					if settings.TelegramNotification {
						shouldSend = true
					}
				case "telegram":
					if settings.TelegramNotification {
						shouldSend = true
					}
				default:
					log.Printf("Unkmown platform found for user %d: %s", task.UserID, platform)
				}

				if shouldSend {
					err := botInstance.SendMessage(chat.ChatID, task.Message)
					if err != nil {
						log.Printf("Error sending message via %s for user %d: %v", platform, task.UserID, err)
					}
				}
			}

			d.Ack(false)
		}
	}()
}
