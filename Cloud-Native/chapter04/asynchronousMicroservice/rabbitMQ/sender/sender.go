package main

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// connect ไปยัง RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Fatiled to open channel : ", err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs", // exchange name
		"fanout",
		true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare exchange: ", err)
	}

	// send message
	body := "Hello from Publisher"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx, "logs", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})

	if err != nil {
		log.Fatal("Failed to publish message : ", err)
	}
	fmt.Printf("Sent message : %s\n", body)
}
