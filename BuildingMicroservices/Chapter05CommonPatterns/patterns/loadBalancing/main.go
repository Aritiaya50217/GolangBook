package main

import (
	"fmt"
	"log"
	"net/http"

	"math/rand"

	"github.com/hashicorp/consul/api"
)

func discoverService(serviceName string) (string, error) {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return "", err
	}

	// ดึงเฉพาะ healthy service
	services, _, err := client.Health().Service(serviceName)
	if err != nil {
		return "", err
	}

	if len(services) == 0 {
		return "", fmt.Errorf("no healthy service found for %s", serviceName)
	}
	// เลือก instance แบบสุ่ม (load balancing เบื้องต้น)
	instance := services[rand.Intn(len(services))].Service
	return fmt.Sprintf("http://%s:%d", instance.Address, instance.Port), nil
}

func main() {
	endpoint, err := discoverService("user-service")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Calling", endpoint)

	resp, err := http.Get(endpoint + "/api/user")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var body = make([]byte, resp.ContentLength)
	resp.Body.Read(body)

	fmt.Println("Response:", string(body))
}
