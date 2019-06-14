package gamesrv

import (
	"fmt"

	"github.com/arnaucube/gogame/constants"
	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/models"
	"github.com/arnaucube/gogame/utils"
	"gopkg.in/mgo.v2/bson"
)

type Service struct {
	db *database.Db
}

func New(db *database.Db) *Service {
	return &Service{
		db,
	}
}

// CreatePlanet is used when a user conquers a planet
func (srv Service) CreatePlanet(userId bson.ObjectId) (*models.SolarSystem, *models.Planet, error) {
	size := int64(250)   // TODO get rand inside a range
	name := "planetname" // TODO get random name

	newPlanet := models.Planet{
		Size:    size,
		Name:    name,
		OwnerId: userId,
	}
	err := srv.db.Planets.Insert(newPlanet)
	if err != nil {
		return nil, nil, err
	}
	var planet *models.Planet
	err = srv.db.Planets.Find(bson.M{"name": name}).One(&planet)
	if err != nil {
		return nil, nil, err
	}

	// now put the planet in a solar system
	// get random solar system
	systemPosition := utils.RandInRange(0, constants.GALAXYSIZE)
	solarSystem, err := srv.PutPlanetInSolarSystem(systemPosition, planet)
	// TODO if error is returned because there is no empty slots for planets in the solar system in systemPosition, get another systemPosition and try again

	return solarSystem, planet, err
}

func (srv Service) PutPlanetInSolarSystem(position int64, planet *models.Planet) (*models.SolarSystem, error) {
	var solarSystem models.SolarSystem
	err := srv.db.SolarSystems.Find(bson.M{"position": position}).One(&solarSystem)
	if err != nil {
		// solar system non existing yet
		// create a solarsystem with empty planets
		var emptyPlanets []string
		for i := 0; i < constants.SOLARSYSTEMSIZE; i++ {
			emptyPlanets = append(emptyPlanets, "empty")
		}
		newSolarSystem := models.SolarSystem{
			Position: position,
			Planets:  emptyPlanets[:15],
		}
		err = srv.db.SolarSystems.Insert(newSolarSystem)
		if err != nil {
			return nil, err
		}
		err := srv.db.SolarSystems.Find(bson.M{"position": position}).One(&solarSystem)

		return &solarSystem, err
	}
	// get free slots in solarSystem
	posInSolarSystem := utils.RandInRange(0, constants.SOLARSYSTEMSIZE)
	if solarSystem.Planets[posInSolarSystem] != "" {
		// not empty slot, take another one TODO
		// if there are no empty slots, return error
		fmt.Println("not empty slot")
	}
	// store planet in solar system
	solarSystem.Planets[posInSolarSystem] = planet.Id.String()
	err = srv.db.SolarSystems.Update(bson.M{"position": position}, solarSystem)

	return &solarSystem, err
}
