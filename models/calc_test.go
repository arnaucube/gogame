package models

import (
	"testing"

	"github.com/arnaucube/gogame/constants"
	"github.com/stretchr/testify/assert"
)

func TestGrowth(t *testing.T) {
	// metal
	assert.Equal(t, int64(33), MetalGrowth(1, 60)/constants.MineVelocity)
	assert.Equal(t, int64(66), MetalGrowth(1, 120)/constants.MineVelocity)
	assert.Equal(t, int64(72), MetalGrowth(2, 60)/constants.MineVelocity)

	// crystal
	assert.Equal(t, int64(22), CrystalGrowth(1, 60)/constants.MineVelocity)
	assert.Equal(t, int64(44), CrystalGrowth(1, 120)/constants.MineVelocity)
	assert.Equal(t, int64(48), CrystalGrowth(2, 60)/constants.MineVelocity)

	// deuterium
	assert.Equal(t, int64(8), DeuteriumGrowth(1, 60)/constants.MineVelocity)
	assert.Equal(t, int64(17), DeuteriumGrowth(1, 120)/constants.MineVelocity)
	assert.Equal(t, int64(19), DeuteriumGrowth(2, 60)/constants.MineVelocity)

	// solar
	assert.Equal(t, int64(22), SolarGrowth(1, 60)/constants.MineVelocity)
	assert.Equal(t, int64(44), SolarGrowth(1, 120)/constants.MineVelocity)
	assert.Equal(t, int64(48), SolarGrowth(2, 60)/constants.MineVelocity)

	// fusion
	assert.Equal(t, int64(31), FusionGrowth(1, 1, 60)/constants.MineVelocity)
	assert.Equal(t, int64(63), FusionGrowth(1, 1, 120)/constants.MineVelocity)
	assert.Equal(t, int64(34), FusionGrowth(2, 2, 60)/constants.MineVelocity)
}

func TestMineCost(t *testing.T) {
	// metalmine
	assert.Equal(t, Resources{Metal: 60, Crystal: 15}, MetalMineCost(1))
	assert.Equal(t, Resources{Metal: 90, Crystal: 22}, MetalMineCost(2))
	assert.Equal(t, Resources{Metal: 17515, Crystal: 4378}, MetalMineCost(15))

	// crystalmine
	assert.Equal(t, Resources{Metal: 48, Crystal: 24}, CrystalMineCost(1))
	assert.Equal(t, Resources{Metal: 76, Crystal: 38}, CrystalMineCost(2))
	assert.Equal(t, Resources{Metal: 34587, Crystal: 17293}, CrystalMineCost(15))

	// deuteriummine
	assert.Equal(t, Resources{Metal: 225, Crystal: 75}, DeuteriumMineCost(1))
	assert.Equal(t, Resources{Metal: 337, Crystal: 112}, DeuteriumMineCost(2))
	assert.Equal(t, Resources{Metal: 65684, Crystal: 21894}, DeuteriumMineCost(15))

	// energymine
	assert.Equal(t, Resources{Metal: 75, Crystal: 30}, EnergyMineCost(1))
	assert.Equal(t, Resources{Metal: 112, Crystal: 45}, EnergyMineCost(2))
	assert.Equal(t, Resources{Metal: 21894, Crystal: 8757}, EnergyMineCost(15))

	// FussionReactorCost
	assert.Equal(t, Resources{Metal: 900, Crystal: 360, Deuterium: 180}, FussionReactorCost(1))
	assert.Equal(t, Resources{Metal: 1620, Crystal: 648, Deuterium: 324}, FussionReactorCost(2))
	assert.Equal(t, Resources{Metal: 3373320, Crystal: 1349328, Deuterium: 674664}, FussionReactorCost(15))

	// roboticsfactory
	assert.Equal(t, Resources{Metal: 400, Crystal: 120, Deuterium: 200}, RoboticsFactoryCost(1))
	assert.Equal(t, Resources{Metal: 800, Crystal: 240, Deuterium: 400}, RoboticsFactoryCost(2))
	assert.Equal(t, Resources{Metal: 6553600, Crystal: 1966080, Deuterium: 3276800}, RoboticsFactoryCost(15))

	// shipyard
	assert.Equal(t, Resources{Metal: 400, Crystal: 200, Deuterium: 100}, ShipyardCost(1))
	assert.Equal(t, Resources{Metal: 800, Crystal: 400, Deuterium: 200}, ShipyardCost(2))
	assert.Equal(t, Resources{Metal: 6553600, Crystal: 3276800, Deuterium: 1638400}, ShipyardCost(15))

	// metalstorage
	assert.Equal(t, Resources{Metal: 1000}, MetalStorageCost(1))
	assert.Equal(t, Resources{Metal: 2000}, MetalStorageCost(2))
	assert.Equal(t, Resources{Metal: 16384000}, MetalStorageCost(15))

	// crystalstorage
	assert.Equal(t, Resources{Metal: 1000, Crystal: 500}, CrystalStorageCost(1))
	assert.Equal(t, Resources{Metal: 2000, Crystal: 1000}, CrystalStorageCost(2))
	assert.Equal(t, Resources{Metal: 16384000, Crystal: 8192000}, CrystalStorageCost(15))

	// deuteriumstorage
	assert.Equal(t, Resources{Metal: 1000, Crystal: 1000}, DeuteriumStorageCost(1))
	assert.Equal(t, Resources{Metal: 2000, Crystal: 2000}, DeuteriumStorageCost(2))
	assert.Equal(t, Resources{Metal: 16384000, Crystal: 16384000}, DeuteriumStorageCost(15))

	// researchlab
	assert.Equal(t, Resources{Metal: 200, Crystal: 400, Deuterium: 200}, RessearchLabCost(1))
	assert.Equal(t, Resources{Metal: 400, Crystal: 800, Deuterium: 400}, RessearchLabCost(2))
	assert.Equal(t, Resources{Metal: 3276800, Crystal: 6553600, Deuterium: 3276800}, RessearchLabCost(15))

	// alliancedepot
	assert.Equal(t, Resources{Metal: 20000, Crystal: 40000}, AllianceDepotCost(1))
	assert.Equal(t, Resources{Metal: 40000, Crystal: 80000}, AllianceDepotCost(2))
	assert.Equal(t, Resources{Metal: 327680000, Crystal: 655360000}, AllianceDepotCost(15))

	// missilesilo
	assert.Equal(t, Resources{Metal: 20000, Crystal: 20000, Deuterium: 1000}, MissileSiloCost(1))
	assert.Equal(t, Resources{Metal: 40000, Crystal: 40000, Deuterium: 2000}, MissileSiloCost(2))
	assert.Equal(t, Resources{Metal: 327680000, Crystal: 327680000, Deuterium: 16384000}, MissileSiloCost(15))

	// spacedock
	assert.Equal(t, Resources{Metal: 20000, Crystal: 20000, Deuterium: 1000}, SpaceDockCost(1))
	assert.Equal(t, Resources{Metal: 40000, Crystal: 40000, Deuterium: 2000}, SpaceDockCost(2))
	assert.Equal(t, Resources{Metal: 327680000, Crystal: 327680000, Deuterium: 16384000}, SpaceDockCost(15))
}

func TestConstructionTime(t *testing.T) {
	// Numbers of the online calculators
	// assert.Equal(t, int64(30), ConstructionTime(Resources{Metal: 60, Crystal: 15}, 1))
	// assert.Equal(t, int64(53), ConstructionTime(Resources{Metal: 90, Crystal: 22}, 1))
	// assert.Equal(t, int64(96), ConstructionTime(Resources{Metal: 135, Crystal: 33}, 1))
	// assert.Equal(t, int64(383340), ConstructionTime(Resources{Metal: 204800, Crystal: 61440}, 10))

	// Numbers following the formulas
	assert.Equal(t, int64(27), ConstructionTime(Resources{Metal: 60, Crystal: 15}, 1))
	assert.Equal(t, int64(40), ConstructionTime(Resources{Metal: 90, Crystal: 22}, 1))
	assert.Equal(t, int64(60), ConstructionTime(Resources{Metal: 135, Crystal: 33}, 1))
	assert.Equal(t, int64(17426), ConstructionTime(Resources{Metal: 204800, Crystal: 61440}, 10))
}
func TestRessearchTime(t *testing.T) {
	assert.Equal(t, int64(1440), RessearchTime(Resources{Metal: 0, Crystal: 800}, 1))
	assert.Equal(t, int64(5236), RessearchTime(Resources{Metal: 12800, Crystal: 3200}, 10))

	// Numbers of the online calculators
	// assert.Equal(t, int64(2880), RessearchTime(Resources{Metal: 12800, Crystal: 6400}, 15))
	// assert.Equal(t, int64(360), RessearchTime(Resources{Metal: 1600, Crystal: 800}, 15))
}
