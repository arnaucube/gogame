package models

import "gopkg.in/mgo.v2/bson"

type Planet struct {
	Id      bson.ObjectId `json:"id", bson:"_id, omitempty"`
	Size    int64         // fields/slots
	Name    string
	OwnerId bson.ObjectId
}
