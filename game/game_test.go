package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testTemplate(t *testing.T, moc *MockRandGenerator, testName string, rand int, humanChoice, botChoice, result string) {
	t.Run(testName, func(t *testing.T) {
		moc.On("Rand", 4).Return(rand).Once()
		bp := NewBot(moc)
		hp := NewHuman(humanChoice)
		res := NewGame(hp, bp).Play()

		assert.Equal(t, humanChoice, res.Choice1.Name())
		assert.Equal(t, botChoice, res.Choice2.Name())
		assert.Equal(t, result, res.Result)
		moc.AssertExpectations(t)
	})
}

func TestGamePlay(t *testing.T) {

	mockRandGen := &MockRandGenerator{}
	testCases := []struct {
		name                           string
		rand                           int
		humanChoice, botChoice, result string
	}{
		{"Rock Vs Rock - match ties", 0, "rock", "rock", "tie"},
		{"Rock Vs Paper - rock loses", 1, "rock", "paper", "lose"},
		{"Rock Vs Scissors - rock wins", 2, "rock", "scissors", "win"},
		{"Rock Vs Well - rock loses", 3, "rock", "well", "lose"},

		{"Paper Vs Rock - Paper wins", 0, "paper", "rock", "win"},
		{"Paper Vs Paper - match ties", 1, "paper", "paper", "tie"},
		{"Paper Vs Scissors - Paper loses", 2, "paper", "scissors", "lose"},
		{"Paper Vs Well - Paper wins", 3, "paper", "well", "win"},

		{"Scissors Vs Rock - Scissors loses", 0, "scissors", "rock", "lose"},
		{"Scissors Vs Paper - Scissors wins", 1, "scissors", "paper", "win"},
		{"Scissors Vs Scissors - match ties", 2, "scissors", "scissors", "tie"},
		{"Scissors Vs Well - Scissors loses", 3, "scissors", "well", "lose"},

		{"Well Vs Rock - Well wins", 0, "well", "rock", "win"},
		{"Well Vs Paper - Well loses", 1, "well", "paper", "lose"},
		{"Well Vs Scissors - Well wins", 2, "well", "scissors", "win"},
		{"Well Vs Well - match ties", 3, "well", "well", "tie"},
	}

	for _, tc := range testCases {
		testTemplate(t, mockRandGen, tc.name, tc.rand, tc.humanChoice, tc.botChoice, tc.result)
	}
}
