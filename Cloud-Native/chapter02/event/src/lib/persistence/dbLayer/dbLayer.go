package dblayer

import (
	"cloud-Native-programming-with-golang/chapter02/event/src/lib/persistence"
	mongolayer "cloud-Native-programming-with-golang/chapter02/event/src/lib/persistence/mongoLayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewePersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
