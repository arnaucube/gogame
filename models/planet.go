package models

import "gopkg.in/mgo.v2/bson"

type Planet struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Size      int64         // fields/slots
	Name      string
	OwnerId   bson.ObjectId
	Buildings map[string]int64
	/*
		Buildings types (in the map, all in lowcase):
		   	MetalMine       int64
		   	CrystalMine     int64
		   	DeuteriumMine   int64
		   	EnergyMine      int64
		   	FusionReactor   int64
		   	RoboticsFactory int64
		   	Shipyard        int64
		   	RessearchLab    int64
	*/
}
