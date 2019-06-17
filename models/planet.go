package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/arnaucube/gogame/database"
	"gopkg.in/mgo.v2/bson"
)

type Process struct {
	// if Title == "", is not active, and can build other buildings/research
	Title     string // building name / research name + level
	Building  string
	Ends      time.Time
	CountDown int64
}

type Resources struct {
	Metal     int64
	Crystal   int64
	Deuterium int64
	Energy    int64
}

type Planet struct {
	Db           *database.Db
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	LastUpdated  time.Time
	Size         int64 // fields/slots
	Name         string
	OwnerId      bson.ObjectId
	Resources    Resources
	Buildings    map[string]int64
	CurrentBuild Process
	Research     Process
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

func NewPlanet(db *database.Db, size int64, name string, ownerId bson.ObjectId) (*Planet, error) {
	newPlanet := Planet{
		Db:      db,
		Size:    size,
		Name:    name,
		OwnerId: ownerId,
		Resources: Resources{
			Metal:     500,
			Crystal:   500,
			Deuterium: 500,
			Energy:    500,
		},
	}

	// in case that wants to start with resources plants
	newPlanet.Buildings = make(map[string]int64)
	newPlanet.Buildings["metalmine"] = 1
	newPlanet.Buildings["crystalmine"] = 1
	newPlanet.Buildings["deuteriummine"] = 1
	newPlanet.Buildings["energymine"] = 1

	err := db.Planets.Insert(newPlanet)
	if err != nil {
		return nil, err
	}

	var planet *Planet
	err = db.Planets.Find(bson.M{"ownerid": newPlanet.OwnerId}).One(&planet)
	if err != nil {
		return nil, err
	}

	return planet, nil
}

func (p *Planet) GetFromDb() error {
	var planet Planet
	err := p.Db.Planets.Find(bson.M{"_id": p.Id}).One(&planet)
	if err != nil {
		return err
	}
	planet.Db = p.Db
	p = &planet
	return nil
}

func (p *Planet) StoreInDb() error {
	p.LastUpdated = time.Now()
	err := p.Db.Planets.Update(bson.M{"_id": p.Id}, p)
	return err

}

// GetResources updates the values of resources and returns the value, also updates the user.Resources
// Resource types: metal, crystal, deuterium, energy
func (p *Planet) GetResources() (*Resources, error) {
	// get current values
	err := p.GetFromDb()
	if err != nil {
		return nil, err
	}
	// calculate Delta time = currentTime - p.LastUpdated
	delta := time.Since(p.LastUpdated).Seconds()

	// get Resource-Plant level in each planet
	// and calculate growth = ResourcePlant.Level for each planet
	var metalGrowth, crystalGrowth, deuteriumGrowth, energyGrowth int64
	metalGrowth = metalGrowth + MetalGrowth(p.Buildings["metalmine"], int64(delta))
	crystalGrowth = crystalGrowth + MetalGrowth(p.Buildings["crystalmine"], int64(delta))
	deuteriumGrowth = deuteriumGrowth + MetalGrowth(p.Buildings["deuteriummine"], int64(delta))
	energyGrowth = energyGrowth + MetalGrowth(p.Buildings["energymine"], int64(delta))

	// calculate newAmount = oldAmount + growth
	p.Resources.Metal = p.Resources.Metal + metalGrowth
	p.Resources.Crystal = p.Resources.Crystal + crystalGrowth
	p.Resources.Deuterium = p.Resources.Deuterium + deuteriumGrowth
	p.Resources.Energy = p.Resources.Energy + energyGrowth

	// store new amount to user db
	err = p.StoreInDb()
	if err != nil {
		return nil, err
	}

	// return user
	return &p.Resources, nil
}

// SpendResources checks if user has enough resources, then substracts the resources, and updates the amounts in the database
func (p *Planet) SpendResources(r Resources) error {
	err := p.GetFromDb()
	if err != nil {
		return err
	}
	if p.Resources.Metal < r.Metal {
		return errors.New("not enough metal resources")
	}
	if p.Resources.Crystal < r.Crystal {
		return errors.New("not enough crystal resources")
	}
	if p.Resources.Deuterium < r.Deuterium {
		return errors.New("not enough deuterium resources")
	}
	if p.Resources.Energy < r.Energy {
		return errors.New("not enough energy resources")
	}

	p.Resources.Metal = p.Resources.Metal - r.Metal
	p.Resources.Crystal = p.Resources.Crystal - r.Crystal
	p.Resources.Deuterium = p.Resources.Deuterium - r.Deuterium
	p.Resources.Energy = p.Resources.Energy - r.Energy

	err = p.StoreInDb()
	if err != nil {
		return err
	}
	return nil
}
func (p *Planet) GetBuildingCost(building string) (Resources, error) {
	switch building {
	case "metalmine":
		return MetalMineCost(p.Buildings["metalmine"] + 1), nil
	case "crystalmine":
		return CrystalMineCost(p.Buildings["crystalmine"] + 1), nil
	case "deuteriummine":
		return DeuteriumMineCost(p.Buildings["deuteriummine"] + 1), nil
	case "energymine":
		return EnergyMineCost(p.Buildings["energymine"] + 1), nil
	case "fusionreactor":
		return FussionReactorCost(p.Buildings["fusionreactor"] + 1), nil
	case "roboticsfactory":
		return RoboticsFactoryCost(p.Buildings["roboticsfactory"] + 1), nil
	case "shipyard":
		return ShipyardCost(p.Buildings["shipyard"] + 1), nil
	case "metalstorage":
		return MetalStorageCost(p.Buildings["metalstorage"] + 1), nil
	case "crystalstorage":
		return CrystalStorageCost(p.Buildings["crystalstorage"] + 1), nil
	case "deuteriumstorage":
		return DeuteriumStorageCost(p.Buildings["deuteriumstorage"] + 1), nil
	case "ressearchlab":
		return RessearchLabCost(p.Buildings["ressearchlab"] + 1), nil
	case "alliancedepot":
		return AllianceDepotCost(p.Buildings["alliancedepot"] + 1), nil
	case "missilesilo":
		return MissileSiloCost(p.Buildings["missilesilo"] + 1), nil
	case "spacedock":
		return SpaceDockCost(p.Buildings["spacedock"] + 1), nil
	default:
		return Resources{}, errors.New("building not found")
	}

}

// CheckCurrentBuild checks if the planet has a ongoing building in process, and if has finished
// in case that has finished, updates it in db
func (p *Planet) CheckCurrentBuild() (bool, error) {
	if p.CurrentBuild.Title != "" {
		// the planet is building something, check if has ended
		if p.CurrentBuild.Ends.Unix() < time.Now().Unix() {
			// upgrade level of building in planet
			p.Buildings[p.CurrentBuild.Building] += 1

			// build end
			p.CurrentBuild.Title = ""
			p.CurrentBuild.Building = ""

			// store in db
			err := p.Db.Planets.Update(bson.M{"_id": p.Id}, p)
			if err != nil {
				return true, err
			}
			return false, nil
		}
		p.CurrentBuild.CountDown = p.CurrentBuild.Ends.Unix() - time.Now().Unix()

		return true, nil
	}
	return false, nil

}
func (p *Planet) UpgradeBuilding(building string) error {
	busy, err := p.CheckCurrentBuild()
	if err != nil {
		return err
	}
	if busy {
		return errors.New("busy")
	}

	// get current building level, and get the needed resources for next level
	resourcesNeeded, err := p.GetBuildingCost(building)
	if err != nil {
		return err
	}
	// get time cost of the build
	timei64 := ConstructionTime(resourcesNeeded, p.Buildings[building]+1)
	endsTime := time.Now().Add(time.Second * time.Duration(timei64))

	// if user have enough resources to upgrade the building, upgrade the building
	err = p.SpendResources(resourcesNeeded)
	if err != nil {
		return err
	}
	// add current task to planet
	p.CurrentBuild.Building = building
	p.CurrentBuild.Title = building + " - Level " + strconv.Itoa(int(p.Buildings[building]))
	p.CurrentBuild.Ends = endsTime
	p.CurrentBuild.CountDown = p.CurrentBuild.Ends.Unix() - time.Now().Unix()

	// store planet in db
	err = p.Db.Planets.Update(bson.M{"_id": p.Id}, p)
	return nil
}
