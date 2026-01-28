// Package pubsub handles internal/external messaging
package pubsub

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	conn *amqp.Connection
	ch   *amqp.Channel
)

func Init() error {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("cannot load .env file")
	// }
	rabbitConnString := os.Getenv("RABBITMQ_URL")
	var err error

	conn, err = amqp.Dial(rabbitConnString)
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare(
		"booking_events",
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func Close() {
	if ch != nil {
		ch.Close()
	}

	if conn != nil {
		conn.Close()
	}
}
