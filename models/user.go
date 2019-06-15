package models

import (
	"errors"
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
	Planets     []bson.ObjectId
}

func NewUser(db *database.Db, name, password, email string) (*User, error) {
	newUser := UserDb{
		Id:          bson.NewObjectId(),
		Name:        name,
		Password:    password,
		Email:       email,
		LastUpdated: time.Now(),
		Resources: Resources{
			Metal:     500,
			Crystal:   500,
			Deuterium: 500,
			Energy:    500,
		},
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
		Planets:     u.Planets,
	}
}

func (u *User) StoreInDb() error {
	err := u.db.Users.Update(bson.M{"_id": u.Id}, bson.M{"$set": bson.M{
		"lastupdated": time.Now(),
		"resources":   u.Resources,
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

// GetResources updates the values of resources and returns the value, also updates the user.Resources
// Resource types: metal, crystal, deuterium, energy
func (u *User) GetResources() (*Resources, error) {
	// get current values
	err := u.GetFromDb()
	if err != nil {
		return nil, err
	}
	// calculate Delta time = currentTime - u.LastUpdated
	delta := time.Since(u.LastUpdated).Seconds()

	// get planets
	planets, err := u.GetPlanets()
	if err != nil {
		return nil, err
	}

	// get Resource-Plant level in each planet
	// and calculate growth = ResourcePlant.Level for each planet
	var metalGrowth, crystalGrowth, deuteriumGrowth, energyGrowth int64
	for _, planet := range planets {
		metalGrowth = metalGrowth + MetalGrowth(planet.Buildings["metalmine"], int64(delta))
		crystalGrowth = crystalGrowth + MetalGrowth(planet.Buildings["crystalmine"], int64(delta))
		deuteriumGrowth = deuteriumGrowth + MetalGrowth(planet.Buildings["deuteriummine"], int64(delta))
		energyGrowth = energyGrowth + MetalGrowth(planet.Buildings["energymine"], int64(delta))
	}
	// calculate newAmount = oldAmount + growth
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
	return &u.Resources, nil
}

// SpendResources checks if user has enough resources, then substracts the resources, and updates the amounts in the database
func (u *User) SpendResources(r Resources) error {
	err := u.GetFromDb()
	if err != nil {
		return err
	}
	if u.Resources.Metal < r.Metal {
		return errors.New("not enough metal resources")
	}
	if u.Resources.Crystal < r.Crystal {
		return errors.New("not enough crystal resources")
	}
	if u.Resources.Deuterium < r.Deuterium {
		return errors.New("not enough deuterium resources")
	}
	if u.Resources.Energy < r.Energy {
		return errors.New("not enough energy resources")
	}

	u.Resources.Metal = u.Resources.Metal - r.Metal
	u.Resources.Crystal = u.Resources.Crystal - r.Crystal
	u.Resources.Deuterium = u.Resources.Deuterium - r.Deuterium
	u.Resources.Energy = u.Resources.Energy - r.Energy

	err = u.StoreInDb()
	if err != nil {
		return err
	}
	return nil
}

func (u *User) GetBuildingCost(planet Planet, building string) (Resources, error) {
	switch building {
	case "metalmine":
		return MetalMineCost(planet.Buildings["metalmine"] + 1), nil
	case "crystalmine":
		return CrystalMineCost(planet.Buildings["crystalmine"] + 1), nil
	case "deuteriummine":
		return DeuteriumMineCost(planet.Buildings["deuteriummine"] + 1), nil
	case "energymine":
		return EnergyMineCost(planet.Buildings["energymine"] + 1), nil
	case "fusionreactor":
		return FussionReactorCost(planet.Buildings["fusionreactor"] + 1), nil
	case "roboticsfactory":
		return RoboticsFactoryCost(planet.Buildings["roboticsfactory"] + 1), nil
	case "shipyard":
		return ShipyardCost(planet.Buildings["shipyard"] + 1), nil
	case "metalstorage":
		return MetalStorageCost(planet.Buildings["metalstorage"] + 1), nil
	case "crystalstorage":
		return CrystalStorageCost(planet.Buildings["crystalstorage"] + 1), nil
	case "deuteriumstorage":
		return DeuteriumStorageCost(planet.Buildings["deuteriumstorage"] + 1), nil
	case "ressearchlab":
		return RessearchLabCost(planet.Buildings["ressearchlab"] + 1), nil
	case "alliancedepot":
		return AllianceDepotCost(planet.Buildings["alliancedepot"] + 1), nil
	case "missilesilo":
		return MissileSiloCost(planet.Buildings["missilesilo"] + 1), nil
	case "spacedock":
		return SpaceDockCost(planet.Buildings["spacedock"] + 1), nil
	default:
		return Resources{}, errors.New("building not found")
	}

}
