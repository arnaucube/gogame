package models

import (
	"fmt"
	"time"

	"github.com/arnaucube/gogame/database"
	"gopkg.in/mgo.v2/bson"
)

type Resources struct {
	Metal     int64
	Crystal   int64
	Deuterium int64
	Energy    int64
}

// UserDb is the data in DB
type UserDb struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	Password    string
	Email       string
	LastUpdated time.Time
	Resources   Resources
	Planets     []bson.ObjectId
}

// User is the data in memory, after getting it from DB
type User struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name        string
	LastUpdated time.Time
	db          *database.Db
	Resources   Resources
}

func NewUser(db *database.Db, name, password, email string) (*User, error) {
	newUser := UserDb{
		Id:          bson.NewObjectId(),
		Name:        name,
		Password:    password,
		Email:       email,
		LastUpdated: time.Now(),
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
		Resources:   u.Resources,
	}
}

func (u *User) StoreInDb() error {
	userDb := UserDb{
		Id:          u.Id,
		Name:        u.Name,
		LastUpdated: u.LastUpdated,
		Resources:   u.Resources,
	}
	err := u.db.Users.Update(bson.M{"_id": u.Id}, userDb)
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
	err := u.db.Users.Find(bson.M{"OwnerId": u.Id}).All(&planets)
	if err != nil {
		return planets, err
	}
	return planets, nil
}

// GetResources updates the values of resources and returns the value
// Resource types: metal, crystal, deuterium, energy
func (u *User) GetResources() (*User, error) {
	// get current values
	err := u.GetFromDb()
	if err != nil {
		return nil, err
	}
	// get u.LastUpdated
	fmt.Println(u.LastUpdated)
	// calculate Delta time = currentTime - u.LastUpdated
	delta := time.Since(u.LastUpdated)

	// get planets
	planets, err := u.GetPlanets()
	if err != nil {
		return nil, err
	}

	// get Resource-Plant level in each planet
	var metalGrowth, crystalGrowth, deuteriumGrowth, energyGrowth int64
	for _, planet := range planets {
		// calculate growth = ResourcePlant.Level for each planet
		// TODO find correct formulas
		metalGrowth = metalGrowth + (planet.Buildings.MetalMine * int64(delta))
		crystalGrowth = crystalGrowth + (planet.Buildings.CrystalMine * int64(delta))
		deuteriumGrowth = deuteriumGrowth + (planet.Buildings.DeuteriumMine * int64(delta))
		energyGrowth = energyGrowth + (planet.Buildings.EnergyMine * int64(delta))
	}
	// calculate newAmount = oldAmount + (growth & DeltaTime)
	u.Resources.Metal = u.Resources.Metal + metalGrowth
	u.Resources.Crystal = u.Resources.Crystal + crystalGrowth
	u.Resources.Deuterium = u.Resources.Deuterium + deuteriumGrowth
	u.Resources.Energy = u.Resources.Energy + energyGrowth

	// store new amount to user db
	err = u.StoreInDb()
	if err != nil {
		return nil, err
	}

	// return user
	return u, nil
}
