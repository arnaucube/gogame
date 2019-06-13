package database

import mgo "gopkg.in/mgo.v2"

type Db struct {
	Users        *mgo.Collection
	Planets      *mgo.Collection
	SolarSystems *mgo.Collection
	Galaxies     *mgo.Collection
}

func New(url string, databaseName string) (*Db, error) {
	session, err := mgo.Dial("mongodb://" + url)
	if err != nil {
		return nil, err
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := Db{}
	db.Users = session.DB(databaseName).C("users")
	db.Planets = session.DB(databaseName).C("planets")
	db.SolarSystems = session.DB(databaseName).C("solarsystems")
	db.Galaxies = session.DB(databaseName).C("galaxies")

	return &db, nil
}
