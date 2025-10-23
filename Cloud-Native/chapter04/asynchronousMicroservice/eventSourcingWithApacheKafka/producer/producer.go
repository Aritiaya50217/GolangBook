package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "bank-events",
	})
	defer writer.Close()

	events := []string{
		`{"type":"Deposit","amount":1000}`,
		`{"type":"Withdraw","amount":200}`,
		`{"type":"Deposit","amount":500}`,
	}
	for _, event := range events {
		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(fmt.Sprintf("event-%d", time.Now().UnixNano())),
			Value: []byte(event),
		})
		if err != nil {
			log.Fatal("Failed to write message : ", err)
		}
		fmt.Println("Sent Event : ", event)
		time.Sleep(time.Second)
	}
}
