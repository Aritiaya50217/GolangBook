package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type OrderCreated struct {
	Event   string  `json:"event"`
	OrderID int     `json:"order_id"`
	Amount  float64 `json:"amount"`
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "order-events",
		GroupID: "payment-service",
	})

	defer reader.Close()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "payment-events",
	})
	defer writer.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Read error: ", err)
			continue
		}

		var event OrderCreated
		_ = json.Unmarshal(msg.Value, &event)
		if event.Event == "OrderCreated" {
			log.Printf("Processing payment for order #%d\n", event.OrderID)

			paymentEvent := `{"event":"PaymentProcessed","order_id":` + string(msg.Key) + `}`
			writer.WriteMessages(context.Background(), kafka.Message{
				Value: []byte(paymentEvent),
			})
			log.Println("Published : ", paymentEvent)
		}

	}
}
