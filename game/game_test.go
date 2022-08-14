package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockRandGen struct {
	num int
}

func (m MockRandGen) Rand(end int) int {
	return m.num
}

func TestGamePlay_RockVsPaper(t *testing.T) {
	mockRandGen := MockRandGen{1}
	bp := NewBot(mockRandGen)
	hp := NewHuman("rock")
	res := NewGame(hp, bp).Play()

	assert.Equal(t, "rock", res.Choice1.Name())
	assert.Equal(t, "paper", res.Choice2.Name())
	assert.Equal(t, "lose", res.Result)
}
