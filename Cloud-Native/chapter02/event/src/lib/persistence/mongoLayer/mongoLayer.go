package mongolayer

import (
	"Cloud-Native/chapter02/event/src/lib/persistence"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	s, err := mgo.Dial(connection)
	return &MongoDBLayer{
		session: s,
	}, err
}

func (mongolayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mongolayer.session.Copy()
}

func (mongolayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	s := mongolayer.getFreshSession()
	defer s.Close()

	if !e.ID.Valid() {
		e.ID = bson.NewObjectId()
	}

	if !e.Location.ID.Valid() {
		e.Location.ID = bson.NewObjectId()
	}

	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e)
}

func (mongolayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mongolayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}

	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e)
	return e, err
}

func (mongolayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	s := mongolayer.getFreshSession()
	defer s.Close()

	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

func (mongolayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	s := mongolayer.getFreshSession()
	defer s.Close()
	events := []persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
	return events, err
}
