package common

import (
	"log"

	"gopkg.in/mgo.v2"
)

func NewMongoSession() *mgo.Session {
	log.Println("Connencting to MongoDB...")
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal("Unable to Dial to MongoDB on 127.0.0.1:27017: ", err)
	}
	session.SetSafe(&mgo.Safe{})
	session.SetMode(mgo.Monotonic, true)
	log.Println("Connected to testdb...")
	return session
}
