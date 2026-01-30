package pubsub

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Publish(eventType string, payload any) error {
	body, err := json.Marshal(struct {
		Type string `json:"type"`
		Data any    `json:"data"`
	}{Type: eventType, Data: payload})
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		"booking_events",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("publish failed: %v", err)
	}

	return err
}
