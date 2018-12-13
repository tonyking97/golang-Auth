package core

import "gopkg.in/mgo.v2"

func MongoDBInit() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

}