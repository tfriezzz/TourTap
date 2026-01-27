package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("cannot load .env file")
	// }
	// rabbitConnString := os.Getenv("RABBITMQ_URL")
	rabbitConnString := "amqp://guest:guest@localhost:5672/"

	fmt.Printf("connection string: %v\n", rabbitConnString)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("recieved a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] waiting for messages")
	<-forever
}
