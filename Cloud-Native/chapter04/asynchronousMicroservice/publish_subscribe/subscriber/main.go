package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)
	q, _ := ch.QueueDeclare("", false, false, true, false, nil)
	ch.QueueBind(q.Name, "", "logs", false, nil)

	msgs, _ := ch.Consume(q.Name, "", true, false, false, false, nil)

	for msg := range msgs {
		log.Printf("Received: %s", msg.Body)
	}
}
