package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-events",
	})
	defer writer.Close()

	orderEvent := `{"event":"OrderCreated","order_id":101,"amount":500.0}`
	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("order-101"),
			Value: []byte(orderEvent),
		})
	if err != nil {
		log.Fatal("Failed to publish: ", err)
	}
	log.Println("Published Event: ", orderEvent)
}
