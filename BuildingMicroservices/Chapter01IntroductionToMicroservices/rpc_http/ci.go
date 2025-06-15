package main

import (
	"fmt"

	"github.com/Aritiaya50217/GolangBook/tree/main/BuildingMicroservices/Chapter01IntroductionToMicroservices/rpc/client"
	"github.com/Aritiaya50217/GolangBook/tree/main/BuildingMicroservices/Chapter01IntroductionToMicroservices/rpc/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
