package consumer

import (
	"log"

	"github.com/MatTwix/RE-minder-Bots/config"
	"github.com/rabbitmq/amqp091-go"
)

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
			log.Printf("Received a message: %v", d.Body)

			d.Ack(false)
		}
	}()
}
