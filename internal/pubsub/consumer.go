package pubsub

import (
	"encoding/json"
	"log"
)

type Event struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func Start(handler func(event Event)) {
	msgs, err := ch.Consume(
		"booking_events",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range msgs {
			var event Event
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("pubsub: failed to unmarshal message: %v (body: %s)", err, string(msg.Body))
				continue
			}
			handler(event)
		}
	}()
}
