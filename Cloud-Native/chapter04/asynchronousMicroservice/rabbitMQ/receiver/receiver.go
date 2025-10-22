package main

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ : ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel : ", err)
	}
	defer ch.Close()

	// ประกาศ exchange เหมือน publisher
	ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)

	// ประกาศ queue (ชื่อสุ่ม)
	q, err := ch.QueueDeclare("", false, false, true, false, nil)
	if err != nil {
		log.Fatal("Failed to declare queue : ", err)
	}

	// bind queue เข้ากับ exchange
	ch.QueueBind(q.Name, "", "logs", false, nil)

	// consume message
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal("Failed to consume: ", err)
	}

	log.Println("Waiting for message ...")
	for msg := range msgs {
		log.Printf("Received : %s\n", msg.Body)
	}
}
