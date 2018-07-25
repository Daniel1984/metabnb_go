package mongolayer

import (
	mgo "github.com/globalsign/mgo"
)

const (
	DB = "metabnb"
)

type MongoDBLayer struct {
	session *mgo.Session
}

func NewMongoDBLayer(connection string) (*MongoDBLayer, error) {
	s, err := mgo.Dial(connection)

	if err != nil {
		return nil, err
	}

	return &MongoDBLayer{session: s}, err
}

func (mgoLayer *MongoDBLayer) getFreshSession() *mgo.Session {
	return mgoLayer.session.Copy()
}
