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

	// researchlab
	assert.Equal(t, Resources{Metal: 200, Crystal: 400, Deuterium: 200}, RessearchLabCost(1))
	assert.Equal(t, Resources{Metal: 400, Crystal: 800, Deuterium: 400}, RessearchLabCost(2))
	assert.Equal(t, Resources{Metal: 3276800, Crystal: 6553600, Deuterium: 3276800}, RessearchLabCost(15))
}
