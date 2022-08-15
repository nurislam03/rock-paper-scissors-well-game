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
		{"Rock Vs Paper - rock loses", 1, "rock", "paper", "lose"},
		{"Rock Vs Scissors - rock wins", 2, "rock", "scissors", "win"},
	}

	for _, tc := range testCases {
		testTemplate(t, mockRandGen, tc.name, tc.rand, tc.humanChoice, tc.botChoice, tc.result)
	}
}
