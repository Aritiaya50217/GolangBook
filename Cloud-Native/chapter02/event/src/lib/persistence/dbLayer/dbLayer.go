package dblayer

import (
	"Cloud-Native/chapter02/event/src/lib/persistence"
	mongolayer "Cloud-Native/chapter02/event/src/lib/persistence/mongoLayer"
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
