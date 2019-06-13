package models

import "gopkg.in/mgo.v2/bson"

type Resource struct {
	Value int64
	Max   int64
}

type User struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string
	Password  string
	Email     string
	Resources struct {
		Metal     Resource
		Crystal   Resource
		Deuterium Resource
		Energy    Resource
	}
}
