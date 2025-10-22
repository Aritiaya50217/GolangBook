package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        string
	Type      string
	Amount    float64
	CreatedAt time.Time
}

var eventStore []Event

func Deposit(amount float64) {
	e := Event{
		ID:        uuid.NewString(),
		Type:      "Deposit",
		Amount:    amount,
		CreatedAt: time.Now(),
	}
	eventStore = append(eventStore, e)
}

func Withdraw(amount float64) {
	e := Event{
		ID:        uuid.NewString(),
		Type:      "Withdraw",
		Amount:    -amount,
		CreatedAt: time.Now(),
	}
	eventStore = append(eventStore, e)
}

func GetCurrentBalance() float64 {
	var balance float64
	for _, e := range eventStore {
		balance += e.Amount
	}
	return balance
}

func main() {
	Deposit(1000)
	Deposit(200)
	Withdraw(50)

	fmt.Println("Event Store : ")
	for _, e := range eventStore {
		fmt.Printf("- %s: %.2f\n", e.Type, e.Amount)
	}
	fmt.Printf("Current Balance: %.2f\n", GetCurrentBalance())
}
