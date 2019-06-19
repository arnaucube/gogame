package models

import (
	"time"

	"github.com/arnaucube/gogame/database"
	"gopkg.in/mgo.v2/bson"
)

// UserDb is the data in DB
type UserDb struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	Password    string
	Email       string
	LastUpdated time.Time
	Planets     []bson.ObjectId
	Points      int64 // real points are this value / 1000
}

// User is the data in memory, after getting it from DB
type User struct {
	db          *database.Db
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	LastUpdated time.Time
	Planets     []bson.ObjectId
	Points      int64
}

func NewUser(db *database.Db, name, password, email string) (*User, error) {
	newUser := UserDb{
		Id:          bson.NewObjectId(),
		Name:        name,
		Password:    password,
		Email:       email,
		LastUpdated: time.Now(),
		Points:      0,
	}
	err := db.Users.Insert(newUser)
	if err != nil {
		return nil, err
	}
	user := UserDbToUser(db, newUser)
	return user, nil
}

func UserDbToUser(db *database.Db, u UserDb) *User {
	return &User{
		Id:          u.Id,
		Name:        u.Name,
		LastUpdated: u.LastUpdated,
		db:          db,
		Planets:     u.Planets,
		Points:      u.Points,
	}
}

func (u *User) StoreInDb() error {
	err := u.db.Users.Update(bson.M{"_id": u.Id}, bson.M{"$set": bson.M{
		"lastupdated": time.Now(),
		"planets":     u.Planets,
	}})
	return err

}

func (u *User) GetFromDb() error {
	var userDb UserDb
	err := u.db.Users.Find(bson.M{"_id": u.Id}).One(&userDb)
	if err != nil {
		return err
	}
	u = UserDbToUser(u.db, userDb)
	return nil
}

func (u *User) GetPlanets() ([]Planet, error) {
	var planets []Planet
	err := u.db.Planets.Find(bson.M{"ownerid": u.Id}).All(&planets)
	if err != nil {
		return planets, err
	}
	return planets, nil
}

func AddPoints(db *database.Db, userId bson.ObjectId, points int64) error {
	var userDb UserDb
	err := db.Users.Find(bson.M{"_id": userId}).One(&userDb)
	if err != nil {
		return err
	}
	newPoints := userDb.Points + points

	err = db.Users.Update(bson.M{"_id": userId}, bson.M{"$set": bson.M{
		"points": newPoints,
	}})
	return err
}
