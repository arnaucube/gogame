package models

import "gopkg.in/mgo.v2/bson"

type SolarSystem struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Position int64         // position of the solar system in the galaxy, the maximum position is GALAXYSIZE-1
	Planets  []string      // array with ids of the planets, if empty is equal to ""
}
