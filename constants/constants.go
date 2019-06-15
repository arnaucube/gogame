package constants

import (
	"github.com/arnaucube/gogame/models"
)

// game constants

const GALAXYSIZE = 50
const SOLARSYSTEMSIZE = 15

// MetalMineLevels contains the constants of productivity for each level, sorted from 0 to X levels
var MetalMineLevels = []int64{
	0,
	1,
	5,
	10,
	// TODO this will be same strategy with all the buildings and research
}
var CrystalMineLevels = []int64{
	0,
	1,
	5,
	10,
}
var DeuteriumMineLevels = []int64{
	0,
	1,
	5,
	10,
}
var EnergyMineLevels = []int64{
	0,
	1,
	5,
	10,
}

// BuildingsNeededResources hold
// map with all the buildings, that each one is a map with the levels of the buildings with the needed ressources
var BuildingsNeededResources = map[string]map[int64]models.Resources{
	"metalplant": map[int64]models.Resources{
		1: models.Resources{
			Metal:     50,
			Crystal:   50,
			Deuterium: 50,
			Energy:    50,
		},
		2: models.Resources{
			Metal:     70,
			Crystal:   70,
			Deuterium: 70,
			Energy:    70,
		},
		3: models.Resources{
			Metal:     90,
			Crystal:   90,
			Deuterium: 90,
			Energy:    90,
		},
	},
}

// extra
const JWTIdKey = "id"
