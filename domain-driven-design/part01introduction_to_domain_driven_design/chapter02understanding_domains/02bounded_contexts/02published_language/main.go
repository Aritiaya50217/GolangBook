package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Published Language (DTO)
type PaymentRequest struct {
	OrderID  string  `json:"order_id"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

// Open Host Service
func paymentHandler(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest
	json.NewDecoder(r.Body).Decode(&req)
	fmt.Println("Payment received : ")
	fmt.Println("OrderID : ", req.OrderID)
	fmt.Println("Amount : ", req.Amount)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"paid"}`))
}

// Order Domain (Client)

type Order struct {
	ID          string
	TotalAmount float64
}

func sendToPayment(order Order) {
	// map -> Published Language
	dto := PaymentRequest{
		OrderID:  order.ID,
		Amount:   order.TotalAmount,
		Currency: "THB",
	}

	body, _ := json.Marshal(dto)
	http.Post("http://localhost:8080/pay",
		"application/json",
		bytes.NewBuffer(body))
}

func main() {
	// start Payment service (OHS)
	http.HandleFunc("/pay", paymentHandler)

	go func() {
		fmt.Println("Payment service running on : 8080")
		http.ListenAndServe(":8080", nil)
	}()

	time.Sleep(1 * time.Second)

	// simulate Order
	order := Order{
		ID:          "ORD-001",
		TotalAmount: 500,
	}

	fmt.Println("Sending order to payment...")
	sendToPayment(order)

	// กันโปรแกรมปิด
	time.Sleep(2 * time.Second)
}
