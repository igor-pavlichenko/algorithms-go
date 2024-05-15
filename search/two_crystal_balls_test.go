package search

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	assert := assert.New(t)

	idx := rand.Intn(10000)
	floors := make([]bool, 10000)

	for i := idx; i < len(floors); i++ {
		floors[i] = true
	}

	index, err := TwoCrystalBalls(floors)
	assert.Equal(idx, index)
	assert.Nil(err)
}

func TestTwoCrystalBalls_firstFloor(t *testing.T) {
	assert := assert.New(t)

	floors := []bool{false, true, true, true, true, true, true, true, true}
	index, err := TwoCrystalBalls(floors)
	assert.Equal(1, index)
	assert.Nil(err)
}

func TestTwoCrystalBalls_zerothFloor(t *testing.T) {
	assert := assert.New(t)

	floors := []bool{true, true, true, true, true, true, true, true, true}
	index, err := TwoCrystalBalls(floors)
	assert.Equal(0, index)
	assert.Nil(err)
}

func TestTwoCrystalBalls_ballsNeverBreak(t *testing.T) {
	assert := assert.New(t)

	floors := []bool{false, false, false, false, false, false, false, false, false}
	index, err := TwoCrystalBalls(floors)
	assert.Equal(-1, index)
	assert.Error(err)
}
