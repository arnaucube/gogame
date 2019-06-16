package gamesrv

import (
	"errors"
	"fmt"
	"strconv"
	"time"

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
	// in case that wants to start with resources plants
	newPlanet.Buildings = make(map[string]int64)
	newPlanet.Buildings["metalmine"] = 1
	newPlanet.Buildings["crystalmine"] = 1
	newPlanet.Buildings["deuteriummine"] = 1
	newPlanet.Buildings["energymine"] = 1

	err := srv.db.Planets.Insert(newPlanet)
	if err != nil {
		return nil, nil, err
	}
	var planet *models.Planet
	err = srv.db.Planets.Find(bson.M{"ownerid": newPlanet.OwnerId}).One(&planet)
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

// CheckCurrentBuild checks if the planet has a ongoing building in process, and if has finished
// in case that has finished, updates it in db
func (srv Service) CheckCurrentBuild(planet *models.Planet) (bool, error) {
	if planet.CurrentBuild.Title != "" {
		// the planet is building something, check if has ended
		if planet.CurrentBuild.Ends.Unix() < time.Now().Unix() {
			// upgrade level of building in planet
			planet.Buildings[planet.CurrentBuild.Building] += 1

			// build end
			planet.CurrentBuild.Title = ""
			planet.CurrentBuild.Building = ""

			// store in db
			err := srv.db.Planets.Update(bson.M{"_id": planet.Id}, planet)
			if err != nil {
				return true, err
			}
			return false, nil
		}
		return true, nil
	}
	return false, nil

}

func (srv Service) GetBuildings(user *models.User, planetid bson.ObjectId) (*models.Planet, error) {
	var planet models.Planet
	err := srv.db.Planets.Find(bson.M{"_id": planetid, "ownerid": user.Id}).One(&planet)
	if err != nil {
		return nil, err
	}
	_, err = srv.CheckCurrentBuild(&planet)
	return &planet, err
}

func (srv Service) UpgradeBuilding(user *models.User, planetid bson.ObjectId, building string) (*models.Planet, error) {
	// get planet
	var planet models.Planet
	err := srv.db.Planets.Find(bson.M{"_id": planetid}).One(&planet)
	if err != nil {
		return nil, err
	}
	busy, err := srv.CheckCurrentBuild(&planet)
	if err != nil {
		return nil, err
	}
	if busy {
		return nil, errors.New("busy")
	}

	// get current building level, and get the needed resources for next level
	resourcesNeeded, err := user.GetBuildingCost(planet, building)
	if err != nil {
		return nil, err
	}
	// get time cost of the build
	timei64 := models.ConstructionTime(resourcesNeeded, planet.Buildings[building]+1)
	endsTime := time.Now().Add(time.Second * time.Duration(timei64))

	// if user have enough resources to upgrade the building, upgrade the building
	err = user.SpendResources(resourcesNeeded)
	if err != nil {
		return nil, err
	}
	// add current task to planet
	planet.CurrentBuild.Building = building
	planet.CurrentBuild.Title = building + " - Level " + strconv.Itoa(int(planet.Buildings[building]))
	planet.CurrentBuild.Ends = endsTime

	// store planet in db
	err = srv.db.Planets.Update(bson.M{"_id": planet.Id}, planet)
	return &planet, nil
}
