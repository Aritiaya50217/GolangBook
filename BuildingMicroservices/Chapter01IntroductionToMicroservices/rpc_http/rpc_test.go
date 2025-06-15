package main

import (
	"testing"

	"github.com/Aritiaya50217/GolangBook/tree/main/BuildingMicroservices/Chapter01IntroductionToMicroservices/rpc/client"
	"github.com/Aritiaya50217/GolangBook/tree/main/BuildingMicroservices/Chapter01IntroductionToMicroservices/rpc/server"
)

func BenchmarkDialHTTP(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c := client.CreateClient()
		c.Close()
	}
}

func BenchmarkHelloHandlerHTTP(b *testing.B) {
	b.ResetTimer()

	c := client.CreateClient()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest(c)
	}
	c.Close()
}

func init() {
	// start the server
	go server.StartServer()
}
