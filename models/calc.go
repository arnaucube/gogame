package models

import (
	"math"

	"github.com/arnaucube/gogame/constants"
)

// formulas
// https://ogame.fandom.com/wiki/Formulas
// https://ogame.fandom.com/wiki/Research

// idelta is time in seconds units
func MetalGrowth(ilvl int64, idelta int64) int64 {
	lvl := float64(ilvl)
	delta := float64(idelta)

	// 30 * L * 1.1^L
	perHour := constants.UniverseAcceleration * 30 * lvl * math.Pow(1.1, lvl)
	r := (perHour / 60) * delta * constants.MineVelocity
	return int64(r)
}
func CrystalGrowth(ilvl int64, idelta int64) int64 {
	lvl := float64(ilvl)
	delta := float64(idelta)

	// 20 * L * 1.1^L
	perHour := constants.UniverseAcceleration * 20 * lvl * math.Pow(1.1, lvl)
	r := (perHour / 60) * delta * constants.MineVelocity
	return int64(r)
}
func DeuteriumGrowth(ilvl int64, idelta int64) int64 {
	lvl := float64(ilvl)
	delta := float64(idelta)

	t := float64(240) // t: max temperature
	// 10 * L * 1.1^L * (−0.002 * T + 1.28))
	perHour := constants.UniverseAcceleration * 10 * lvl * math.Pow(1.1, lvl) * ((-0.002)*t + 1.28)
	r := (perHour / 60) * delta * constants.MineVelocity
	return int64(r)
}
func SolarGrowth(ilvl int64, idelta int64) int64 {
	lvl := float64(ilvl)
	delta := float64(idelta)

	// 20 * L * 1.1^L
	perHour := constants.UniverseAcceleration * 20 * lvl * math.Pow(1.1, lvl)
	r := (perHour / 60) * delta * constants.MineVelocity
	return int64(r)
}
func FusionGrowth(ilvl int64, ilvlTech int64, idelta int64) int64 {
	lvl := float64(ilvl)
	lvlTech := float64(ilvlTech)
	delta := float64(idelta)

	// 30 * L * (1.05 + lvlTech * 0.01)^lvl
	perHour := constants.UniverseAcceleration * 30 * math.Pow((1.05+lvlTech*0.01), lvl)
	r := (perHour / 60) * delta * constants.MineVelocity
	return int64(r)
}

// https://ogame.fandom.com/wiki/Buildings

// https://ogame.fandom.com/wiki/Metal_Mine
func MetalMineCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     60,
		Crystal:   15,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 1.5^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(1.5, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(1.5, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(1.5, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(1.5, lvl-1))
	return cost
}

// https://ogame.fandom.com/wiki/Crystal_Mine
func CrystalMineCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     48,
		Crystal:   24,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 1.6^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(1.6, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(1.6, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(1.6, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(1.6, lvl-1))
	return cost
}

// https://ogame.fandom.com/wiki/Deuterium_Synthesizer
func DeuteriumMineCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     225,
		Crystal:   75,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 1.5^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(1.5, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(1.5, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(1.5, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(1.5, lvl-1))
	return cost
}

func EnergyMineCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     75,
		Crystal:   30,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 1.5^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(1.5, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(1.5, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(1.5, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(1.5, lvl-1))
	return cost
}

func FussionReactorCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     900,
		Crystal:   360,
		Deuterium: 180,
		Energy:    0,
	}
	// cost = base * 1.8^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(1.8, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(1.8, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(1.8, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(1.8, lvl-1))
	return cost
}
func RoboticsFactoryCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     400,
		Crystal:   120,
		Deuterium: 200,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}
func ShipyardCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     400,
		Crystal:   200,
		Deuterium: 100,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}
func MetalStorageCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     1000,
		Crystal:   0,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = 0
	cost.Deuterium = 0
	cost.Energy = 0
	return cost
}
func CrystalStorageCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     1000,
		Crystal:   500,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = 0
	cost.Energy = 0
	return cost
}
func DeuteriumStorageCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     1000,
		Crystal:   1000,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}

func RessearchLabCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     200,
		Crystal:   400,
		Deuterium: 200,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}
func AllianceDepotCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     20000,
		Crystal:   40000,
		Deuterium: 0,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}
func MissileSiloCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     20000,
		Crystal:   20000,
		Deuterium: 1000,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}
func SpaceDockCost(ilvl int64) Resources {
	lvl := float64(ilvl)
	base := Resources{
		Metal:     20000,
		Crystal:   20000,
		Deuterium: 1000,
		Energy:    0,
	}
	// cost = base * 2^(lvl-1)
	cost := Resources{}
	cost.Metal = int64(float64(base.Metal) * math.Pow(2, lvl-1))
	cost.Crystal = int64(float64(base.Crystal) * math.Pow(2, lvl-1))
	cost.Deuterium = int64(float64(base.Deuterium) * math.Pow(2, lvl-1))
	cost.Energy = int64(float64(base.Energy) * math.Pow(2, lvl-1))
	return cost
}

// TODO ConstructionTime and ResearchTime are following the formulas from https://ogame.fandom.com/wiki/Formulas
// but are not giving exact same numbers than in online calculators
func ConstructionTime(r Resources, roboticsLvl int64) int64 {
	naniteLvl := float64(1)
	// T(h) = (metal + crystal) / (2500 * (1+roboticsLvl) * 2^naniteLvl * universespeed)
	tHours := float64(r.Metal+r.Crystal) / (float64(2500) * float64(1+roboticsLvl) * math.Pow(2, naniteLvl) * constants.UniverseAcceleration)
	return int64(tHours * 3600)
}
func RessearchTime(r Resources, researchLvl int64) int64 {
	// T(h) = (metal + crystal) / (1000 * (1+researchLvl * universespeed)
	tHours := float64(r.Metal+r.Crystal) / (float64(1000) * float64(1+researchLvl*constants.UniverseAcceleration))
	return int64(tHours * 3600)
}
