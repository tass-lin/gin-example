package models

import "gopkg.in/mgo.v2/bson"

type SampleGet struct {
	Name string `bson:"name" json:"name"`
	ListenUrl string `bson:"listen_url" json:"listenUrl"`
	Summary string `bson:"summary" json:"summary"`
	Space string `bson:"space" json:"space"`
	Description string `bson:"description" json:"description"`
}

type SampleInsert struct {
	Name string `bson:"name" json:"name"`
	ListenUrl string `bson:"listen_url" json:"listenUrl"`
	Summary string `bson:"summary" json:"summary"`
	Space string `bson:"space" json:"space"`
	Description string `bson:"description" json:"description"`
}

type SamplePatch struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `bson:"name" json:"name"`
	ListenUrl string `bson:"listen_url" json:"listenUrl"`
	Summary string `bson:"summary" json:"summary"`
	Space string `bson:"space" json:"space"`
	Description string `bson:"description" json:"description"`
}

type SampleDelete struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

