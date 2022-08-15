package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testTemplate(t *testing.T, testName, input string, result bool) {
	t.Run(testName, func(t *testing.T) {
		assert.Equal(t, result, validateGameCharacter(input))
	})
}

func TestValidateGameCharacter(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		result bool
	}{
		{"for input 'rock' it should return true", "rock", true},
		{"for input 'Rock' it should return true", "Rock", true},
		{"for input 'scissors' it should return true", "scissors", true},
		{"for input 'paper' it should return true", "paper", true},
		{"for input 'well' it should return true", "well", true},
		{"for input 'welcome' it should return false", "welcome", false},
		{"for input 'PaPeR' it should return true", "PaPeR", true},
		{"for input 'papers' it should return false", "papers", false},
		{"for input as 'empty string' it should return false", " ", false},
		{"for input '12345' it should return false", "12345", false},
	}

	for _, tc := range testCases {
		testTemplate(t, tc.name, tc.input, tc.result)
	}
}
