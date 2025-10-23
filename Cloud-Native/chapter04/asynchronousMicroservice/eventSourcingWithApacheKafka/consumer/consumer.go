package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// Event คือ โครงสร้าง event ที่จะถูกอ่านจาก kafka
type Event struct {
	Type   string  `json"type"`
	Amount float64 `json:"amount"`
}

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "bank-events",
		GroupID: "bank-group",
	})

	defer reader.Close()

	var balance float64

	fmt.Println("Listening for events...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message : ", err)
			continue
		}
		var event Event
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Println("Invalid JSON : ", err)
			continue
		}

		// Event Sourcing
		switch event.Type {
		case "Deposit":
			balance += event.Amount
		case "Withdraw":
			balance -= event.Amount
		}
		fmt.Printf("Received Event : %-10s -> Balance = %.2f\n", event.Type, balance)
	}
}
