package main

import (
	"MicroserviceProjectSetup/config"
	"MicroserviceProjectSetup/internal/adapters/db"
	"MicroserviceProjectSetup/internal/adapters/grpc"
	"MicroserviceProjectSetup/internal/application/core/api"
	"log"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error : %v ", err)
	}
	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
