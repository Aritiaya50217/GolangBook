package main

import (
	"Cloud-Native/chapter02/event/src/lib/configuration"
	dblayer "Cloud-Native/chapter02/event/src/lib/persistence/dbLayer"
	"Cloud-Native/chapter02/event/src/service/rest"
	"flag"
	"fmt"
	"log"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewePersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
